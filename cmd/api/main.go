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

func main() {
	var fileName string
	flag.StringVar(&fileName, "filename", "problem.csv", "Quiz questions and answers")
	flag.Parse()
	correctCount, total := quiz(fileName)
	fmt.Printf("You have answered %d questions correctly out of %d\n", correctCount, total)
}

func quiz(fileName string) (int, int) {
	var (
		correctCount   = 0
		incorrectCount = 0
	)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(records[0])
		buf := bufio.NewReader(os.Stdin)
		val, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		val = strings.TrimSpace(val)
		if val != records[1] {
			incorrectCount += 1
		} else {
			correctCount += 1
		}
	}
	return correctCount, correctCount + incorrectCount
}
