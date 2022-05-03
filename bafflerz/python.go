package bafflerz

import (
	"bufio"
	"fmt"
	"strings"
)

//BafflePythonFile to setup and return Python baffler... is this needed?
func BafflePythonFile(scanner *bufio.Scanner) []string {
	var output []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		switch {
		case strings.Contains(line, "import "):
			newLine := updateImportLine(&line, 32)
			output = append(output, newLine)
		default:
			newLine := line
			output = append(output, newLine)
		}
	}

	return output
}

// UpdateImportLine update import statements in a python file
func updateImportLine(line *string, n int) string {
	newID := RandomString(n)
	if strings.Contains(*line, " as ") {
		return fmt.Sprintf("%s as %s", strings.Split(*line, " as ")[0], newID)
	}
	return fmt.Sprintf("%s as %s", *line, newID)
}
