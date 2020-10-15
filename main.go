package main

import "fmt"

func main() {
	var n int
	var x int
	var s []int // slice
	var resultado int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		s = append(s, x)
	}
	for _, k := range s {
		//fmt.Println(k)
		resultado += k
	}
	fmt.Println(resultado)
}
