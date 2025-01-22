package stringanalyzer

import (
	"encoding/json"
)

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

func (mas *MinimalAnalyzedString) GetJson() ([]byte, error) {
	return json.Marshal(mas)
}

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

func (as *AnalyzedStringWithOriginalText) GetJson() ([]byte, error) {
	return json.Marshal(as)
}

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

func (as *AnalyzedStringWithCombinations) GetJson() ([]byte, error) {
	return json.Marshal(as)
}

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

func (as FullAnalyzedString) GetJson() ([]byte, error) {
	return json.Marshal(as)
}
