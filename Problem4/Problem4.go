package main

import "fmt"

func MaxSequence(arr []int) int {
	Max := arr[0]

	for i := 0; i < len(arr); i++ { // i=0, -2 | i=1
		var sum = 0
		for j := i; j < len(arr); j++ { // j=0, -2 | j=1, 1 | j>8
			sum += arr[j]

			if Max <= sum {
				Max = sum
			}
		}
	}
	return Max
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   //7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   //7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   //8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    //12
}
