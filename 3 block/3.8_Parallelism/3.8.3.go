package main

import (
	"fmt"
	"time"
)

/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов.
В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)
Функция должна называться removeDuplicates()
Выводить или вводить ничего не нужно!
*/
func removeDuplicates(inputStream, outputStream chan string) {
	var p string

	for v := range inputStream {
		if p != v {
			outputStream <- v
			p = v
		}

	}
	close(outputStream)
}

func chanPrint(c chan string) {
	for {
		reception := <-c
		fmt.Println(reception)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go chanPrint(outputStream)

	for _, v := range "11111cc222pp33344hh56" {
		inputStream <- string(v)
	}

}
