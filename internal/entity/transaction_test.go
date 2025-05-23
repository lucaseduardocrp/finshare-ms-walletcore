package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Client Name", "client@mail.com")
	account1 := NewAccount(client1)

	client2, _ := NewClient("Client Name 2", "client2@mail.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 500)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 500.0, account1.Balance)
	assert.Equal(t, 1500.0, account2.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := NewClient("Client Name", "client@mail.com")
	account1 := NewAccount(client1)

	client2, _ := NewClient("Client Name 2", "client2@mail.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 2000)

	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Error(t, err, "insufficient founds")
}
