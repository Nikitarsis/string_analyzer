package args_test

import (
	"args"
	"testing"
)

func getBasicsElement() (func(...string), string, bool, []string) {
	callback := func(s ...string) {

	}
	name := "name"
	isNecessary := false
	pseudonyms := []string{"n"}
	pseudonyms[0] = "aaa"
	return callback, name, isNecessary, pseudonyms
}

func TestAddElement(t *testing.T) {
	builder := args.InitParserBuilder()
	callback, name, isNecessary, pseudonyms := getBasicsElement()
	checker := func(i uint) bool {
		return i%2 == 0
	}
	_, err := builder.AddElement(callback, checker, name, isNecessary, pseudonyms...).Construct()
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}
