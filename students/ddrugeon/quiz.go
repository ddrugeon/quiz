package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parseQuizFile(file string) (map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	quiz := map[string]string{}

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			return quiz, nil
		}

		question := strings.TrimSpace(row[0])
		answer := strings.TrimSpace(row[1])
		answer = strings.ToLower(answer)
		quiz[question] = answer
	}
}

func main() {
	var file = flag.String("file", "problems.csv", "Path of problems files (CSV Format)")
	flag.Parse()

	quiz, err := parseQuizFile(*file)
	if err != nil {
		log.Fatal(*file, err)
	}

	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Quiz!")
	fmt.Println("Reading questions from file: ", *file)
	for k, v := range quiz {
		fmt.Printf("%s?\n", k)
		fmt.Print("Your answer: ")

		scanner.Scan()
		text := scanner.Text()

		text = strings.TrimSpace(text)
		text = strings.ToLower(text)
		if strings.Compare(text, v) == 0 {
			score = score + 1
		}
	}

	fmt.Printf("\nYour final score: %d / %d\n", score, len(quiz))
}
