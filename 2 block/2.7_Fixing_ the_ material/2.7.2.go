package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами
	(перед первой буквой и после последней символ ‘*’ добавлять не нужно).
	Входные данные
	Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.
	Выходные данные
	Вывести строку, которая получится после добавления символов '*'.
	Sample Input:
	LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
	Sample Output:
	L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
	Sample Input:
	LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
	Sample Output:
	L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/

func main() {
	var err error
	var a string
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	if utf8.RuneCountInString(a) > 1000 {
		fmt.Println("в строке более 1000 символов")
		return
	}
	b := strings.Split(a, "")         // Разбивает строку согласно разделителю
	fmt.Println(strings.Join(b, "*")) // объединяет массив строк через символ
	// короче
	fmt.Println((strings.Join(strings.Split(a, ""), "*")))

	// вариант через цикл
	c := []rune(a)

	fmt.Print(string(c[0])) // вывод с 1 индекса

	for i := 1; i < len(c); i++ {
		fmt.Print("*", string(c[i]))
	}

	// вариант через Replace
	d := strings.Replace(a, "", "*", -1)
	fmt.Println(d[1 : len(d)-1])

}
