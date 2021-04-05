package main

import (
	"fmt"
)

func sort_odd(arr []int) {
	var temp int
	for i, n := range arr {
		if n%2 == 0 {
			continue
		}
		for j := (i + 1); j < len(arr); j++ {
			if arr[j]%2 != 0 {
				if arr[i] > arr[j] {
					temp = arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	var numbers = []int{9, 4, 2, 4, 1, 5, 3, 0}
	sort_odd(numbers)
}
