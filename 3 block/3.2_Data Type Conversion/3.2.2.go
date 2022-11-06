package main

import (
	"strconv"
	"strings"
	"unicode"
)

func main() {
	adding("%^80", "hhhhh20&&&&nd")
}

/*
Представьте что вы работаете в большой компании где используется модульная архитектура.
Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.

Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки.
Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего импортировать не нужно!

Считывать и выводить ничего не нужно!
Ваша функция должна называться adding() !
Считайте что функция и пакет main уже объявлены!

Sample Input:
%^80 hhhhh20&&&&nd
Sample Output:
100
*/
func replacement(text string) int64 {
	str := strings.Builder{}
	for _, i := range text {
		if unicode.Is(unicode.Digit, i) {
			str.WriteRune(i)
		}
	}
	num, err := strconv.ParseInt(str.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	//fmt.Println(num)
	//fmt.Printf("%T \n", num)
	return num
}
func adding(string1, string2 string) int64 {
	sum := replacement(string1) + replacement(string2)
	//fmt.Println(sum)
	//fmt.Printf("%T \n", sum)
	return sum
}
