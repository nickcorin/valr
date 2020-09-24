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
	r.HandleFunc("/public/currencies", makeHandler("currencies.json"))
	r.HandleFunc("/public/pairs", makeHandler("currencyPairs.json"))
	r.HandleFunc("/public/{currencyPair}/orderbook",
		makeHandler("orderBook.json"))
	r.HandleFunc("/public/time", makeHandler("serverTime.json"))
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
