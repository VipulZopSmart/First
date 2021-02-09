package handler

import (
	"awesomeProject/catalog/model"
	serviceInterface "awesomeProject/catalog/service"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type handlerr struct {
	servicehandle serviceInterface.Product
}

type responseerror struct {
	statuscode int   `json:"statuscode"`
	err        error `json:"err"`
}


func(e responseerror)Error() string{
	return fmt.Sprintf("Error %v - Status %v",e.err,e.statuscode)
}



func New(s serviceInterface.Product) handlerr {
	return handlerr{s}
}

func (s handlerr) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.Get(w, r)
	}
	if r.Method == http.MethodPost {
		s.Post(w, r)
	}
	if r.Method == http.MethodDelete {
		s.Delete(w, r)
	}

}

func (s handlerr) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		resp := responseerror{statuscode: http.StatusBadRequest, err: errors.New("invalid perimeter id")}
		r, _ := json.Marshal(resp.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(r)

		return
		//log.Fatal(err)
	}

	product, err := s.servicehandle.Getbyid(i)
	if err != nil {
		resp := responseerror{statuscode: 500, err: errors.New("couldn't retrieve data")}
		r, _ := json.Marshal(resp.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(r)
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

		resp := responseerror{statuscode: 400, err: errors.New("invalid perimeter id")}
		r, _ := json.Marshal(resp.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(r)

		return
	}

	resp, err := s.servicehandle.Create(p)
	if err != nil {

		resp := responseerror{statuscode: 500, err: errors.New("couldn't retrieve data")}
		r, _ := json.Marshal(resp.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(r)
		return



	}
	w.WriteHeader(http.StatusCreated)
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)

}



func (s handlerr) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		resp := responseerror{statuscode: 400, err: errors.New("invalid perimeter id")}
		r, _ := json.Marshal(resp.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(r)

		return

	}

	product, err := s.servicehandle.Deletebyid(i)
	if err != nil || product==0 {
		resp := responseerror{statuscode: 500, err: errors.New("couldn't retrieve data")}
		r, _ := json.Marshal(resp.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(r)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("DELETED!!"))

}
