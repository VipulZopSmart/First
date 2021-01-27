package main

import (
	handler "awesomeProject/catalog/handler/product"
	"awesomeProject/catalog/service/product"
	"awesomeProject/catalog/store/brand"
	"awesomeProject/catalog/store/product"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)




const (
	username = "newuser"
	password = "Password_123"
	hostname = "127.0.0.1"
	dbname   = "CATALOGUE"
)


func dsn(s string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}



func main(){
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	productt:=product.New(db)
	brandd:=brand.New(db)
	ser:=service.New(productt,brandd)
	h:=handler.New(ser)
	http.HandleFunc("/product",h.Handler)

	fmt.Println(http.ListenAndServe(":8080",nil))

}
