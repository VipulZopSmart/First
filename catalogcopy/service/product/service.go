package service

import (
	"awesomeProject/catalog/model"
	serviceInterface "awesomeProject/catalog/service"
	"awesomeProject/catalog/store"
	"log"
)

type service struct {
	productInterface store.Product
	brandInterface   store.Brand
}

func New(productStore store.Product, brandstore store.Brand) serviceInterface.Product {
	return service{productStore, brandstore}

}

func (s service) Getbyid(id int) (model.Product, error) {
	productdetails, err := s.productInterface.GetbyProductid(id)

	if err != nil {
		return model.Product{}, err
	}
	branddetails, err := s.brandInterface.GetbyBrandid(productdetails.Brand.Bid)

	if err != nil {
		
		return model.Product{}, err
	}
	productdetails.Brand.Name = branddetails.Name
	return productdetails, nil

}

func (s service) Create(p model.Product) (model.Product, error) {

	bID, err := s.brandInterface.Check(p.Brand.Name)
	if bID == 0 {
		branddetails, err := s.brandInterface.CreateB(p.Brand)
		if err != nil {
			log.Fatal(err)
		}
		bID = branddetails.Bid

	}

	p.Brand.Bid = bID
	res, err := s.productInterface.CreateP(p)
	if err != nil {
		return model.Product{}, err
	}
	return res, err

}
