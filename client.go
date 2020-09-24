package valr

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nickcorin/snorlax"
)

// PrivateClient contains methods that require authentication in order to access
// and have more relaxed rate limiting rules.
type PrivateClient interface {
}

// PublicClient contains methods that do not require authentication in order to
// access and have stricter rate limiting rules.
type PublicClient interface {
	// Currencies returns a list of currencies supported by VALR.
	Currencies(ctx context.Context) ([]Currency, error)

	// CurrencyPairs returns a list of all the currency pairs supported by VALR.
	CurrencyPairs(ctx context.Context) ([]CurrencyPair, error)

	// MarketSummary reutns a market summary for all supported currency pairs.
	MarketSummary(ctx context.Context) ([]MarketSummary, error)

	// MarketSummaryForCurrency returns the market summary for a given currency
	// pair.
	MarketSummaryForCurrency(ctx context.Context, pair string) (*MarketSummary,
		error)

	// OrderBook returns a list of the top 20 bids and asks in the order book.
	// Ask orders are sorted by price ascending. Bid orders are sorted by price
	// descending. Orders of the same price are aggregated.
	OrderBook(ctx context.Context, pair string) (*OrderBook, error)

	// OrderTypes returns a map of all the order types supported for all
	// currency pairs. A 2D map is returned with the first key being the
	// currency pair and the second key being the order type. The value of the
	// map is a boolean denoting whether the order type is supported. You can
	// only place an order that is supported by that currency pair.
	//
	// Example:
	//
	// if orderTypes["BTCZAR"][OrderTypeSimple] {
	//     /* The simple order type is supported for BTCZAR. */
	// } else {
	//     /* The simple order type is not supported for BTCZAR. */
	// }
	OrderTypes(ctx context.Context) (map[string]map[OrderType]bool, error)

	// OrderTypesForCurrency returns a map of the order types supported for a
	// given currency pair. A map is returned with  key being the currency pair
	// and the value being a boolean denoting whether the order type is
	// supported. You can only place an order that is supported by that currency
	// pair.
	//
	// Example:
	//
	// if orderTypes[OrderTypeSimple] {
	//     /* The simple order type is supported. */
	// } else {
	//     /* The simple order type is not supported. */
	// }
	OrderTypesForCurrency(ctx context.Context, pair string) (map[OrderType]bool,
		error)

	// ServerTime returns the time on the server.
	ServerTime(ctx context.Context) (*ServerTime, error)

	// Status returns the current status of VALR.
	Status(ctx context.Context) (Status, error)
}

// Client is an HTTP client wrapper for the VALR REST API. It is a combination
// of the PublicClient and PrivateClient.
type Client interface {
	PrivateClient
	PublicClient
}

var defaultBaseURL = "https://api.valr.com/api/v1"

// DefaultClient is a VALR client initialized with default values. This should
// be sufficient for callers only using the
var DefaultClient PublicClient = &client{
	apiKey:    "",
	apiSecret: "",
	httpClient: snorlax.DefaultClient.
		SetBaseURL(defaultBaseURL).
		SetHeader(http.CanonicalHeaderKey("Content-Type"), "application/json"),
}

// NewClient returns a Client.
func NewClient(key, secret string) Client {
	c := *DefaultClient.(*client)
	c.apiKey = key
	c.apiSecret = secret

	// Sign requests to private endpoints.
	c.httpClient.AddRequestHook(authenticationHook(key, secret))

	return &c
}

// NewClientForTesting returns a Client with a custom base URL for testing
// purposes.
func NewClientForTesting(_ *testing.T, baseURL string) Client {
	c := NewClient("", "").(*client)
	c.httpClient.SetBaseURL(baseURL)

	return c
}

// NewPublicClient returns a PublicClient.
func NewPublicClient() PublicClient {
	return NewClient("", "")
}

type client struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient snorlax.Client
}

func authenticationHook(key, secret string) snorlax.RequestHook {
	return func(c snorlax.Client, r *http.Request) error {
		// If we are requesting a public endpoint, we do not need to sign the
		// request.
		if strings.Contains(r.URL.Path, "public") {
			return nil
		}

		bodyReader, err := r.GetBody()
		if err != nil {
			return fmt.Errorf("failed to get request body: %w", err)
		}

		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			return fmt.Errorf("failed to read request body: %w", err)
		}

		timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		signature := generateAuthSignature(secret, timestamp, r.Method,
			r.URL.Path, body)

		r.Header.Set("X-VALR-API-KEY", key)
		r.Header.Set("X-VALR-SIGNATURE", signature)
		r.Header.Set("X-VALR-TIMESTAMP", timestamp)

		return nil
	}
}

func generateAuthSignature(secret string, timestamp, method, path string,
	body []byte) string {

	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte(strings.ToUpper(method)))
	mac.Write([]byte(path))
	mac.Write(body)

	return hex.EncodeToString(mac.Sum(nil))
}

// ToPrivateClient converts a PublicClient to a PrivateClient.
func ToPrivateClient(c PublicClient, key, secret string) PrivateClient {
	client := c.(*client)
	client.apiKey = key
	client.apiSecret = secret

	return client
}

// ToPublicClient converts a PrivateClient to a PublicClient.
func ToPublicClient(c PrivateClient) PublicClient {
	client := c.(*client)
	client.apiKey = ""
	client.apiSecret = ""

	return client
}
