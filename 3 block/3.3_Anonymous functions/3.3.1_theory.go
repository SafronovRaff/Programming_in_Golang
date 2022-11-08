package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
									Анонимные функции

	# 	Функции как объекты первого класса

В Go функции являются объектами первого класса, это значит, что в этом языке программирования функцию можно передать в качестве аргумента другой функции или же вернуть функцию в качестве значения.
Подробнее об этом можно почитать в Википедии здесь https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%8A%D0%B5%D0%BA%D1%82_%D0%BF%D0%B5%D1%80%D0%B2%D0%BE%D0%B3%D0%BE_%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B0 и
https://ru.wikipedia.org/wiki/%D0%A4%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D0%B8_%D0%BF%D0%B5%D1%80%D0%B2%D0%BE%D0%B3%D0%BE_%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B0 здесь.

Рассмотрим передачу функции в качестве аргумента другой функции на примере Map из уже знакомого нам пакета strings.
Эта функция выглядит так:
func Map(mapping func(rune) rune, s string) string

Функция Map в качестве первого аргумента получает функцию вида func (rune) rune,
производящей какие-то действия с символом Unicode и возвращающей в качестве результата символ Unicode.
Из описания этой функции следует, что переданная в качестве аргумента функция будет применена к каждому символу строки,
переданной в качестве второго аргумента функции Map, получившаяся строка будет возвращена в качестве результата.

	Создадим такую функцию:
*/
func invert(r rune) rune {
	// Если буква строчная, то она возвращается заглавной
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	// Иначе возвращается строчной
	return unicode.ToLower(r)

}

// А теперь используем ее:
func main() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	// Обратите внимание, что скобки после имени функции используются только при ее вызове
	src = strings.Map(invert, src)

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)

	// Output:
	// Инвертированная строка: AbCdEfG. Результат: true.
}

//Аналогично мы можем вернуть функцию в качестве значения:

func returnFunction() func(rune) rune {
	return invert
}
