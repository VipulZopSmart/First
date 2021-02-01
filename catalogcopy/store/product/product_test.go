package product

import (
	"awesomeProject/catalog/model"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
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





	mock.ExpectQuery("SELECT id, name,bid FROM PRODUCT WHERE id = ?").WithArgs(3).WillReturnError(errors.New("NOT FOUND"))

	dbh1 := dbStore{db}
	output1,err1 := dbh1.Getbyid(3)
	expErr:=errors.New("NOT FOUND")
	if !reflect.DeepEqual(err1, expErr){
		t.Errorf("Expected %v error but got %v",expErr,err1)
	}
	expected1 := model.Product{}

	if expected1 != output1 {
		t.Errorf("Expected %v but got %v",expected1,output1)
	}

}





/*
func TestDeletebyid(t *testing.T) {
	db, mock := newmock()
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "bid"})
	rows.AddRow("3", "BISCUIT", "1")

	mock.ExpectQuery("SELECT id, name,bid FROM PRODUCT WHERE id = ?").WithArgs(3).WillReturnRows(rows)

	dbh := dbStore{db}
	output,err := dbh.Deletebyid(3)
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



*/















func TestCreate(t *testing.T) {
	db, mock := newmock()
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "bid"})
	rows.AddRow(5, "MAGGIE", 5)

	mock.ExpectExec("INSERT INTO PRODUCT").WithArgs("MAGGIE",5).WillReturnResult(sqlmock.NewResult(5,1))

	created:=model.Product{Id: 5,Name: "MAGGIE",Brand:model.Brand{Bid: 5}}
	dbh:= dbStore{db}
	output,err := dbh.Create(created)
	if err!=nil{
		log.Fatal(err)
	}
	expected := model.Product{
		Id: 5,Name: "MAGGIE", Brand: model.Brand{Bid: 5},
	}

	if expected != output {
		t.Errorf("Not Matched!")
	}




	mock.ExpectExec("INSERT INTO PRODUCT").WithArgs(5).WillReturnError(errors.New("NOTCREATEDD"))

	created1:=model.Product{Id: 5,Name: "MAGGIE",Brand:model.Brand{Bid: 5}}
	dbh1:= dbStore{db}
	output1,err1 := dbh1.Create(created1)
	if err1!=nil{
		return
	}
	expected1 := model.Product{
		Id: 5,Name: "MAGGIE", Brand: model.Brand{Bid: 5},
	}

	if expected1 != output1 {
		t.Errorf("Not Matched!")
	}


}



