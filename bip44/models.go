package bip44

import (
	"github.com/btcsuite/btcutil/hdkeychain"
)

type Addresses []Address

type Address struct {
	Address string
	Pubkey  string
	Privkey string
}

type Account struct {
	Coin       string
	CoinType   int
	Key        *hdkeychain.ExtendedKey // bip44 extended key (m/44'/cointype'/0')
	External   *hdkeychain.ExtendedKey // external extended key (m/44'/cointype'/0'/0)
	Addresses  Addresses
	PrivateKey string
}

type Wallet struct {
	Seed      []byte                  // bip39 64byte (512bits) bip39
	Seedwords string                  // bip39 mnemonic
	Masterkey *hdkeychain.ExtendedKey // bip32 master (root) key
	Accounts  []Account               // different coin accounts
}

type Altcoin struct {
	Name             string
	PubKeyHashAddrID byte
	PrivateKeyID     byte
	CoinType         int
}

type Coin string

const (
	Ether    Coin = "Ether"
	Bitcoin  Coin = "Bitcoin"
	Litecoin Coin = "Litecoin"
)

var CoinList = map[Coin]*Altcoin{
	Ether:    {"Ether", 0xff, 0xff, 60},
	Bitcoin:  {"Bitcoin", 0x00, 0x80, 0},
	Litecoin: {"Litecoin", 0x30, 0xb0, 2},
}
