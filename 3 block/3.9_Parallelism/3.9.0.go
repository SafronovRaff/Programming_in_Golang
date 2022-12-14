package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) { // 	3) запускаем fibonacci

	x, y := 0, 1
	for {
		select { // 4) fibonacci выполняет первый for select:
		case c <- x: //- т.к. в канале "c" ещё ничего нет то запись не заблокирована, мы можем выполнить первый case (c <- x)
			x, y = y, x+y
		case <-quit: // - т.к. в канале "quit" ещё ничего нет то чтение из него заблокировано, второй селект сделать мы не можем
			fmt.Println("quit")
			return
		}
	}
	time.Sleep(1 * time.Second)
}

//   n) анонимная рутина выполняет запись в канал "quit", анонимная рутина завершает своё выполнение
// n+1) fibonacci выполняет очередной for:
// - т.к. в канале "с" уже есть значение которое не забрали на очередном цикле анонимной рутины (т.к. for закончил своё выполнение)
//то запись в "с" заблокирована и первый case не выполняется
//- т.к. в канале "quit" есть значение, то чтение не заблокировано и выполняется второй case, идёт вывод на экран сообщения и завершение функции fibonacci через return
// main завершает своё выполнение

func main() { // 1) запускаем main
	c := make(chan int)
	quit := make(chan int)
	go func() { //2) запускаем анонимную рутину
		for i := 0; i < 10; i++ { // 5) анонимная рутина выполняет for с i == 0, и выводит из "c" значение записанное в п.4
			fmt.Println(<-c) // Пункты 4 и 5 повторяются пока i удовлетворяет условиям for
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
