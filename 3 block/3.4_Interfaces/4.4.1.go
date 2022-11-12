package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	fmt.Print(doMathOperation(readTask()))
}

func doMathOperation(value1, value2, operation interface{}) interface{} {
	if _, ok := value1.(float64); !ok {
		return fmt.Sprintf("value=%v: %T", value1, value1)
	}
	if _, ok := value2.(float64); !ok {
		return fmt.Sprintf("value=%v: %T", value2, value2)
	}
	switch operation {
	case "+":
		return fmt.Sprintf("%.4f", value1.(float64)+value2.(float64))
	case "-":
		return fmt.Sprintf("%.4f", value1.(float64)-value2.(float64))
	case "*":
		return fmt.Sprintf("%.4f", value1.(float64)*value2.(float64))
	case "/":
		return fmt.Sprintf("%.4f", value1.(float64)/value2.(float64))
	default:
		return fmt.Sprint("неизвестная операция")
	}
}
