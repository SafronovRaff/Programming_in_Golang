package main

import "fmt"

/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().

Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
*/
func main() {

	a := "string"
	c := make(chan string, 5)

	go task2(c, a)

	for i := 0; i < 5; i++ {
		fmt.Print(<-c)
	}

}

func task2(c chan string, a string) {
	b := " "
	for i := 0; i < 5; i++ {
		c <- a + b
	}

}
