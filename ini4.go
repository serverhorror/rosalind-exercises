package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	c = make(chan int64, 0)
)

func genOdd(min, max int64) {
	for i := min; i <= max; i++ {
		if i%2 == 1 {
			log.Print(i)
			c <- i
		}
	}
	close(c)
}

func main() {
	var min, max, sum int64
	sum = 0
	fd, err := os.Open("rosalind_ini4.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fd)
	r.Comma=' '
	records, err := r.ReadAll()
record := records[0]
	min, err = strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	max, err = strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	go genOdd(min, max)
	for elem := range c {
		sum += elem
	}
	fmt.Println(sum)
}
