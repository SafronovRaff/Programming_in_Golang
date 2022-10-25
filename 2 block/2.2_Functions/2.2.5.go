package main

import "fmt"

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)
Результат: 2, 1
*/
func main() {

	a, b := sumInt(2, 1, 1, 1, 1, 1, 10)
	fmt.Println(a, b)
}

func sumInt(num ...int) (int, int) {
	a, b := 0, 0
	for i, n := range num {
		a = i + 1
		b += n
	}
	return a, b
}
