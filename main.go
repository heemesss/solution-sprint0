package main

import (
	"fmt"
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
		fmt.Println(a[i])
	}
	return 0, nil

}

func main() {
	fmt.Println(Calc("2+2*222222"))
	fmt.Println(Calc("2*2+2"))
	fmt.Println(Calc("(2+2)*2"))
}
