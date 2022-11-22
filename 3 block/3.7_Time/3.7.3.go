package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s
*/
const (
	T = "02.01.2006 15:04:05"
)

func main() {
	var res time.Duration // String returns a string representing the duration in the form "72h3m0.5s".
	// Leading zero units are omitted. As a special case, durations less than one second format use a smaller unit (milli-, micro-, or nanoseconds)
	//to ensure that the leading digit is non-zero. The zero duration formats as 0s.

	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("I can't read the data")
	}
	/*
		// на stepic только через io.EOF  или ioutil.ReadAll
		if err != nil && err != io.EOF {
			panic(err)
		}  */

	splitedBuf := strings.Split(buf, ",") // splitting the string into 2 parts

	splitedBuf[0], splitedBuf[1] = strings.TrimSpace(splitedBuf[0]), strings.TrimSpace(splitedBuf[1]) // removing carriage transfer for further parsing

	firstTime, err := time.Parse(T, splitedBuf[0])
	if err != nil {
		fmt.Println("I can't convert the first value to time")
	}
	secondTime, err := time.Parse(T, splitedBuf[1])
	if err != nil {
		fmt.Println("I can't convert the second value to time")
	}

	if firstTime.After(secondTime) { // true if later
		res = time.Since(secondTime).Round(time.Second) - time.Since(firstTime).Round(time.Second)
		//Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).

		fmt.Print(res.Round(time.Second))
	} else {
		res = time.Since(firstTime).Round(time.Second) - time.Since(secondTime).Round(time.Second)
		fmt.Print(res.Round(time.Second))

	}

	v2()
}
func v2() {
	splitedBuf, err := csv.NewReader(os.Stdin).Read() //csv format by default, the column separator is the symbol ","
	if err != nil {
		fmt.Println("I can't read the data")
	}
	firstTime, err := time.Parse(T, splitedBuf[0])
	if err != nil {
		fmt.Println("I can't convert the first value to time")
	}
	secondTime, err := time.Parse(T, splitedBuf[1])
	if err != nil {
		fmt.Println("I can't convert the second value to time")
	}

	fmt.Println(time.Duration(math.Abs(float64(firstTime.Sub(secondTime)))).Round(time.Second)) // func Abs returns the absolute value of x.
}
