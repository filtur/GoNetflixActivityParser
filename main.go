package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/filtur/GoNetflixActivityParser/models"
)

type CsvHeader int

const (
	Title CsvHeader = iota
	WatchDate
)

func main() {
	args := os.Args[1:]
	movies := make(map[string]models.Movie)

	if len(args) == 0 {
		fmt.Println("Usage: GoNetflixActivityParser <input file path>")
		return
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

func handleInput(input string, movies map[string]models.Movie) {
	switch input {
	case "/stats":
		fmt.Printf("Movies watched: %d\n", len(movies))
	case "/quit":
		os.Exit(3)
	case "/movies":
		for _, movie := range movies {
			fmt.Println(movie)
		}
	}
}

func parseMovie(line []string, movies map[string]models.Movie) {
	m := models.Movie{Title: line[Title], WatchDate: getFormattedDate(line[WatchDate])}
	movies[m.Title] = m
}

func getFormattedDate(date string) time.Time {
	d, err := time.Parse("1/2/06", date)
	if err != nil {
		panic(err)
	}

	return d
}
