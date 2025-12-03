package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const SEQUENCE_LENGTH = 12

var batchList []batch

func prepareData(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currBatch := batch{}

		line := scanner.Text()
		numberStr := line[:]
		numberStr = strings.TrimSpace(numberStr)

		currBatch.strToSlice(numberStr)

		currBatch.findSequenceOfLength(SEQUENCE_LENGTH)
		batchList = append(batchList, currBatch)
		fmt.Println(currBatch)
	}
}

type batch struct {
	sequence  []int
	topDigits []int
	currIdx   int
}

func (b *batch) findSequenceOfLength(iter int) {
	currIdx := -1
	for j := iter; j > 0; j-- {
		start := currIdx + 1
		end := len(b.sequence) - j + 1
		temp := b.sequence[start:end]
		maxDigit := 0
		relativeIdx := -1
		for i, digit := range temp {
			if digit > maxDigit {
				maxDigit = digit
				relativeIdx = i
			}
		}
		currIdx = start + relativeIdx
		b.topDigits = append(b.topDigits, maxDigit)
	}
}

func (b *batch) strToSlice(s string) {
	digits := make([]int, 0, len(s))

	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}
	b.sequence = digits
}

func concatInt(list []int) int64 {
	var sb strings.Builder
	for _, digit := range list {
		sb.WriteString(strconv.Itoa(digit))
	}
	num, err := (strconv.ParseInt(sb.String(), 10, 64))
	check(err)
	return num
}

func main() {
	currPath, err := os.Getwd()
	check(err)

	file, err := os.Open(filepath.Join(currPath, "/day3/input.txt"))
	check(err)
	defer file.Close()

	prepareData(file)

	var tot int64 = 0
	for _, b := range batchList {
		tot += concatInt(b.topDigits)
	}
	fmt.Println(tot)
}
