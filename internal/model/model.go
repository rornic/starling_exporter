package model

type AccountsResponse struct {
	Accounts []AccountV2 `json:"accounts"`
}

type AccountV2 struct {
	AccountUid      string `json:"accountUid"`
	DefaultCategory string `json:"defaultCategory"`
}

type BalanceResponse struct {
	TotalClearedBalance   CurrencyAndAmount `json:"totalClearedBalance"`
	TotalEffectiveBalance CurrencyAndAmount `json:"totalEffectiveBalance"`
	PendingTransations    CurrencyAndAmount `json:"pendingTransactions"`
}

type FeedItemsResponse struct {
	FeedItems []FeedItem `json:"FeedItems"`
}

type FeedItem struct {
	FeedItemUid      string            `json:"feedItemUid"`
	Amount           CurrencyAndAmount `json:"amount"`
	CounterPartyName string            `json:"counterPartyName"`
	SpendingCategory string            `json:"spendingCategory"`
	Direction        string            `json:"direction"`
}

type CurrencyAndAmount struct {
	Currency   string
	MinorUnits int64
}
