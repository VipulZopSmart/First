package product

import (
	"awesomeProject/catalog/model"
	"awesomeProject/catalog/store"
	"database/sql"
	"errors"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.Product {
	return &dbStore{db}
}

func (s *dbStore) GetbyProductid(id int) (model.Product, error) {

	res, err := s.db.Query("SELECT id, name,bid FROM PRODUCT WHERE id = ?", id)

	var p model.Product
	if err != nil {
		return p, err

	}

	if res.Next() {
		err := res.Scan(&p.Id, &p.Name, &p.Brand.Bid)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

func (s *dbStore) Deletebyid(id int) (int, error) {

	res, err := s.db.Exec("DELETE FROM PRODUCT WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()
	if int(rows) == 0 {
		return 0, errors.New("invalid id")
	}
	return int(rows), err

}

func (s *dbStore) CreateP(p model.Product) (model.Product, error) {
	//var p model.Product
	res, err := s.db.Exec("INSERT INTO PRODUCT(NAME,BID) VALUES(?,?)", p.Name, p.Brand.Bid)
	if err != nil {
		return model.Product{}, err
	}
	i, _ := res.LastInsertId()
	p.Id = int(i)

	return p, nil
}
