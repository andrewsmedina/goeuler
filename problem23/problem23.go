package main

import (
	"math"
	"fmt"
)

func properDivisors(number int) []int {
	divisors := make([]int, 0)
	for i := 1; i < number; i++ {
		if math.Mod(float64(number), float64(i)) == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func sumReduce(numbers []int) int {
	sum := 0
	for _, value:= range numbers {
		sum += value
	}
	return sum
}

func isAbundant(number int) bool {
	return sumReduce(properDivisors(number)) > number
}

func main() {
	fmt.Println(properDivisors(28))
}
