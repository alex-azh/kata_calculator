package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var rimskie = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}
var arabs = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}
var arabToRim = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
}

type Calculator struct{}

func (c Calculator) operationFunc(op string) func(int, int) int {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "-":
		return func(a, b int) int {
			return a - b
		}
	case "/":
		return func(a, b int) int {
			return a / b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		panic("Операция не определена!")
	}
}
func (c Calculator) rimResultToArabNumber(result int) string {
	if result <= 10 {
		return arabToRim[result]
	} else {
		razryad1 := result % 10
		var part1 string = arabToRim[razryad1]
		if razryad1 == 9 {
			desyatok, _ := arabToRim[(result-razryad1)+1]
			return part1 + desyatok
		} else {
			desyatok, _ := arabToRim[result-razryad1]
			return desyatok + part1
		}
	}
}
func (c Calculator) Calculate(values []string) string {
	if len(values) != 3 {
		panic("Выражение не соответствует правилам.")
	} else {
		operation := c.operationFunc(values[1])
		v1, ok1 := arabs[values[0]]
		v2, ok2 := arabs[values[2]]
		if ok1 == ok2 && ok1 == true {
			// арабский режим
			return strconv.Itoa(operation(v1, v2))
		} else {
			v1, ok1 := rimskie[values[0]]
			v2, ok2 := rimskie[values[2]]
			if ok1 == ok2 {
				// римский режим
				if v1 <= 0 || v2 <= 0 || v1 > 10 || v2 > 10 {
					panic("Числа от 0 до 10 должно быть.")
				}
				result := operation(v1, v2)
				if result <= 0 {
					panic("Римский ответ не может быть меньше 1.")
				}
				return c.rimResultToArabNumber(result)
			} else {
				panic("Числа разные или не соответствуют.")
			}
		}
	}
}

func InputFromCLI() string {
	reader := bufio.NewReader(os.Stdin)
	promt, _ := reader.ReadString('\n')
	return strings.TrimRight(promt, "\n")
}

func main() {
	// хотел убрать это сравнение ненужное, а оно даже и неправильное по текущей логике, но не стал исправлять.
	var promt string
	for promt != "q" {
		line := InputFromCLI()
		if line == "q" {
			break
		}
		calculator := Calculator{}
		values := strings.Split(line, " ")
		result := calculator.Calculate(values)
		println(result)
	}
}
