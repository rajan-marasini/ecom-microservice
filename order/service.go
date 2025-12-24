package order

import (
	"context"
	"time"
)

type Service interface {
	PostOrder()
	GetOrdersForAccount()
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice float64
	AccountID  string
	Products   []OrderedProduct
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       float32
	Quantity    uint32
}

type orderService struct {
	repository Repository
}

func NewService(r *Repository) Service {
	return nil
}

func (s orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	return nil, nil
}

func (s orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	return nil, nil
}
