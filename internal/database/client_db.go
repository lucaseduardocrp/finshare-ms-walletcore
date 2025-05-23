package database

import (
	"database/sql"

	"github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&client.Id, &client.Name, &client.Email, &client.CreatedAt)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientDB) Create(client *entity.Client) error {
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(client.Id, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
