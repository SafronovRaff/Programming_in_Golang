package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	{
		fmt.Println("не могу считать данные")
	}

	//	buf = strings.TrimSpace(buf) //strings.TrimRight(buf, "\n")

	splitedBuf := strings.Split(buf, ",") //разбиваем строку на 2 части

	splitedBuf[0], splitedBuf[1] = strings.TrimSpace(splitedBuf[0]), strings.TrimSpace(splitedBuf[1]) // убираем перенос коретки для дальнейшего парсинга

}
