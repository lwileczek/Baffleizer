package bafflerz

import (
	"fmt"
	"math/rand"
	"regexp"
)

//RandomString Create a random string starting with an underscore
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, (n - 1))
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("_%s", string(s))
}

//getParams Take groups from a regular expression and turn the results in a map where the group is the key
func getParams(regEx, lineInFile string) (paramsMap map[string]string) {

	var re = regexp.MustCompile(regEx)
	match := re.FindStringSubmatch(lineInFile)

	paramsMap = make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
