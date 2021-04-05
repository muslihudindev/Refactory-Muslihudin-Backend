package main

import (
	"fmt"
	"strings"
)

func main() {

	Words := `Aku ingin begini
Aku ingin begitu
Ingin ini itu banyak sekali

Semua semua semua
Dapat dikabulkan
Dapat dikabulkan
Dengan kantong ajaib

Aku ingin terbang bebas
Di angkasa
Hei… baling baling bambu

La... la... la...
Aku sayang sekali
Doraemon

La... la... la...
Aku sayang sekali`

	fmt.Println("aku : ", strings.Count(strings.ToLower(Words), "aku"))
	fmt.Println("ingin : ", strings.Count(strings.ToLower(Words), "ingin"))
	fmt.Println("dapat : ", strings.Count(strings.ToLower(Words), "dapat"))
}
