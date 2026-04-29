package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleepRandomly() {
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
}

func routine(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("🟡 Горутина %d начала работу\n", index)
	sleepRandomly()
	fmt.Printf("🟢 Горутина %d завершила работу\n", index)
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go routine(i, &wg)
	}

	wg.Wait()
}
