package main

import (
	"flag"
	"fmt"
	"math/rand"
	// "sync"
	"time"
)

func main() {
	var m, n uint
	parsFlags(&m, &n)
	
	fmt.Println(generateRandomTime(&m))
}

func parsFlags(m, n *uint) {
	flag.UintVar(n, "n", 0, "Number of goroutines")
	flag.UintVar(m, "m", 0, "Time-sleep")

	flag.Parse()
}

func generateRandomTime(m *uint) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(int(*m))

	return randomNumber
}

