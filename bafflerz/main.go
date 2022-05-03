package bafflerz

import (
	"fmt"
	"math/rand"
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
