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
			b, err := strconv.ParseFloat(a[i-1], 64)
			if err != nil {
				return 0, err
			}
			c, err := strconv.ParseFloat(a[i+1], 64)
			if err != nil {
				return 0, err
			}
			a = append(a[:i-1], a[i+1:]...)

			if c == 0.0 {
				return 0, fmt.Errorf("division by zero")
			}

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
			b, err := strconv.ParseFloat(a[i-1], 64)
			if err != nil {
				return 0, err
			}
			c, err := strconv.ParseFloat(a[i+1], 64)
			if err != nil {
				return 0, err
			}
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b+c)
			i -= 2
		case "-":
			b, err := strconv.ParseFloat(a[i-1], 64)
			if err != nil {
				return 0, err
			}
			c, err := strconv.ParseFloat(a[i+1], 64)
			if err != nil {
				return 0, err
			}
			a = append(a[:i-1], a[i+1:]...)

			a[i-1] = fmt.Sprintf("%f", b-c)
			i -= 2
		}
	}
	if len(a) > 1 {
		return 0, fmt.Errorf("wrong procedure")
	}
	return strconv.ParseFloat(a[len(a)-1], 64)
}

func ParserForCalc(expression string) ([]string, error) {
	answer := []string{}
	num := ""
	for i := 0; i < len(expression); i++ {
		if !unicode.IsSpace(rune(expression[i])) && !unicode.IsDigit(rune(expression[i])) && expression[i] != '(' && expression[i] != ')' &&
			expression[i] != '+' && expression[i] != '-' && expression[i] != '*' && expression[i] != '/' {
			return answer, fmt.Errorf("symbol \"%c\" is unknown", expression[i])
		}
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
	return answer, nil
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
	a, err := ParserForCalc(expression)
	if err != nil {
		return 0, err
	}
	// end of Parse

	/*for i := range a {
		fmt.Println(a[i])
	}*/
	// ToDo
	for i := slices.Index(a, ")"); i != -1; i = slices.Index(a, ")") {
		// fmt.Println("-------", i, "----------")

		j := IndexRight(a[:i], "(")
		if j == -1 || j+1 == i {
			return 0, fmt.Errorf("error with parentheses")
		}
		// fmt.Println("-------", j, "----------")

		num, _ := CalcSimple(a[j+1 : i])
		// fmt.Println(a[j:i], num)
		a = append(append(a[:j], fmt.Sprintf("%f", num)), a[i+1:]...)
		// Calc
	}

	return CalcSimple(a)

}

func main() {

	//     деление на ноль     скобки
	// fmt.Println(Calc("2 + a"))
	//fmt.Println(Calc("3 / 0"))
	//fmt.Println(Calc("3 3 0"))
	//fmt.Println(Calc("())5 + 3()))"))
	//fmt.Println(Calc("14f4"))

}
