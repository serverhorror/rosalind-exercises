package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("rosalind_rna.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		in := scanner.Text()
		switch in {
		case "T":
			fmt.Print("U")
		default:
			fmt.Print(in)
		}
	}
}
