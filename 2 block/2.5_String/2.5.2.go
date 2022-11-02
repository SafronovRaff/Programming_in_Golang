package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.
Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
Sample Input:
топот
Sample Output:
Палиндром
*/
func main() {
	tex, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	tex = strings.Trim(tex, "\n \r")
	t := []rune(tex)
	var revers []rune
	for i := len(t) - 1; i >= 0; i-- {
		revers = append(revers, t[i])
	}

	//fmt.Printf("%s", revers)

	// Работает, https://pkg.go.dev/reflect@master#DeepEqual DeepEqual сообщает, являются ли x и y «глубоко равными»,
	/*
		if reflect.DeepEqual(t, revers) == true {
					fmt.Println("Палиндром")
				} else {
					fmt.Println("Нет")
				}
	*/

	if string(t) == string(revers) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

	ver2(tex)
}

func ver2(tex string) {
	t := []rune(tex)
	mid := len(t) / 2
	for i := 0; i <= mid; i++ {
		if t[i] != t[len(t)-1-i] {
			fmt.Println("нет")
			return

		}
	}
	fmt.Println("Палиндром")
}
