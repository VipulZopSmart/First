package brand

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
	rows := sqlmock.NewRows([]string{"bid", "name"})
	rows.AddRow("1", "OREO")

	mock.ExpectQuery("SELECT bid,name FROM BRAND WHERE BID = ?").WithArgs(1).WillReturnRows(rows)

	dbh := dbStore{db}
	output, err := dbh.GetbyBrandid(1)
	if err != nil {
		log.Fatal(err)
	}
	expected := model.Brand{
		Bid: 1, Name: "OREO",
	}

	if expected != output {
		t.Errorf(model.Error)
	}

	mock.ExpectQuery("SELECT bid,name FROM BRAND WHERE BID = ?").WillReturnError(errors.New("NOT FOUND"))

	dbh1 := dbStore{db}
	output1, err1 := dbh1.GetbyBrandid(1)
	expErr := errors.New(model.Idnotfound)
	if !reflect.DeepEqual(err1, expErr) {
		t.Errorf("Expected %v error but got %v", expErr, err1)
	}
	expected1 := model.Brand{}

	if expected1 != output1 {
		t.Errorf("Expected %v but got %v", expected1, output1)
	}






}

func TestCreate(t *testing.T) {
	db, mock := newmock()
	defer db.Close()
	rows := sqlmock.NewRows([]string{"bid", "name"})
	rows.AddRow(5, "YIPPEE")

	mock.ExpectExec("INSERT INTO BRAND").WithArgs("YIPPEE").WillReturnResult(sqlmock.NewResult(5, 1))

	created := model.Brand{Bid: 5, Name: "YIPPEE"}
	dbh := dbStore{db}
	output, err := dbh.CreateB(created)
	if err != nil {
		log.Fatal(err)
	}
	expected := model.Brand{
		Bid: 5, Name: "YIPPEE",
	}

	if expected != output {
		t.Errorf(model.NotMatched)
	}

	mock.ExpectExec("INSERT INTO BRAND").WithArgs("YIPPEE").WillReturnError(errors.New("NOT FOUND"))

	created1 := model.Brand{Bid: 5, Name: "YIPPEE"}
	dbh1 := dbStore{db}
	output1, err1 := dbh1.CreateB(created1)
	if err1 != nil {
		return
	}
	expected1 := model.Brand{
		Bid: 5, Name: "YIPPEE",
	}

	if expected1 != output1 {
		t.Errorf(model.NotMatched)
	}
}
