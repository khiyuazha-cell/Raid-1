package main

func renderRow(row, busy string, important string, busyColor string, importantColor string) string {
	line := "|"
	const cellWidth = 9
	const redBg = "\033[41m"
	const yellowBg = "\033[43m"
	const blueBg = "\033[44m"
	const greenBg = "\033[42m"
	const reset = "\033[0m"

	switch busyColor {
	case "red":
		busyColor = redBg
	case "yellow":
		busyColor = yellowBg
	case "blue":
		busyColor = blueBg
	case "green":
		busyColor = greenBg
	default:
		busyColor = ""
	}

	switch importantColor {
	case "red":
		importantColor = redBg
	case "yellow":
		importantColor = yellowBg
	case "blue":
		importantColor = blueBg
	case "green":
		importantColor = greenBg
	default:
		importantColor = ""
	}

	for _, ch := range row {
		switch ch {
		case '0':
			line += centerCell(" ", cellWidth) + "|"
		case '1':
			line += busyColor + centerCell(busy, cellWidth) + reset + "|"
		case '2':
			line += importantColor + centerCell(important, cellWidth) + reset + "|"
		}
	}
	return line
}
