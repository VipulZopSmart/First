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
	output,err := dbh.GetbyProductid(3)
	if err!=nil{
		log.Fatal(model.Error)
	}
	expected := model.Product{
		Id: 3,Name: "BISCUIT", Brand: model.Brand{Bid: 1},
	}

	if expected != output {
		t.Errorf(model.Error)
	}





	mock.ExpectQuery("SELECT id, name,bid FROM PRODUCT WHERE id = ?").WithArgs(3).WillReturnError(errors.New("NOT FOUND"))

	dbh1 := dbStore{db}
	output1,err1 := dbh1.GetbyProductid(3)
	expErr:=errors.New(model.Idnotfound)
	if !reflect.DeepEqual(err1, expErr){
		t.Errorf("Expected %v error but got %v",expErr,err1)
	}
	expected1 := model.Product{}

	if expected1 != output1 {
		t.Errorf("Expected %v but got %v",expected1,output1)
	}

}





func TestDeletebyid(t *testing.T) {
	db, mock := newmock()
	defer db.Close()

	mock.ExpectExec("DELETE FROM PRODUCT WHERE id = ?").WithArgs(3).WillReturnResult(sqlmock.NewResult(0,1))

	dbh := dbStore{db}
	output,err := dbh.Deletebyid(3)
	if err!=nil{
		errors.New(model.Idnotfound)
		return
	}
	expected := 1

	if !reflect.DeepEqual(output,expected)  {
		t.Errorf(model.Error)
		return
	}



	mock.ExpectPrepare("DELETE FROM PRODUCT WHERE id = ?").ExpectExec().WithArgs(3).WillReturnResult(sqlmock.NewResult(0,0))

	dbh1 := dbStore{db}
	output1,err1 := dbh1.Deletebyid(3)
	if err1.Error()!=errors.New("invalid id").Error(){
		errors.New(model.Idnotfound)
		return
	}
	expected1 := 0

	if !reflect.DeepEqual(output1,expected1)  {
		t.Errorf(model.Error)
		return
	}


}



















func TestCreate(t *testing.T) {
	db, mock := newmock()
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "bid"})
	rows.AddRow(5, "MAGGIE", 5)

	mock.ExpectExec("INSERT INTO PRODUCT").WithArgs("MAGGIE",5).WillReturnResult(sqlmock.NewResult(5,1))

	created:=model.Product{Id: 5,Name: "MAGGIE",Brand:model.Brand{Bid: 5}}
	dbh:= dbStore{db}
	output,err := dbh.CreateP(created)
	if err!=nil{
		log.Fatal(model.Error)
	}
	expected := model.Product{
		Id: 5,Name: "MAGGIE", Brand: model.Brand{Bid: 5},
	}

	if expected != output {
		t.Errorf(model.NotMatched)
	}




	mock.ExpectExec("INSERT INTO PRODUCT").WithArgs(5).WillReturnError(errors.New("NOTCREATEDD"))

	created1:=model.Product{Id: 5,Name: "MAGGIE",Brand:model.Brand{Bid: 5}}
	dbh1:= dbStore{db}
	output1,err1 := dbh1.CreateP(created1)
	if err1!=nil{
		return
	}
	expected1 := model.Product{
		Id: 5,Name: "MAGGIE", Brand: model.Brand{Bid: 5},
	}

	if expected1 != output1 {
		t.Errorf(model.NotMatched)
	}


}



