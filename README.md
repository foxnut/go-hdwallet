## go-hdwallet

A multi-cryptocurrency HD wallet implementated by golang.

## supported coins

- BTC
- LTC
- DOGE
- DASH
- ETH
- ETC
- BCH
- QTUM
- USDT
- IOST
- USDC

## install

```sh
go get -v -u github.com/foxnut/go-hdwallet
```

## example

```go
func main() {
package main

import (
  "fmt"

  "github.com/foxnut/go-hdwallet"
)

var (
  mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

func main() {
  master, err := hdwallet.NewKey(
    hdwallet.Mnemonic(mnemonic),
  )
  if err != nil {
    panic(err)
  }

  // BTC: 1AwEPfoojHnKrhgt1vfuZAhrvPrmz7Rh4
  wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC))
  address, _ := wallet.GetAddress()
  fmt.Println("BTC: ", address)

  // BCH: 1CSBT18sjcCwLCpmnnyN5iqLc46Qx7CC91
  wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.BCH))
  address, _ = wallet.GetAddress()
  fmt.Println("BCH: ", address)

  // LTC: LLCaMFT8AKjDTvz1Ju8JoyYXxuug4PZZmS
  wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.LTC))
  address, _ = wallet.GetAddress()
  fmt.Println("LTC: ", address)

  // DOGE: DHLA3rJcCjG2tQwvnmoJzD5Ej7dBTQqhHK
  wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.DOGE))
  address, _ = wallet.GetAddress()
  fmt.Println("DOGE:", address)

  // ETH: 0x37039021cBA199663cBCb8e86bB63576991A28C1
  wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
  address, _ = wallet.GetAddress()
  fmt.Println("ETH: ", address)

  // ETC: 0x480C69E014C7f018dAbF17A98273e90f0b0680cf
  wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETC))
  address, _ = wallet.GetAddress()
  fmt.Println("ETC: ", address)
}
```
