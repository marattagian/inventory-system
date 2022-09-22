package repository

import (
	context "context"

	entity "github.com/marattagian/inventory-system/internal/entity"
)

const (
	queryInsertProduct = `
    insert into products (name, description, price, created_by)
    values ($1, $2, $3, $4);
  `

	queryGetAllProducts = `
    select
      id,
      name,
      description,
      price,
      created_by
    from products
  `
	queryGetProductById = `
    select
      id,
      name,
      description,
      price,
      created_by
    from products
    where id = $1
  `
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {

	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)

	return err
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {

	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, queryGetProductById, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {

	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, pp, queryGetAllProducts)
	if err != nil {
		return nil, err
	}

	return pp, nil
}
