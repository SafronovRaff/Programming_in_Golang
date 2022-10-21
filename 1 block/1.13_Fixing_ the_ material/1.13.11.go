package main

import "fmt"

/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).
Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy,
например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
Sample Input:
10
Sample Output:
10 korov
*/

/*
1: "корова"
от 2 до 4: "коровы"
от 5 до 20: "коров"
от 21: если число заканчивается на "1", то "корова"
если число заканчивается на 2, 3, 4, то "коровы"
НО  с 11 по 14 "коров"
*/
func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a%10 == 1 && a%100 != 11:
		fmt.Printf("%d korova", a)
	case a%10 == 2 && a%100 != 12:
		fallthrough
	case a%10 == 3 && a%100 != 13:
		fallthrough
	case a%10 == 4 && a%100 != 14:
		fmt.Printf("%d korovy", a)
	default:
		fmt.Printf("%d korov", a)

	}
}

/*
func main() {
	var a int
	fmt.Scan(&a)

	cows := [100]string{"1 korova", "2 korovy", "3 korovy", "4 korovy", "5 korov", "6 korov",
"7 korov", "8 korov", "9 korov", "10 korov", "11 korov", "12 korov", "13 korov", "14 korov",
"15 korov", "16 korov", "17 korov", "18 korov", "19 korov", "20 korov", "21 korova", "22 korovy",
"23 korovy", "24 korovy", "25 korov", "26 korov", "27 korov", "28 korov", "29 korov", "30 korov",
"31 korova", "32 korovy", "33 korovy", "34 korovy", "35 korov", "36 korov", "37 korov", "38 korov",
"39 korov", "40 korov", "41 korova", "42 korovy", "43 korovy", "44 korovy", "45 korov", "46 korov",
"47 korov", "48 korov", "49 korov", "50 korov", "51 korova", "52 korovy", "53 korovy", "54 korovy",
"55 korov", "56 korov", "57 korov", "58 korov", "59 korov", "60 korov", "61 korova", "62 korovy", "63 korovy",
"64 korovy", "65 korov", "66 korov", "67 korov", "68 korov", "69 korov", "70 korov", "71 korova", "72 korovy",
"73 korovy", "74 korovy", "75 korov", "76 korov", "77 korov", "78 korov", "79 korov", "80 korov", "81 korova",
"82 korovy", "83 korovy", "84 korovy", "85 korov", "86 korov", "87 korov", "88 korov", "89 korov", "90 korov",
"91 korova", "92 korovy", "93 korovy", "94 korovy", "95 korov", "96 korov", "97 korov", "98 korov", "99 korov", "100 korov"}

	fmt.Println(cows[a-1])

}
*/
