package services

import (
	models "store/models/product"
	providers "store/repositories/providers"
)

var repository = providers.GetProductRespository()

func GetAll() []models.Product {
	return repository.GetAll()
}

func Store(product models.Product) {
	product.SetCreatedAt()
	product.SetUpdatedAt()
	repository.Store(product)
}

func Update(product models.Product) {
	product.SetUpdatedAt()
	repository.Update(product)
}

func DeleteById(id int) {
	product := models.Product{}
	product.SetDeletedAt()
	repository.DeleteById(id, product)
}

func FindById(id int) models.Product {
	return repository.FindById(id)
}
