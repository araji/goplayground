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

var (
	filename string
	limit    int
)

//TODO: redo with io/util instead of bufio
func main() {
	flag.StringVar(&filename, "filename", "problems.csv", "exercise filename")
	flag.IntVar(&limit, "limit", 5, "time limit")
	flag.Parse()
	fmt.Printf("using  %s with a time limit of %d (s) \n", filename, limit)
	//read csv file
	fid, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open %s", filename)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	right := 0
	wrong := 0
	tstart := time.Now()
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		l := strings.Split(scanner.Text(), ",")
		if quiz(l[0], l[1]) {
			right++
		} else {
			wrong++
		}
	}
	tend := time.Now()
	fmt.Printf("Total score : correct = %d , wrong = %d  was achieved under %.2f seconds\n", right, wrong, time.Duration(tend.Sub(tstart)).Seconds())
}

func quiz(question, expected string) bool {

	fmt.Fprintf(os.Stdout, "%s =  ", question)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", err)
	}
	answer = strings.TrimSpace(answer)
	if answer != expected {
		fmt.Printf("Expected <%s> but got <%s>\n", expected, answer)
		return false
	}
	return true
}
