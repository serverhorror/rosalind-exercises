package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	out := make(map[string]int)
	file, err := os.Open("rosalind_dna.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		out[scanner.Text()] += 1
	}

// fmt.Println("%#v", out)
fmt.Printf("%d %d %d %d\n", out["A"], out["C"], out["G"], out["T"])
}
