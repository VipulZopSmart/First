package factorial

import "testing"


func Test_fact(t *testing.T) {
	testcases:=[]struct {
		in  int
		out int
		error bool
	}{
		{1,1,false},
		{0,1,false},
		{4,24,false},
		{-5,0,true},
	}
		for _,use:=range testcases{
			res,err:=fact(use.in)
			if err!=nil && !use.error{
				t.Errorf("Error %v",err)
			}
			if res!=use.out {
				t.Errorf("expected %v,got %v", use.out, res)
			}
		}

}
