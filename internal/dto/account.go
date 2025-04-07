package dto

import (
	"time"

	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
)

type CreateAccountDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromAccount(input *domain.Account) AccountOutputDTO {
	return AccountOutputDTO{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Balance:   input.Balance,
		APIKey:    input.APIKey,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ToAccount(input CreateAccountDTO) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}
