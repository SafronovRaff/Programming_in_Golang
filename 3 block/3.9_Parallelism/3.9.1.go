package main

import (
	"log"
	"time"
)

/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
Функция work() ничего не принимает и не возвращает.
*/
var d string

func main() {

	done := make(chan struct{})

	go func() {
		work()
		close(done)
	}()
	<-done
	//

	if d != "nil" {
		log.Fatal("Error")
	}
	log.Print("All right")
}

func work() {
	time.Sleep(3 * time.Second)
	d = "nil"
}
