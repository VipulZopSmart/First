package brand

import (
	"awesomeProject/catalog/model"
	"database/sql"
	"log"
)

type dbStore struct {
	db *sql.DB
}


func New(db *sql.DB) Store{
	return &dbStore{db}
}




func (s *dbStore) Getbyid(id int) (model.Brand,error) {
	db:=s.db
	res,err := db.Query("SELECT BID,NAME from BRAND WHERE BID = ?",id)
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	var p model.Brand

	if res.Next(){
		err:=res.Scan(&p.Bid,&p.Name)
		if err!=nil{
			log.Fatal(err)
		}
		return p,nil
	}else {
		log.Fatal(err)
		return p,err
	}
}
