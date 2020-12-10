package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var completedInstructions []int
var accumulator = 0

func contains(arr []int, instruction int) bool {
	for _, a := range arr {
		if a == instruction {
			return true
		}
	}
	return false
}

func parseTxt() (instructionsArray []string) {
	f, _ := ioutil.ReadFile("input.txt")
	str := string(f)
	instructionsArray = strings.Split(str, "\n")
	return
}

func day8() (accumulator int) {
	i := 0
	instructionsArray := parseTxt()
	for {
		command := strings.Split(instructionsArray[i], " ")[0]
		intValue, _ := strconv.Atoi(strings.Split(instructionsArray[i], " ")[1])
		if !contains(completedInstructions, i) {
			completedInstructions = append(completedInstructions, i)
			if command == "acc" {
				accumulator += intValue
			} else if command == "jmp" {
				i += intValue
				continue
			}
			i++
		} else {
			return
		}
	}
}

func main() {
	fmt.Println(day8())
}
