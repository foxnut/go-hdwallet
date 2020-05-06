package hdwallet

func init() {
	coins[DOGE] = newDOGE
}

type doge struct {
	*btc
}

func newDOGE(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "Dogecoin"
	token.symbol = "DOGE"
	token.key.Opt.Params = &DOGEParams

	return &doge{btc: token}
}
