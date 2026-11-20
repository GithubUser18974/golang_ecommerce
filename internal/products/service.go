package products

import (
	"context"
	"database/sql"
)

type Service interface {
	Listproducts(ctx context.Context) ([]string, error)
}

type svc struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &svc{db: db}
}

func (s *svc) Listproducts(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT name FROM products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		products = append(products, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
