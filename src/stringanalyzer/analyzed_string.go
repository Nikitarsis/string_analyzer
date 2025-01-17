package stringanalyzer

type AnalyzedString struct {
	innerString    *string
	sizeOfString   int
	symbolMap      map[rune]uint
	combinationMap map[[2]rune]uint
	flagMap        map[string]bool
}

func (as AnalyzedString) GetString() *string {
	return as.innerString
}

func (as AnalyzedString) GetSize() int {
	return as.sizeOfString
}

func (as AnalyzedString) GetSymbolMap() *map[rune]uint {
	return &as.symbolMap
}

func (as AnalyzedString) GetFlagMap() *map[string]bool {
	return &as.flagMap
}

func constructAnalyzedString(
	str *string,
	flags ...struct {
		key string
		b   bool
	}) AnalyzedString {

	flagMap := make(map[string]bool)
	for _, flag := range flags {
		flagMap[flag.key] = flag.b
	}
	symbolMap, combinationMap := constructSymMap([]rune(*str))
	return AnalyzedString{str, len(*str), symbolMap, combinationMap, flagMap}
}

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
