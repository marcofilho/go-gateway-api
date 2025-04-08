package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
	CVV            string
}

func NewInvoice(accountID, description, paymentType string, amount float64, creditCard *CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrorInvalidAmount
	}

	if accountID == "" {
		return nil, ErrorAccountIdNotInformed

	}

	cardLastDigits := creditCard.Number[len(creditCard.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         Pending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: cardLastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000.00 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() < 0.7 {
		newStatus = Approved
	} else {
		newStatus = Rejected
	}

	i.Status = newStatus
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != Pending {
		return ErrorInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
