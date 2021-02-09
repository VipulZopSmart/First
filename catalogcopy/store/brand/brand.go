package brand

import (
	"awesomeProject/catalog/model"
	"awesomeProject/catalog/store"
	"database/sql"
	"log"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.Brand {
	return &dbStore{db}
}

func (s *dbStore) GetbyBrandid(id int) (model.Brand, error) {
	res, err := s.db.Query("SELECT bid,name FROM BRAND WHERE BID = ?", id)
	//defer db.Close()
	var p model.Brand
	if err != nil {
		return p, err
	}

	if res.Next() {
		err := res.Scan(&p.Bid, &p.Name)
		if err != nil {
			return p, err
		}

	}
	return p, nil
}

func (s *dbStore) CreateB(p model.Brand) (model.Brand, error) {
	db := s.db
	_, err := db.Exec("INSERT INTO BRAND(NAME) VALUES(?)", p.Name)
	if err != nil {
		return model.Brand{}, err
	}

	return p, nil
}

func (s *dbStore) Check(bname string) (int, error) {
	resp, err := s.db.Query("SELECT BID FROM BRAND WHERE NAME=?", bname)
	if err != nil {
		log.Fatal(err)
	}
	var brandId int
	for resp.Next() {
		err := resp.Scan(&brandId)
		if err != nil {
			return brandId, err
		}
	}
	return brandId, nil

}
