package main

import (
	"fmt"
	"math"
)

func CheckPrime(number int) bool {
	if number < 2 {
		return false
	}

	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func PrimaAfterNumber(number int) int {
	number++

	for !CheckPrime(number) {
		number++
	}
	return number
}

func PrimaSegiEmpat(wide, high, start int) string {

	var Jum int
	var res string
	var number = start

	for i := 1; i <= high; i++ {
		for j := 1; j <= wide; j++ { // Prosesmendapatkan bilangan Prima setela start
			number = PrimaAfterNumber(number)
			res += fmt.Sprintf("%v\t", number)
			Jum += number
		}
		res += "\n"
	}
	res += fmt.Sprintf("%v", Jum)
	return res
}
func main() {

	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	/*
		17 19
		23 29
		31 37
		156
	*/
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	/*
		2 3 5 7 11
		13 17 19 23 29
		129
	*/
}
