package main

import (
	"fmt"
)

func PrimeX(number int) interface{} {

	var res interface{}

	deret := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	if number > 0 && number < 11 {
		res = deret[number-1]
	} else {
		res = "Silahkan masukkan dari angka 1 - 10"
	}
	return res
}

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29
}
