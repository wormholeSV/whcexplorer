package view

import (
	"github.com/shopspring/decimal"
)

type BurnInfo struct {
	TxHash        string
	BlockTime     int
	Address       string
	Whc           *decimal.Decimal
	Bch           *decimal.Decimal
	TxBlockHeight uint32
	Process       string
	MatureTime    int
}

type BurnSummary struct {
	Total *decimal.Decimal
	Avail *decimal.Decimal
}

type FeeRateSummary struct {
	MinFeeRate *decimal.Decimal `json:"min_fee_rate"`
	AvgFeeRate *decimal.Decimal `json:"avg_fee_rate"`
	MaxFeeRate *decimal.Decimal `json:"max_fee_rate"`
}

type AddressBalance struct {
	Address          string
	BalanceAvailable *decimal.Decimal `json:"balance_available"`
	Status           int
}