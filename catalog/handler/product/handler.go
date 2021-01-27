package handler

import (
	service "awesomeProject/catalog/service/product"
	"encoding/json"
	"net/http"
	"strconv"
)

type handlerr struct {
	servicehandle service.Service

}

func New(s service.Service)handlerr{
	return handlerr{s}
}


func (s handlerr)Handler(w http.ResponseWriter,r *http.Request ){
	if r.Method=="GET"{
		s.get(w,r)
	}
}




func (s handlerr)get(w http.ResponseWriter,r *http.Request){
	id:=r.URL.Query().Get("id")
	i,err:=strconv.Atoi(id)
	if err!=nil{
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)

		return
		//log.Fatal(err)
	}

	product,err:=s.servicehandle.Getbyid(i)
	if err!=nil{
		_, _ = w.Write([]byte("could not retrieve data"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	p,_:=json.Marshal(product)
	//return p
	w.WriteHeader(200)
	_,_=w.Write(p)



}