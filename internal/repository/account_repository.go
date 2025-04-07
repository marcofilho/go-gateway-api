package repository

import (
	"database/sql"
	"time"

	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.DB.Prepare(`
	INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.DB.QueryRow(`
	SELECT id, name, email, api_key, balance, created_at, updated_at 
	FROM accounts 
	WHERE id = $1`,
		id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorAccountNotFound
		}
		return nil, err
	}
	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

func (r *AccountRepository) FindAll() ([]*domain.Account, error) {
	rows, err := r.DB.Query(`
	SELECT id, name, email, api_key, balance, created_at, updated_at 
	FROM accounts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*domain.Account

	for rows.Next() {
		var account domain.Account

		if err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Email,
			&account.APIKey,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.DB.QueryRow(`
	SELECT id, name, email, api_key, balance, created_at, updated_at 
	FROM accounts 
	WHERE api_key = $1`,
		apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorAccountNotFound
		}
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64

	err = tx.QueryRow(`
	SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, account.ID).
		Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return domain.ErrorAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE accounts
	SET balance = $1, updated_at = $2
	WHERE id = $3`,
		account.Balance,
		time.Now(),
		account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
