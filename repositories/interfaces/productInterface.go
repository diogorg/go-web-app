package interfaces

import models "store/models/product"

type ProductRepository interface {
	GetAll() []models.Product
	Store(product models.Product)
	DeleteById(id int, product models.Product)
	FindById(id int) models.Product
	Update(product models.Product)
}
