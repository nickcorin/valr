package valr

import (
	"context"
	"fmt"
)

// DepositAddress contains the default deposit address for a crypto wallet.
type DepositAddress struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

// DepositAddress returns the default deposit address associated with a
// currency.
func (c *client) DepositAddress(ctx context.Context, currency string) (
	*DepositAddress, error) {
	res, err := c.httpClient.Get(ctx,
		fmt.Sprintf("/wallet/crypto/%s/deposit/address", currency), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get default deposit address: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch default deposit address: %d "+
			"status code received", res.StatusCode)
	}

	var addr DepositAddress
	if err = res.JSON(&addr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit address: %w", err)
	}

	return &addr, nil
}
