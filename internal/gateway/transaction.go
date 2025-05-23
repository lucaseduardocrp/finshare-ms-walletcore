package gateway

import "github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
