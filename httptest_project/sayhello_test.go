package main

import (
	"fmt"
	"io/ioutil"
//	"net/http"
	"net/http/httptest"
	"testing"
)

func Testsayhell(t *testing.T){
	testcases:=[]struct {
		id int
		output string
	}{
		{0,"this is zero"},
		{1,"this is one"},
		{-1,"this is minusone"},
		{5,"this is invalid one"},
	}
	for i,testc:=range testcases{
		link := "/hello?id=%v"
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest("GET",fmt.Sprintf(link,testc.id),nil)
		sayhell(w,req)
		check:=w.Result()
		resbytess,err:=ioutil.ReadAll(check.Body)
		if err!=nil{
			t.Errorf("Expected %v Got %v error %v",testc.output,string(resbytess),i)
		}
       if string(resbytess) != testc.output {
       		t.Errorf("Expected %v Got %v error %v",testc.output,string(resbytess),i)
	   }


	}

}





