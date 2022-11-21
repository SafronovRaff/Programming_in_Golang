package main

import (
	"fmt"
	"time"
)

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
		buf = strings.TrimSpace(buf)
	*/

	firstTime, err := time.Parse(RFC3339, buf)
	if err != nil {
		fmt.Println("I can't convert the data")
	}

	fmt.Print(firstTime.Format(UnixDate))
}
