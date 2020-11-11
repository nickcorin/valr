package valr

import (
	"time"
)

// Public API
// -----------------------------------------------------------------------------

// OrderBookRequest contains the request parameters for obtaining the order book
// for a given currency, or for a non-aggregated full order book.
//
// GET /public/{pair}/orderbook
// GET /marketdata/{pair}/orderbook/full
type OrderBookRequest struct {
	Pair string
}

// MarketSummaryRequest contains the request parameters for obtaining a market
// summary for a given currency pair.
//
// GET /public/{pair}/marketsummary
type MarketSummaryRequest struct {
	Pair string
}

// OrderTypesRequest contains the request parameters for obtaining a list of
// supported order types for a currency pair.
//
// GET /public/{pair}/ordertypes
type OrderTypesRequest struct {
	Pair string
}

// Accounts
// -----------------------------------------------------------------------------

// AccountTradeHistoryRequest contains the request parameters for obtaining the
// trade history for a given currency pair for your account, or for the market
// in general.
//
// GET /account/{pair}/tradehistory
// GET /marketdata/{pair}/tradehistory
type TradeHistoryRequest struct {
	Pair  string
	Limit int `schema:"limit,omitempty"`
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

// TransactionHistoryRequest contains the request parameters for obtaining the
// transaction history for your account.
//
// GET /account/transactionhistory
type TransactionHistoryRequest struct {
	BeforeID  string            `schema:"beforeId,omitempty"`
	Currency  string            `schema:"currency,omitempty"`
	EndTime   time.Time         `schema:"endTime,omitempty"`
	Limit     int               `schema:"skip,omitempty"`
	Offset    int               `schema:"limit,omitempty"`
	StartTime time.Time         `schema:"startTime,omitempty"`
	Types     []TransactionType `schema:"transactionTypes,omitempty"`
}

// Crypto
// -----------------------------------------------------------------------------

// DepositAddressRequest contains the request paremeters for getting the default
// deposit address associated with a given currency.
//
// GET /wallet/crypto/{currency}/deposit/address
type CryptoDepositAddressRequest struct {
	Currency string
}

// CryptoDepositHistoryRequest contains the request parameters for getting the
// deposit history records for a given currency.
//
// GET /wallet/crypto/{currency}/deposit/history
type CryptoDepositHistoryRequest struct {
	Currency string
	Limit    int `schema:"limit,omitempty"`
	Offset   int `schema:"skip,omitempty"`
}

// CryptoWithdrawalHistory contains the request parameters for getting the
// withdrawal history records for a given currency.
//
// GET /wallet/crypto/{currency}/withdraw/history
type CryptoWithdrawalHistoryRequest struct {
	Currency string
	Limit    int `schema:"limit,omitempty"`
	Offset   int `schema:"skip,omitempty"`
}

// CryptoWithdrawalInfoRequest contains the request parameters for getting
// information about withdrawing a given currency from your VALR account.
//
// GET /wallet/crypto/{currency}/withdraw
type CryptoWithdrawalInfoRequest struct {
	Currency string
}

// CryptoWithdrawalRequest contains the request parameters for creating a
// crypto withdrawals.
//
// POST /wallet/crypto/{currency}/withdrawal
type CryptoWithdrawalRequest struct {
	Amount   string `json:"amount"`
	Address  string `json:"address"`
	Currency string
}

// CryptoWithdrawalStatusRequest contains the request paremeters for getting the
// status of a crypto withdrawal.
//
// GET /wallet/crypto/{currency}/withdrawal/{id}
type CryptoWithdrawalStatusRequest struct {
	Currency string
	ID       string
}

// Exchange
// -----------------------------------------------------------------------------

// CancelOrderRequest contains the request parameters for cancelling an open
// order on the exchange. Only one of the order ID types should be provided.
//
// DELETE /orders/order
type CancelOrderRequest struct {
	CustomerOrderID string `json:"customerOrderId"`
	OrderID         string `json:"orderId"`
	Pair            string `json:"pair"`
}

// LimitOrderRequest contains the request parameters for placing a limit order
// on the exchange.
//
// POST /orders/limit
type LimitOrderRequest struct {
	CustomerOrderID string `json:"customerOrderId"`
	Pair            string `json:"pair"`
	PostOnly        bool   `json:"postOnly"`
	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	Side            string `json:"side"`
}

// MarketOrderRequest contains the request parameters for placing a market
// order on the exchange.
//
// POST /orders/market
type MarketOrderRequest struct {
	BaseAmount      string `json:"baseAmount"`
	CustomerOrderID string `json:"customerOrderId"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
}

// OrderHistoryRequest contains the request parameters for getting your order
// history.
//
// GET /orders/history
type OrderHistoryRequest struct {
	Limit  int
	Offset int
}

// OrderHistoryDetailRequest contains the request parameters for getting the
// detailed history of an order's statuses ordered by time descending. Only
// one of the order ID types should be provided.
//
// GET /orders/history/detail/customerorderid/{orderId}
// GET /orders/history/detail/orderid/{orderId}
type OrderHistoryDetailRequest struct {
	CustomerOrderID string
	OrderID         string
}

// OrderHistorySummaryRequest contains the request parameters for getting a
// summary for an order with one of the following statuses: "Filled",
// "Cancelled" or "Failed". Only one of the order ID types should be provided.
//
// GET /orders/history/summary/customerorderid/{orderId}
// GET /orders/history/summary/orderid/{orderId}
type OrderHistorySummaryRequest struct {
	CustomerOrderID string
	OrderID         string
}

// OrderStatusRequest contains the request parameters for getting the status of
// an order on the exchange, queried by the order ID. Only one of the order ID
// types should be provided.
//
// GET /orders/{pair}/customerorderid/{orderId}
// GET /orders/{pair}/{orderId}
type OrderStatusRequest struct {
	CustomerOrderID string
	OrderID         string
	Pair            string
}

// Fiat
// -----------------------------------------------------------------------------

// BankAccountsRequest contains the request parameters for getting a list of
// bank accounts that are linked to your VALR account.
//
// GET /wallet/fiat/{currency}/accounts
type BankAccountsRequest struct {
	Currency string
}

// FiatWithdrawalRequest contains the request parameters for withdrawing your
// ZAR funds into one of your linked bank accounts.
//
// POST /wallet/fiat/{currency}/withdraw
type FiatWithdrawalRequest struct {
	Amount      string `json:"amount"`
	BankAccount string `json:"linkedBankAccountId"`
	Currency    string
	Fast        bool `json:"fast"`
}

// Simple Buy / Sell
// -----------------------------------------------------------------------------

// SimpleQuoteRequest contains the request parameters for generating a simple
// buy or sell quote.
//
// POST /simple/{pair}/quote
type SimpleQuoteRequest struct {
	Amount        string `json:"payAmount"`
	Pair          string
	QuoteCurrency string `json:"payInCurrency"`
	Side          string `json:"side"`
}

// SimpleOrderRequest contains the request parameters for placing a simple buy
// or sell order.
//
// POST /simple/{pair}/order
type SimpleOrderRequest struct {
	Amount        string `json:"payAmount"`
	Pair          string
	QuoteCurrency string `json:"payInCurrency"`
	Side          string `json:"side"`
}

// SimpleOrderStatusRequest contains the request parameters for getting the
// current status of a simple buy or sell.
//
// GET /simple/{pair}/order/{id}
type SimpleOrderStatusRequest struct {
	OrderID string `json:"orderId"`
	Pair    string
}
