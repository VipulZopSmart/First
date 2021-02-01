package handler

import (
	"awesomeProject/catalog/model"
	serviceInterface "awesomeProject/catalog/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type handlerr struct {
	servicehandle serviceInterface.Product
}

func New(s serviceInterface.Product) handlerr {
	return handlerr{s}
}





func (s handlerr) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.Get(w, r)
	}
	if r.Method == "POST" {
		s.Post(w, r)
	}

}


func (s handlerr) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)

		return
		//log.Fatal(err)
	}

	product, err := s.servicehandle.Getbyid(i)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve data"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	p, _ := json.Marshal(product)
	//return p
	w.WriteHeader(200)
	_, _ = w.Write(p)

}





func (s handlerr) Post(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &p)
	if err != nil {
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := s.servicehandle.Create(p)
	if err != nil {
		_, _ = w.Write([]byte("could not create product"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.WriteHeader(http.StatusCreated)
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)


}
