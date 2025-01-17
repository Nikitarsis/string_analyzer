package args

type argEntity struct {
	ArgName            string
	AssociatedFunction func([]string)
	Pseudonyms         []string
	NargsChecker       func(uint) bool
}

func constructEntityChecker(function func([]string), name string, nargsChecker func(uint) bool, pseudonyms ...string) argEntity {
	return argEntity{name, function, pseudonyms, nargsChecker}
}

func constructEntity(function func([]string), name string, nargs uint, pseudonyms ...string) argEntity {
	return constructEntityChecker(function, name, func(x uint) bool { return x == nargs }, pseudonyms...)
}
