package valr

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type clientTestSuite struct {
	suite.Suite
	client Client
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func (suite *clientTestSuite) SetupSuite() {
	suite.client = NewClient("", "")
}

func (suite *clientTestSuite) TestNewClient() {
	key, secret := "key", "s3cret"
	c := NewClient(key, secret)
	suite.Require().NotNil(c)

	c, ok := c.(*client)
	suite.Require().True(ok)
	suite.Require().Equal(key, c.(*client).apiKey)
	suite.Require().Equal(secret, c.(*client).apiSecret)
}

func (suite *clientTestSuite) TestNewPublicClient() {
	c := NewPublicClient()
	suite.Require().NotNil(c)

	c, ok := c.(*client)
	suite.Require().True(ok)
	suite.Require().Equal("", c.(*client).apiKey)
	suite.Require().Equal("", c.(*client).apiSecret)
}

func (suite *clientTestSuite) TestAuthenticationHook() {
	testcases := []struct {
		key  string
		path string
		body []byte
		err  bool
	}{
		{
			key:  "myKey",
			path: "/v1/account/balances",
			body: nil,
			err:  false,
		},
	}

	for _, test := range testcases {
		suite.T().Run("", func(t *testing.T) {
			hook := authenticationHook(test.key, "")
			suite.Require().NotNil(hook)

			r, err := http.NewRequest(http.MethodGet, test.path,
				bytes.NewBuffer(test.body))
			suite.Require().NoError(err)
			suite.Require().NotNil(r)

			err = hook(suite.client.(*client).httpClient, r)
			if test.err {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotEmpty(r.Header.Get("X-VALR-API-KEY"))
				suite.Require().NotEmpty(r.Header.Get("X-VALR-SIGNATURE"))
				suite.Require().NotEmpty(r.Header.Get("X-VALR-TIMESTAMP"))
			}

		})
	}
}

func (suite *clientTestSuite) TestGenerateAuthSignature() {
	testcases := []struct {
		path      string
		method    string
		body      []byte
		timestamp string
		secret    string
		signature string
	}{
		{
			path:      "/v1/account/balances",
			method:    http.MethodGet,
			body:      nil,
			timestamp: "1558014486185",
			secret: "4961b74efac86b25cce8fbe4c9811c4c7a787b7a5996660afcc2e287" +
				"ad864363",
			signature: "9d52c181ed69460b49307b7891f04658e938b21181173844b5018" +
				"b2fe783a6d4c62b8e67a03de4d099e7437ebfabe12c56233b73c6a0cc0f7" +
				"ae87e05f6289928",
		},
		{
			path:   "/v1/orders/market",
			method: http.MethodPost,
			body: []byte("{\"customerOrderId\":\"ORDER-000001\",\"pair\":\"BT" +
				"CZAR\",\"side\":\"BUY\",\"quoteAmount\":\"80000\"}"),
			timestamp: "1558017528946",
			secret: "4961b74efac86b25cce8fbe4c9811c4c7a787b7a5996660afcc2e287" +
				"ad864363",
			signature: "be97d4cd9077a9eea7c4e199ddcfd87408cb638f2ec2f7f74dd44" +
				"aef70a49fdc49960fd5de9b8b2845dc4a38b4fc7e56ef08f042a3c78a3af" +
				"9aed23ca80822e8",
		},
	}

	for _, test := range testcases {
		suite.T().Run("", func(t *testing.T) {
			signature := generateAuthSignature(test.secret, test.timestamp,
				test.method, test.path, test.body)
			suite.Require().Equal(test.signature, signature)
		})
	}
}
