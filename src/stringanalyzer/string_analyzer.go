package stringanalyzer

// String Analyzer is getting string and returning AnalyzedString
type StringAnalyzer struct {
	flags     map[string]bool
	criteries map[string]func(*string) bool
}

// Main function of String Analyzer that returns AnalyzedString
func (sa StringAnalyzer) AnalyzeString(s *string) IAnalyzedString {

	checkMap := make(map[string]bool)
	_, saveStrings := sa.flags["save_string"]
	_, saveCombos := sa.flags["save_combinations"]
	for name, function := range sa.criteries {
		checkMap[name] = function(s)
	}
	symbolMap, combinationMap := constructSymMap([]rune(*s))
	minimalAS := MinimalAnalyzedString{len(*s), symbolMap, checkMap}
	if !saveCombos && !saveStrings {
		return &minimalAS
	}
	if saveCombos && !saveStrings {
		return &AnalyzedStringWithCombinations{&minimalAS, combinationMap}
	}
	stringsAS := AnalyzedStringWithOriginalText{&minimalAS, s}
	if !saveCombos && saveStrings {
		return &stringsAS
	}
	return &FullAnalyzedString{&stringsAS, combinationMap}
}

// returns maps of symbols and combinations
func constructSymMap(slice []rune) (map[rune]uint, map[[2]rune]uint) {
	retOne := map[rune]uint{}
	retTwo := map[[2]rune]uint{}
	var previousSymbol rune
	for i, symbol := range slice {
		retOne[symbol]++
		if i == 0 {
			previousSymbol = symbol
			continue
		}
		retTwo[[2]rune{previousSymbol, symbol}]++
		previousSymbol = symbol
	}
	return retOne, retTwo
}
