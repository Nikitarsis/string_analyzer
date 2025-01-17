package args

import (
	"slices"
)

type ArgsParserBuilder struct {
	entities []argEntity
}

func NewParserBuilder() ArgsParserBuilder {
	return ArgsParserBuilder{}
}

func (apb *ArgsParserBuilder) AddElement(
	function func(...string),
	checker func(uint) bool,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	add := constructEntityChecker(function, name, checker, isNecessary, pseudonyms...)
	apb.entities = append(apb.entities, add)
	return apb
}

func (apb *ArgsParserBuilder) AddElementSlice(
	function func(...string),
	possibleArguments []uint,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	checker := func(x uint) bool {
		return slices.Contains(possibleArguments, x)
	}
	return apb.AddElement(function, checker, name, isNecessary, pseudonyms...)
}

func (apb *ArgsParserBuilder) AddElimentSingle(
	function func(...string),
	argument uint,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	checker := func(x uint) bool {
		return argument == x
	}
	return apb.AddElement(function, checker, name, isNecessary, pseudonyms...)
}

func (apb *ArgsParserBuilder) AddElementAtLeast(
	function func(...string),
	minNumArguments uint,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	checker := func(x uint) bool {
		return x >= minNumArguments
	}
	return apb.AddElement(function, checker, name, isNecessary, pseudonyms...)
}

func (apb *ArgsParserBuilder) AddElementAtMost(
	function func(...string),
	maxNumArguments uint,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	checker := func(x uint) bool {
		return x <= maxNumArguments
	}
	return apb.AddElement(function, checker, name, isNecessary, pseudonyms...)
}

func (apb *ArgsParserBuilder) AddElementBetween(
	function func(...string),
	minNumArguments uint,
	maxNumArguments uint,
	name string,
	isNecessary bool,
	pseudonyms ...string) *ArgsParserBuilder {

	checker := func(x uint) bool {
		return x <= maxNumArguments && x >= minNumArguments
	}
	return apb.AddElement(function, checker, name, isNecessary, pseudonyms...)
}

func (apb ArgsParserBuilder) Construct() (*ArgsParser, error) {
	return constructParser(apb.entities...)
}
