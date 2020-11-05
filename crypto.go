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

// DepositAddress satisfies the PrivateClient interface.
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

// WithdrawalInfo contains information about withdrawing from your VALR account.
type WithdrawalInfo struct {
	Currency            string `json:"currency"`
	IsActive            bool   `json:"isActive"`
	MinWithdrawalAmount string `json:"minimumWithdrawAmount"`
	SupportsPaymentRef  bool   `json:"supportsPaymentReference"`
	WithdrawalCost      string `json:"withdrawCost"`
}

// WithdrawalInfo satisfies the PrivateClient interface.
func (c *client) WithdrawalInfo(ctx context.Context, currency string) (
	*WithdrawalInfo, error) {

	res, err := c.httpClient.Get(ctx,
		fmt.Sprintf("/wallet/crypto/:%s/withdraw", currency), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get withdrawal info: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch withdrawal info: %d status "+
			"code received", res.StatusCode)
	}

	var info WithdrawalInfo
	if err = res.JSON(&info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit address: %w", err)
	}

	return &info, nil
}
