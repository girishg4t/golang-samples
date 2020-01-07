package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	data, err := readCSVfile("./problems.csv")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println(data[0][1])
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
