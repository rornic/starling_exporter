package metrics

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rornic/starlingexporter/internal/client"
)

var (
	balanceGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "starling_balance",
		Help: "The current balance of a Starling account",
	}, []string{"account", "balance_type"})
	transactionHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "starling_transactions",
		Help:    "A histogram of transactions on a Starling account",
		Buckets: []float64{1, 5, 10, 50, 100, 1000, 10000},
	}, []string{"account", "counter_party", "spending_category", "direction"})
)

func Record(client client.StarlingClient) {
	go func() {
		for {
			accounts, err := client.Accounts()
			if err != nil {
				slog.Error(fmt.Sprintf("error getting accounts: %v", err))
				time.Sleep(1 * time.Minute)
				continue
			}

			for _, account := range accounts.Accounts {
				accountId := account.AccountUid
				recordBalance(accountId, client)
				recordTransactions(accountId, account.DefaultCategory, client)
			}

			time.Sleep(15 * time.Second)
		}
	}()
}

func recordBalance(account string, client client.StarlingClient) {
	balance, err := client.Balance(account)
	if err != nil {
		slog.Error(fmt.Sprintf("error getting balance: %v", err))
		return
	}
	balanceGauge.WithLabelValues(account, "cleared").Set(float64(balance.TotalClearedBalance.MinorUnits) / 100.0)
	balanceGauge.WithLabelValues(account, "available").Set(float64(balance.TotalEffectiveBalance.MinorUnits) / 100.0)
	balanceGauge.WithLabelValues(account, "pending").Set(float64(balance.PendingTransations.MinorUnits) / 100.0)
}

func startOfMonth() time.Time {
	today := time.Now().UTC()
	return today.AddDate(0, 0, -today.Day()+1)
}

func recordTransactions(account string, category string, client client.StarlingClient) {
	feedItems, err := client.FeedItems(account, category, startOfMonth())
	if err != nil {
		slog.Error(fmt.Sprintf("error getting feed items: %v", err))
		return
	}

	transactionHistogram.Reset()
	for _, feedItem := range feedItems.FeedItems {
		transactionHistogram.WithLabelValues(account, feedItem.CounterPartyName, feedItem.SpendingCategory, feedItem.Direction).Observe(float64(feedItem.Amount.MinorUnits) / 100.0)
	}
}
