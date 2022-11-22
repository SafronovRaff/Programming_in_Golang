package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	L = "2006-01-02 15:04:05"
)

/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
Sample Input:
2020-05-15 08:00:00
Sample Output:
2020-05-15 08:00:00
*/
func main() {
	var hourTime int
	deadline := 13 // remove with version 2

	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("I can't read the data")
	}
	buf = strings.TrimSpace(buf)

	firstTime, err := time.Parse(L, buf)
	if err != nil {
		fmt.Println("I can't convert the data")
	}

	//version 1
	hourTime = firstTime.Hour()
	if hourTime < deadline {
		fmt.Println(firstTime.Format(L))

	} else {
		fmt.Println(firstTime.Add(time.Hour * 24).Format(L))

	}

	/*
		// version 2
		deadLine := time.Date(firstTime.Year(), firstTime.Month(), firstTime.Day(), 13, 00, 0, 0, time.Local)
		if firstTime.After(deadLine) {
			deadLine = deadLine.Add(time.Hour * 24)
		}
			fmt.Println(firstTime.Format(L))

	*/
}
