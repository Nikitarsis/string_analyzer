package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type TypeOfString string

const (
	CLASSICAL_O TypeOfString = "class"
	REFORMED_O               = "reformed"
	TRASH_O                  = "trash"
	UNKNOWN_O                = "unknown"
)

type AnalyzedString struct {
	innerString  string
	sizeOfString int
	numOfString  uint32
	symbolMap    map[rune]uint
	stringType   TypeOfString
	hasYo        bool
}

func GetType(s *string) TypeOfString {
	CLASS := regexp.MustCompile(`[ѢѣІіѲѳѴѵ]|([ВКСфкцнгшщзхфвпрлджчсмтб]ъ[ ,.;:?!\-"'])`)
	REFORM := regexp.MustCompile(`([иИ][яеёоыеиюэ])|([ВКСфкцнгшщзхфвпрлджчсмтб][ ,.;:?!\-"'])`)
	TRASH := regexp.MustCompile(`.{,5}`)
	if TRASH.MatchString(*s) {
		return TRASH_O
	}
	isClass := CLASS.MatchString(*s)
	isReform := REFORM.MatchString(*s)
	if isReform && isClass {
		return UNKNOWN_O
	}
	if isReform {
		return REFORMED_O
	}
	if isClass {
		return CLASSICAL_O
	}
	return UNKNOWN_O
}

func CreateAnalyzedString(numOfString uint32, s *string) AnalyzedString {
	symbolMap := getMap([]rune(*s))
	sizeOfString := len(*s)
	hasYo := checkYo(s)
	typeOfStr := GetType(s)
	return AnalyzedString{*s, sizeOfString, numOfString, symbolMap, typeOfStr, hasYo}
}

func getMap(slice []rune) map[rune]uint {
	ret := map[rune]uint{}
	for _, symbol := range slice {
		ret[symbol]++
	}
	return ret
}

func checkYo(s *string) bool {
	return strings.ContainsAny(*s, "ёЁ")
}

func loadFromFile(fpath *string) []string {
	file, err := os.Open(*fpath)
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var ret []string
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			return ret
		}
		if err != nil {
			l := log.New(os.Stderr, "", 0)
			l.Println(err)
		}
		ret = append(ret, str)
	}
}

type ArgsParser struct {
	pseudonyms     map[string]string
	possibleArgs   map[string]uint
	functionMapper map[string]func([]string)
}

func ConstructParser(entities ...ArgEntity) ArgsParser {

}

type ArgEntity struct {
	ArgName            string
	AssociatedFunction func([]string)
	Pseudonyms         []string
}

func ConstructEntity(function func([]string), name string, pseudonyms ...string) ArgEntity {
	return ArgEntity{name, function, pseudonyms}
}

func parseArgs(args []string) map[string][]string {
	result := make(map[string][]string)
	var currentKey string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			currentKey = strings.TrimLeft(arg, "--")
			if currentKey == "" {
				panic("Invalid argument key")
			}
			_, exists := result[currentKey]
			if exists {
				panic(fmt.Sprintf("Duplicate argument key: %s", currentKey))
			}

			result[currentKey] = []string{}
			continue
		}
		if strings.HasPrefix(arg, "-") {
			currentKey = strings.TrimLeft(arg, "-")
			if currentKey == "" {
				panic("Invalid argument key")
			}

			if _, exists := result[currentKey]; exists {
				panic(fmt.Sprintf("Duplicate argument key: %s", currentKey))
			}

			result[currentKey] = []string{}
			continue
		}
		if currentKey == "" {
			panic(fmt.Sprintf("Value provided without a key: %s", arg))
		}
		result[currentKey] = append(result[currentKey], arg)
	}

	return result
}

func main() {
	args := os.Args
	message := "Hello world"
	fmt.Println(message)
}
