package main

import (
	"fmt"
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

type MyString struct {
	innerString  string
	sizeOfString int
	numOfString  uint32
	symbolMap    map[rune]uint
	stringType   TypeOfString
	hasYo        bool
}

func GetType(s *string) TypeOfString {
	CLASS := regexp.MustCompile(`[ѢѣІіѲѳѴѵ]|([ВКСфкцнгшщзхфвпрлджчсмтб]ъ[ ,.;:?!-"'])`)
	REFORM := regexp.MustCompile(`([иИ][яеёоыеиюэ])|([ВКСфкцнгшщзхфвпрлджчсмтб][ ,.;:?!-"'])`)
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

func CreateMyString(numOfString uint32, s *string) MyString {
	symbolMap := getMap([]rune(*s))
	sizeOfString := len(*s)
	hasYo := checkYo(s)
	typeOfStr := GetType(s)
	return MyString{*s, sizeOfString, numOfString, symbolMap, typeOfStr, hasYo}
}

func getMap(slice []rune) map[rune]uint {
	ret := map[rune]uint{}
	for _, symbol := range slice {
		/* _, ok := ret[symbol]
		if ok {
			ret[symbol] = 0
		} */
		ret[symbol]++
	}
	return ret
}

func checkYo(s *string) bool {
	return strings.ContainsAny(*s, "ёЁ")
}

func main() {
	message := "Hello world"
	fmt.Println(message)
}
