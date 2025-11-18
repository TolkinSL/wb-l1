package main

import "fmt"

// any — это синоним interface{}
func myType(v any) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int, chan string, chan bool, chan any:
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Тип: unknown type")
	}
}

func main() {
	myType(7)
	myType("Go")
	myType(false)

	ch := make(chan int)
	myType(ch)

	myType(3.14)

}
