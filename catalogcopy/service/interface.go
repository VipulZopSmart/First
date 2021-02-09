package service

import "awesomeProject/catalog/model"

type Product interface {
	Getbyid(id int) (model.Product,error)
	Create(p model.Product)(model.Product,error)
	Deletebyid(id int)(int,error)
}
