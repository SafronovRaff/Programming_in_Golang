package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	T = "02.01.2006 15:04:05"
)

func main() {
	var res time.Duration

	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("I can't read the data")
	}
	// на stepic
	/*if err != nil && err != io.EOF {
		panic(err)
	}   или ioutil.ReadAll  */

	buf = strings.TrimSpace(buf)

	splitedBuf := strings.Split(buf, ",") //разбиваем строку на 2 части

	splitedBuf[0], splitedBuf[1] = strings.TrimSpace(splitedBuf[0]), strings.TrimSpace(splitedBuf[1]) // убираем перенос коретки для дальнейшего парсинга

	firstTime, err := time.Parse(T, splitedBuf[0])
	if err != nil {
		fmt.Println("I can't convert the first value to time")
	}
	secondTime, err := time.Parse(T, splitedBuf[1])
	if err != nil {
		fmt.Println("I can't convert the second value to time")
	}
	res = time.Since(firstTime) - time.Since(secondTime)

	fmt.Print(res.Round(time.Hour))

}
