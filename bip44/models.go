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
	Ethereum Coin = "Ethereum"
	Energi   Coin = "Energi"
	Bitcoin  Coin = "Bitcoin"
	Litecoin Coin = "Litecoin"
	Dash     Coin = "Dash"
	Dogecoin Coin = "Dogecoin"
)

var CoinList = map[Coin]*Altcoin{
	Ethereum: {"Ethereum", 0xff, 0xff, 60},
	Energi:   {"Energi", 0xff, 0xff, 39797},
	Bitcoin:  {"Bitcoin", 0x00, 0x80, 0},
	Litecoin: {"Litecoin", 0x30, 0xb0, 2},
	Dash:     {"Dash", 0x4c, 0xcc, 5},
	Dogecoin: {"Dogecoin", 0x1e, 0x9e, 3},
}
