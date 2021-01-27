package factorial

import "errors"

func fact(n int)(int,error){
	res:=1
	if n<0{
		return 0,errors.New("Neg number")
	}else if n==0 || n==1{
		return 1,nil
	}else{
		for i:=2;i<n+1;i++{
			res=res*i
		}
		return res,nil
	}

	return -1,nil



}
