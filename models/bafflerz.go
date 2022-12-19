package models

// Baffler A structure to mystify a file
type Baffler struct {
	Name      string
	Rules     []RuleSet
	Content   *[]byte
	Injection map[string]string
}

// RuleSet A pair of functions to identify rules and update the text with the rule
type RuleSet interface {
	Find(c *[]byte)
	Update(c *[]byte)
}
