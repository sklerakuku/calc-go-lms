package calculation

import (
	"errors"
	"math"
	"strconv"
)

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Calc(expression string) (float64, error) {

	var operandStack []float64
	var operatorStack []rune

	priority := map[rune]int{ // Приоритет операторов
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	calculate := func(a, b float64, op rune) (float64, error) {
		switch op {
		case '+':
			return a + b, nil
		case '-':
			return a - b, nil
		case '*':
			return a * b, nil
		case '/':
			if b == 0 {
				return 0, errors.New("division by zero")
			}
			return a / b, nil
		default:
			return 0, errors.New("unsupported operator")
		}
	}

	for i := 0; i < len(expression); {

		//удаление пробелов теперь в application cmd ver.
		if expression[i] == ' ' {
			i++
			continue
		}

		if expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.' {
			numStr := ""
			for ; i < len(expression); i++ {
				if expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.' {
					numStr += string(expression[i])
				} else {
					break
				}
			}
			num, err := parseFloat(numStr)
			if err != nil {
				return 0, err
			}
			operandStack = append(operandStack, num)
			continue
		}

		if expression[i] == '(' {
			operatorStack = append(operatorStack, rune(expression[i]))
			i++
			continue
		}

		if expression[i] == ')' {
			for operatorStack[len(operatorStack)-1] != '(' {
				if len(operatorStack) == 0 || len(operandStack) < 2 {
					return 0, errors.New("invalid expression")
				}
				op := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if len(operandStack) < 2 {
					return 0, errors.New("invalid expression")
				}
				b := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				a := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				result, err := calculate(a, b, op)
				if err != nil {
					return 0, err
				}
				operandStack = append(operandStack, result)
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
			i++
			continue
		}

		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' && priority[rune(expression[i])] <= priority[operatorStack[len(operatorStack)-1]] {
				op := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if len(operandStack) < 2 {
					return 0, errors.New("invalid expression")
				}
				b := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				a := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				result, err := calculate(a, b, op)
				if err != nil {
					return 0, err
				}
				operandStack = append(operandStack, result)
			}
			operatorStack = append(operatorStack, rune(expression[i]))
			i++
		}
	}

	for len(operatorStack) > 0 {
		if len(operatorStack) == 0 || len(operandStack) < 2 {
			return 0, errors.New("invalid expression")
		}
		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		if len(operandStack) < 2 {
			return 0, errors.New("invalid expression")
		}
		b := operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-1]
		a := operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-1]
		result, err := calculate(a, b, op)
		if err != nil {
			return 0, err
		}
		operandStack = append(operandStack, result)
	}

	if len(operandStack) != 1 {
		return 0, errors.New("invalid expression")
	}

	if len(operandStack) > 0 {
		roundedResult := roundFloat(operandStack[0], 2)
		return roundedResult, nil
	} else {
		return 0, errors.New("invalid expression")
	}
}
