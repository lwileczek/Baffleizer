package models

import (
	"errors"
	"fmt"
	"strings"
)

type void struct{}

var empty void

// Config Model structure for a config object
type Config struct {
	lang   string
	letter string
	Output string
	Input  string
	Format string
}

//SetLetter Setter function to enforce length of string
func (cfg *Config) SetLetter(val string) error {
	if len(val) > 1 {
		return errors.New("too long")
	}
	cfg.letter = val
	return nil
}

//Letter Getter function for private letter
func (cfg *Config) Letter() string {
	return cfg.letter
}

//SetLang Setter function to enforce length of string
func (cfg *Config) SetLang(val string) error {
	var err error
	approvedLanguages := map[string]void{
		"python": empty,
	}
	lower := strings.ToLower(val)
	_, ok := approvedLanguages[lower]
	if ok {
		cfg.lang = val
		err = nil
	} else {
		err = fmt.Errorf("%s is not an approved language to baffle, please select from %v", val, approvedLanguages)
	}
	return err
}

//Lang Getter function for private lang
func (cfg *Config) Lang() string {
	return cfg.lang
}
