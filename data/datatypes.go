package data

const NMAX int = 100          // Maximum number of cryptocurrencies
const NMAXHistory int = 10000 // Maximum number of transactions

// tCrypto represents a cryptocurrency with its Name and price
type TCrypto struct {
	Name  string
	Price float64
}

// tCryptoList represents a list of cryptocurrencies
type TabCrypto [NMAX]TCrypto

type THistory struct {
	Name            string
	BuyPrice        float64
	Quantity        int
	TotalPrice      float64
	TransactionType string // "buy" or "sell"
}

type TabHistory [NMAXHistory]THistory

type TPorto struct {
	Name        string
	Quantity    int
	BuyPrice    float64
	MarketPrice float64
}
type TabPorto [NMAX]TPorto
