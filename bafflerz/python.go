package bafflerz

import (
	"fmt"
	"regexp"

	"github.com/lwileczek/Bafflizer/models"
)

//BafflePythonFile to setup and return Python baffler... is this needed?
func BafflePythonFile(content *[]byte) {
	test := 0
	sir := standardImportRule{
		passwdLength: 30,
		MatchCount:   &test,
		dictionary:   make(map[string][][]byte),
	}
	commentRemover := comments{}
	b := models.Baffler{
		Name:    "Pythonalizer",
		Content: content,
		Rules: []models.RuleSet{
			sir,
			commentRemover,
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

//TODO: Change RuleSets to check for errors?
type functionName struct {
	varLength    int
	functionList [][]byte
}

func (fn functionName) Find(text *[]byte) {
	re := regexp.MustCompile(`def\s+(?P<funcName>[a-zA-Z][\w_]*)\s*\(`)
	matches := re.FindAllSubmatch(*text, -1)
	if len(matches) == 0 {
		return
	}
	//TODO: loop through matches and flatten to a single slice
	//fn.functionList = append(fn.functionList, matches)
}

type comments struct{}

func (c comments) Find(text *[]byte) {}
func (c comments) Update(text *[]byte) {
	lineComments := regexp.MustCompile(`#.*`) // what if # i within a string?
	*text = lineComments.ReplaceAll(*text, []byte(""))
	multiLineComents := regexp.MustCompile(`""".*"""`)
	*text = multiLineComents.ReplaceAll(*text, []byte(""))
}

//TODO:
// Create replacement functions for each of the following, and try to run them on an entire file not line-by-line
//  - imports
//  - function
//  - global variables
//  - function parameters
//  - classes
//  - ...
