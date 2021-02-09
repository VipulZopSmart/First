package store

import "awesomeProject/catalog/model"

type Brand interface {
	GetbyBrandid(id int) (model.Brand,error)
	CreateB(p model.Brand)(model.Brand,error)
	Check(bname string) (int,error)

}


type Product interface {
	GetbyProductid(id int) (model.Product,error)
	CreateP(p model.Product)(model.Product,error)
	Deletebyid(id int)(int,error)
}
