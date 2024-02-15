package main

import (
	"bufio"
	"os"
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
func (c Calculator) validateRimResult(result *int) {
	if *result <= 0 || *result > 10 {
		panic("Результат<0 или >10.")
	}
}
func (c Calculator) Calculate(values []string) int {
	if len(values) != 3 {
		panic("Выражение не соответствует правилам.")
	} else {
		operation := c.operationFunc(values[1])
		v1, ok1 := arabs[values[0]]
		v2, ok2 := arabs[values[2]]
		if ok1 == ok2 && ok1 == true {
			// арабский режим
			return operation(v1, v2)
		} else {
			v1, ok1 := rimskie[values[0]]
			v2, ok2 := rimskie[values[2]]
			if ok1 == ok2 {
				// римский режим
				result := operation(v1, v2)
				c.validateRimResult(&result)
				return result
			} else {
				panic("Числа разные или не соответствуют.")
			}
		}
	}
}

func main() {
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

func InputFromCLI() string {
	reader := bufio.NewReader(os.Stdin)
	promt, _ := reader.ReadString('\n')
	return strings.TrimRight(promt, "\n")
}
