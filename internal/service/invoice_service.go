package service

import (
	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
	"github.com/marcofilho/go-api-payment-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) CreateInvoice(input dto.CreateInvoiceDTO) (*dto.InvoiceOutputDTO, error) {
	accountOutput, err := s.accountService.FindAccountByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.Approved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil

}

func (s *InvoiceService) GetById(id, apiKey string) (*dto.InvoiceOutputDTO, error) {
	accountOutput, err := s.accountService.FindAccountByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrorUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) GetByAccountID(accountID string) ([]*dto.InvoiceOutputDTO, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	var output []*dto.InvoiceOutputDTO

	for _, invoice := range invoices {
		output = append(output, dto.FromInvoice(invoice))
	}

	return output, nil
}

func (s *InvoiceService) GetByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutputDTO, error) {
	accountOutput, err := s.accountService.FindAccountByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoices, err := s.invoiceRepository.FindByAccountID(accountOutput.ID)
	if err != nil {
		return nil, err
	}

	var output []*dto.InvoiceOutputDTO

	for _, invoice := range invoices {
		output = append(output, dto.FromInvoice(invoice))
	}

	return output, nil
}
