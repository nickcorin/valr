package valr_test

import (
	"context"
	"testing"
	"time"

	"github.com/nickcorin/valr"
	"github.com/nickcorin/valr/mock"
	"github.com/stretchr/testify/suite"
)

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

func (suite *PublicTestSuite) TestPublicClient_OrderBook() {
	book, err := suite.client.OrderBook(context.TODO(), "BTCZAR")
	suite.Require().NoError(err)
	suite.Require().NotNil(book)
	suite.Require().Len(book.Asks, 9)
	suite.Require().Len(book.Bids, 11)
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

func TestPublicTestSuite(t *testing.T) {
	suite.Run(t, new(PublicTestSuite))
}
