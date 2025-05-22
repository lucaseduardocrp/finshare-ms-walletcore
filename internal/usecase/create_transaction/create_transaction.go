package createtransaction

import (
	"github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"
	"github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	Id string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.Get(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.Get(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	response := &CreateTransactionOutputDTO{
		Id: transaction.Id,
	}

	return response, nil
}
