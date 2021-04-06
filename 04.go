package main

import (
	"fmt"
	"math"
	"strings"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func main() {
	var even []int
	var odd []int
	var multiply []int
	var prime []int
	var primelessthan100 []int

	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
		if i%5 == 0 {
			multiply = append(multiply, i)
		}

		if IsPrime(i) {
			prime = append(prime, i)
		}
		if IsPrime(i) && i < 100 {
			primelessthan100 = append(primelessthan100, i)
		}
	}

	fmt.Println("1) even : ", arrayToString(even, ", "))
	fmt.Println("\n2) odd : ", arrayToString(odd, ", "))
	fmt.Println("\n3) numbers multiplies by 5 : ", arrayToString(multiply, ", "))
	fmt.Println("\n4) prime numbers : ", arrayToString(prime, ", "))
	fmt.Println("\n5) prime numbers less than 100 : ", arrayToString(primelessthan100, ", "))
}
