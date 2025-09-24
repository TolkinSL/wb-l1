package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d: увеличивает число %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите параметр количество воркеров")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка в преобразовании в число")
		return
	}

	jobs := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs)
	}

	for i := 1; ; i++ {
		jobs <- i
		time.Sleep(500 * time.Millisecond)
	}
}
