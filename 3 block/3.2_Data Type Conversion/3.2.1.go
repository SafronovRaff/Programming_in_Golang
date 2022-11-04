package main

import "fmt"

/*
Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.

Считывать и выводить ничего не нужно!

Считайте что функция main и пакет main уже объявлены!

Sample Input:

Sample Output:
*/
func main() {
	var num int64 = 12
	fmt.Printf("num = %v %T  ", convert(num), convert(num))
}

func convert(x int64) uint16 {
	return uint16(x)

}
