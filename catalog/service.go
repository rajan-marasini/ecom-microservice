package catalog

import (
	"context"
	"fmt"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostProduct(ctx context.Context, name, descriptin string, price float32) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	GetProductByID(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type catalogService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &catalogService{r}
}

func (r *catalogService) PostProduct(ctx context.Context, name, description string, price float32) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       fmt.Sprintf("%f", price),
		ID:          ksuid.New().String(),
	}
	if err := r.repo.PutProduct(ctx, *p); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *catalogService) GetProduct(ctx context.Context, id string) (*Product, error) {

	return r.repo.GetProductByID(ctx, id)
}

func (r *catalogService) GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	if take > 10 || (skip == 0 && take == 0) {
		take = 10
	}

	return r.repo.ListProducts(ctx, skip, take)
}
func (r *catalogService) GetProductByID(ctx context.Context, ids []string) ([]Product, error) {

	return r.repo.ListProductsWithIDs(ctx, ids)
}
func (r *catalogService) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	if take > 10 || (skip == 0 && take == 0) {
		take = 10
	}
	return r.repo.SearchProducts(ctx, query, skip, take)
}
