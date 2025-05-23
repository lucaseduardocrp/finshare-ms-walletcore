package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, err := NewClient("Client Name", "client@mail.com")
	account := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.Id, account.Client.Id)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Client Name", "client@mail.com")
	account := NewAccount(client)
	account.Credit(100)

	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Client Name", "client@mail.com")
	account := NewAccount(client)

	account.Credit(100)
	account.Debit(50)

	assert.Equal(t, float64(50), account.Balance)
}
