package handler

import (
	"awesomeProject/catalog/model"
	service2 "awesomeProject/catalog/service"
	"bytes"
	"strconv"

	//"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetbyid(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := service2.NewMockProduct(ctrl)

	testcases := []struct {
		id int
		expected model.Product
		//expectedErr error
		err error
	}{
		{3, model.Product{Id: 3, Name: "BISCUIT", Brand: model.Brand{Bid: 1, Name: "OREO"}}, nil},
		{22, model.Product{Id: 22, Name: "CAR", Brand: model.Brand{Bid: 7, Name: "BMW"}}, nil},
		//{"abc",model.Product{},errors.New("invalid perimeter id")},
	}

	for _, tc := range testcases {
		ps.EXPECT().Getbyid(tc.id).Return(tc.expected, nil)
		serv := New(ps)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/product?id="+strconv.Itoa(tc.id), nil)

		serv.Get(w, r)
		result := w.Result()
		res, _ := ioutil.ReadAll(result.Body)

		actual, _ := json.Marshal(tc.expected)

		if !reflect.DeepEqual(res, actual) {
			t.Errorf(model.Notequal)
		}
	}

}

func TestHandlerrPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := service2.NewMockProduct(ctrl)

	//testcases
	testcases := []struct {
		body     []byte
		resp     model.Product
		p        model.Product
		err      error
		expected string
	}{
		{[]byte(`{"name": "ONEPLUS6","brand": {"name": "ONE PLUS"}}`), model.Product{Id: 24, Name: "ONEPLUS6", Brand: model.Brand{Bid: 4, Name: "ONE PLUS"}}, model.Product{Name: "ONEPLUS6", Brand: model.Brand{Name: "ONE PLUS"}}, nil, `{"id":24,"name":"ONEPLUS6","brand":{"bid":4,"name":"ONE PLUS"}}`},
	}

	for _, tc := range testcases {
		ps.EXPECT().Create(tc.p).Return(tc.resp, tc.err)
		serv := New(ps)
		req := httptest.NewRequest("POST", "/product", bytes.NewBuffer(tc.body))
		w := httptest.NewRecorder()

		serv.Post(w, req)
		r := w.Result()
		actual, _ := ioutil.ReadAll(r.Body)

		if !reflect.DeepEqual(string(actual), tc.expected) {
			t.Errorf(model.Idnotfound)
		}
	}
}
