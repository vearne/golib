package utils

import (
	"fmt"
	"strings"
)

/*
	0 <= progress <= 100
*/
func progressStr(progress int) string {
	if progress <= 0 {
		progress = 0
	}
	if progress >= 100 {
		progress = 100
	}
	var sb strings.Builder
	for i := 0; i < 100/5; i++ {
		if i < progress/5 {
			sb.WriteString("#")
		} else {
			sb.WriteString(" ")
		}
	}
	return fmt.Sprintf("[%v] %v ", sb.String(), progress)
}

/*
	output: 
	[####                ] 20 %
*/
func PrintProgress(progress int) {
	fmt.Printf(progressStr(progress) + "%% \r")
}
