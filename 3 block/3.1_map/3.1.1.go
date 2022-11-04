package main

import "fmt"

/*нутри функции main (объявлять функцию не нужно) необходимо написать программу:

На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться). Для чтения из стандартного ввода импортирован пакет fmt.

Вам необходимо вычислить результат выполнения функции work для каждого из полученных чисел. Функция work имеет следующий вид:

func work(x int) int
Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.

Однако работа функции work занимает слишком много времени. Выделенного вам времени выполнения не хватит на последовательную обработку каждого числа,
поэтому необходимо реализовать кэширование уже готовых результатов и использовать их в работе.

После завершения работы программы результат выполнения будет дополнен информацией о соблюдении установленного лимита времени выполнения.

Sample Input:

3 1 5 2 3 5 3 0 3 4
Sample Output:

2 0 6 1 2 6 2 -1 2 5 time limit ok*/

func work(x int) int {
	return x - 1
}

func main() {
	m := map[int]int{}

	var num int
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)

		if _, inMap := m[num]; !inMap {
			m[num] = work(num)
		}

		num = m[num]
		fmt.Printf("%d ", num)
	}

	/*_, inMap := m[num]
	if !inMap {
		m[num] = work(num)
		fmt.Print(m[num], " ")
	} else {
		fmt.Print(m[num], " ")
	}*/

}
