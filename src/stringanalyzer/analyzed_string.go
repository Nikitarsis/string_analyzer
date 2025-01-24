package stringanalyzer

import (
	"encoding/json"
)

/*
String with only symbol map and checklist
*/
type MinimalAnalyzedString struct {
	sizeOfString int
	symbolMap    map[rune]uint
	flagMap      map[string]bool
}

func (mas *MinimalAnalyzedString) GetSize() int {
	return mas.sizeOfString
}

func (mas *MinimalAnalyzedString) GetSymbolMap() map[rune]uint {
	return mas.symbolMap
}

func (mas *MinimalAnalyzedString) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(mas)
	return &ret, err
}

/*
Contains all from [[MinimalAnalyzedString]] and source string
*/
type AnalyzedStringWithOriginalText struct {
	*MinimalAnalyzedString
	innerString *string
}

func (as *AnalyzedStringWithOriginalText) GetSize() int {
	return as.sizeOfString
}

func (as *AnalyzedStringWithOriginalText) GetSymbolMap() map[rune]uint {
	return as.symbolMap
}

func (as *AnalyzedStringWithOriginalText) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}

/*
Contains all from [[MinimalAnalyzedString]] plus combinationMap
*/
type AnalyzedStringWithCombinations struct {
	*MinimalAnalyzedString
	combinationMap map[[2]rune]uint
}

func (as *AnalyzedStringWithCombinations) GetSize() int {
	return as.sizeOfString
}

func (as *AnalyzedStringWithCombinations) GetSymbolMap() map[rune]uint {
	return as.symbolMap
}

func (as *AnalyzedStringWithCombinations) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}

/*
Full Analyzed String with sorce string and combinations
*/
type FullAnalyzedString struct {
	*AnalyzedStringWithOriginalText
	combinationMap map[[2]rune]uint
}

func (as FullAnalyzedString) GetSize() int {
	return as.sizeOfString
}

func (as FullAnalyzedString) GetSymbolMap() map[rune]uint {
	return as.symbolMap
}

func (as FullAnalyzedString) GetJson() (*[]byte, error) {
	ret, err := json.Marshal(as)
	return &ret, err
}
