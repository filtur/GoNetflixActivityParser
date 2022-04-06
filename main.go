package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type CsvHeader int

const (
	Title CsvHeader = iota
	WatchDate
)

type Movie struct {
	Title     string
	WatchDate time.Time
}

func main() {
	args := os.Args[1:]
	movies := make(map[string]Movie)

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

	for _, line := range data[1:] {
		if !strings.Contains(line[Title], "Season") {
			parseMovie(line, movies)
		}
	}

	fmt.Println("Enter a command")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleInput(scanner.Text(), movies)
	}
}

func handleInput(input string, movies map[string]Movie) {
	switch input {
	case "/stats":
		fmt.Printf("Items watched: %d\n", len(movies))
	case "/quit":
		os.Exit(3)
	case "/movies":
		for _, movie := range movies {
			fmt.Println(movie)
		}
	}
}

func parseMovie(line []string, movies map[string]Movie) {
	m := Movie{Title: line[Title], WatchDate: getFormattedDate(line[WatchDate])}
	movies[m.Title] = m
}

func getFormattedDate(d string) time.Time {
	fd, err := time.Parse("1/2/06", d)
	if err != nil {
		panic(err)
	}

	return fd
}
