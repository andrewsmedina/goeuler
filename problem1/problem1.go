package main

import "fmt"

func sumOfMultiplesOf3And5Bellow(number int) int {
	sum := 0
	for i := 3; i < number; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfMultiplesOf3And5Bellow(1000))
}
