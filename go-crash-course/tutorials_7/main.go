package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// 1 + 1 + 1 + 1 +1
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	// calls wg.Done five times
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v\n", results)
	m.RUnlock()
}
