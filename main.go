package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {

	// Parse
	a := []string{}
	num := ""
	for i := 0; i < len(expression); i++ {
		if unicode.IsDigit(rune(expression[i])) {
			num += string(expression[i])
		} else {
			if num != "" {
				a = append(a, num)
				num = ""
			}
			switch expression[i] {
			case '(':
				a = append(a, "(")
			case ')':
				a = append(a, ")")
			case '+':
				a = append(a, "+")
			case '-':
				a = append(a, "-")
			case '/':
				a = append(a, "/")
			case '*':
				a = append(a, "*")
			case ' ':
			}
		}
	}
	if num != "" {
		a = append(a, num)
	}
	// end of Parse

	for i := 0; i < len(a); i++ {
		switch a[i] {
		case "+":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a[i+1] = fmt.Sprintf("%f", b+c)
		case "-":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a[i+1] = fmt.Sprintf("%f", b-c)
		case "/":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a[i+1] = fmt.Sprintf("%f", b/c)
		case "*":
			b, _ := strconv.ParseFloat(a[i-1], 64)
			c, _ := strconv.ParseFloat(a[i+1], 64)
			a[i+1] = fmt.Sprintf("%f", b*c)
		}
	}

	return strconv.ParseFloat(a[len(a)-1], 64)
}

func main() {
	fmt.Println(Calc("34 + 5"))
	fmt.Println(Calc("2*2+2"))
	fmt.Println(Calc("(2+2)*2"))
}
