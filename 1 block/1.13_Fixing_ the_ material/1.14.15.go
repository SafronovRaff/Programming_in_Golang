package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Из натурального числа удалить заданную цифру.
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.
Sample Input:
38012732
3
Sample Output:
801272
*/

func main() {
	var num, sample string
	fmt.Scanln(&num)
	fmt.Scanln(&sample)
	fmt.Println(strings.ReplaceAll(num, sample, ""))

	// конвертирование string в int
	n, _ := strconv.Atoi(num)
	// fmt.Printf("%T \n %v", n, n) проверка типа
	s, _ := strconv.Atoi(sample)
	// fmt.Printf("%T \n %v", s, s)

	Ver2(n, s)
}

func Ver2(n int, s int) {
	var arr []int
	counter := 0
	for ; n > 0; n /= 10 {
		if n%10 != s {
			arr = append(arr, n%10)
			counter++
		}
	}
	for i := counter - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}

}
