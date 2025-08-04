package domain

type PaymentSummaryOutput struct {
	TotalAmount   float32 `json:"totalAmount"`
	TotalRequests uint    `json:"totalRequests"`
}

type PaymentSummaryResponseOutput struct {
	DefaultOutput  PaymentSummaryOutput `json:"default"`
	FallbackOutput PaymentSummaryOutput `json:"fallback"`
}
