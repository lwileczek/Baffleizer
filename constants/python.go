package constants

// what is this, Java? C?
type void struct{}

var empty void

//PythonReservedWords A set of words that cannot be changed in a python file
var PythonReservedWords = map[string]void{
	"and":       empty,
	"else":      empty,
	"in":        empty,
	"return as": empty,
	"except":    empty,
	"is":        empty,
	"True":      empty,
	"assert":    empty,
	"finally":   empty,
	"lambda":    empty,
	"try":       empty,
	"break":     empty,
	"false":     empty,
	"nonlocal":  empty,
	"with":      empty,
	"class":     empty,
	"for":       empty,
	"None":      empty,
	"while":     empty,
	"continue":  empty,
	"from":      empty,
	"not":       empty,
	"yield":     empty,
	"def":       empty,
	"global":    empty,
	"or":        empty,
	"del":       empty,
	"if":        empty,
	"pass":      empty,
	"elif":      empty,
	"import":    empty,
	"raise":     empty,
}
