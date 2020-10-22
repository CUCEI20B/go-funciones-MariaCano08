package main
func Burbuja(s []int64) []int64{
	var aux int64;
	var t,j,i int;
	t=len(s);
	for  i=1; i<t ;i++{
		for j=0;j< (t-2);j++{
			if s[j]> s[j+1]{
				aux=s[j];
				s[j]=s[j+1]
				s[j+1]=aux;
			}
		}
	}
	return s;
}

func main()  {
	
}