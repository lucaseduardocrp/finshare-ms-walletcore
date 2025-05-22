package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Client Name", "client@mail.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Client Name", client.Name)
	assert.Equal(t, "client@mail.com", client.Email)
}

func TestCreateNewClientWithArgsInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
	assert.Error(t, err, "email is required")
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Client Name", "client@mail.com")
	err := client.Update("Client Name Updated", "clientupdated@mail.com")

	assert.Nil(t, err)
	assert.Equal(t, "Client Name Updated", client.Name)
	assert.Equal(t, "clientupdated@mail.com", client.Email)
}

func TestUpdateClientWithArgsInvalid(t *testing.T) {
	client, _ := NewClient("Client Name", "client@mail.com")
	err := client.Update("", "")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
	assert.Error(t, err, "email is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Client Name", "client@mail.com")
	account := NewAccount(client)
	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
