package main

import (
	"fmt"
	"os"

	"github.com/lwileczek/Bafflizer/bafflerz"
	"github.com/lwileczek/Bafflizer/config"
)

func main() {

	cfg := config.SetupEnv()

	file, err := os.ReadFile(cfg.Input) // []bytes
	check(err)
	bafflerz.BafflePythonFile(&file)

	if printOutput := true; printOutput {
		fmt.Println(string(file))
	}

	if writeFile := false; writeFile {
		WriteFile(&file, cfg.Output)
	}
	fmt.Println("Done.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//WriteFile Write out the contents of the new file
func WriteFile(fileContent *[]byte, outputFile string) {
	f, err := os.Create(outputFile)
	check(err)
	defer f.Close()
	f.Write(*fileContent)
}
