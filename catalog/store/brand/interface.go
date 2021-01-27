package brand

import "awesomeProject/catalog/model"

type Store interface {
	Getbyid(id int) (model.Brand,error)
}


