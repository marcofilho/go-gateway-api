package domain

import "errors"

var (
	ErrorAccountNotFound  = errors.New("account not found")
	ErrorDuplicatedAPIKey = errors.New("api key already exists")
	ErrorInvoiceNotFound  = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
