package valr

import (
	"time"
)

// CryptoWithdrawalResponse contains the response values returned from creating
// a new crypto withdrawal.
//
// POST /wallet/crypto/{currency}/withdrawal
type CryptoWithdrawalResponse struct {
	ID string `json:"id"`
}

// CryptoWithdrawalStatus contains the response values returned when creating
// a new crypto withdrawal.
//
// GET /wallet/crypto/{currency}/withdrawal/{id}
type CryptoWithdrawalStatusResponse struct {
	Address         string    `json:"address"`
	Amount          string    `json:"amount"`
	Confirmations   string    `json:"confirmations"`
	CreatedAt       time.Time `json:"createdAt"`
	Currency        string    `json:"currency"`
	Fees            string    `json:"feeAmount"`
	ID              string    `json:"id"`
	LastConfirmedAt time.Time `json:"lastConfirmedAt"`
	Status          string    `json:"status"`
	TransactionHash string    `json:"transactionHash"`
	Verified        bool      `json:"verified"`
}
