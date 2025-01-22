package stringanalyzer

import "fmt"

type StringAnalyzerBuilder struct {
	flags     map[string]bool
	criteries map[string](func(*string) bool)
}

func (sab *StringAnalyzerBuilder) AddChecker(str string, function func(*string) bool) *StringAnalyzerBuilder {
	if _, ok := sab.criteries[str]; ok {
		panic(fmt.Sprintf("%s already exists", str))
	}
	sab.criteries[str] = function
	return sab
}

func (sab *StringAnalyzerBuilder) AddCheckers(slice []struct {
	str     string
	checker func(*string) bool
}) *StringAnalyzerBuilder {
	for _, element := range slice {
		sab.AddChecker(element.str, element.checker)
	}
	return sab
}

func (sab *StringAnalyzerBuilder) SaveOriginalString() *StringAnalyzerBuilder {
	sab.flags["save_string"] = true
	return sab
}

func (sab *StringAnalyzerBuilder) SaveCombinations() *StringAnalyzerBuilder {
	sab.flags["save_combinations"] = true
	return sab
}

func (sab *StringAnalyzerBuilder) Construct() *StringAnalyzer {
	ret := &StringAnalyzer{sab.flags, sab.criteries}
	return ret
}

func CreateSABuilder() StringAnalyzerBuilder {
	return StringAnalyzerBuilder{make(map[string]bool), make(map[string](func(*string) bool))}
}
