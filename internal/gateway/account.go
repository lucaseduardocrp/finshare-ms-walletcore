package gateway

import "github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Create(account *entity.Account) error
}
