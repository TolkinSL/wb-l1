package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]string
}

func (s *SafeMap) Write(key int, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Read(key int) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {
	safeMap := SafeMap{m: make(map[int]string)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			safeMap.Write(n, fmt.Sprintf("Значение - %d", n))
		}(i)
	}

	wg.Wait()

	for k, v := range safeMap.m {
		fmt.Printf("%d: %s\n", k, v)
	}
}
