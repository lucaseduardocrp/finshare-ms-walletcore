package createclient

import (
	"testing"

	"github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "Client Name",
		Email: "client@mail.com",
	})

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.Equal(t, "Client Name", output.Name)
	assert.Equal(t, "client@mail.com", output.Email)
}
