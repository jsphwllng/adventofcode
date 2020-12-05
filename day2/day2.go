package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func partOne() (answer int) {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		count := strings.Split(line[0], "-")
		lowerBound, _ := strconv.Atoi(count[0])
		upperBound, _ := strconv.Atoi(count[1])
		letterRange := makeRange(lowerBound, upperBound)
		letter := strings.Replace(line[1], ":", "", -1)
		password := line[2]
		letterCount := strings.Count(password, letter)
		for _, val := range letterRange {
			if letterCount == val {
				answer++
			}
		}
	}
	return
}

func partTwo() (answer int) {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		score := 0
		line := strings.Split(scanner.Text(), " ")
		count := strings.Split(line[0], "-")
		lowIndex, _ := strconv.Atoi(count[0])
		highIndex, _ := strconv.Atoi(count[1])
		letter := strings.Replace(line[1], ":", "", -1)
		password := line[2]
		for i, v := range password {
			if (i+1 == highIndex || i+1 == lowIndex) && string(v) == letter {
				score++
			}
		}
		if score == 1 {
			answer++
		}
	}
	return
}

func main() {
	start := time.Now()
	fmt.Println(partOne())
	fmt.Println(partTwo())
	fmt.Println("time taken: ", time.Since(start))
}
