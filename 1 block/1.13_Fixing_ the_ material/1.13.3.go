package main

import (
	"fmt"
	"time"
)

/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если

k=13257=3*3600+40*60+57,

то h=3 и m=40.

Входные данные
На вход программе подается целое число k (0 < k < 86399).
Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
Sample Input:
13257
Sample Output:
It is 3 hours 40 minutes.
*/
func main() {

	var Sek, h, m int
	fmt.Scan(&Sek)
	h = Sek / 3600
	Sek %= 3600
	m = Sek / 60

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
	V2(Sek) // надо разбирать, на тестах не выводит d.Hour() "It is 0 hours 40 minutes."

}
func V2(Sek int) {

	d := time.Date(0, 0, 0, 0, 0, Sek, 0, time.UTC)

	fmt.Printf("It is %d hours %d minutes.", d.Hour(), d.Minute())

}
