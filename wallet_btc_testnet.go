package hdwallet

func init() {
	coins[BTCTestnet] = newBTCTestnet
}

type btcTestnet struct {
	name   string
	symbol string
	key    *Key
}

func newBTCTestnet(key *Key) Wallet {
	key.Opt.Params = &BTCTestnetParams
	return &btcTestnet{
		name:   "Bitcoin Testnet",
		symbol: "BTCTestnet",
		key:    key,
	}
}

func (c *btcTestnet) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *btcTestnet) GetName() string {
	return c.name
}

func (c *btcTestnet) GetSymbol() string {
	return c.symbol
}

func (c *btcTestnet) GetKey() *Key {
	return c.key
}

func (c *btcTestnet) GetAddress() (string, error) {
	return c.key.AddressBTC()
}
