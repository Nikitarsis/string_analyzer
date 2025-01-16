package args

type ArgEntity struct {
	ArgName            string
	AssociatedFunction func([]string)
	Pseudonyms         []string
	NargsChecker       func(uint) bool
}

func ConstructEntityChecker(function func([]string), name string, nargsChecker func(uint) bool, pseudonyms ...string) ArgEntity {
	return ArgEntity{name, function, pseudonyms, nargsChecker}
}

func ConstructEntity(function func([]string), name string, nargs uint, pseudonyms ...string) ArgEntity {
	return ConstructEntityChecker(function, name, func(x uint) bool { return x == nargs }, pseudonyms...)
}
