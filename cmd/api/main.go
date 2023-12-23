package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		fileName string
		t        time.Duration
		shuffle  bool
	)

	flag.StringVar(&fileName, "filename", "problem.csv", "Quiz questions and answers")
	flag.DurationVar(&t, "timeout", 30*time.Second, "Timer set on the quiz")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the quiz questons")

	flag.Parse()

	correctCount, total := quiz(fileName, t, shuffle)
	fmt.Printf("You have answered %d questions correctly out of %d\n", correctCount, total)
}

func quiz(fileName string, t time.Duration, shuffle bool) (int, int) {
	var (
		correctCount = 0
		totalCount   = 0
	)

	records := readFile(fileName)
	totalCount = len(records)

	if shuffle {
		records = shuffleRecord(records)
	}

	fmt.Println("Press any key to start the game:")
	buf := bufio.NewReader(os.Stdin)
	_, err := buf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ti := time.NewTimer(t)
	done := make(chan bool)
	go func() {
		<-ti.C
		fmt.Println("Timer stopped")
		done <- true
	}()

	for _, record := range records {
		answerCh := make(chan string)
		go func() {
			fmt.Println(record[0])
			val, err := buf.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			val = strings.TrimSpace(val)
			answerCh <- val
		}()
		select {
		case <-done:
			return correctCount, totalCount
		case val := <-answerCh:
			if val == record[1] {
				correctCount += 1
			}
		}
	}
	return correctCount, totalCount

}
