package hdwallet

func init() {
	coins[BCH] = newBCH
}

type bch struct {
	*btc
}

func newBCH(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "Bitcoin Cash"
	token.symbol = "BCH"
	token.key.Opt.Params = &BCHParams

	return &bch{btc: token}
}
