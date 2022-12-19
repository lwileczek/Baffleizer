package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// How to Regex ...
	file, err := os.ReadFile("ex.py")
	if err != nil {
		panic(err)
	}

	// Ignore import x from ... Standard import
	expression := `(?m)^\s*?import\s(?P<package>\w+)(?P<fullAlias>\sas\s(?P<alias>\w+))?`
	var re = regexp.MustCompile(expression)
	matches := re.FindAllSubmatch(file, -1)

	fmt.Println(matches, "\n\n ")
	if len(matches) == 0 {
		fmt.Println("No matches")
		return
	}

	params := make(map[string][][]byte)
	for match := 0; match < len(matches); match++ {
		fmt.Println(matches[match])
		for i, name := range re.SubexpNames() {
			if i > 0 {
				params[name] = append(params[name], matches[match][i])
			}
		}
	}

	fmt.Println(params)
	fmt.Println("Let's render package names:")
	for _, pkg := range params["fullAlias"] {
		fmt.Println(string(pkg))
	}
}
