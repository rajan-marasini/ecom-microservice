package order

import (
	"github.com/rajan-marasini/ecom-microservice/account"
	"github.com/rajan-marasini/ecom-microservice/catalog"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		return err
	}

	return err

}
