package adapters

import (
	"log"
	"net/http"
	models "store/models/product"
	"strconv"
)

func GetInsertModel(r *http.Request) models.Product {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println(err)
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println(err)
	}

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	return product
}
