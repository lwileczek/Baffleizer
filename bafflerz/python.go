package bafflerz

import (
	"fmt"
	"regexp"
)

//BafflePythonFile to setup and return Python baffler... is this needed?
func BafflePythonFile(content *[]byte) {
	//injective := map[string]string{}
	// regex = fmt.Sprintf(`([=\s\(,\[])%s`, k)
	sir := standardImportRule{
		passwdLength: 30,
		dictionary:   make(map[string][][]byte),
	}
	b := Baffler{
		Name:    "Pythonalizer",
		Content: content,
		Rules:   []RuleSet{sir},
	}

	for _, rule := range b.Rules {
		rule.Find(b.Content)
		rule.Update(b.Content)
	}
}

type standardImportRule struct {
	passwdLength int
	matchCount   int
	dictionary   map[string][][]byte
}

func (i standardImportRule) Find(text *[]byte) {
	expression := `(?m)^\s*?import\s(?P<package>\w+)(?P<fullAlias>\sas\s(?P<alias>\w+))?`
	var re = regexp.MustCompile(expression)
	matches := re.FindAllSubmatch(*text, -1)
	if len(matches) == 0 {
		return
	}

	for match := 0; match < len(matches); match++ {
		for k, name := range re.SubexpNames() {
			if k > 0 {
				i.dictionary[name] = append(i.dictionary[name], matches[match][k])
			}
		}
	}
	i.matchCount = len(matches)

}

// UpdateImportLine update import statements in a python file
func (i standardImportRule) Update(text *[]byte) {
	//I need to update the beginning lines to use the alias and then replace all the instances after.
	for z := 0; z < i.matchCount; z++ {
		newID := RandomString(i.passwdLength)
		if key := i.dictionary["alias"][z]; key != nil {
			//step 1 import line
			currentImportLine := fmt.Sprintf("import %s%s", i.dictionary["package"][z], i.dictionary["fullAlias"][z])
			newImportLine := fmt.Sprintf("import %s as %s", i.dictionary["package"][z], i.dictionary["alias"][z])
			re := regexp.MustCompile(currentImportLine)
			*text = re.ReplaceAll(*text, []byte(newImportLine))

			//step 2 all instances
			regex := fmt.Sprintf(`(\s*)?%s\.`, i.dictionary["alias"][z])
			re2 := regexp.MustCompile(regex)
			*text = re2.ReplaceAll(*text, []byte("${1}"+newID+"."))
		} else {
			currentImportLine := fmt.Sprintf("import %s", i.dictionary["package"][z])
			newImportLine := fmt.Sprintf("import %s as %s", i.dictionary["package"][z], newID)
			re := regexp.MustCompile(currentImportLine)
			*text = re.ReplaceAll(*text, []byte(newImportLine))

			regex := fmt.Sprintf(`(\s*)?%s\.`, i.dictionary["package"][z])
			re2 := regexp.MustCompile(regex)
			*text = re2.ReplaceAll(*text, []byte("${1}"+newID+"."))
		}
	}
}

func findFunctionNames(text *string, n int, dictionary map[string]string) error {
	re := regexp.MustCompile(`def\s+(?P<funcName>[a-zA-Z][\w_]*)\s*\(`)
	matches := re.FindAllStringSubmatch(*text, -1)
	if len(matches) == 0 {
		return nil
	}
	for _, arr := range matches {
		replacement := RandomString(n)
		dictionary[arr[1]] = replacement
	}
	return nil
}

//TODO:
// Create replacement functions for each of the following, and try to run them on an entire file not line-by-line
//  - imports
//  - function
//  - global variables
//  - function parameters
//  - classes
//  - ...
