package dto

import (
	"time"

	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
)

const (
	Pending  = string(domain.Pending)
	Approved = string(domain.Approved)
	Rejected = string(domain.Rejected)
)

type CreateInvoiceDTO struct {
	APIKey         string
	AccountID      string  `json:"account_id"`
	Amount         float64 `json:"amount"`
	Description    string  `json:"description" `
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardHolderName string  `json:"card_holder_name"`
}

type InvoiceOutputDTO struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	Status         string    `json:"status"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceDTO, accountID string) (*domain.Invoice, error) {
	card := &domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
		CardHolderName: input.CardHolderName,
	}

	return domain.NewInvoice(
		accountID,
		input.Description,
		input.PaymentType,
		input.Amount,
		card,
	)
}

func FromInvoice(domain *domain.Invoice) *InvoiceOutputDTO {
	return &InvoiceOutputDTO{
		ID:             domain.ID,
		AccountID:      domain.AccountID,
		Amount:         domain.Amount,
		Status:         string(domain.Status),
		Description:    domain.Description,
		PaymentType:    domain.PaymentType,
		CardLastDigits: domain.CardLastDigits,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
