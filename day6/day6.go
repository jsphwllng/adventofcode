package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var testingArray []string
var answerOne int = 0
var answerArray []int
var groupsAnswer string
var answerTwo int

func main() {
	fmt.Println(day6One())
	fmt.Println(day6Two())
}

func contains(arr []string, char string) bool {
	for _, a := range arr {
		if a == char {
			return true
		}
	}
	return false
}

func day6One() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//reset the array
		testingArray = nil
		guestAnswer := scanner.Text()
		if guestAnswer != "" {
			groupsAnswer += guestAnswer
		} else {
			//if a character doesn't exsist in the groups answers then add to it
			for _, character := range groupsAnswer {
				if !contains(testingArray, string(character)) {
					testingArray = append(testingArray, string(character))
				}
			}
			//increase the total by the individual characters
			answerOne += len(testingArray)
			groupsAnswer = ""
		}
	}
	return answerOne
}

func day6Two() int {
	f, _ := ioutil.ReadFile("input.txt")
	//as ReadFile returns bytes
	str := string(f)
	//divide up the groups
	strArray := strings.Split(str, "\n\n")
	for _, group := range strArray {
		//to get individual answers
		group := strings.Split(group, "\n")
		if len(group) == 1 {
			//as no duplicates allowed
			answerTwo += len(string(group[0]))
		} else {
			answerMap := make(map[string]int)
			// make a map for example
			// a: 1
			// b: 2
			// c: 3
			// c: 4
			for _, individualAnswer := range group {
				answers := strings.Split(individualAnswer, "")
				for _, question := range answers {
					answerMap[string(question)]++
				}
			}
			fmt.Println(answerMap)
			for _, v := range answerMap {
				if v == len(group) {
					answerTwo++
				}
			}
			//reset the map
			answerMap = nil
		}
	}
	return answerTwo
}
