// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	fmt.Println("Race Codition")

// 	wg := &sync.WaitGroup{}
// 	mut := &sync.RWMutex{}

// 	var score = []int{0}

// 	wg.Add(3)

// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("One R")
// 		mut.Lock()
// 		score = append(score, 1)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Two R")
// 		mut.Lock()
// 		score = append(score, 2)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Three R")
// 		mut.Lock()
// 		score = append(score, 3)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Three R")
// 		mut.RLock()
// 		fmt.Println(score)
// 		mut.RUnlock()
// 		wg.Done()
// 	}(wg, mut)

// 	wg.Wait()

// 	fmt.Println(score)
// }

package main

import (
    "fmt"
    "sync"
)

var counter int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
    defer wg.Done()
    mu.Lock()
    counter = counter + 1
    mu.Unlock()
}

func main() {
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment()
    }

    wg.Wait() // Wait till all goroutines complete

    fmt.Println("Final counter:", counter)
}
