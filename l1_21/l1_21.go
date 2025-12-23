/* В этом примере структура OldInput имеет метод Scan(),
а новый код работает с интерфейсом Reader, с методом Read().

Для согласования этих интерфейсов используется структура OldInputAdapter,
которая реализует интерфейс Reader.

Плюсы паттерна Adapter:
- Старый код (OldInput) используется без изменений.
- Функция myInput работает только с интерфейсом Reader и не зависит от конкретной реализации.

Минусы паттерна Adapter:
- Для адаптации требуется дополнительная структура-обвертка.
*/

package main

import "fmt"

type Reader interface {
	Read() string
}

type OldInput struct{}

func (o *OldInput) Scan() string {
	var text string
	// fmt.Scanf("%s", &text)
	text = "Hello Go"
	return text
}

type OldInputAdapter struct {
	oldInput *OldInput
}

func (a *OldInputAdapter) Read() string {
	return a.oldInput.Scan()
}

func myInput(r Reader) {
	text := r.Read()
	fmt.Println("Вывод:", text)
}

func main() {
	oldInput := &OldInput{}
	adapter := &OldInputAdapter{oldInput: oldInput}

	myInput(adapter)
}
