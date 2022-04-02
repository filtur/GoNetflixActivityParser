package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: GoNetflixActivityParser <input file path>")
	}

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter a command")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleInput(scanner.Text(), data)
	}
}

func handleInput(input string, data [][]string) {
	switch input {
	case "/stats":
		fmt.Printf("Items watched: %d\n", len(data))
	case "/quit":
		os.Exit(3)
	}
}
