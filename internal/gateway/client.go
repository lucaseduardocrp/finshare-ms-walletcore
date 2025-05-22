package gateway

import "github.com.br/lucaseduardocrp/finshared-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Create(client *entity.Client) error
}
