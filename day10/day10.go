package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

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

func partOne(input []int) int {
	sort.Ints(input)
	ones, threes := 1, 1
	for i := range input[:len(input)-1] {
		plug := input[i+1] - input[i]
		if plug == 3 {
			threes++
		} else if plug == 1 {
			ones++
		}
	}
	return threes * ones
}

func partTwo(input []int) {
	sort.Ints(input)
	input = append([]int{0}, append(input, input[len(input)-1]+3)...)
	m := map[int]uint64{}
	var next func(int) uint64
	next = func(idx int) uint64 {
		if m[idx] != 0 {
			return m[idx]
		}
		if idx == len(input)-1 {
			return 1
		}
		var sum uint64
		for i := idx + 1; i < len(input); i++ {
			if input[i]-input[idx] > 3 {
				break
			}
			n := next(i)
			sum += n
			m[i] = n
		}
		return sum
	}
	fmt.Println(next(0))
}

func main() {
	fmt.Println(partOne(parseTxt()))
	partTwo(parseTxt())
}
