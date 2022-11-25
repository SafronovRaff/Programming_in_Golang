package main

import (
	"fmt"
	"sync"
)

/*
Внутри функции main (функцию объявлять не нужно),
вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
*/

func main() {

	wg := new(sync.WaitGroup) // определяем группу в виде переменной

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			work2()
		}()
	}
	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")

}
func work2() {

	fmt.Println("Горутина начала выполнение")

}