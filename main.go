package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lwileczek/Bafflizer/bafflerz"
	"github.com/lwileczek/Bafflizer/config"
)

func main() {

	cfg := config.SetupEnv()
	fmt.Println(cfg)
	// TODO: Read number of lines and then preallocate memory

	file, err := os.Open("examples/ex.py")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := bafflerz.BafflePythonFile(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if printOutput := false; printOutput {
		for _, eachLine := range output {
			fmt.Println(eachLine)
		}
	}

	if writeFile := true; writeFile {
		WriteFile(&output)
	}
}

//WriteFile Write out the contents of the new file
func WriteFile(file *[]string) {
	f, err := os.Create("examples/out.py")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, line := range *file {
		f.Write([]byte(line))
	}
}
