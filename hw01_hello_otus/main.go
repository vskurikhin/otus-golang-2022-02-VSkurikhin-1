package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// функция обёртка для тестирования.
func reverse(s string) string {
	return stringutil.Reverse(s)
}

func main() {
	// переворот строки в переменную hello
	hello := stringutil.Reverse("Hello, OTUS!")
	// вывод строки на stdout
	fmt.Println(hello)
}
