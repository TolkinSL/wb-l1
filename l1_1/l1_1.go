package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) hello() string {
	return fmt.Sprintf("Привет %s!", h.Name)
}

type Action struct {
	Human
	Role string
}

func main() {
	a := Action{
		Human: Human{
			Name: "Петя",
			Age:  30,
		},
		Role: "Admin",
	}

	fmt.Println(a.hello())
	fmt.Printf("Возраст %d лет\n", a.Age)
	fmt.Printf("%s имеет аккаунт с правами: %s\n", a.Name, a.Role)
}
