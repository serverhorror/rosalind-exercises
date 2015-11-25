package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	payload, err := ioutil.ReadFile("rosalind_ini3.txt")
	data := strings.Split(string(payload), "\n")
	numSlice := strings.Split(data[1], " ")
	for idx, elem := range numSlice {
		numSlice[idx] = strings.TrimSpace(elem)

	}

	if err != nil {
		log.Fatalf("s1: %s", err)
	}

	s1, err := strconv.ParseInt(numSlice[0], 10, 64)
	if err != nil {
		log.Fatal("s1: %s", err)
	}

	e1, err := strconv.ParseInt(numSlice[1], 10, 64)
	if err != nil {
		log.Fatal("e1: %s", err)
	}

	s2, err := strconv.ParseInt(numSlice[2], 10, 64)
	if err != nil {
		log.Fatal("s2: %s", err)
	}

	e2, err := strconv.ParseInt(numSlice[3], 10, 64)
	if err != nil {
		log.Fatal("e2: %s", err)
	}

	fmt.Println(s1, e1, s2, e2)

	fmt.Printf("%s %s\n", data[0][s1:e1+1], data[0][s2:e2+1])

}
