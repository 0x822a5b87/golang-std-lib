package _func

import (
	"errors"
	"fmt"
)

var (
	IntegerDividedByZeroError = errors.New("integer divide by zero")
)

func Math(left, right int64, operator rune) (int64, error) {
	switch operator {
	case '+':
		return left + right, nil
	case '-':
		return left - right, nil
	case '*':
		return left * right, nil
	case '/':
		if right == 0 {
			return 0, IntegerDividedByZeroError
		} else {
			return left / right, nil
		}
	default:
		panic(fmt.Sprint("not supported operator : ", operator))
	}
}
