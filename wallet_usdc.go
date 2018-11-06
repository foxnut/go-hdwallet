package hdwallet

func init() {
	coins[USDC] = newUSDC
}

type usdc struct {
	*eth
}

func newUSDC(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "USD Coin"
	token.symbol = "USDC"
	token.contract = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"

	return &usdc{eth: token}
}
