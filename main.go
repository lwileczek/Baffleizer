package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	// TODO: Read number of lines and then preallocate memory
	var output []string

	file, err := os.Open("examples/ex.py")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords) // use scanwords
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

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, eachLine := range output {
		fmt.Println(eachLine)
	}

	WriteFile(&output)
}

func updateImportLine(line *string, n int) string {
	newID := RandomString(n)
	if strings.Contains(*line, " as ") {
		return fmt.Sprintf("%s as %s", strings.Split(*line, " as ")[0], newID)
	}
	return fmt.Sprintf("%s as %s", *line, newID)
}

//RandomString Create a random string starting with an underscore
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, (n - 1))
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("_%s", string(s))
}

//WriteFile Write out the contents of the new file
func WriteFile(file *[]string) {
	f, err := os.Create("example_out.py")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, line := range *file {
		f.Write([]byte(line))
	}
}
