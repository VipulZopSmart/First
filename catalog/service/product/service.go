package service

import (
	"awesomeProject/catalog/model"
	"awesomeProject/catalog/store/brand"
	"awesomeProject/catalog/store/product"
	//"awesomeProject/catalog/store/product"
	"fmt"
	"log"
)

type sstore struct {
	productInterface product.Store
	brandInterface brand.Store
}

func New(productstore product.Store,brandstore brand.Store)Service{
	return sstore{productstore,brandstore}

}

func (s sstore)  Getbyid(id int)(model.Product,error){
	productdetails,err:=s.productInterface.Getbyid(id)
	if err!=nil{
		log.Fatal(err)
	}
	branddetails,err:=s.brandInterface.Getbyid(productdetails.Brand.Bid)
	if err!=nil{
		log.Fatal(err)
	}
	productdetails.Brand.Name = branddetails.Name
	fmt.Println(productdetails,branddetails)

	return productdetails,nil

}
