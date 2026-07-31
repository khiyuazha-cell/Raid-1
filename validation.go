package main

// ValidateInput проверяет три условия:
// 1. Пустое расписание (нет строк или первая строка пуста)
// 2. Строки разной длины
// 3. Недопустимые символы (не 0, 1 или 2)
//
// Возвращает ошибку в виде строки. Если всё ОК — возвращает пустую строку.
func ValidateInput(rows []string) string {
	// Проверка 1: пустое расписание
	if len(rows) == 0 {
		return "error: empty schedule"
	}

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

	// Если всё прошло проверку
	return ""
}
