package gateway

import "github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"

type AccountGateway interface {
	FindById(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
