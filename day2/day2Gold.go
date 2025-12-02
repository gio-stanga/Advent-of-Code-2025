package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type interval struct {
	initial int
	final   int
}

var intervalList []interval

func prepareData(file *os.File) {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	delim := regexp.MustCompile(`[-,]`)
	numbers := delim.Split(line, -1)

	currInterval := interval{}

	for i, num := range numbers {
		extracted, err := strconv.Atoi(num)
		check(err)

		if i%2 == 0 {
			currInterval.initial = extracted
			continue
		}
		currInterval.final = extracted
		intervalList = append(intervalList, currInterval)
		currInterval = interval{}
	}
}

func intToSlice(n int) []int {
	s := strconv.Itoa(n)

	digits := make([]int, 0, len(s))
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}
	return digits
}

func sliceToInt(s []int) int {
	n := 0
	for _, digit := range s {
		n = n*10 + digit
	}
	return n
}

func isValidCode(slice []int) bool {

	isValid := false
	for period := 1; period <= len(slice)/2; period++ {
		isValid = false
		if len(slice)%period != 0 {
			continue
		}
		for i := 0; i < period; i++ {
			for k := i + period; k < len(slice); k += period {
				if slice[i] != slice[k] {
					isValid = true
					break
				}
			}

		}
		if isValid == false {
			return isValid
		}
	}
	return true
}

func iterateInterval(initial, final int) int {
	tot := 0
	for i := initial; i <= final; i++ {
		if !isValidCode(intToSlice(i)) {
			tot += i
		}
	}
	return tot
}

func main() {
	currPath, err := os.Getwd()
	check(err)
	file, err := os.Open(filepath.Join(currPath, "day2/input.txt"))
	check(err)
	defer file.Close()
	prepareData(file)

	tot := 0
	for _, currInterval := range intervalList {
		tot += iterateInterval(currInterval.initial, currInterval.final)
	}
	fmt.Println(tot)
}
