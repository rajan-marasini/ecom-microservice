package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Repository interface {
	Close()
	PutOrder(ctx context.Context, o Order) error
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) PutOrder(ctx context.Context, o Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO orders(id, created_at, account_id, total_price) VALUES ($1, $2, $3, $4)",
		o.ID,
		o.CreatedAt,
		o.AccountID,
		o.TotalPrice,
	)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, pq.CopyIn("order_products", "order_id", "product_id", "quantity"))
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range o.Products {
		_, err = stmt.ExecContext(ctx, o.ID, p.ID, p.Quantity)
		if err != nil {
			return err
		}
	}

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT 
			o.id, 
			o.created_at, 
			o.account_id, 
			o.total_price::numeric::float8, 
			op.product_id, 
			op.quantity 
		FROM orders o 
		JOIN order_products op ON (o.id = op.order_id) 
		WHERE o.account_id = $1 
		ORDER BY o.id`,
		accountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	var lastOrder *Order

	for rows.Next() {
		var (
			id         string
			createdAt  time.Time
			accID      string
			totalPrice float64
			productID  string
			productQty uint32
		)

		if err := rows.Scan(
			&id,
			&createdAt,
			&accID,
			&totalPrice,
			&productID,
			&productQty,
		); err != nil {
			return nil, err
		}

		if lastOrder == nil || lastOrder.ID != id {
			if lastOrder != nil {
				orders = append(orders, *lastOrder)
			}
			lastOrder = &Order{
				ID:         id,
				CreatedAt:  createdAt,
				AccountID:  accID,
				TotalPrice: totalPrice,
			}
		}

		lastOrder.Products = append(lastOrder.Products, OrderedProduct{
			ID:       productID,
			Quantity: productQty,
		})
	}

	if lastOrder != nil {
		orders = append(orders, *lastOrder)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
