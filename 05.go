package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func duplicateCount(s1 string) {
	data := make(map[string]interface{})
	word := strings.Split(s1, "")
	for _, r := range word {
		replacement := strings.Repeat("*", strings.Count(strings.ToLower(s1), strings.ToLower(r)))
		if r == " " {
			continue
		}

		data[strings.ToLower(r)] = replacement

	}
	j, _ := json.Marshal(data)
	fmt.Println(string(j))

}

func main() {
	var text_1 = "Mammals"
	var text_2 = "Bruiser build"

	duplicateCount(text_1)
	duplicateCount(text_2)
}
