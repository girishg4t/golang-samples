package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	argsWithProg := os.Args[1:]
	var limit int = 30
	inputArgs := strings.Split(argsWithProg[0], "=")

	if i, err := strconv.Atoi(inputArgs[1]); err == nil && inputArgs[0] == string("limit") {
		limit = i
	}

	data, err := readCSVfile("./problems.csv")
	if err != nil {
		panic(err)
	}
	correctAns := getCorrectAnsCount(data, limit)

	fmt.Printf("\nTotal number of questions correct are : %d out off %d\n", correctAns, len(data))

}
func getCorrectAnsCount(data [][]string, limit int) int {
	var correctAnsCount int = 0

	for i := 0; i < len(data); i++ {
		ans := askQues(data[i][0])
		select {
		case res := <-ans:
			if strings.TrimRight(res, "\n") == data[i][1] {
				correctAnsCount++
			}
		case <-time.After(time.Duration(limit) * time.Second):
			return correctAnsCount
		}
	}
	return correctAnsCount
}

func askQues(ques string) <-chan string {
	reader := bufio.NewReader(os.Stdin)
	c1 := make(chan string, 1)
	fmt.Printf("What is %s : ", ques)
	go func() {
		input, _ := reader.ReadString('\n')
		c1 <- input
	}()
	return c1
}

func readCSVfile(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, err
}
