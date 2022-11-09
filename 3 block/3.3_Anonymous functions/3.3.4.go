package main

import (
	"fmt"
	"strconv"
)

/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
*/

func main() {

	fn := func(num uint) uint {
		numString := strconv.FormatUint(uint64(num), 10)
		t := make([]rune, 0, len(numString))

		for _, i := range numString {
			if i%2 == 0 && i != 48 { // 0 - 48 эемент rune
				t = append(t, i)
			}
		}
		numString = string(t)

		conclusion, _ := strconv.Atoi(numString)

		if conclusion == 0 {
			conclusion = 100
		}

		return uint(conclusion)

	}

	fmt.Println(fn(404040)) // 444
}
