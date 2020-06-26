package bip44

import (
	"encoding/hex"

	"github.com/Pantani/pool-party/bip39"

	"github.com/Pantani/errors"
	log "github.com/Pantani/logger"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
)

// bip44 creates a bip44 with count receive addresses, based on pkb (bip39 key)
// m/44'/cointype'/0'/0/i
func bip44(coin *Altcoin, start, qty int, pkb []byte) (*Account, error) {
	account := &Account{Coin: coin.Name, CoinType: coin.CoinType}
	net := &chaincfg.MainNetParams

	// cointype as specified in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	net.PubKeyHashAddrID = coin.PubKeyHashAddrID
	net.PrivateKeyID = coin.PrivateKeyID

	// See https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
	ext, err := hdkeychain.NewMaster(pkb, net)
	if err != nil {
		return nil, err
	}

	// m/44'

	// Child returns a derived child extended key at the given index.  When this
	// extended key is a private extended key (as determined by the IsPrivate
	// function), a private extended key will be derived.  Otherwise, the derived
	// extended key will be also be a public extended key.
	purpose, err := ext.Child(44 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, err
	}

	// m/44'/altcointype'
	coinType, err := purpose.Child(uint32(coin.CoinType) + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, err
	}

	// m/44'/altcointype'/0'
	acct0, err := coinType.Child(0 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, err
	}

	// Account extended private key (eg to import in electrum)
	account.Key = acct0

	// m/44'/altcointype'/0'/0
	// 0 = external accounts for receive addresses
	// 1 = internal accounts for change
	acct0External, err := acct0.Child(0)
	if err != nil {
		return nil, err
	}
	account.External = acct0External

	for i := start; i < (start + qty); i++ {
		receive, err := acct0External.Child(uint32(i))
		if err != nil {
			log.Error(err, "Failed to create receive address", log.Params{"i": i})
			continue
		}
		// ECPrivKey converts the extended key to a btcec private key and returns it.
		// As you might imagine this is only possible if the extended key is a private
		// extended key (as determined by the IsPrivate function).  The ErrNotPrivExtKey
		// error will be returned if this function is called on a public extended key.
		privk, err := receive.ECPrivKey()
		if err != nil {
			log.Error(err, "converts the extended key to a private key failed", log.Params{"i": i, "receive": receive})
			continue
		}

		// ECPubKey converts the extended key to a btcec public key and returns it.
		pubk, err := receive.ECPubKey()
		if err != nil {
			log.Error(err, "converts the extended key to a public key failed", log.Params{"i": i, "receive": receive})
			continue
		}

		// Ethereum and Energi addresses are handle differently
		if coin.Name == "Ethereum" || coin.Name == "Energi" {
			// create our address from the publickey
			address := crypto.PubkeyToAddress(*pubk.ToECDSA())

			// add the address to our addresses, with the pub and privkey as a string (compressed)
			account.Addresses = append(account.Addresses,
				Address{
					Address: address.String(),
					Pubkey:  "0x" + hex.EncodeToString(pubk.SerializeCompressed()),
					Privkey: "0x" + hex.EncodeToString(privk.Serialize()),
				})

			continue
		}

		// Address converts the extended key to a standard bitcoin pay-to-pubkey-hash
		// address for the passed network.
		address, err := receive.Address(net)
		if err != nil {
			log.Error(err, "address conversion failed", log.Params{"i": i, "receive": receive, "net": net})
			continue
		}

		// NewWIF creates a new WIF structure to export an address and its private key
		// as a string encoded in the Wallet Import Format.  The compress argument
		// specifies whether the address intended to be imported or exported was created
		// by serializing the public key compressed rather than uncompressed.
		wif, err := btcutil.NewWIF(privk, net, true)
		if err != nil {
			log.Error(err, "creates a new WIF structure failed", log.Params{"i": i, "receive": receive, "net": net})
			continue
		}
		account.Addresses = append(account.Addresses,
			Address{
				Address: address.String(),
				Pubkey:  hex.EncodeToString(pubk.SerializeCompressed()),
				Privkey: wif.String(),
			})
	}
	return account, nil
}

func GenerateWallets(coin Coin, mnemonic, passphrase string, start, qty int) (*Account, error) {
	if _, ok := CoinList[coin]; !ok {
		return nil, errors.E("Invalid coin", errors.Params{"coin": coin})
	}
	var (
		pk  *btcec.PrivateKey
		err error
	)
	wallet := &Wallet{Seedwords: mnemonic}

	// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
	// to create binary bip39 from the mnemonic, we use the PBKDF2 function with a mnemonic sentence (in UTF-8 NFKD)
	// used as the passphrase and the string "mnemonic" + passphrase (again in UTF-8 NFKD) used as the salt. i
	// The iteration count is set to 2048 and HMAC-SHA512 is used as the pseudo-random function.
	// The length of the derived key is 512 bits (= 64 bytes).
	// don't use an extra passphrase for the bip39
	wallet.Seed, err = bip39.NewSeedWithErrorChecking(mnemonic, passphrase)
	if err != nil {
		return nil, err
	}
	// hexencode the bip39 and give it to pkstr
	pkstr := hex.EncodeToString(wallet.Seed)

	// we have a user specified 256bits private key in hex OR we have a 512bits bip39 key in hex
	// set the key back to bytes
	pkb, err := hex.DecodeString(pkstr)
	if err != nil {
		return nil, err
	}

	// PrivKeyFromBytes returns a private and public key for `curve' based on the
	// private key passed as an argument as a byte slice.
	pk, _ = btcec.PrivKeyFromBytes(btcec.S256(), pkb)
	// we have to specify a net (use the mainnet (bitcoin) for now)
	net := &chaincfg.MainNetParams

	// NewMaster creates a new master node (bip32 root key) for use in creating a hierarchical
	// deterministic key chain.
	// see https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#Master_key_generation
	// pkb is the bip39 bip39 key
	if _, err := hdkeychain.NewMaster(pkb, net); err != nil {
		return nil, err
	}
	account, err := bip44(CoinList[coin], start, qty, pkb)
	if err != nil {
		return nil, err
	}
	account.PrivateKey = hex.EncodeToString(pk.Serialize())

	return account, err
}
