package valr

import (
	"context"
	"fmt"
	"time"
)

// Currency is a fiat or crypto currency supported by VALR.
type Currency struct {
	IsActive  bool   `json:"isActive"`
	LongName  string `json:"longName"`
	ShortName string `json:"shortName"`
	Symbol    string `json:"symbol"`
}

// Currencies satisfies the PublicClient interface.
func (c *client) Currencies(ctx context.Context) ([]Currency, error) {
	res, err := c.httpClient.Get(ctx, "/public/currencies", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currencies: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch currencies: %d status code "+
			"received", res.StatusCode)
	}

	var currencies []Currency
	if err = res.JSON(&currencies); err != nil {
		return nil, fmt.Errorf("failed to unmarshal currencies: %w", err)
	}

	return currencies, nil
}

// CurrencyPair is a fiat/crypto or crypto/crypto pair supported by VALR.
type CurrencyPair struct {
	Active         bool   `json:"active"`
	BaseCurrency   string `json:"baseCurrency"`
	MinBaseAmount  string `json:"minBaseAmount"`
	MaxBaseAmount  string `json:"maxBaseAmount"`
	MinQuoteAmount string `json:"minQuoteAmount"`
	MaxQuoteAmount string `json:"maxQuoteAmount"`
	QuoteCurrency  string `json:"quoteCurrency"`
	ShortName      string `json:"shortName"`
	Symbol         string `json:"symbol"`
}

// CurrencyPairs satisfies the PublicClient interface.
func (c *client) CurrencyPairs(ctx context.Context) ([]CurrencyPair, error) {
	res, err := c.httpClient.Get(ctx, "/public/pairs", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currency pairs: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch currency pairs: %d status "+
			"code received", res.StatusCode)
	}

	var pairs []CurrencyPair
	if err = res.JSON(&pairs); err != nil {
		return nil, fmt.Errorf("failed to fetch currency pairs: %w", err)
	}

	return pairs, nil
}

// MarketSummary contains a summary of a particular market pair on the
// exchange.
type MarketSummary struct {
	AskPrice           string    `json:"askPrice"`
	BaseVolume         string    `json:"baseVolume"`
	BidPrice           string    `json:"bidPrice"`
	ChangeFromPrevious string    `json:"changeFromPrevious"`
	CreatedAt          time.Time `json:"created"`
	CurrencyPair       string    `json:"currencyPair"`
	HighPrice          string    `json:"highPrice"`
	LastTradedPrice    string    `json:"lastTradedPrice"`
	LowPrice           string    `json:"lowPrice"`
	PreviousClosePrice string    `json:"previousClosePrice"`
}

// MarketSummary satisfies the PublicClient interface.
func (c *client) MarketSummary(ctx context.Context) ([]MarketSummary, error) {
	res, err := c.httpClient.Get(ctx, "/public/marketsummary", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the market summaries: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch market summary: %d status "+
			"code received", res.StatusCode)
	}

	var summaries []MarketSummary
	if err = res.JSON(&summaries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal market summaries: %w", err)
	}

	return summaries, nil
}

// MarketSummaryForCurrency satisfies the PublicClient interface.
func (c *client) MarketSummaryForCurrency(ctx context.Context, pair string) (
	*MarketSummary, error) {
	res, err := c.httpClient.Get(ctx, fmt.Sprintf("/public/%s/marketsummary",
		pair), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the market summaries: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch market summaries: %d status "+
			"code received", res.StatusCode)
	}

	var summary MarketSummary
	if err = res.JSON(&summary); err != nil {
		return nil, fmt.Errorf("failed to unmarshal market summaries: %w", err)
	}

	return &summary, nil
}

// OrderBook contains a list of the top 20 bids and asks. Ask orders are sorted
// by priceascending. Bid orders are sorted by price descending. Orders of the
// same price are aggregated.
type OrderBook struct {
	Asks []OrderBookEntry `json:"Asks"`
	Bids []OrderBookEntry `json:"Bids"`
}

