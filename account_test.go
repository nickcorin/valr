package valr_test

import (
	"context"
	"testing"
	"time"

	"github.com/nickcorin/valr"
	"github.com/nickcorin/valr/mock"
	"github.com/stretchr/testify/suite"
)

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}

type accountTestSuite struct {
	suite.Suite
	client valr.Client
	server *mock.Server
}

func (suite *accountTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = valr.NewClientForTesting(suite.T(), suite.server.URL)
}

func (suite *accountTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *accountTestSuite) TestPrivateClient_Balances() {
	balances, err := suite.client.Balances(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(balances)

	eth := valr.Balance{
		Available: "0.01626594758",
		Currency:  "ETH",
		Reserved:  "0.49",
		Total:     "0.50626594758",
		UpdatedAt: time.Date(2020, 5, 31, 5, 10, 16, 522000000, time.UTC),
	}

	suite.Require().EqualValues(eth, balances[0])
}

func (suite *accountTestSuite) TestPrivateClient_TradeHistory() {
	history, err := suite.client.TradeHistory(context.TODO(),
		"BTCZAR")
	suite.Require().NoError(err)
	suite.Require().NotNil(history)

	trades := []valr.Trade{
		{
			CurrencyPair: "BTCZAR",
			ID:           10634,
			Price:        "87000",
			Quantity:     "0.0001",
			Side:         "buy",
			TradedAt: time.Date(2019, 5, 13, 15, 14, 48, 422000000,
				time.UTC),
		},
	}

	suite.Require().Contains(history, trades[0])
}

func (suite *accountTestSuite) TestPrivateClient_TransactionHistory() {
	history, err := suite.client.TransactionHistory(context.TODO(),
		&valr.TransactionHistoryRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(history)

	transactions := []valr.Transaction{
		{
			CreditCurrency: "BTC",
			CreditValue:    "0.0000003",
			EventAt: time.Date(2019, 5, 7, 10, 55, 9, 949000000,
				time.UTC),
			TypeInfo: &valr.TransactionTypeInfo{
				Type:        valr.TransactionTypeReferralRebate,
				Description: "Referral Rebate",
			},
		},
	}

	suite.Require().Contains(history, transactions[0])
}
