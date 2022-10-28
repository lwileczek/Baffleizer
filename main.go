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
	// TODO: Read number of lines and then preallocate memory

	file, err := os.Open(cfg.Input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := bafflerz.BafflePythonFile(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if printOutput := true; printOutput {
		for _, eachLine := range output {
			fmt.Println(eachLine)
		}
	}

	if writeFile := false; writeFile {
		WriteFile(&output, cfg.Output)
	}
	fmt.Println("Done.")
}

//WriteFile Write out the contents of the new file
func WriteFile(file *[]string, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, line := range *file {
		f.Write([]byte(line))
	}
}
