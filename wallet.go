package hdwallet

var coins = make(map[uint32]func(*Key) Wallet)

// Wallet interface
type Wallet interface {
	GetType() uint32
	GetName() string
	GetSymbol() string
	GetKey() *Key
	GetAddress() (string, error)
}
