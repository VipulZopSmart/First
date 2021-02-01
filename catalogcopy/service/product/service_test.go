package service

import (
	"awesomeProject/catalog/model"
	"awesomeProject/catalog/store"
	"log"
	"reflect"

	//"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetbyid(t *testing.T) {

	ctrl := gomock.NewController(t)
	bs := store.NewMockBrand(ctrl)
	ps := store.NewMockProduct(ctrl)

	ser := New(ps, bs)

	testcases := []struct {
		id            int
		productdata   model.Product
		branddata     model.Brand
		expected      model.Product
		producterror  error
		branderror    error
		expectederror error
	}{
		{3, model.Product{Id: 3, Name: "BISCUIT", Brand: model.Brand{Bid: 1, Name: ""}}, model.Brand{Bid: 1, Name: "OREO"}, model.Product{Id: 3, Name: "BISCUIT", Brand: model.Brand{Bid: 1, Name: "OREO"}}, nil, nil, nil},
		{2, model.Product{Id: 2, Name: "SHAMPOO", Brand: model.Brand{Bid: 2, Name: ""}}, model.Brand{Bid: 2, Name: "BEER"}, model.Product{Id: 2, Name: "SHAMPOO", Brand: model.Brand{Bid: 2, Name: "BEER"}}, nil, nil, nil},
	}

	for _, tc := range testcases {
		ps.EXPECT().GetbyProductid(tc.productdata.Id).Return(tc.productdata, tc.producterror)
		bs.EXPECT().GetbyBrandid(tc.productdata.Brand.Bid).Return(tc.branddata, tc.branderror)

		res, err := ser.Getbyid(tc.id)
		if err != nil {
			log.Fatal(err)
		}

		if res != tc.expected {
			log.Fatal("Error")
		}

		//t.Log(res, err)

	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := store.NewMockProduct(ctrl)
	bs := store.NewMockBrand(ctrl)


	testcases := []struct {

		bid int
		perr error
		cerr error
		products model.Product
		brandout model.Brand
		productout model.Product
		berr error
		expected model.Product
		expectederror error
	}{
		{5,nil,nil,model.Product{Id: 5,Name: "MAGGIE",Brand:model.Brand{Bid: 5,Name: "YIPPEE"}},model.Brand{Name: "YIPPEE"},model.Product{Id: 5,Name: "MAGGIE",Brand: model.Brand{Bid: 5,Name: "YIPPEE"}},nil,model.Product{Id: 5,Name: "MAGGIE",Brand: model.Brand{Bid: 5,Name: "YIPPEE"}},nil},

	}
		for i,tc:=range testcases{
			bs.EXPECT().Check(tc.products.Brand.Name).Return(tc.bid,tc.cerr)
			if tc.bid==0{
				bs.EXPECT().CreateB(tc.products.Brand).Return(tc.brandout,tc.berr)
				tc.bid=tc.brandout.Bid

			}
			tc.products.Brand.Bid=tc.bid
			ps.EXPECT().CreateP(tc.products).Return(tc.productout,tc.perr)
			serv := New(ps, bs)
			res,err:=serv.Create(tc.productout)

			if !reflect.DeepEqual(res,tc.expected){
				t.Errorf("test failed %v got %v but expected %v", i, res, tc.expected)
			}
			if !reflect.DeepEqual(err, tc.expectederror){
				t.Errorf("test failed %v got %v but expected %v", i, err, tc.expectederror)
			}
		}
}



