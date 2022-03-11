package repositories

import (
	postgres "store/db/postgres"
	models "store/models/product"
	support "store/support"
)

type PostgresProductRepository struct {
}

func (c PostgresProductRepository) GetAll() []models.Product {
	dbProducts := postgres.ExecQuery(
		"SELECT id, name, description, price, quantity from products WHERE deleted_at IS NULL",
	)
	products := []models.Product{}
	for dbProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := dbProducts.Scan(&id, &name, &description, &price, &quantity)
		support.PanicError(err)
		newProduct := models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}

		products = append(products, newProduct)
	}

	return products
}

func (c PostgresProductRepository) Store(product models.Product) {
	postgres.ExecPreparedQuery(
		"INSERT INTO products(name, description, price, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CreatedAt,
		product.UpdatedAt,
	)
}

func (c PostgresProductRepository) DeleteById(id int, product models.Product) {
	postgres.ExecPreparedQuery(
		"UPDATE products SET deleted_at=$1 WHERE id=$2",
		product.DeletedAt,
		id,
	)
}

func (c PostgresProductRepository) FindById(id int) models.Product {

	dbProduct := postgres.ExecQuery(
		"SELECT id, name, description, price, quantity from products WHERE id=$1",
		id,
	)
	foundProduct := models.Product{}
	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		scanErr := dbProduct.Scan(&id, &name, &description, &price, &quantity)
		support.PanicError(scanErr)
		foundProduct = models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}
	}
	return foundProduct
}

func (c PostgresProductRepository) Update(product models.Product) {
	postgres.ExecPreparedQuery(
		"UPDATE products SET name=$1, description=$2, price=$3, quantity=$4, updated_at=$5 WHERE id=$6",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.UpdatedAt,
		product.Id,
	)
}
