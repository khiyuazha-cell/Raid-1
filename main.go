package main

import "fmt"

func main() {

	printSchedule([]string{"0102", "1100", "0010", "2001"}, "*", "$", "blue", "green")

}

func printSchedule(rows []string, busy string, important string, busyColor string, importantColor string) {
	if ValidateInput(rows, busyColor, importantColor) != "" {
		fmt.Println(ValidateInput(rows, busyColor, importantColor))
		return
	}
	rowCount := len(rows)
	fmt.Println(renderBorder(rowCount))
	for _, row := range rows {
		fmt.Println(renderRow(row, busy, important, busyColor, importantColor))
		fmt.Println(renderBorder(rowCount))
	}
}
