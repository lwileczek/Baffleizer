package bafflerz

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

//BafflePythonFile to setup and return Python baffler... is this needed?
func BafflePythonFile(scanner *bufio.Scanner) []string {
	var output []string
	injective := map[string]string{}
	//TODO: loop over file once to get mapping, then loop over mapping to update the file
	// Currently we're probably doing it slowly.
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		switch {
		case strings.Contains(line, "import "):
			newLine, m := updateImportLine(&line, 32)
			output = append(output, newLine, "\n")
			for k, v := range m {
				injective[k] = v
			}
		default:
			var regex string
			var re *regexp.Regexp
			for k, v := range injective {
				regex = fmt.Sprintf(`[\s\(](%s)[\s\)\.]`, k)
				fmt.Println("regex expression was:", regex)
				re = regexp.MustCompile(regex)
				line = re.ReplaceAllString(line, v)
			}
			newLine := line
			output = append(output, newLine, "\n")
		}
	}

	fmt.Println("File Mapping")
	fmt.Println(injective)
	return output
}

// UpdateImportLine update import statements in a python file
func updateImportLine(line *string, n int) (string, map[string]string) {
	groupRegex := `import\s(?P<package>\w+)(?P<fullAlias>\sas\s(?P<alias>\w+))?`
	groups := getParams(groupRegex, *line)
	fmt.Println("These were the named groups found")
	fmt.Println(groups)
	newID := RandomString(n)
	switch groups["alias"] {
	case "":
		key := groups["package"]
		mapping := map[string]string{
			key: newID,
		}
		return fmt.Sprintf("%s as %s", *line, newID), mapping
	default:
		mapping := map[string]string{
			groups["alias"]: newID,
		}
		return fmt.Sprintf("%s as %s", strings.Split(*line, " as ")[0], newID), mapping
	}
}
