package product

import (
	"awesomeProject/catalog/model"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)



func newmock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}





func TestFetchbyid(t *testing.T) {
	db, mock := newmock()
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "bid"})
	rows.AddRow("3", "BISCUIT", "1")

	mock.ExpectQuery("SELECT id, name,bid FROM PRODUCT WHERE id = ?").WithArgs(3).WillReturnRows(rows)

	dbh := dbStore{db}
	output,err := dbh.Getbyid(3)
	if err!=nil{
		log.Fatal(err)
	}
	expected := model.Product{
		Id: 3,Name: "BISCUIT", Brand: model.Brand{Bid: 1},
	}

	if expected != output {
		t.Errorf("error")
	}

}
