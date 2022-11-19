package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

/*
Поэтапный поиск данных
Данная задача в основном ориентирована на изучение типа bufio.Reader,
поскольку этот тип позволяет считывать данные постепенно.

В тестовом файле, который вы можете скачать из нашего репозитория на github.com,
содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/

func main() {

	urlDownload := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_files/task_sep_1/task.data"
	fiel, err := http.Get(urlDownload)
	if err != nil {
		fmt.Println(err)
		return
	}

	rd := bufio.NewReader(fiel.Body)

	i := 0
	for {
		s, err := rd.ReadString(';')
		if err == io.EOF {
			break
		}

		if s == "0;" {
			fmt.Println(i + 1)
		}
		i++
	}

	forCSV()
}

func forCSV() {
	urlDownload := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_files/task_sep_1/task.data"
	f, err := http.Get(urlDownload)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := csv.NewReader(f.Body)
	reader.Comma = ';'
	lin, err := reader.ReadAll()

	for i, elem := range lin[0] {
		if elem == "0" {
			fmt.Println(i + 1)
		}
	}

}
