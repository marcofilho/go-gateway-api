package service

import (
	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
	"github.com/marcofilho/go-api-payment-gateway/internal/dto"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountDTO) (*dto.AccountOutputDTO, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.accountRepository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrorAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrorDuplicatedAPIKey
	}

	err = s.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindAccountByID(id string) (*dto.AccountOutputDTO, error) {
	account, err := s.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindAccountByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutputDTO, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	account.AddBalance(amount)
	err = s.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindAllAccounts() ([]dto.AccountOutputDTO, error) {
	accounts, err := s.accountRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []dto.AccountOutputDTO

	for _, account := range accounts {
		output = append(output, dto.FromAccount(account))
	}

	return output, nil
}
