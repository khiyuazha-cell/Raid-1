package main

// ValidateInput проверяет три условия:
// 1. Пустое расписание (нет строк или первая строка пуста)
// 2. Строки разной длины
// 3. Недопустимые символы (не 0, 1 или 2)
//
// Возвращает ошибку в виде строки. Если всё ОК — возвращает пустую строку.
func ValidateInput(rows []string, busy string, important string, busyColor string, importantColor string) string {
	// Проверка 1: пустое расписание — нет строк вообще
	if len(rows) == 0 {
		return "error: empty schedule"
	}

	// Проверка 1б: первая строка пуста
	if len(rows[0]) == 0 {
		return "error: empty schedule"
	}

	// Проверка 2: все строки должны быть одной длины
	firstLen := len(rows[0])
	for _, row := range rows {
		if len(row) != firstLen {
			return "error: rows have different lengths"
		}
	}

	// Проверка 3: каждый символ должен быть только 0, 1 или 2
	for _, row := range rows {
		for _, ch := range row {
			if ch != '0' && ch != '1' && ch != '2' {
				return "error: invalid symbol (only 0/1/2 allowed)"
			}
		}
	}

	// Проверка 4: цвета не должны совпадать
	if busyColor == importantColor {
		return "error: busyColor and importantColor cannot be the same"
	}

	// Проверка 5: цвета не должны быть пустыми
	if busyColor == "" || importantColor == "" {
		return "error: busyColor and importantColor cannot be empty"
	}

	const cellWidth = 9

	if len(busy)-cellWidth > 0 || len(important)-cellWidth > 0 {
		return "error: content is wider than cell width"
	}

	// Если всё прошло проверку
	return ""
}