// OrderBookEntry is a single entry in an order book.
type OrderBookEntry struct {
	CurrencyPair string `json:"currencyPair"`
	OrderCount   int    `json:"orderCount"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
	Side         string `json:"side"`
}

// OrderBook satisfies the PublicClient interface.
func (c *client) OrderBook(ctx context.Context, pair string) (*OrderBook,
	error) {
	res, err := c.httpClient.Get(ctx, fmt.Sprintf("/public/%s/orderbook", pair),
		nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order book: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch order book: %d status code "+
			"received", res.StatusCode)
	}

	var book OrderBook
	if err = res.JSON(&book); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order book: %w", err)
	}

	return &book, nil
}

// OrderType describes the kind of an order.
type OrderType string

const (
	// Place a limit order on the exchange that will either be added to the
	// order book or, should it match, be cancelled completely.
	OrderTypePostOnly OrderType = "post-only limit"

	// Place a limit order on the exchange.
	OrderTypeLimit OrderType = "limit"

	// Place a market order on the exchange (only crypto-to-ZAR pairs).
	OrderTypeMarket OrderType = "market"

	// Similar to market order, but allows for crypto-to-crypto pairs.
	OrderTypeSimple OrderType = "simple"
)

// OrderTypes satisfies the PublicClient interface.
func (c *client) OrderTypes(ctx context.Context) (map[string]map[OrderType]bool,
	error) {
	res, err := c.httpClient.Get(ctx, "/public/ordertypes", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order types: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch order types: %d status code "+
			"received", res.StatusCode)
	}

	types := []struct {
		Pair       string      `json:"currencyPair"`
		OrderTypes []OrderType `json:"orderTypes"`
	}{}

	if err = res.JSON(&types); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order types: %w", err)
	}

	typesMap := make(map[string]map[OrderType]bool)
	for _, pair := range types {
		typesMap[pair.Pair] = make(map[OrderType]bool)
		for _, typ := range pair.OrderTypes {
			typesMap[pair.Pair][typ] = true
		}
	}

	return typesMap, nil
}

// OrderTypesForCurrency satisfies the PublicClient interface.
func (c *client) OrderTypesForCurrency(ctx context.Context, pair string) (
	map[OrderType]bool, error) {
	res, err := c.httpClient.Get(ctx, fmt.Sprintf("/public/%s/ordertypes",
		pair), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order types for pair: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch order types: %d status code "+
			"received", res.StatusCode)
	}

	var orderTypes []OrderType
	if err = res.JSON(&orderTypes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order types: %w", err)
	}

	typesMap := make(map[OrderType]bool)
	for _, typ := range orderTypes {
		typesMap[typ] = true
	}

	return typesMap, nil
}

// Server time contains the time on VALRs servers.
type ServerTime struct {
	Epoch int64     `json:"epochTime"`
	Time  time.Time `json:"time"`
}

// ServerTime satisfies the PublicClient interface.
func (c *client) ServerTime(ctx context.Context) (*ServerTime, error) {
	res, err := c.httpClient.Get(ctx, "/public/time", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch server time: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch server time: %d status code "+
			"received", res.StatusCode)
	}

	var serverTime ServerTime
	if err = res.JSON(&serverTime); err != nil {
		return nil, fmt.Errorf("failed to unmarshal server time: %w", err)
	}

	return &serverTime, nil
}

// Status describes the current status of VALR.
type Status string

const (
	// StatusUnknown implies that the call to VALR failed and we are unable to
	// determine VALRs current status.
	StatusUnknown Status = "unknown"

	// StatusOnline implies that all functionality is available.
	StatusOnline Status = "online"

	// StatusReadOnly implies that only GET and OPTIONS requests are accepted.
	// All other requests in read-only mode will respond with a 503 status code.
	StatusReadOnly Status = "read-only"
)

// Status satisfies the PublicClient interface.
func (c *client) Status(ctx context.Context) (Status, error) {
	res, err := c.httpClient.Get(ctx, "/public/status", nil)
	if err != nil {
		return StatusUnknown, fmt.Errorf("failed to fetch status: %w", err)
	}

	if !res.IsSuccess() {
		return StatusUnknown, fmt.Errorf("failed to fetch status: %d status "+
			"code received", res.StatusCode)
	}

	var status Status
	if err = res.JSON(&status); err != nil {
		return StatusUnknown, fmt.Errorf("failed to unmarshal status: %w", err)
	}

	return status, nil
}
