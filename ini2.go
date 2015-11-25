package main

import (
	"encoding/csv"
	"fmt"
	"log"
	// "math"
	"os"
	"strconv"
)

func main() {
	csvfile, err := os.Open("rosalind_ini2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.Comma = ' '
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		a, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		result := (a*a)+(b*b)
		fmt.Printf("%f\n", result)
		return
	}
}
