package main

import "fmt"

/*
											Замыкания

- это такие функции, которые вы можете создавать в рантайме и им будет доступно текущее окружение,
в рамках которого они были созданы. Другими словами, функции, определенные как замыкания, "запоминают" окружение, в котором они были созданы.

Перед тем как мы начнем разбираться с замыканиями, давайте определимся с функциями и анонимными функциями.

										Анонимные функции

Функции, у которых есть имя - это именованные функции. Функции, которые могут быть созданы без указания имени - это анонимные функции.
	Все просто :) Как можно видет в приведенном ниже коде,
можно создать анонимную функцию и непосредственно вызвать или можно присвоить функцию некоторой переменной и вызвать с указанием этой переменной.
	В общем случае, для замыканий используются анонимные функции. В Go у вас есть возможность создать анонимную функцию и передать ее как параметр в другую функцию,
таким образом мы используем функции высшего порядка.

*/

// В printfunc сохраняется анонимная функция, которая затем вызывается.
func printMessage(message string) {
	fmt.Println(message)
}

// Функция getPrintMessage создает анонимную функцию и возвращает ее.
func getPrintMessage() func(string) {
	// Возвращаем анонимную функцию
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	// Именованная функция
	printMessage("Hello function!")

	// Анонимная фукция объявляется и вызывается
	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	// Получаем анонимную функцию и вызываем ее
	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller!")

	// Output:
	// Hello function!
	// Hello anonymous function!
	//Hello anonymous function using caller!

}
