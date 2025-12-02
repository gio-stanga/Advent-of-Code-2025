package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DIALSIZE = 100
const MIN = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type move struct {
	Value  int
	IsLeft bool
}

func prepareData(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		currentMove := move{}

		if strings.HasPrefix(line, "L") {
			currentMove.IsLeft = true
		} else {
			currentMove.IsLeft = false
		}
		numberStr := line[1:]
		numberStr = strings.TrimSpace(numberStr)

		num, err := strconv.Atoi(numberStr)
		check(err)

		currentMove.Value = num
		movesList = append(movesList, currentMove)
	}
}

func mod(a int) int {
	m := a % DIALSIZE
	if m < 0 {
		m += DIALSIZE
	}
	return m
}

var movesList []move

func main() {
	currPath, err := os.Getwd()
	check(err)

	file, err := os.Open(filepath.Join(currPath, "day1/input.txt"))
	check(err)
	defer file.Close()

	prepareData(file)

	var secretCode int = 0
	var pointer int = 50
	for _, currentMove := range movesList {
		var rotations int = 0

		rotations = currentMove.Value / DIALSIZE

		remainder := currentMove.Value % DIALSIZE

		if currentMove.IsLeft {
			if pointer > 0 && pointer-remainder <= 0 {
				rotations++
			}
			pointer = mod(pointer - currentMove.Value)
		} else {
			if pointer+remainder >= DIALSIZE {
				rotations++
			}
			pointer = mod(pointer + currentMove.Value)
		}
		secretCode += rotations
	}
	fmt.Println(secretCode)
}
