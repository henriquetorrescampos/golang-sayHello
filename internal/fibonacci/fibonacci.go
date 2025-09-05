package fibonnaci

import (
	"fmt"
)

func Fibonnaci(num int) (int, error) {
	if num <= 0 {
		fmt.Println("Digite um número positivo.")
	}

	if num == 0 {
		return 0, fmt.Errorf("Fibonnaci de 0 é 0.")
	}

	if num == 2 {
		return 1, fmt.Errorf("Fibonnaci de 2 é 1.")
	}

	// fib1, fib2 := 0, 1
	a, b := 0, 1

	for i := 2; i <= num; i++ {
		a, b = b, a+b
		// nextFib := fib1 + fib2
		// fib1 = fib2
		// fib2 = nextFib
	}

	return b, nil
}
