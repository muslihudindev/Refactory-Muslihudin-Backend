package main

import "fmt"

func masking(str string) {
	rs := []rune(str)
	for i := 0; i < len(rs)-3; i++ {
		rs[i] = '*'
	}
	fmt.Println(string(rs))
}

func main() {
	var secret_text = "23dn3ir30fd2eddd"
	masking(secret_text)
}
