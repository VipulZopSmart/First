package product

import "awesomeProject/catalog/model"

type Store interface {
	Getbyid(id int) (model.Product,error)
}


