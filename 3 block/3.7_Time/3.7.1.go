package main

import (
	"fmt"
	"time"
)

/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:
1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
Wed Apr 16 05:20:00 +0600 1986
Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.
Sample Input:

1986-04-16T05:20:00+06:00
Sample Output:

Wed Apr 16 05:20:00 +0600 1986
*/
const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"

	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func main() {
	var buf string
	_, err := fmt.Scanln(&buf)
	if err != nil {
		fmt.Println("I can't read the data")
	}

	/*
		buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("не могу считать данные")
		}
	*/

	firstTime, err := time.Parse(RFC3339, buf)
	if err != nil {
		fmt.Println("I can't convert the data")
	}

	fmt.Print(firstTime.Format(UnixDate))
}
