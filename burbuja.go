package main

func Burbuja(s []int64) []int64 {
	var aux int64
	var j, i int
	for i = 0; i < len(s)-1; i++ {
		for j = 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
	return s
}

func main() {

}
