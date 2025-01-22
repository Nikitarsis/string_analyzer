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

func getParser() *args.ArgsParser {
	builder := args.NewParserBuilder()
	builder.AddElementAtLeast(func(s ...string) {}, 1, "inputFile", false, "i")
	builder.AddElementAtLeast(func(s ...string) {}, 1, "outputFile", false, "o")
	builder.AddElementAtMost(func(s ...string) {}, 0, "saveStrings", false, "s")
	builder.AddElementAtMost(func(s ...string) {}, 0, "countCombinations", false, "c")
	builder.AddElementAtMost(func(s ...string) {}, 0, "nopipeline", false, "n")
	builder.AddElementAtMost(func(s ...string) {}, 0, "debug", false, "d")
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
	parser := getParser()
	parser.ParseArgs(os.Args...)
	message := "Hello world"
	fmt.Println(message)
}
