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
	fmt.Println("limit", limit)
	var correctAnsCount int = 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(data); i++ {
		fmt.Printf("What is %s : ", data[i][0])
		c1 := make(chan string, 1)
		go func() {
			ans, _ := reader.ReadString('\n')
			c1 <- ans
		}()
		select {
		case res := <-c1:
			if strings.TrimRight(res, "\n") == data[i][1] {
				correctAnsCount++
			}
		case <-time.After(time.Duration(limit) * time.Second):
			return correctAnsCount
		}
	}
	return correctAnsCount
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
