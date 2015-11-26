package main

import (
	"github.com/serverhorror/rog-go/reverse"

	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("rosalind_revc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := reverse.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		in := scanner.Text()
		switch in {
		case "A":
			fmt.Print("T")
		case "C":
			fmt.Print("G")
		case "G":
			fmt.Print("C")
		case "T":
			fmt.Print("A")
		}
	}
	fmt.Println()
}
