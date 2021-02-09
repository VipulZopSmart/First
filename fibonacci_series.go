package factorial

import (
	"fmt"
)

func solve() func() int {
	n:=0
	t1:=0
	t2:=1
	next:=t1+t2
	return func() int {
		result:=0
		if n==0 {
			n++
			result = 0
		}else if n==1 {
			n++
			result=1
		}else{
			result=next
			t1=t2
			t2=next
			next=t1+t2
		}
		return result
	}
}


func main() {
	fibo:=solve()
	for i:=0; i<=10; i++ {
		fmt.Println(fibo())
	}
}
