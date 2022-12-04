package dtos

import "time"

type InitializeTransactionDto struct {
	AuthorizationURL string `json:"authorization_url,omitempty"`
	AccessCode       string `json:"access_code,omitempty"`
	Reference        string `json:"reference,omitempty"`
}

type VerifyTransactionDto struct {
	ID              int       `json:"id"`
	Domain          string    `json:"domain"`
	Status          string    `json:"status"`
	Reference       string    `json:"reference"`
	Amount          int       `json:"amount"`
	GatewayResponse string    `json:"gateway_response"`
	PaidAt          time.Time `json:"paid_at"`
	CreatedAt       time.Time `json:"created_at"`
	Channel         string    `json:"channel"`
	Currency        string    `json:"currency"`
	Customer        Customer  `json:"customer"`
	IPAddress       string    `json:"ip_address"`
	Metadata        string    `json:"metadata"`
	Fees            int       `json:"fees"`
	RequestedAmount int       `json:"requested_amount"`
	TransactionDate time.Time `json:"transaction_date"`
}

type Customer struct {
	ID                       int    `json:"id"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	Email                    string `json:"email"`
	CustomerCode             string `json:"customer_code"`
	Phone                    string `json:"phone"`
	Metadata                 string `json:"metadata"`
	RiskAction               string `json:"risk_action"`
	InternationalFormatPhone string `json:"international_format_phone"`
}
