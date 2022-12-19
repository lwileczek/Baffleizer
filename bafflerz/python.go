package bafflerz

import (
	"fmt"
	"regexp"

	"github.com/lwileczek/Bafflizer/models"
)

//BafflePythonFile to setup and return Python baffler... is this needed?
func BafflePythonFile(content *[]byte, passLength int) {
	test := 0
	localFunctions := &functionName{
		varLength: passLength,
	}
	sir := standardImportRule{
		passwdLength: passLength,
		MatchCount:   &test,
		dictionary:   make(map[string][][]byte),
	}
	commentRemover := comments{}
	multiLineReturnRemover := lineEndings{}
	b := models.Baffler{
		Name:    "Pythonalizer",
		Content: content,
		Rules: []models.RuleSet{
			sir,
			localFunctions,
			commentRemover,
			multiLineReturnRemover,
		},
	}

	for _, rule := range b.Rules {
		rule.Find(b.Content)
		rule.Update(b.Content)
	}
}

type standardImportRule struct {
	passwdLength int
	MatchCount   *int
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
		*i.MatchCount++
	}
}

// UpdateImportLine update import statements in a python file
func (i standardImportRule) Update(text *[]byte) {
	for z := 0; z < *i.MatchCount; z++ {
		newID := RandomString(i.passwdLength)
		if key := i.dictionary["alias"][z]; key != nil {
			//step 1 import line
			currentImportLine := fmt.Sprintf("import %s%s", i.dictionary["package"][z], i.dictionary["fullAlias"][z])
			newImportLine := fmt.Sprintf("import %s as %s", i.dictionary["package"][z], newID)
			re := regexp.MustCompile(currentImportLine)
			*text = re.ReplaceAll(*text, []byte(newImportLine))

			//step 2 all instances
			regex := fmt.Sprintf(`(\W)%s\.`, i.dictionary["alias"][z])
			re2 := regexp.MustCompile(regex)
			*text = re2.ReplaceAll(*text, []byte("${1}"+newID+"."))
		} else {
			currentImportLine := fmt.Sprintf("import %s", i.dictionary["package"][z])
			newImportLine := fmt.Sprintf("import %s as %s", i.dictionary["package"][z], newID)
			re := regexp.MustCompile(currentImportLine)
			*text = re.ReplaceAll(*text, []byte(newImportLine))

			regex := fmt.Sprintf(`(\W)%s\.`, i.dictionary["package"][z])
			re2 := regexp.MustCompile(regex)
			*text = re2.ReplaceAll(*text, []byte("${1}"+newID+"."))
		}
	}
}

type functionName struct {
	varLength    int
	functionList [][]byte
}

func (fn *functionName) Find(text *[]byte) {
	//re := regexp.MustCompile(`def\s+(?P<funcName>[a-zA-Z][\w_]*)\s*\(`)
	re := regexp.MustCompile(`def\s+(?P<funcName>[a-zA-Z][\w_]*)\b`)
	matches := re.FindAllSubmatch(*text, -1)
	if len(matches) == 0 {
		return
	}
	for match := 0; match < len(matches); match++ {
		for k := range re.SubexpNames() {
			if k > 0 {
				fn.functionList = append(fn.functionList, matches[match][k])
			}
		}
	}
}
func (fn functionName) Update(text *[]byte) {
	for _, wrd := range fn.functionList {
		newName := RandomString(fn.varLength)
		findFunc := fmt.Sprintf(`([\s\.])%s\(`, wrd)
		re := regexp.MustCompile(findFunc)
		*text = re.ReplaceAll(*text, []byte("${1}"+newName+"("))
	}
}

type comments struct{}

func (c comments) Find(text *[]byte) {
	// You have to make sure the octothorp is not within a string
	lineComment := regexp.MustCompile(`^\s+\#`)
	matches := lineComment.FindAllSubmatch(*text, -1)
	if len(matches) == 0 {
		return
	}
}
func (c comments) Update(text *[]byte) {
	//Not able to handle if a pound sign is within a quote
	none := []byte("")
	//Trusting openning and closing quotes are the same. Probably bad.
	multiLineComents := regexp.MustCompile(`(?m)^\s*(['"]){3}[\w\W]*?(['"]){3}`)
	*text = multiLineComents.ReplaceAll(*text, none)
	lineComments := regexp.MustCompile(`^\s*#.*`)
	*text = lineComments.ReplaceAll(*text, none)
}

type lineEndings struct{}

func (le lineEndings) Find(text *[]byte) {}
func (le lineEndings) Update(text *[]byte) {
	multiLineComents := regexp.MustCompile(`(?m)[\n\r][\r\n]+`)
	*text = multiLineComents.ReplaceAll(*text, []byte("\n"))
}

//TODO:
// Create replacement functions for each of the following, and try to run them on an entire file not line-by-line
//  - [x] imports
//  - [x] function
//  - global variables
//  - function parameters
//  - classes
//  - ...
