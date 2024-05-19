package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

const fileName = "problems.csv"

var (
	correctAnswer, totalQues int
	rowRecords               [][]string
	channel1                 = make(chan bool)
)

func main() {

	openFileName := flag.String("f", fileName, "determine a specific file")
	flag.Parse()
	openAndGetRe(*openFileName)
	timer := time.NewTimer(3 * time.Second)
	go startQuiz()
	select {

	case <-timer.C:
		fmt.Printf("\nyour result %v / %v", correctAnswer, totalQues)
	case <-channel1:
		fmt.Printf("your result %v / %v", correctAnswer, totalQues)

	}
}
func startQuiz() {

	for in, record := range rowRecords {
		ques, answer := record[0], record[1]
		var input string
		fmt.Printf("#problem %v: %v = ", in+1, ques)
		fmt.Scan(&input)
		if answer == input {
			correctAnswer++
		}
	}

	channel1 <- true
}
func openAndGetRe(openFileName string) {
	// open csv file
	file, err := os.Open(openFileName)
	if err != nil {
		fmt.Println("can't open this file", err.Error())
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// Read all records from the CSV file
	records, _ := reader.ReadAll()
	rowRecords = records
	totalQues = len(records)
}
