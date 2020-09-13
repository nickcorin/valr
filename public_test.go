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

func (suite *PublicTestSuite) TestPublicClient_OrderBook() {
	book, err := suite.client.OrderBook(context.TODO(), "BTCZAR")
	suite.Require().NoError(err)
	suite.Require().NotNil(book)
	suite.Require().Equal(9, len(book.Asks))
	suite.Require().Equal(11, len(book.Bids))
}

func (suite *PublicTestSuite) TestPublicCLient_ServerTime() {
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
