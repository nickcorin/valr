package valr

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type Balance struct {
	Available string    `json:"available"`
	Currency  string    `json:"currency"`
	Reserved  string    `json:"reserved"`
	Total     string    `json:"total"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Balances satisfies the PrivateClient interface.
func (c *client) Balances(ctx context.Context) ([]Balance, error) {
	res, err := c.httpClient.Get(ctx, "/account/balances", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account balances: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch account balances: %d status "+
			"code received", res.StatusCode)
	}

	var balances []Balance
	if err = res.JSON(&balances); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account balances: %w", err)
	}

	return balances, nil
}

// Trade contains information regarding a single trade which has been executed.
type Trade struct {
	CurrencyPair string    `json:"currencyPair"`
	ID           int64     `json:"tradeId"`
	Price        string    `json:"price"`
	Quantity     string    `json:"quantity"`
	Side         string    `json:"side"`
	TradedAt     time.Time `json:"tradedAt"`
}

// TradeHistory gets the last 100 recent trades for a given currency pair for
// your account.
func (c *client) TradeHistory(ctx context.Context, pair string) ([]Trade,
	error) {
	res, err := c.httpClient.Get(ctx, fmt.Sprintf("/account/%s/tradehistory",
		pair), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction history: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch transaction history: %d "+
			"status code recevied", res.StatusCode)
	}

	var trades []Trade
	if err = res.JSON(&trades); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trades: %w", err)
	}

	return trades, nil
}

// Transaction contains information regarding certain activities of your
// wallets.
type Transaction struct {
	AdditionalInfo *TransactionInfo     `json:"additionalInfo,omitempty"`
	CreditCurrency string               `json:"creditCurrency,omitempty"`
	CreditValue    string               `json:"creditValue,omitempty"`
	DebitCurrency  string               `json:"debitCurrency,omitempty"`
	DebitValue     string               `json:"debitValue,omitempty"`
	FeeCurrency    string               `json:"feeCurrency,omitempty"`
	FeeValue       string               `json:"feeValue,omitempty"`
	EventAt        time.Time            `json:"eventAt,omitempty"`
	TypeInfo       *TransactionTypeInfo `json:"transactionType,omitempty"`
}

// TransactionInfo contains additional information regarding Transactions.
type TransactionInfo struct {
	CostPerCoin        float64 `json:"costPerCoin,omitempty"`
	CostPerCoinSymbol  string  `json:"costPerCoinSymbol,omitempty"`
	CurrencyPairSymbol string  `json:"currencyPairSymbol,omitempty"`
	OrderID            string  `json:"orderID,omitempty"`
}

// TransactionTypeInfo contains additional information regarding the type of a
// transaction.
type TransactionTypeInfo struct {
	Type        TransactionType `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
}

// TransactionHistory satisfies the PrivateClient interface.
func (c *client) TransactionHistory(ctx context.Context,
	req *TransactionHistoryRequest) ([]Transaction, error) {

	params := make(url.Values)
	if err := c.encoder.Encode(req, params); err != nil {
		return nil, fmt.Errorf("failed to encode request params: %w", err)
	}

	res, err := c.httpClient.Get(ctx, "/account/transactionhistory", params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction history: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch transaction history: %d "+
			"status code recevied", res.StatusCode)
	}

	var transactions []Transaction
	if err = res.JSON(&transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transactions: %w", err)
	}

	return transactions, nil
}
