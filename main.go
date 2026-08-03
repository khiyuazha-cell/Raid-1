package main

import "fmt"

func main() {

	printSchedule([]string{"010", "211", "002", "201"}, "Busy", "Importants", "red", "green")

}

func printSchedule(rows []string, busy string, important string, busyColor string, importantColor string) {
	if ValidateInput(rows, busy, important, busyColor, importantColor) != "" {
		fmt.Println(ValidateInput(rows, busy, important, busyColor, importantColor))
		return
	}
	colCount := len(rows[0])
	fmt.Println(renderBorder(colCount))
	for _, row := range rows {
		fmt.Println(renderRow(row, busy, important, busyColor, importantColor))
		fmt.Println(renderBorder(colCount))
	}
}
