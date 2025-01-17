package stringanalyzer

type AnalyzedString struct {
	innerString  *string
	sizeOfString int
	symbolMap    map[rune]uint
	flagMap      map[string]bool
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
	symbolMap := constructSymMap([]rune(*str))
	return AnalyzedString{str, len(*str), symbolMap, flagMap}
}

func constructSymMap(slice []rune) map[rune]uint {
	ret := map[rune]uint{}
	for _, symbol := range slice {
		ret[symbol]++
	}
	return ret
}
