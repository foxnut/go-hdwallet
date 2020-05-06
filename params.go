package hdwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// init net params
var (
	BTCParams        = chaincfg.MainNetParams
	BTCTestnetParams = chaincfg.TestNet3Params
	LTCParams        = chaincfg.MainNetParams
	DOGEParams       = chaincfg.MainNetParams
	DASHParams       = chaincfg.MainNetParams
	BCHParams        = chaincfg.MainNetParams
	QTUMParams       = chaincfg.MainNetParams
	USDTParams       = chaincfg.MainNetParams
)

func init() {
	// ltc net params
	// https://github.com/litecoin-project/litecoin/blob/master/src/chainparams.cpp
	LTCParams.Bech32HRPSegwit = "ltc"
	LTCParams.PubKeyHashAddrID = 0x30 // 48
	LTCParams.ScriptHashAddrID = 0x32 // 50
	LTCParams.PrivateKeyID = 0xb0     // 176

	// doge net params
	// https://github.com/dogecoin/dogecoin/blob/master/src/chainparams.cpp
	DOGEParams.PubKeyHashAddrID = 0x1e // 30
	DOGEParams.ScriptHashAddrID = 0x16 // 22
	DOGEParams.PrivateKeyID = 0x9e     // 158

	// dash net params
	// https://github.com/dashpay/dash/blob/master/src/chainparams.cpp
	DASHParams.PubKeyHashAddrID = 0x4c // 76
	DASHParams.ScriptHashAddrID = 0x10 // 16
	DASHParams.PrivateKeyID = 0xcc     // 204

	// bch net params
	// https://github.com/Bitcoin-ABC/bitcoin-abc/blob/master/src/chainparams.cpp
	BCHParams.PubKeyHashAddrID = 0x00 // 0
	BCHParams.ScriptHashAddrID = 0x05 // 5
	BCHParams.PrivateKeyID = 0x80     // 128

	// qtum net params
	// https://github.com/qtumproject/qtum/blob/master/src/chainparams.cpp
	QTUMParams.PubKeyHashAddrID = 0x3a // 58
	QTUMParams.ScriptHashAddrID = 0x32 // 50
	QTUMParams.PrivateKeyID = 0x80     // 128

	// usdt net params
	// https://github.com/OmniLayer/omnicore/blob/master/src/chainparams.cpp
	USDTParams.PubKeyHashAddrID = 0x00 // 0
	USDTParams.ScriptHashAddrID = 0x05 // 5
	USDTParams.PrivateKeyID = 0x80     // 128
}
