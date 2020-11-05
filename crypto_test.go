package valr_test

import (
	"context"
	"testing"

	"github.com/nickcorin/valr"
	"github.com/nickcorin/valr/mock"
	"github.com/stretchr/testify/suite"
)

type cryptoTestSuite struct {
	suite.Suite
	client valr.Client
	server *mock.Server
}

func TestCryptoTestSuite(t *testing.T) {
	suite.Run(t, new(cryptoTestSuite))
}

func (suite *cryptoTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = valr.NewClientForTesting(suite.T(), suite.server.URL)
}

func (suite *cryptoTestSuite) TestDepositAddress() {
	addr, err := suite.client.DepositAddress(context.TODO(), "ETH")
	suite.Require().NoError(err)
	suite.Require().NotNil(addr)

	eth := valr.DepositAddress{
		Address:  "0xA7Fae2Fd50886b962d46FF4280f595A3982aeAa5",
		Currency: "ETH",
	}

	suite.Require().EqualValues(&eth, addr)
}

func (suite *cryptoTestSuite) TestWithdrawalInfo() {
	info, err := suite.client.WithdrawalInfo(context.TODO(), "BTC")
	suite.Require().NoError(err)
	suite.Require().NotNil(info)

	btc := valr.WithdrawalInfo{
		Currency:            "BTC",
		IsActive:            true,
		MinWithdrawalAmount: "0.0002",
		SupportsPaymentRef:  false,
		WithdrawalCost:      "0.0004",
	}

	suite.Require().EqualValues(&btc, info)
}
