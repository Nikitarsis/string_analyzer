package stringanalyzer

type StringAnalyzer struct {
	criteriaMapper []struct {
		str     string
		checker func(*string) bool
	}
}

func (sa StringAnalyzer) AnalyzeString(s *string) AnalyzedString {

	var flags []struct {
		key string
		b   bool
	}
	for _, criteria := range sa.criteriaMapper {
		flag := criteria.checker(s)
		flags = append(flags, struct {
			key string
			b   bool
		}{criteria.str, flag})
	}
	return constructAnalyzedString(s, flags...)
}

func getMap(slice []rune) map[rune]uint {
	ret := map[rune]uint{}
	for _, symbol := range slice {
		ret[symbol]++
	}
	return ret
}
