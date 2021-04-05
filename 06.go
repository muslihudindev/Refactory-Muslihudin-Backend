package main

import (
	"fmt"
)

func missing(a, b []string) string {
	ma := make(map[string]bool, len(a))
	for _, ka := range a {
		ma[ka] = true
	}
	for _, kb := range b {
		if !ma[kb] {
			return kb
		}
	}
	return ""
}

func stringInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findmising(arr []rune) string {

	for r := arr[0]; r <= arr[len(arr)-1]; r++ {
		if !stringInSlice(r, arr) {
			return string(r)
		}
	}
	return ""
}

func main() {
	var list_letters_first = []rune{'c', 'd', 'e', 'g', 'h'}
	var list_letters_second = []rune{'X', 'Z'}

	fmt.Println("list_letters_first = ", findmising(list_letters_first))
	fmt.Println("list_letters_second = ", findmising(list_letters_second))

}
