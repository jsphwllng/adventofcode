// acc increases accumulator
// jump jumps to the x code
// nop does nothing

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

func day8() (accumulator int) {
	var i int
	f, _ := ioutil.ReadFile("input.txt")
	//as ReadFile returns bytes
	str := string(f)
	//divide up the instructions
	instructionsArray := strings.Split(str, "\n")
	for {
		// divide up line into the command and the value
		command := strings.Split(instructionsArray[i], " ")[0]
		intValue, _ := strconv.Atoi(strings.Split(instructionsArray[i], " ")[1])
		if !contains(completedInstructions, i) {
			completedInstructions = append(completedInstructions, i)
			if command == "acc" {
				accumulator += intValue
				i++
			} else if command == "jmp" {
				i += intValue
			} else {
				i++
			}
		} else {
			return
		}
	}
}

func main() {
	fmt.Println(day8())
}
