package valr_test

import (
	"context"
	"testing"
	"time"

	"github.com/nickcorin/valr"
	"github.com/nickcorin/valr/mock"
	"github.com/stretchr/testify/suite"
)

func TestPublicTestSuite(t *testing.T) {
	suite.Run(t, new(PublicTestSuite))
}

type PublicTestSuite struct {
	suite.Suite
	client valr.PublicClient
	server *mock.Server
}


func (suite *PublicTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = valr.NewClientForTesting(suite.T(), suite.server.URL)
}

func (suite *PublicTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *PublicTestSuite) TestPublicClient_Currencies() {
	currencies, err := suite.client.Currencies(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(currencies)
	suite.Require().Len(currencies, 57)

	zar := valr.Currency{
		Symbol:    "R",
		IsActive:  true,
		ShortName: "ZAR",
		LongName:  "Rand",
	}

	suite.Require().EqualValues(zar, currencies[0])
}

func (suite *PublicTestSuite) TestPublicClient_CurrencyPairs() {
	pairs, err := suite.client.CurrencyPairs(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(pairs)
	suite.Require().Len(pairs, 91)

	btczar := valr.CurrencyPair{
		Active:         true,
		BaseCurrency:   "BTC",
		MaxBaseAmount:  "2",
		MaxQuoteAmount: "100000",
		MinBaseAmount:  "0.0001",
		MinQuoteAmount: "10",
		ShortName:      "BTC/ZAR",
		Symbol:         "BTCZAR",
		QuoteCurrency:  "ZAR",
	}

	suite.Require().EqualValues(btczar, pairs[0])
}

func (suite *PublicTestSuite) TestPublicClient_MarketSummary() {
	summary, err := suite.client.MarketSummary(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(summary)

	btczar := valr.MarketSummary{
		AskPrice: "10000",
		BaseVolume: "0.16065663",
		BidPrice: "7005",
		ChangeFromPrevious: "0",
		CurrencyPair: "BTCZAR",
		CreatedAt: time.Date(2019, 4, 20, 13, 02, 03, 228000000, time.UTC),
		HighPrice: "10000",
		LastTradedPrice: "7005",
		LowPrice: "7005",
		PreviousClosePrice: "7005",
	}

	suite.Require().EqualValues(btczar, summary[0])
}

func (suite *PublicTestSuite) TestPublicClient_MarketSummaryForCurrency() {
	summary, err := suite.client.MarketSummaryForCurrency(context.TODO(),
		"BTCZAR")
	suite.Require().NoError(err)
	suite.Require().NotNil(summary)

	btczar := &valr.MarketSummary{
		AskPrice: "10000",
		BaseVolume: "0.16065663",
		BidPrice: "7005",
		ChangeFromPrevious: "0",
		CurrencyPair: "BTCZAR",
		CreatedAt: time.Date(2019, 4, 20, 13, 03, 03, 230000000, time.UTC),
		HighPrice: "10000",
		LastTradedPrice: "7005",
		LowPrice: "7005",
		PreviousClosePrice: "7005",
	}

	suite.Require().EqualValues(btczar, summary)
}

func (suite *PublicTestSuite) TestPublicClient_OrderBook() {
	book, err := suite.client.OrderBook(context.TODO(), "BTCZAR")
	suite.Require().NoError(err)
	suite.Require().NotNil(book)
	suite.Require().Len(book.Asks, 9)
	suite.Require().Len(book.Bids, 11)

	ask := valr.OrderBookEntry{
		CurrencyPair: "BTCZAR",
		OrderCount:   1,
		Price:        "9000",
		Quantity:     "0.101",
		Side:         "sell",
	}

	bid := valr.OrderBookEntry{
		CurrencyPair: "BTCZAR",
		OrderCount:   1,
		Price:        "8802",
		Quantity:     "0.1",
		Side:         "buy",
	}

	suite.Require().Equal(ask, book.Asks[0])
	suite.Require().Equal(bid, book.Bids[0])

}

func (suite *PublicTestSuite) TestPublicClient_OrderTypes() {
	types, err := suite.client.OrderTypes(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(types)
	suite.Require().Len(types, 91)

	btczar := map[valr.OrderType]bool{
		valr.OrderTypePostOnly: true,
		valr.OrderTypeLimit:    true,
		valr.OrderTypeMarket:   true,
	}

	for k, v := range btczar {
		suite.Require().Equal(v, types["BTCZAR"][k])
	}
}

func (suite *PublicTestSuite) TestPublicClient_ServerTime() {
	res, err := suite.client.ServerTime(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	serverTime := valr.ServerTime{
		Epoch: 1555513811,
		Time:  time.Date(2019, 04, 17, 15, 10, 11, 956000000, time.UTC),
	}

	suite.Require().EqualValues(&serverTime, res)
}

func (suite *PublicTestSuite) TestPublicClient_Status() {
	status, err := suite.client.Status(context.TODO())
	suite.Require().NoError(err)
	suite.Require().Equal(valr.StatusOnline, status)
}
