package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var m, n uint
	// var list = make(*sync.Map, int(m))
	var list sync.Map
	parsFlags(&m, &n, &list)
}

func parsFlags(m, n *uint, list *sync.Map) {
	flag.UintVar(n, "n", 0, "Number of goroutines")
	flag.UintVar(m, "m", 0, "Time-sleep")

	flag.Parse()
	launchGoroutines(m, n, list)
}

func generateRandomTime(m *uint) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(int(*m))

	return randomNumber
}

func launchGoroutines(m, n *uint, list *sync.Map) {
	var wg sync.WaitGroup

	for i := 1; i <= int(*n); i++ {
        wg.Add(1)
        go func(i int) {
			sleepGoroutines(i, &wg, m, n, list)
		}(i)
    }

	wg.Wait()
	printInfo(list, int(*n))
}

func sleepGoroutines(id int, wg *sync.WaitGroup, m, n *uint, list *sync.Map) {
	defer wg.Done()

	randomTime := generateRandomTime(m)
	time.Sleep(time.Duration(randomTime) * time.Millisecond)
	recInfo(list, id, randomTime)
}

func recInfo(list *sync.Map, id, randomTime int) {
	list.Store(id, randomTime)
}

func printInfo(list *sync.Map, n int) {
	for i := n; i > 0; i-- {
		if timeSleep, ok := list.Load(n); ok {
			fmt.Println("Номер горутины: ", i, "Время сна: ", timeSleep)
		}
	}
}