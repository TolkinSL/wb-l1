package main

import (
	"fmt"
	"runtime"
	"time"
)

func byGoexit() {
	go func() {
		defer fmt.Println("Дефер выполнится перед Goexit()")
		fmt.Println(" - Работает")
		time.Sleep(500 * time.Millisecond)
		runtime.Goexit()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("\nГлавная горутина завершена")
}

func main() {
	fmt.Println("Горутины - Остановка через runtime.Goexit()")
	byGoexit()
}
