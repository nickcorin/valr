package mock

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Server struct {
	*httptest.Server
}

// NewServer returns a mock server to be used for unit testing. All enpoints
// return static response data read from JSON files in the testdata directory.
func NewServer() *Server {
	r := mux.NewRouter()
	registerRoutes(r)

	s := httptest.NewServer(r)
	return &Server{s}
}

func registerRoutes(r *mux.Router) {
	// Accounts.
	r.HandleFunc("/account/balances", makeHandler("accountBalances.json"))
	r.HandleFunc("/account/{pair}/tradehistory", makeHandler("tradehistory.json"))
	r.HandleFunc("/account/transactionhistory",
		makeHandler("transactionHistory.json"))

	// Public.
	r.HandleFunc("/public/currencies", makeHandler("currencies.json"))
	r.HandleFunc("/public/pairs", makeHandler("currencyPairs.json"))
	r.HandleFunc("/public/{pair}/orderbook", makeHandler("orderBook.json"))
	r.HandleFunc("/public/marketsummary", makeHandler("marketSummaries.json"))
	r.HandleFunc("/public/{pair}/marketsummary",
		makeHandler("marketSummary.json"))
	r.HandleFunc("/public/ordertypes", makeHandler("orderTypes.json"))
	r.HandleFunc("/public/{pair}/ordertypes",
		makeHandler("orderTypesForCurrency.json"))
	r.HandleFunc("/public/status", makeHandler("status.json"))
	r.HandleFunc("/public/time", makeHandler("serverTime.json"))

	// Crypto
	r.HandleFunc("/wallet/crypto/{currency}/deposit/address",
		makeHandler("depositaddress.json"))
	r.HandleFunc("/wallet/crypto/{currency}/withdraw",
		makeHandler("withdrawinfo.json"))
}

const testDir = "mock/testdata"

func makeHandler(responseFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := readResponseFile(filepath.Join(testDir, responseFile))
		if err != nil {
			serverError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func readResponseFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func serverError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
