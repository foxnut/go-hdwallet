package hdwallet

func init() {
	coins[LTC] = newLTC
}

type ltc struct {
	*btc
}

func newLTC(key *Key) Wallet {
	key.Opt.Params = &LTCParams
	token := newBTC(key).(*btc)
	token.name = "Litecoin"
	token.symbol = "LTC"

	return &ltc{btc: token}
}
