/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные

На вход программе подается целое число d (0 < d < 360).

Выходные данные

Выведите на экран фразу:

It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

Sample Input:

90
Sample Output:

It is 3 hours 0 minutes.
*/

package main

import "fmt"

func main() {

	var d, h, m int
	fmt.Scan(&d)
	h = d / 30
	m = (d % 30) * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}
