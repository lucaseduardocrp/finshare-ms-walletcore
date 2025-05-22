package gateway

import "github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"

type AccountGateway interface {
	GetById(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
