package main

import (
	"fmt"
	"strings"
)

// printSchedule(rows []string, busy string, important string)
func centerCell(content string, cellWidth int) string {
	padding := cellWidth - len(content)
	if padding < 0 {
		return "error: content is wider than cell width"
	}
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftPadding), content, strings.Repeat(" ", rightPadding))
}
func renderRow(row, busy string, important string) string {
	line := "|"
	for _, ch := range row {
		switch ch {
		case '0':
			line += centerCell(" ", 3) + "|"
		case '1':
			line += centerCell(busy, 3) + "|"
		case '2':
			line += centerCell(important, 3) + "|"
		default:
			return "error: invalid symbol (only 0/1/2 allowed)"
		}
	}
	return line
}
