package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	var num1, num2 float64
	var err error
	num1, err = checkValue(value1)
	if err != nil {
		panic(err)
	}

	num2, err = checkValue(value2)
	if err != nil {
		panic(err)
	}

	oper, err := checkOper(operation)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%.4f\n", oper(num1, num2))
}

func checkValue(value interface{}) (float64, error) {
	if val, ok := value.(float64); ok {
		return val, nil
	}
	return 0, errors.New(fmt.Sprintf("value=%v: %T", value, value))

}

func checkOper(operation interface{}) (fn func(a, b float64) float64, err error) {
	v, ok := operation.(string)
	if !ok {
		return nil, errors.New("неизвестная операция")
	}
	switch v {
	case "+":
		fn = func(a, b float64) float64 {
			return a + b
		}
	case "-":
		fn = func(a, b float64) float64 {
			return a - b
		}
	case "*":
		fn = func(a, b float64) float64 {
			return a * b
		}

	case "/":
		fn = func(a, b float64) float64 {
			return a / b
		}
	default:
		err = errors.New("неизвестная операция")

	}
	return

}
