package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

func readFile(fileName string) []problem {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return parseLines(records)
}

func shuffleRecords(records []problem) []problem {
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(records) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		records[i], records[j] = records[j], records[i]
	}
	return records
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
