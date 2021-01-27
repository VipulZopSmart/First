package product

import (
	"awesomeProject/catalog/model"
	"database/sql"
	"errors"
	"log"
)

type dbStore struct {
	db *sql.DB
}


func New(db *sql.DB) Store{
	return &dbStore{db}
}




func (s *dbStore) Getbyid(id int) (model.Product,error) {
	db:=s.db
	res,err := db.Query("SELECT ID,NAME,BID from PRODUCT WHERE ID = ?",id)
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	var p model.Product

	if res.Next(){
		err:=res.Scan(&p.Id,&p.Name,&p.Brand.Bid)
		if err!=nil{
			log.Fatal(err)
		}
		return p,nil
	}else {

		return p,errors.New("NotFound")
	}
}
