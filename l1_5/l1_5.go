package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите параметр количество секунд")
		return
	}
	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка в преобразовании в число")
		return
	}

	jobs := make(chan int)

	go func() {
		for i := 1; ; i++ {
			jobs <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	timeout := time.After(time.Duration(seconds) * time.Second)

	for {
		select {
		case v := <-jobs:
			fmt.Println("Получено из канала:", v)
		case <-timeout:
			fmt.Printf("Программа завершена через %d секунд!", seconds)
			return
		}
	}
}
