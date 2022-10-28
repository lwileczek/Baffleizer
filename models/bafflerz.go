package models

//Baffler Define universal methods for a Baffler
type Baffler interface {
	ProcessFile(file string, rules []UpdateRule) []string
}

//UpdateRule Each change to a file has a condition and function to perform that update
//For example condition: import is in line, then update the import statements by aliasing them
type UpdateRule interface {
	Condition(s *string) bool
	Update(s *string) string
}

//BaseBaffler All bafflers, regardless of language or obfuscation method will have
type BaseBaffler struct {
	Injection   map[string]string
	ChangeCount uint32
	rules       []UpdateRule
}
