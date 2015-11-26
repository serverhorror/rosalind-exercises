package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counter := make(map[string]int)

	file, err := os.Open("rosalind_ini6.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counter[scanner.Text()] += 1
	}
	// fmt.Print(counter)
	for key, val := range counter {
		fmt.Printf("%s %d\n", key, val)
	}
}
