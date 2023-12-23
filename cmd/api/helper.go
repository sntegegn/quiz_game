package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

func readFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func shuffleRecord(record [][]string) [][]string {
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(record) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		record[i], record[j] = record[j], record[i]
	}
	return record
}
