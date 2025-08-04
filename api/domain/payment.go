package domain

import "github.com/google/uuid"

type PaymentEntryDTO struct {
	Amount        float32   `binding:required; json:"amount"`
	CorrelationID uuid.UUID `binding:required; json:"correlationId"`
}
