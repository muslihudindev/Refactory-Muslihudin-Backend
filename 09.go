package main

import (
	"fmt"
	"sort"
)

func sortMethod(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr)
}

func main() {
	var arrayList = []int{3, 12, 4, 5, 8, 9}
	sortMethod(arrayList)
}
