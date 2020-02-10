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
	"time"
)

var (
	filename string
	limit    int
)

/**
func timer(timeout int, done chan bool) {
	tstart := time.Now()
	fmt.Printf("Timer of %d started \n", timeout)
	for {
		if time.Now().Sub(tstart).Seconds() >= float64(timeout) {
			break

		}
		//done <- false
	}
	fmt.Println("TIME's UP\n ")
	done <- true
}

*/
func main() {
	flag.StringVar(&filename, "filename", "problems.csv", "exercise filename")
	flag.IntVar(&limit, "limit", 10, "time limit")
	flag.Parse()
	//done := make(chan bool)
	answerChan := make(chan bool)
	//read csv file
	fid, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open %s", filename)
	}
	defer fid.Close()
	csvReader := csv.NewReader(fid)
	right := 0
	wrong := 0
	fmt.Fprintf(os.Stdout, "Hit Enter to start")
	reader := bufio.NewReader(os.Stdin)
	_, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", err)
	}

	timer1 := time.NewTimer(time.Duration(limit * int(time.Second)))
	defer timer1.Stop()

	tstart := time.Now()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%s", err)
		}
		go quiz(strings.TrimSpace(record[0]), strings.TrimSpace(record[1]), answerChan, *timer1)
		select {
		case <-timer1.C:
			fmt.Println("TIMEOUT RECEIVED")
			break
		case msg2 := <-answerChan:
			fmt.Println("ANSWER RECEIVED")
			if msg2 == true {
				right++
			} else {
				wrong++
			}
		}

	}
	//fmt.Println("Out of Select")
	//break

	tend := time.Now()
	fmt.Printf("Total score : correct = %d , wrong = %d  was achieved under %.2f seconds\n", right, wrong, time.Duration(tend.Sub(tstart)).Seconds())
}

func quiz(question, expected string, answerChan chan<- bool, timeout time.Timer) {
	fmt.Fprintf(os.Stdout, "%s =  ", question)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", err)
	}
	answer = strings.TrimSpace(answer)
	if answer != expected {
		//fmt.Printf("Expected %s but got %s\n", expected, answer)
		answerChan <- false
	} else {
		answerChan <- true
	}

}
