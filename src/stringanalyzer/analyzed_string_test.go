package string_analyzer_test

import (
	"fmt"
	"strings"
	"testing"

	stringanalyzer "github.com/Nikitarsis/string_analyzer"
)

func addDefaultCheckers(builder *stringanalyzer.StringAnalyzerBuilder) *stringanalyzer.StringAnalyzerBuilder {
	builder.AddChecker("not_null", func(s *string) bool { return *s != "" })
	builder.AddChecker("e", func(s *string) bool { return strings.Contains(*s, "e") })
	builder.AddChecker("Bruh", func(s *string) bool { return strings.HasPrefix(*s, "Bruh") })
	builder.AddChecker("Lol", func(s *string) bool { return strings.HasSuffix(*s, "Lol") })
	return builder
}

func TestBasicFunctions(t *testing.T) {
	builder := stringanalyzer.CreateSABuilder()
	addDefaultCheckers(&builder)
	parser := builder.Construct()
	str := "Bruh assfqageeeee"
	as := parser.AnalyzeString(&str)
	rune_map := as.GetSymbolMap()
	flag_map := as.GetFlagMap()
	if rune_map["a"] != 2 {
		t.Error("incorrect number of letters a")
	}
	if rune_map["s"] != 2 {
		t.Error("incorrect number of letters s")
	}
	if rune_map["B"] != 1 {
		t.Error("incorrect number of letters B")
	}
	if !flag_map["not_null"] {
		t.Error("incorrect flag not_null")
	}
	if !flag_map["e"] {
		t.Error("incorrect flag e")
	}
	if !flag_map["Bruh"] {
		t.Error("incorect flag Bruh")
	}
	if flag_map["Lol"] {
		t.Error("incorect flag Lol")
	}
}

func TestSerialization(t *testing.T) {
	builder := stringanalyzer.CreateSABuilder()
	builder.AddChecker("is_big", func(s *string) bool { return len(*s) >= 250 })
	builder.AddChecker("contains_God", func(s *string) bool { return strings.Contains(*s, "Бог") })
	builder.AddChecker("is_small", func(s *string) bool { return len(*s) < 50 })
	builder.SaveCombinations()
	builder.SaveOriginalString()
	parser := builder.Construct()
	str := "Чего бы глаза мои ни пожелали, я не отказывал им, не возбранял сердцу моему никакого веселья, потому что сердце мое радовалось во всех трудах моих, и это было моею долею от всех трудов моих."
	ret, err := parser.AnalyzeString(&str).GetJson()
	if err != nil {
		t.Errorf("Error of %s", err.Error())
	}
	ret_str := string(*ret)
	fmt.Println(ret_str)
}
