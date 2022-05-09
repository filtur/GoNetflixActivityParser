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
	var movies []models.Movie

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
			movies = append(movies, parseMovie(line))
		}
	}

	fmt.Println("Enter a command")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		t := scanner.Text()

		if strings.HasPrefix(t, "/search") {
			fmt.Println("Enter Search Term")
			c := make(chan string)
			scanner.Scan()
			s := scanner.Text()
			go searchMovies(movies[:len(movies)/2], s, c)
			go searchMovies(movies[len(movies)/2:], s, c)
			fmt.Println(<-c)
		} else {
			handleInput(t, movies)
		}

		fmt.Println("Enter a command")
	}
}

func handleInput(input string, movies []models.Movie) {
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

func searchMovies(movies []models.Movie, searchTerm string, ch chan string) {
	s := strings.ToLower(searchTerm)
	r := ""
	for _, m := range movies {
		if strings.Contains(strings.ToLower(m.Title), s) {
			r += fmt.Sprintf("%s\n", m.Title)
		}
	}
	ch <- r
}

func parseMovie(line []string) models.Movie {
	m := models.Movie{Title: line[Title], WatchDate: getFormattedDate(line[WatchDate])}
	return m
}

func getFormattedDate(date string) time.Time {
	d, err := time.Parse("1/2/06", date)
	if err != nil {
		panic(err)
	}

	return d
}
