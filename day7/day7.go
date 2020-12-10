package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(rand.Intn(100) * time.Millisecond)
	fmt.Println("I have no idea how to invert a binary tree, sorry!")
	fmt.Println("time taken: ", time.Since(start))
}
