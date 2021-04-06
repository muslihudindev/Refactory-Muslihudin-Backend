package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	rupiah "github.com/yudapc/go-rupiah"
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

	var tdaftar []Datas
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		feetFloat, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		formatRupiah := rupiah.FormatRupiah(float64(feetFloat))
		tdaftar = append(tdaftar, Datas{
			Name:     strings.TrimSpace(record[0]),
			Category: strings.TrimSpace(record[1]),
			Price:    formatRupiah,
		})
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
