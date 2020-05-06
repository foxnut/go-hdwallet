package hdwallet

func init() {
	coins[QTUM] = newQTUM
}

type qtum struct {
	*btc
}

func newQTUM(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "Qtum"
	token.symbol = "QTUM"
	token.key.Opt.Params = &QTUMParams

	return &qtum{btc: token}
}
