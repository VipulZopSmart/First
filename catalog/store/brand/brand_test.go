package brand

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
	rows := sqlmock.NewRows([]string{"bid","name" })
	rows.AddRow("1", "OREO")

	mock.ExpectQuery("SELECT bid,name FROM BRAND WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	dbh := dbStore{db}
	output,err := dbh.Getbyid(1)
	if err!=nil{
		log.Fatal(err)
	}
	expected := model.Brand{
		Bid: 1,Name: "OREO",
	}

	if expected != output {
		t.Errorf("error")
	}

}
