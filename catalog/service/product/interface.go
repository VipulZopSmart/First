package service

import "awesomeProject/catalog/model"

type Service interface {
	Getbyid(id int) (model.Product,error)

}