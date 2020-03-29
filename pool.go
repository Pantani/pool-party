package pool_party

import (
	"github.com/Pantani/errors"
	"github.com/Pantani/pool-party/bip39"
	"github.com/Pantani/pool-party/bip44"
)

type Pool struct {
	coin       bip44.Coin
	mnemonic   string
	passphrase string
}

func NewPool(coin bip44.Coin) *Pool {
	return &Pool{
		coin: coin,
	}
}

func NewPoolWithSecret(coin bip44.Coin, mnemonic, passphrase string) *Pool {
	return &Pool{
		coin:       coin,
		mnemonic:   mnemonic,
		passphrase: passphrase,
	}
}

// GenerateMnemonic generates a new mnemonic key based in the bit size
// It returns an error if occurs
func (p *Pool) GenerateMnemonic(bitSize int, passphrase string) error {
	// NewEntropy will create random entropy bytes
	// Return non-zero first byte, unless all random zeros occurs.
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return err
	}
	// generate (english) seed words based on the entropy
	p.mnemonic, err = bip39.NewMnemonic(entropy)
	p.passphrase = passphrase
	return err
}

// GenerateAddressPool generates the address pool based in the index and length
// It returns the generated addresses and an error if occurs
func (p *Pool) GenerateAddressPool(start, length int) (bip44.Addresses, error) {
	if len(p.mnemonic) == 0 {
		return nil, errors.E("empty mnemonic")
	}
	account, err := bip44.GenerateWallets(p.coin, p.mnemonic, p.passphrase, start, length)
	if err != nil {
		return nil, errors.E(err, "error to generate bip44 wallets", errors.Params{"coin": p.coin, "start": start, "length": length})
	}
	return account.Addresses, nil
}
