package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: завершает работу\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Worker %d: увеличивает число %d\n", id, job)
		}
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Все Worker завершены!")
}
