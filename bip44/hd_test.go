package bip44

import (
	"encoding/hex"
	"testing"

	"github.com/Pantani/pool-party/bip39"
)

func Test_bip44(t *testing.T) {
	type args struct {
		coin      *Altcoin
		start     int
		qty       int
		pkb       []byte
		seedwords string
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Test try the bip44 algorithm",
			args: args{
				coin:      CoinList["Ether"],
				start:     0,
				qty:       10,
				pkb:       nil,
				seedwords: "unknown art reflect enrich zero gloom success sunset slide end enroll again village approve upgrade rib obvious there license inspire plunge radar pull glue",
			},
			wantErr: false,
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
				Addresses: Addresses{
					{
						Address: "0x85fd0d299C08cffF429D48e62EAE71521743f4a1",
						Pubkey:  "0x02a6a7c54b710c130e8def7d5f1790e83d5f409790ace321cfbb0d6250474a880b",
						Privkey: "0xc01ce836cc9b8ac99c88d1dc52df9f2d5360fb5f691389927fa0bd5863a98763",
					},
					{
						Address: "0xD19F9555d918C3A66d4b59644eF24a2903382c80",
						Pubkey:  "0x03401b2173a4db38409ecee4c576ddfe39593792e099e5276e3b7bc83414d17da9",
						Privkey: "0xb321588c6adbc5164b794f792f26d7086bb211bdf4b6d5b472a07c71ae23a724",
					},
					{
						Address: "0xc0016B83a00e2B7b3C8fD63D514DF0d36Ad0a58C",
						Pubkey:  "0x034f356fd67975b2f4041ae8ccf73d3ba69bc5093181814b39e3448681d186d018",
						Privkey: "0x58c6b499a28b38b252f9156650dee959793d15303b2f04203a7029afbb851d79",
					},
					{
						Address: "0x7d74421391B6ae808C381E6A4507a4361C73Ae6d",
						Pubkey:  "0x0214c994f331d0c3463a92ba816020ca4b97de6a599645d573438dc64559fb0898",
						Privkey: "0xa32fbcf4da9bccc9439fa8fe2f7d7b2bc7fb9c6d525808c949399341d7a6d068",
					},
					{
						Address: "0x0c173A9e7f6AD1850792938dCA952BdE7925415C",
						Pubkey:  "0x026cf727ff28c34fdd005e85dc20deb07e542897c271e56a31189d686eec1623a5",
						Privkey: "0xc5d12ebdc4471747fbb95112f6437abd6a814c4fa9d740a0ebc40f0abcb26ba5",
					},
					{
						Address: "0x03B52Bd366f1421D585aaC6AEbC0F91ea1815607",
						Pubkey:  "0x02b0ed0a5eba834b5811b8d1d28b62b833f3635df26d5cd358b29df464bc4ecf2f",
						Privkey: "0xf0a5ac3b8902cd913838eac5fe9bab27559c1b725e27cb3e2d3a7250f2cb0ba3",
					},
					{
						Address: "0x48DF7A87206e5775BCdDA22A39431318f23577de",
						Pubkey:  "0x02a15f88f4edde4093ad6501a07637bdbeea3af4c66f00c649db3c5cc96c101c0c",
						Privkey: "0xb2cb565269792fee39acbdc1be99c17ddc98e0cbd1b5413006a3b5a8e14ae6b9",
					},
					{
						Address: "0x196215155a246C2173773a6BcfF3a4B878076510",
						Pubkey:  "0x033df6ef161d16cc38377b3ec49383f750d75ee4025c39224376c10deec2a9c7d4",
						Privkey: "0x49f7ec8bdca1fa7e36593e184f1001b1a824b71d7a18ae473ee2e5c5631741c3",
					},
					{
						Address: "0xfDc9F7fAdb7E87ed0cEF95Cb33d154f172b17e6F",
						Pubkey:  "0x03c741897ef72131f9d3067860dd0f8634501bda74e1664539beea65d6231b4cde",
						Privkey: "0xc4cb6c394713b2d477316651bcf62c963538ac96d52b9920f57c45ba10148612",
					},
					{
						Address: "0x852EfA8A045B53D8A5161cA48741e25Deac0BA2F",
						Pubkey:  "0x02fb27f84113d0380fb27429cd6f3a7c13ac1f55472271362ba40d61dec811ae9c",
						Privkey: "0x1027e4d0521853935a0348fd1323169ac9dd09f021688a68123a19a706185627",
					},
				},
			},
		},
		{
			name: "Test try the bip44 algorithm",
			args: args{
				coin:      CoinList["Ether"],
				start:     0,
				qty:       10,
				pkb:       nil,
				seedwords: "rent slogan lemon nerve soup annual depend shift olympic similar bounce wait often fury slush fish crazy bring police level economy crush can energy",
			},
			wantErr: false,
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
				Addresses: Addresses{
					{
						Address: "0x1D74465b33E8ca340d613704d5F1Dc301c6e2F44",
						Pubkey:  "0x0349d5bbcf174b5492e3b6fe10bb62aa23c0a860763c18d4b29eaa0fede84ebdb9",
						Privkey: "0x481b7b3193e885df653f510239d03e949d3733380d969705ec7a259eeee067c3",
					},
					{
						Address: "0xb358e6dd727CE7b49C318A6082D1A3DBfbc35Dd0",
						Pubkey:  "0x024e48f7b75910bdbf04ca81c4e91eaafe0a887e89b9b4f90ffaf714a318aed967",
						Privkey: "0x016ffe58df9562f092589e25607a94c12c98883946a2eb6fd9ee99f2e66a0604",
					},
					{
						Address: "0x3D95C140Ff29F6947a9183573E7DB6133e5fF9EC",
						Pubkey:  "0x02394b433de9a7837b0415261602d10f7f1bc554a6e68c1575df69394e67894f07",
						Privkey: "0xad2e59f2ecadd94d13ef3031c91259edf8bdb65e08b14662fc6ed7440c7512e5",
					},
					{
						Address: "0x7088CB866eF882357f32a6C83989FB837981c0F1",
						Pubkey:  "0x020453162c8a5939de66955e9fc622d4239ac5bf4943be3e3e6408eb78dcfb5dbd",
						Privkey: "0xeb6427d66c0759fed5e26935e4d6e6beeeafc6c41e746024d0a5165047a92221",
					},
					{
						Address: "0x8ec8C5AC7297a9338f29Df94772611C35DD424Fb",
						Pubkey:  "0x03db3d8c9ca8a41e0345495b39ca379658953e90b1f323a6f62f800c40f318d5c1",
						Privkey: "0x32a6927631a5c0931d5bb120c1732d5914110f451b90f1bb8ea5c2fab68df335",
					},
					{
						Address: "0x8383fb29EFe19d7cbdA4D4fBD8c407ab5168bBc0",
						Pubkey:  "0x0245b7fba0dfce6a38e9b8db379a5f4aeab1791354774081ca97824e4a1587fbd9",
						Privkey: "0x74d363c12cf31d3096af9d0c28dbf9c688b2d60d8569030a29340de19076521f",
					},
					{
						Address: "0x7b7Ac7Ac609944E78126A438211e6252CDC4C222",
						Pubkey:  "0x0394284b2379953985378c236894f9ea7958d5a85060fe95f2957b489c81436e0f",
						Privkey: "0x34623351a25e1be4badf4ab56f83dc8a05d3caee06326b84270657f2dd5066b6",
					},
					{
						Address: "0x544fAA7A491B9fEb267A7C53D58915E5094a033C",
						Pubkey:  "0x036666da1fa95aedfb3015fe64c07cccd60bb1c1f0639d2ea0861348db91c01310",
						Privkey: "0x18b626316a8ee660fc020698fbaac4d6d8336cb6f945171942658f802a3676f3",
					},
					{
						Address: "0x9Ce6c4E8d5e1d8C8628b7ff2d504D736E013E7AB",
						Pubkey:  "0x0242d18abdf8b31d7bb52fa124878e42afc72805cbc6de8e314ad87387427cb4cc",
						Privkey: "0xb60e4a796c32f29a89bedea9d9e3431c51b99497b17544d266c77532ee308727",
					},
					{
						Address: "0x03B7217843CCA41c5EC3dA1dEfDDC4a20715096d",
						Pubkey:  "0x02d03f3aabaa90981f7b08968e9cd1262c6fa1b8f222ac2bb3067d77230b141236",
						Privkey: "0xde1fccb3f3b89e958cdda78d9e32007ee8ea404d9bf852cbba8bed1ae58f77a7",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		wallet := &Wallet{}
		wallet.Seedwords = tt.args.seedwords
		wallet.Seed, _ = bip39.NewSeedWithErrorChecking(wallet.Seedwords, "")
		pkstr := hex.EncodeToString(wallet.Seed)
		tt.args.pkb, _ = hex.DecodeString(pkstr)
		t.Run(tt.name, func(t *testing.T) {
			got, err := bip44(tt.args.coin, tt.args.start, tt.args.qty, tt.args.pkb)
			if (err != nil) != tt.wantErr {
				t.Errorf("bip44() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.CoinType != tt.want.CoinType {
				t.Errorf("bip44() CoinType = %d, want %d", got.CoinType, tt.want.CoinType)
			}
			for i := 0; i < len(tt.want.Addresses)-1; i++ {
				gotAddress := tt.want.Addresses[i]
				wantAddress := tt.want.Addresses[i]
				if gotAddress != wantAddress {
					t.Errorf("bip44() Addresses = %v, want %v", gotAddress, wantAddress)
				}
			}
		})
	}
}
func Test_bip44WithWrongMnemonic(t *testing.T) {
	type args struct {
		coin      *Altcoin
		start     int
		qty       int
		pkb       []byte
		seedwords string
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Test bip44 algorithm with wrong mnemonic",
			args: args{
				coin:      CoinList["Ether"],
				start:     0,
				qty:       10,
				pkb:       nil,
				seedwords: "slogan lemon nerve soup annual depend shift olympic similar bounce wait often fury slush fish crazy bring police level economy crush can energy",
			},
			wantErr: true,
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
			},
		},
	}
	for _, tt := range tests {

		wallet := &Wallet{}
		wallet.Seedwords = tt.args.seedwords
		wallet.Seed, _ = bip39.NewSeedWithErrorChecking(wallet.Seedwords, "")
		pkstr := hex.EncodeToString(wallet.Seed)
		tt.args.pkb, _ = hex.DecodeString(pkstr)
		t.Run(tt.name, func(t *testing.T) {
			_, err := bip44(tt.args.coin, tt.args.start, tt.args.qty, tt.args.pkb)
			if (err != nil) != tt.wantErr {
				t.Errorf("bip44() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

const (
	mnemonic = "rent slogan lemon nerve soup annual depend shift olympic similar bounce wait often fury slush fish crazy bring police level economy crush can energy"
)

func TestGenerateWallets(t *testing.T) {
	type args struct {
		seedwords string
		password  string
		coin      Coin
		start     int
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Test try to generate hd wallets",
			args: args{
				start: 0,
				qty:   10,
			},
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
				Addresses: Addresses{
					{
						Address: "0x85fd0d299C08cffF429D48e62EAE71521743f4a1",
						Pubkey:  "0x02a6a7c54b710c130e8def7d5f1790e83d5f409790ace321cfbb0d6250474a880b",
						Privkey: "0xc01ce836cc9b8ac99c88d1dc52df9f2d5360fb5f691389927fa0bd5863a98763",
					},
					{
						Address: "0xD19F9555d918C3A66d4b59644eF24a2903382c80",
						Pubkey:  "0x03401b2173a4db38409ecee4c576ddfe39593792e099e5276e3b7bc83414d17da9",
						Privkey: "0xb321588c6adbc5164b794f792f26d7086bb211bdf4b6d5b472a07c71ae23a724",
					},
					{
						Address: "0xc0016B83a00e2B7b3C8fD63D514DF0d36Ad0a58C",
						Pubkey:  "0x034f356fd67975b2f4041ae8ccf73d3ba69bc5093181814b39e3448681d186d018",
						Privkey: "0x58c6b499a28b38b252f9156650dee959793d15303b2f04203a7029afbb851d79",
					},
					{
						Address: "0x7d74421391B6ae808C381E6A4507a4361C73Ae6d",
						Pubkey:  "0x0214c994f331d0c3463a92ba816020ca4b97de6a599645d573438dc64559fb0898",
						Privkey: "0xa32fbcf4da9bccc9439fa8fe2f7d7b2bc7fb9c6d525808c949399341d7a6d068",
					},
					{
						Address: "0x0c173A9e7f6AD1850792938dCA952BdE7925415C",
						Pubkey:  "0x026cf727ff28c34fdd005e85dc20deb07e542897c271e56a31189d686eec1623a5",
						Privkey: "0xc5d12ebdc4471747fbb95112f6437abd6a814c4fa9d740a0ebc40f0abcb26ba5",
					},
					{
						Address: "0x03B52Bd366f1421D585aaC6AEbC0F91ea1815607",
						Pubkey:  "0x02b0ed0a5eba834b5811b8d1d28b62b833f3635df26d5cd358b29df464bc4ecf2f",
						Privkey: "0xf0a5ac3b8902cd913838eac5fe9bab27559c1b725e27cb3e2d3a7250f2cb0ba3",
					},
					{
						Address: "0x48DF7A87206e5775BCdDA22A39431318f23577de",
						Pubkey:  "0x02a15f88f4edde4093ad6501a07637bdbeea3af4c66f00c649db3c5cc96c101c0c",
						Privkey: "0xb2cb565269792fee39acbdc1be99c17ddc98e0cbd1b5413006a3b5a8e14ae6b9",
					},
					{
						Address: "0x196215155a246C2173773a6BcfF3a4B878076510",
						Pubkey:  "0x033df6ef161d16cc38377b3ec49383f750d75ee4025c39224376c10deec2a9c7d4",
						Privkey: "0x49f7ec8bdca1fa7e36593e184f1001b1a824b71d7a18ae473ee2e5c5631741c3",
					},
					{
						Address: "0xfDc9F7fAdb7E87ed0cEF95Cb33d154f172b17e6F",
						Pubkey:  "0x03c741897ef72131f9d3067860dd0f8634501bda74e1664539beea65d6231b4cde",
						Privkey: "0xc4cb6c394713b2d477316651bcf62c963538ac96d52b9920f57c45ba10148612",
					},
					{
						Address: "0x852EfA8A045B53D8A5161cA48741e25Deac0BA2F",
						Pubkey:  "0x02fb27f84113d0380fb27429cd6f3a7c13ac1f55472271362ba40d61dec811ae9c",
						Privkey: "0x1027e4d0521853935a0348fd1323169ac9dd09f021688a68123a19a706185627",
					},
				},
				PrivateKey: "439c690c5efaa2d3943fd05aec4ccad6e45aa4d357ebd4b1f4f9402aba6e45985842448716cce80ed6b8485e7c31bbefbf6e0db0ff9fa23011092d16e200bd69",
			},
			wantErr: false,
		},
		{
			name: "Test try to generate hd wallets",
			args: args{
				start: 0,
				qty:   10,
			},
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
				Addresses: Addresses{
					{
						Address: "0x1D74465b33E8ca340d613704d5F1Dc301c6e2F44",
						Pubkey:  "0x0349d5bbcf174b5492e3b6fe10bb62aa23c0a860763c18d4b29eaa0fede84ebdb9",
						Privkey: "0x481b7b3193e885df653f510239d03e949d3733380d969705ec7a259eeee067c3",
					},
					{
						Address: "0xb358e6dd727CE7b49C318A6082D1A3DBfbc35Dd0",
						Pubkey:  "0x024e48f7b75910bdbf04ca81c4e91eaafe0a887e89b9b4f90ffaf714a318aed967",
						Privkey: "0x016ffe58df9562f092589e25607a94c12c98883946a2eb6fd9ee99f2e66a0604",
					},
					{
						Address: "0x3D95C140Ff29F6947a9183573E7DB6133e5fF9EC",
						Pubkey:  "0x02394b433de9a7837b0415261602d10f7f1bc554a6e68c1575df69394e67894f07",
						Privkey: "0xad2e59f2ecadd94d13ef3031c91259edf8bdb65e08b14662fc6ed7440c7512e5",
					},
					{
						Address: "0x7088CB866eF882357f32a6C83989FB837981c0F1",
						Pubkey:  "0x020453162c8a5939de66955e9fc622d4239ac5bf4943be3e3e6408eb78dcfb5dbd",
						Privkey: "0xeb6427d66c0759fed5e26935e4d6e6beeeafc6c41e746024d0a5165047a92221",
					},
					{
						Address: "0x8ec8C5AC7297a9338f29Df94772611C35DD424Fb",
						Pubkey:  "0x03db3d8c9ca8a41e0345495b39ca379658953e90b1f323a6f62f800c40f318d5c1",
						Privkey: "0x32a6927631a5c0931d5bb120c1732d5914110f451b90f1bb8ea5c2fab68df335",
					},
					{
						Address: "0x8383fb29EFe19d7cbdA4D4fBD8c407ab5168bBc0",
						Pubkey:  "0x0245b7fba0dfce6a38e9b8db379a5f4aeab1791354774081ca97824e4a1587fbd9",
						Privkey: "0x74d363c12cf31d3096af9d0c28dbf9c688b2d60d8569030a29340de19076521f",
					},
					{
						Address: "0x7b7Ac7Ac609944E78126A438211e6252CDC4C222",
						Pubkey:  "0x0394284b2379953985378c236894f9ea7958d5a85060fe95f2957b489c81436e0f",
						Privkey: "0x34623351a25e1be4badf4ab56f83dc8a05d3caee06326b84270657f2dd5066b6",
					},
					{
						Address: "0x544fAA7A491B9fEb267A7C53D58915E5094a033C",
						Pubkey:  "0x036666da1fa95aedfb3015fe64c07cccd60bb1c1f0639d2ea0861348db91c01310",
						Privkey: "0x18b626316a8ee660fc020698fbaac4d6d8336cb6f945171942658f802a3676f3",
					},
					{
						Address: "0x9Ce6c4E8d5e1d8C8628b7ff2d504D736E013E7AB",
						Pubkey:  "0x0242d18abdf8b31d7bb52fa124878e42afc72805cbc6de8e314ad87387427cb4cc",
						Privkey: "0xb60e4a796c32f29a89bedea9d9e3431c51b99497b17544d266c77532ee308727",
					},
					{
						Address: "0x03B7217843CCA41c5EC3dA1dEfDDC4a20715096d",
						Pubkey:  "0x02d03f3aabaa90981f7b08968e9cd1262c6fa1b8f222ac2bb3067d77230b141236",
						Privkey: "0xde1fccb3f3b89e958cdda78d9e32007ee8ea404d9bf852cbba8bed1ae58f77a7",
					},
				},
				PrivateKey: "439c690c5efaa2d3943fd05aec4ccad6e45aa4d357ebd4b1f4f9402aba6e45985842448716cce80ed6b8485e7c31bbefbf6e0db0ff9fa23011092d16e200bd69",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateWallets(Ether, mnemonic, "", tt.args.start, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateWallets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.CoinType != tt.want.CoinType {
				t.Errorf("GenerateWallets() CoinType = %d, want %d", got.CoinType, tt.want.CoinType)
			}
			if got.PrivateKey != tt.want.PrivateKey {
				t.Errorf("GenerateWallets() PrivateKey = %v, want %v", got.PrivateKey, tt.want.PrivateKey)
			}
			for i := 0; i < len(tt.want.Addresses)-1; i++ {
				gotAddress := tt.want.Addresses[i]
				wantAddress := tt.want.Addresses[i]
				if gotAddress != wantAddress {
					t.Errorf("GenerateWallets() Addresses = %v, want %v", gotAddress, wantAddress)
				}
			}
		})
	}
}

func TestGenerateWalletsWithWrongMnemonic(t *testing.T) {
	type args struct {
		mnemonic   string
		passphrase string
		coin       Coin
		start      int
		qty        int
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Test try to generate hd wallets with wrong mnemonic",
			args: args{
				mnemonic:   "slogan lemon nerve soup annual depend shift olympic similar bounce wait often fury slush fish crazy bring police level economy crush can energy",
				passphrase: "",
				coin:       Ether,
				start:      0,
				qty:        10,
			},
			want: &Account{
				Coin:     "Ether",
				CoinType: 60,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateWallets(tt.args.coin, tt.args.mnemonic, tt.args.passphrase, tt.args.start, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateWallets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
