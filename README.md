# Pool-Party

Multi-coin address pool generation.

- Initialization:
```go
pool := pool_party.NewPoolWithSecret(bip44.Ether, "simple habit juice brush blush derive biology busy clown sister maple recipe", "1234")
logger.Info("pool with secret", logger.Params{"params": pool})

// OR

pool = pool_party.NewPool(bip44.Ether)

// Generate new mnemonic with 128 bits
err := pool.GenerateMnemonic(128, "")
if err != nil {
    logger.Panic(err)
}
logger.Info("mnemonic generated", logger.Params{"params": pool})

// Generate new mnemonic with 256 bits
err = pool.GenerateMnemonic(256, "")
if err != nil {
    logger.Panic(err)
}
logger.Info("mnemonic generated", logger.Params{"params": pool})
```

- Generate addresses:
```go
// Generate 100 address starting by index 50 (50 - 150)
start := 50
length := 100
result, err := pool.GenerateAddressPool(start, length)
if err != nil {
    logger.Panic(err)
}

logger.Info("address pool generated", logger.Params{"length": len(result)})
for i, addr := range result {
    logger.Info("new address", logger.Params{"index": start + i, "address": addr.Address})
}
```