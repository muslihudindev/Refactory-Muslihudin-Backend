package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Datas struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

func main() {
	csv_file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	r.Read()
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var daftar Datas
	var tdaftar []Datas

	for _, rec := range records {
		daftar.Name = strings.TrimSpace(rec[0])
		daftar.Category = strings.TrimSpace(rec[1])
		daftar.Price = strings.TrimSpace(rec[2])
		tdaftar = append(tdaftar, daftar)
	}

	sort.Slice(tdaftar, func(i, j int) bool {
		return tdaftar[i].Price < tdaftar[j].Price
	})

	json_data, err := json.Marshal(tdaftar)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(json_data))

}
