package main

import (
	"fmt"
)

/* ВАЖНО!!!
Необходимо запомнить, что в качестве второго значения range возвращает копию элемента массива,
это может быть важно, если в цикле мы хотим изменить массив.
В этом случае мы должны обращаться к элементам массива по индексу:
*/

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for _, elem := range a {
		elem = 100
		fmt.Println(elem)
		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [1 2 3 4 5]

	for idx := range a {
		a[idx] = 100
		fmt.Println(a[idx])
		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [100 100 100 100 100]
}

/*
Представьте, что существует способ преодолеть преграды пространства и времени,
что позволяет объединять миры и путешествовать в мгновение ока? Использование стандартной библиотеки Go в совокупности
с некоторой изобретательностью позволяет функции hyperspace из Листинга 4 модифицировать срез worlds,
убрав отступы между мирами разных планет.

import (
"fmt"
"strings"
)

// hyperspace убирает отступы между мирами планет
func hyperspace(worlds []string) { // Данный аргумент является срезом, а не массивом
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Венера   ", "Земля  ", " Марс"} // Планеты, разделенные друг от друга пробелами
	hyperspace(planets)

	fmt.Println(strings.Join(planets, "")) // Выводит: ВенераЗемляМарс
}
*/