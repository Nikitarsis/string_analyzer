package main

import (
	"args"
	"fmt"
	"os"
	"regexp"
	"stringanalyzer"
)

//CLASS := regexp.MustCompile(`[ѢѣІіѲѳѴѵ]|([ВКСфкцнгшщзхфвпрлджчсмтб]ъ[ ,.;:?!\-"'])`)
//REFORM := regexp.MustCompile(`([иИ][яеёоыеиюэ])|([ВКСфкцнгшщзхфвпрлджчсмтб][ ,.;:?!\-"'])`)
//TRASH := regexp.MustCompile(`.{,5}`)

var configSingleton Config

func getParser() *args.ArgsParser {
	builder := args.NewParserBuilder()
	builder.AddElementAtLeast(func(s ...string) { configSingleton.SetReadingFiles(s...) }, 1, "inputFile", false, "i")
	builder.AddElementAtLeast(func(s ...string) { configSingleton.SetOutputFiles(s...) }, 1, "outputFile", false, "o")
	builder.AddElementAtMost(func(s ...string) { configSingleton.SaveString() }, 0, "saveStrings", false, "s")
	builder.AddElementAtMost(func(s ...string) { configSingleton.CountCombo() }, 0, "countCombinations", false, "c")
	builder.AddElementAtMost(func(s ...string) { configSingleton.DoNotPipeline() }, 0, "nopipeline", false, "n")
	ret, err := builder.Construct()
	if err != nil {
		panic(err)
	}
	return ret
}

func getStringAnalyzer(saveStr bool, countComb bool) *stringanalyzer.StringAnalyzer {
	CLASS := regexp.MustCompile(`[ѢѣІіѲѳѴѵ]|([ВКСфкцнгшщзхфвпрлджчсмтб]ъ[ ,.;:?!\-"'])`)
	REFORM := regexp.MustCompile(`([иИ][яеёоыеиюэ])|([ВКСфкцнгшщзхфвпрлджчсмтб][ ,.;:?!\-"'])`)
	TRASH := regexp.MustCompile(`.{,5}`)

	builder := stringanalyzer.CreateSABuilder()
	if countComb {
		builder.SaveCombinations()
	}
	if saveStr {
		builder.SaveOriginalString()
	}
	builder.AddChecker("isClassical", func(s *string) bool { return CLASS.MatchString(*s) })
	builder.AddChecker("isReformed", func(s *string) bool { return REFORM.MatchString(*s) })
	builder.AddChecker("isTrash", func(s *string) bool { return TRASH.MatchString(*s) })
	return builder.Construct()
}

func main() {
	configSingleton = GetConfig()
	parser := getParser()
	parser.ParseArgs(os.Args...)
	analyzer := getStringAnalyzer(configSingleton.ShouldSaveString(), configSingleton.ShouldCountCombo())
	analyzeFunc := func(s *string) ([]byte, bool) {
		ret, err := analyzer.AnalyzeString(s).GetJson()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, false
		}
		return ret, true
	}
	message := "Hello World"
	fmt.Println(analyzeFunc(&message))
}
