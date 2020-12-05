package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var seatArray []int
var highest int

func contains(arr []int, seat int) bool {
	for _, a := range arr {
		if a == seat {
			return true
		}
	}
	return false
}

func day5() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.Replace(input, "F", "0", -1)
		input = strings.Replace(input, "L", "0", -1)
		input = strings.Replace(input, "B", "1", -1)
		input = strings.Replace(input, "R", "1", -1)
		binary, _ := strconv.ParseInt(input, 2, 64)
		seatArray = append(seatArray, int(binary))
		if int(binary) > highest {
			highest = int(binary)
		}
	}
	fmt.Println("The highest seat number: ", highest)
	for i := 0; i < highest; i++ {
		if !contains(seatArray, i) && contains(seatArray, i+1) && contains(seatArray, i-1) {
			fmt.Println("My seat: ", i)
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println(start)
	day5()
	fmt.Println("Ho ho ho!")
	fmt.Println(time.Since(start))
}
