package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

/*
Пришло время для задач, где вы сможете применить полученные знания на практике.

Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(), которая возвращает 3 значения типа пустой интерфейс.
Эта функция использует пакеты encoding/json, fmt, и os - не удаляйте их из импорта. Скорее всего, вам понадобится пакет "fmt",
но вы можете использовать любой другой пакет для записи в стандартный вывод не удаляя fmt.

Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64.
Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная математическая операция).
Но такие идеальные случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем значении может не относится к одной из указанных математических операций.

Результат выполнения программы должен быть таков:

в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции с точностью до 4 цифры после запятой (fmt.Printf(%.4f));
если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке вида: value=полученное_значение: тип_значения (например: value=true: bool)
если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям, сообщение об ошибке должно иметь вид: неизвестная операция
Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения увидели,
что оно содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных аргументов значения уже не имеет,
а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.

Удачи!
*/
func main() {
	fmt.Print(Operation(readTask()))
}

func Operation(value1, value2, operation interface{}) interface{} {
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

func readTask() (interface{}, interface{}, interface{}) {

	return 6.5, 4.7, "+"

}
