package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const size int = 138

// idea: find a way to surround the matrix with a frame of all '.'
// this will make the check waaaay cleaner
var matrix [1 + size + 1][1 + size + 1]rune

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prepareData(file *os.File) {
	scanner := bufio.NewScanner(file)

	// create the frame of '.'
	for colIdx := 0; colIdx < 1+size+1; colIdx++ {
		matrix[0][colIdx] = '.'
		matrix[size+1][colIdx] = '.'
	}
	for rowIdx := 0; rowIdx < 1+size+1; rowIdx++ {
		matrix[rowIdx][0] = '.'
		matrix[rowIdx][size+1] = '.'
	}

	rowIdx := 1
	for scanner.Scan() {
		line := scanner.Text()
		polishedLine := strings.TrimSpace(line)
		for colIdx, elem := range polishedLine {
			matrix[rowIdx][colIdx+1] = elem
		}
		rowIdx++
	}
}

func closeRollsCount(rowIdx, colIdx int) int {
	count := 0
	for i := rowIdx - 1; i <= rowIdx+1; i++ {
		for j := colIdx - 1; j <= colIdx+1; j++ {
			if i == rowIdx && j == colIdx {
				continue
			} else if matrix[i][j] == '@' {
				count++
			}
		}
	}
	return count
}

func isAccessible(rowIdx, colIdx int) int {
	if matrix[rowIdx][colIdx] == '.' {
		return 0
	}
	if closeRollsCount(rowIdx, colIdx) < 4 {
		matrix[rowIdx][colIdx] = 'X'
		return 1
	}
	return 0
}

func cleanMatrix() {
	for colIdx := 1; colIdx < size+1; colIdx++ {
		for rowIdx := 1; rowIdx < size+1; rowIdx++ {
			if matrix[colIdx][rowIdx] == 'X' {
				matrix[colIdx][rowIdx] = '.'
			}
		}
	}
}

func main() {
	currPath, err := os.Getwd()
	check(err)

	file, err := os.Open(filepath.Join(currPath, "/day4/input.txt"))
	check(err)
	defer file.Close()

	prepareData(file)

	secretCode := 0
	precCode := 0

	// kind naive but it way a quick change from the silver star solution so why not
	for true {
		for colIdx := 1; colIdx < size+1; colIdx++ {
			for rowIdx := 1; rowIdx < size+1; rowIdx++ {
				secretCode += isAccessible(rowIdx, colIdx)
			}
		}
		if precCode == secretCode {
			break
		}
		precCode = secretCode
		cleanMatrix()
	}
	fmt.Println(secretCode)
}
