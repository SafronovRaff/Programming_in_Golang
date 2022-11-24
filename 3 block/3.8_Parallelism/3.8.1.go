package main

import "fmt"

/*
Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().

Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
*/
func main() {
	var num int
	fmt.Println("ENTER A NUMBER")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("input error")
	}

	c := make(chan int)
	go task(c, num)

	fmt.Println(<-c)
}

func task(c chan int, a int) {
	c <- a + 1
}
