package providers

import (
	interfaces "store/repositories/interfaces"
	product "store/repositories/product"
)

func GetProductRespository() interfaces.ProductRepository {
	return product.PostgresProductRepository{}
}
