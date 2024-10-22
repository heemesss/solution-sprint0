package main

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

func CalcSimple(a []string) (float64, error) {
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case "/":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b/c)
			i -= 2
		case "*":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b*c)
			i -= 2
		}
	}

	for i := 0; i < len(a); i++ {
		switch a[i] {
		case "+":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b+c)
			i -= 2
		case "-":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b-c)
			i -= 2
		}
	}
	return strconv.ParseFloat(a[len(a)-1], 64)
}

func ParserForCalc(expression string) []string {
	answer := []string{}
	num := ""
	for i := 0; i < len(expression); i++ {
		if unicode.IsDigit(rune(expression[i])) {
			num += string(expression[i])
		} else {
			if num != "" {
				answer = append(answer, num)
				num = ""
			}
			switch expression[i] {
			case '(':
				answer = append(answer, "(")
			case ')':
				answer = append(answer, ")")
			case '+':
				answer = append(answer, "+")
			case '-':
				answer = append(answer, "-")
			case '/':
				answer = append(answer, "/")
			case '*':
				answer = append(answer, "*")
			default:
			}
		}
	}
	if num != "" {
		answer = append(answer, num)
	}
	return answer
}

func IndexRight(str []string, s string) int {
	for i := len(str) - 1; i > -1; i-- {
		if str[i] == s {
			return i
		}
	}
	return -1
}

func Calc(expression string) (float64, error) {

	// Parse
	a := ParserForCalc(expression)
	// end of Parse

	for i := range a {
		fmt.Println(a[i])
	}
	// ToDo
	for i := slices.Index(a, ")"); i != -1; i = slices.Index(a, ")") {
		fmt.Println("-------", i, "----------")

		j := IndexRight(a[:i], "(")
		fmt.Println("-------", j, "----------")

		num, _ := CalcSimple(a[j+1 : i])
		fmt.Println(a[j:i], num)
		a = append(append(a[:j], fmt.Sprintf("%f", num)), a[i+1:]...)
		// Calc
	}

	return CalcSimple(a)

}

func main() {
	//fmt.Println(Calc("3 * 2 + 2 * 2 / 2"))
	//fmt.Println(Calc("555 + 1 * 2"))
	//fmt.Println(Calc("(12 + 24 * 3 * (45 - 44))"))
}
