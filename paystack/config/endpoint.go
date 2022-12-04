package config

const (
	BaseUrl                       = "https://api.paystack.co"
	InitializeTransactionEndpoint = "/transaction/initialize"
	VerifyTransactionEndpoint     = "/transaction/verify"
	TransactionEndpoint           = "/transaction"
	ChargeAuthorizationEndpoint   = "/transaction/charge_authorization"
	CheckAuthorizationEndpoint    = "/transaction/check_authorization"
	TransactionTimelineEndpoint   = "/transaction/timeline"
	TransactionTotalEndpoint      = "/transaction/totals"
	ExportTransactionEndpoint     = "/transaction/export"
	PartialDebtEndpoint           = "/transaction/partial_debit"
)
