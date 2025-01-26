package string_analyzer

import "fmt"

/*
Builder of string analyzer
*/
type StringAnalyzerBuilder struct {
	flags     map[string]bool
	criteries map[string](func(*string) bool)
}

// Add Checker to analyzer
func (sab *StringAnalyzerBuilder) AddChecker(name string, checker func(*string) bool) *StringAnalyzerBuilder {
	if _, ok := sab.criteries[name]; ok {
		panic(fmt.Sprintf("%s already exists", name))
	}
	sab.criteries[name] = checker
	return sab
}

// Add some checkers to analyzers
func (sab *StringAnalyzerBuilder) AddCheckers(slice []struct {
	name    string
	checker func(*string) bool
}) *StringAnalyzerBuilder {
	for _, element := range slice {
		sab.AddChecker(element.name, element.checker)
	}
	return sab
}

// Analyzed String will save original text
func (sab *StringAnalyzerBuilder) SaveOriginalString() *StringAnalyzerBuilder {
	sab.flags["save_string"] = true
	return sab
}

// Analyzed String will count
func (sab *StringAnalyzerBuilder) SaveCombinations() *StringAnalyzerBuilder {
	sab.flags["save_combinations"] = true
	return sab
}

// Return String Analyzer
func (sab *StringAnalyzerBuilder) Construct() *StringAnalyzer {
	ret := &StringAnalyzer{sab.flags, sab.criteries}
	return ret
}

// Init builder
func CreateSABuilder() StringAnalyzerBuilder {
	return StringAnalyzerBuilder{make(map[string]bool), make(map[string](func(*string) bool))}
}
