package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
								Анонимная функция

Ранее, объявляя функцию, мы давали этой функции имя. Такое объявление можно сделать только за пределами других функций (на уровне пакета).
Объявляя функцию на уровне пакета мы имеем возможность написать для такой функции необходимые тесты и удобно изменять ее реализацию.
В этом нам в т.ч. помогает определенный уровень изоляции области видимости такой функции.

Однако в ряде случаев нам необходимо выполнить определенную задачу на месте, возможно предоставив функции доступ к области видимости вызывающей функции,
как быть в этом случае? Язык Go позволяет нам использовать анонимные функции в любом выражении.
Литерал такой функции записывается как объявление функции, но без имени после ключевого слова func.

Давайте немного изменим пример с функцией Map из пакета strings:
*/

func main2() {
	src := "aBcDeFg"
	test := "AbCdEfG"
	// Обратите внимание, что скобки после имени функции используются только при ее вызове
	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, src)

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)

	// Output:
	// Инвертированная строка: AbCdEfG. Результат: true.
}

/*
В этом примере мы передали функции Map в качестве аргумента анонимную функцию, общий результат работы от этого не изменился.

Анонимные функции могут быть объявлены в другой функции, присвоены переменной или вызваны на месте:
*/
func main() {
	// Присваиваем переменной значение анонимной функции
	fn := func(a, b int) int { return a + b }

	// Выполняем анонимную функцию на месте
	// Обратите внимание на использование скобок при вызове функции
	func(a, b int) {
		fmt.Println(a + b)
	}(12, 34)

	fmt.Println(fn(17, 15))

	// Output:
	// 46
	// 32

	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}

/*
В примере мы присвоили переменной fn функцию вида func(int, int) int, затем выполнили другую анонимную функцию, а затем выполнили функцию,
присвоенную переменной fn.
	Обратите внимание на использование скобок в примерах - вызов функции требует наличия скобок,
в которых указываются передаваемые функции аргументы (если аргументы не передаются - скобки пустые).
*/
