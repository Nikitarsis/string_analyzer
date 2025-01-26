package string_analyzer

import (
	"encoding/json"
)

/*
String with only symbol map and checklist
*/
type minimalAnalyzedString struct {
	SizeOfString int             `json:"size_of_string"`
	SymbolMap    map[string]uint `json:"symbol_map"`
	FlagMap      map[string]bool `json:"flag_map"`
}

func (mas *minimalAnalyzedString) GetSize() int {
	return mas.SizeOfString
}

func (mas *minimalAnalyzedString) GetSymbolMap() map[string]uint {
	return mas.SymbolMap
}

func (mas *minimalAnalyzedString) GetFlagMap() map[string]bool {
	return mas.FlagMap
}

func (mas *minimalAnalyzedString) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(mas)
	return &ret, err
}

/*
Contains all from [[MinimalAnalyzedString]] and source string
*/
type analyzedStringWithOriginalText struct {
	*minimalAnalyzedString
	InnerString *string `json:"inner_string"`
}

func (as *analyzedStringWithOriginalText) GetSize() int {
	return as.SizeOfString
}

func (as *analyzedStringWithOriginalText) GetSymbolMap() map[string]uint {
	return as.SymbolMap
}

func (as *analyzedStringWithOriginalText) GetFlagMap() map[string]bool {
	return as.FlagMap
}

func (as *analyzedStringWithOriginalText) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}

/*
Contains all from [[MinimalAnalyzedString]] plus combinationMap
*/
type analyzedStringWithCombinations struct {
	*minimalAnalyzedString
	CombinationMap map[string]uint `json:"combination_map"`
}

func (as *analyzedStringWithCombinations) GetSize() int {
	return as.SizeOfString
}

func (as *analyzedStringWithCombinations) GetSymbolMap() map[string]uint {
	return as.SymbolMap
}

func (as *analyzedStringWithCombinations) GetFlagMap() map[string]bool {
	return as.FlagMap
}

func (as *analyzedStringWithCombinations) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}

/*
Full Analyzed String with sorce string and combinations
*/
type fullAnalyzedString struct {
	*analyzedStringWithOriginalText
	CombinationMap map[string]uint `json:"combination_map"`
}

func (as fullAnalyzedString) GetSize() int {
	return as.SizeOfString
}

func (as fullAnalyzedString) GetSymbolMap() map[string]uint {
	return as.SymbolMap
}

func (as fullAnalyzedString) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}
