package utils

import (
	"bufio"
	"fmt"
	"os"
)

const (
	defaultTextOutput = "result.txt"
)

var (
	TextOutput = defaultTextOutput
)

// 导出为text文本格式
func ExportText(data []CloudflareIPData) {
	if noOutput() || len(data) == 0 {
		return
	}

	file, err := os.OpenFile(TextOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer from the file.
	writer := bufio.NewWriter(file)

	// Write each line to the file with a newline character
	for _, line := range data {
		_, err := writer.WriteString(fmt.Sprintf("%s:443#%s\n", line.IP.IP, "HKG"))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
	}

	fmt.Println("Lines written to file successfully.")

}
