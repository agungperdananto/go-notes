package main

import (
	"errors"
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func getMultiValue(x int) (low, high int) {
	low = x - 2
	high = x + 2
	return
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't take 0 value")
	}
	return dividend / divisor, nil
}

func main() {

	fmt.Println("=============================================================================")

	fmt.Println(concat("こんにちは", "お元気ですか"))

	fmt.Println("=============================================================================")

	fmt.Println(getMultiValue(10))

	fmt.Println("=============================================================================")
	res, errors := divide(10, 3)
	fmt.Println(res)
	fmt.Println(errors)
}
