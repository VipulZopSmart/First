package main

import (
	"net/http"
	"strconv"
)


func sayhell(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
	iid:=r.URL.Query()["id"]
	iin,err:=strconv.Atoi(iid[0])
	if err!=nil{
		w.WriteHeader(500)
	}
	if iin==0{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("this is zero"))
	}
	if iin==-1{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("this is mone"))

	}
	if iin==1{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("this is one"))
	}
	if iin==5{
		w.WriteHeader(500)
		w.Write([]byte("Invalid"))
	}

}




