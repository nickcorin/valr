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

// TransactionType defines the kind of a transaction.
type TransactionType string

// TransactionType constants which may be used to filter lists of transactions.
const (
	TransactionTypeLimitBuy               TransactionType = "LIMIT_BUY"
	TransactionTypeLimitSell              TransactionType = "LIMIT_SELL"
	TransactionTypeMarketBuy              TransactionType = "MARKET_BUY"
	TransactionTypeMarketSell             TransactionType = "MARKET_SELL"
	TransactionTypeSimpleBuy              TransactionType = "SIMPLE_BUY"
	TransactionTypeSimpleSell             TransactionType = "SIMPLE_SELL"
	TransactionTypeMakerReward            TransactionType = "MAKER_REWARD"
	TransactionTypeBlockchainReceive      TransactionType = "BLOCKCHAIN_RECEIVE"
	TransactionTypeBlockchainSend         TransactionType = "BLOCKCHAIN_SEND"
	TransactionTypeFiatDeposit            TransactionType = "FIAT_DEPOSIT"
	TransactionTypeFiatWithdrawal         TransactionType = "FIAT_WITHDRAWAL"
	TransactionTypeReferralRebate         TransactionType = "REFERRAL_REBATE"
	TransactionTypeReferralReward         TransactionType = "REFERRAL_REWARD"
	TransactionTypePromotionalRebate      TransactionType = "PROMOTIONAL_REBATE"
	TransactionTypeInternalTransfer       TransactionType = "INTERNAL_TRANSFER"
	TransactionTypeFiatWithdrawalReversal TransactionType = "FIAT_WITHDRAWAL_REVERSAL"
)

// TransactionHistoryRequest contains the parameters available for querying an
// account's transaction history. All parameters are optional.
type TransactionHistoryRequest struct {
	// BeforeID indicates to only include transactions which occurred before the
	// transaction with this ID.
	BeforeID string

	// Currency indicates to only include transactions of this currency.
	Currency string `schema:"currency,omitempty"`

	// EndTime indicates to only include transactions before this ISO 8601 end
	// time.
	EndTime time.Time `schema:"endTime,omitempty"`

	// Limit indicates the number of items to be returned. Maximum value of 200.
	Limit int `schema:"limit,omitempty"`

	// Skip indicates the amount of transactions to skip when compiling the
	// list. The offset for paginated queries.
	Skip int `schema:"skip,omitempty"`

	// StartTime indicates to only include transactions after this ISO 8601
	// start time.
	StartTime time.Time `schema:"startTime,omitempty"`

	// Types indicates to include transaction of any type included in this list.
	Types []TransactionType `schema:"types,omitempty"`
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
