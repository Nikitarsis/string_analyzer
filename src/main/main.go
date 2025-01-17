package main

import (
	"args"
	"fmt"
	"os"
)

//CLASS := regexp.MustCompile(`[ѢѣІіѲѳѴѵ]|([ВКСфкцнгшщзхфвпрлджчсмтб]ъ[ ,.;:?!\-"'])`)
//REFORM := regexp.MustCompile(`([иИ][яеёоыеиюэ])|([ВКСфкцнгшщзхфвпрлджчсмтб][ ,.;:?!\-"'])`)
//TRASH := regexp.MustCompile(`.{,5}`)

func getParser() *args.ArgsParser {
	builder := args.NewParserBuilder()
	builder.AddElementAtLeast(func(s ...string) {}, 1, "input", true, "i")
	builder.AddElementAtLeast(func(s ...string) {}, 1, "output", true, "o")
	builder.AddElementAtMost(func(s ...string) {}, 0, "pipeline", false, "p")
	builder.AddElementAtMost(func(s ...string) {}, 0, "debug", false, "d")
	ret, err := builder.Construct()
	if err != nil {
		panic(err)
	}
	return ret
}

func main() {
	parser := getParser()
	parser.ParseArgs(os.Args...)
	message := "Hello world"
	fmt.Println(message)
}
