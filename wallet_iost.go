package hdwallet

func init() {
	coins[IOST] = newIOST
}

type iost struct {
	*eth
}

func newIOST(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "Internet of Services"
	token.symbol = "IOST"
	token.contract = "0xfa1a856cfa3409cfa145fa4e20eb270df3eb21ab"

	return &iost{eth: token}
}
