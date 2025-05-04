package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(id int) (*Account, error)
}

type PostresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank port=5434 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostresStore{
		db: db,
	}, nil
}

func (s *PostresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		number SERIAL,
		balance SERIAL,
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostresStore) CreateAccount(acc *Account) error {
	query := `
		INSERT INTO account 
		(first_name, last_name, number, balance, created_at)
		VALUES ($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account(nil)
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}
