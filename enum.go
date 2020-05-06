package hdwallet

// mnemonic language
const (
	English            = "english"
	ChineseSimplified  = "chinese_simplified"
	ChineseTraditional = "chinese_traditional"
)

// zero is deafult of uint32
const (
	Zero      uint32 = 0
	ZeroQuote uint32 = 0x80000000
	BTCToken  uint32 = 0x10000000
	ETHToken  uint32 = 0x20000000
)

// wallet type from bip44
const (
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md#registered-coin-types
	BTC        = ZeroQuote + 0
	BTCTestnet = ZeroQuote + 1
	LTC        = ZeroQuote + 2
	DOGE       = ZeroQuote + 3
	DASH       = ZeroQuote + 5
	ETH        = ZeroQuote + 60
	ETC        = ZeroQuote + 61
	BCH        = ZeroQuote + 145
	QTUM       = ZeroQuote + 2301

	// btc token
	USDT = BTCToken + 1

	// eth token
	IOST = ETHToken + 1
	USDC = ETHToken + 2
)

var coinTypes = map[uint32]uint32{
	USDT: BTC,
	IOST: ETH,
	USDC: ETH,
}
