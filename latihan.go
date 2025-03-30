package main
import "fmt"

func konsekutif(n int)bool{
	var d1, d2 int
	var b bool 
	b = false
	for n > 0{
		d1 = n%10
		d2 = n%100
		if d1 - d2 == 0 && d2 - d1 == 0{
			b = true
		}
		n = n/10
	}
	return b
}
func main(){
	var n int 
	fmt.Scan(&n)
	fmt.Print(konsekutif(n))
}



