package main

import "fmt"

func main() {
	// Тест 1: валидный ввод
	valid := []string{"0102", "1100", "0010", "2001"}
	err := ValidateInput(valid)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("✓ Валидация прошла успешно")
	}

	// Тест 2: пустой ввод
	empty := []string{}
	fmt.Println(ValidateInput(empty))

	// Тест 3: разные длины
	ragged := []string{"01", "0102"}
	fmt.Println(ValidateInput(ragged))

	// Тест 4: недопустимый символ
	invalid := []string{"0102", "1X00"}
	fmt.Println(ValidateInput(invalid))
}
