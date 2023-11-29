package repositories

import (
	"authing/internal/db/models"

	"github.com/jmoiron/sqlx"
)

type Accounts struct {
	db *sqlx.DB
}

func NewAccountsRepository(db *sqlx.DB) *Accounts {
	return &Accounts{
		db: db,
	}
}

func (r *Accounts) Add(username string, passwordHash string) (*models.Account, error) {
	query := `INSERT INTO accounts (username, password) VALUES (?, ?)`
	result, err := r.db.Exec(query, username, passwordHash)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetByID(int(id))
}

func (r *Accounts) GetByID(id int) (*models.Account, error) {
	var account models.Account
	query := `SELECT id, username, password FROM accounts WHERE id = ?`
	err := r.db.Get(&account, query, id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Accounts) Get(username string) (*models.Account, error) {
	var account models.Account
	query := `SELECT id, username, password FROM accounts WHERE username = ?`
	err := r.db.Get(&account, query, username)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Accounts) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	query := `SELECT id, username, password FROM accounts`
	err := r.db.Select(&accounts, query)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *Accounts) Delete(id int) error {
	query := `DELETE FROM accounts WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
