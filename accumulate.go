package factorial

func Accumulate(giv[] string,opr func(string)string)[]string{
	for i,val:= range giv {
		giv[i]=opr(val)
	}
	return giv

}
