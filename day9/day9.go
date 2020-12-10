package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

var answer = []int{}

func parseTxt() (intArray []int) {
	f, _ := ioutil.ReadFile("input.txt")
	str := string(f)
	stringArray := strings.Split(str, "\n")
	for _, number := range stringArray {
		newNumber, _ := strconv.Atoi(number)
		intArray = append(intArray, newNumber)
	}
	return
}

func partOne(numArray []int) (answer int) {
	for i := 25; i < len(numArray); i++ {
		addsUp, target := false, numArray[i+1]
		for _, a := range numArray[i-25 : i+1] {
			for _, b := range numArray[i-25 : i+1] {
				if a+b == target && a != b {
					addsUp = true
				}
			}
		}
		if !addsUp {
			return numArray[i+1]
		}
	}
	return
}

func partTwo(goal int, numArray []int) (highest, lowest int) {
	for indexOne, a := range numArray {
		total, answer := a, append(answer, a)
		for _, b := range numArray[indexOne+1:] {
			total += b
			answer = append(answer, b)
			if total == goal {
				sort.Ints(answer)
				lowest := answer[0]
				highest := answer[len(answer)-1]
				fmt.Println("highest:", highest, "lowest:", lowest, "total:", highest+lowest)
				return highest, lowest
			}
		}
		answer = nil
	}
	return
}

func main() {
	start := time.Now()
	numArray := parseTxt()
	partOne := partOne(numArray)
	partTwo(partOne, numArray)
	fmt.Println("time taken: ", time.Since(start))
}
