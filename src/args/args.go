package args

import (
	"errors"
	"fmt"
	"strings"
)

func DivideArgs(args []string) (map[string][]string, error) {
	result := make(map[string][]string)
	var currentKey string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			currentKey = strings.TrimLeft(arg, "--")
			if currentKey == "" {
				return nil, errors.New("Invalid argument key")
			}
			_, exists := result[currentKey]
			if exists {
				return nil, errors.New(fmt.Sprintf("Duplicate argument key: %s", currentKey))
			}

			result[currentKey] = []string{}
			continue
		}
		if strings.HasPrefix(arg, "-") {
			currentKey = strings.TrimLeft(arg, "-")
			if currentKey == "" {
				return nil, errors.New("Invalid argument key")
			}

			if _, exists := result[currentKey]; exists {
				return nil, errors.New(fmt.Sprintf("Duplicate argument key: %s", currentKey))
			}

			result[currentKey] = []string{}
			continue
		}
		if currentKey == "" {
			return nil, errors.New(fmt.Sprintf("Value provided without a key: %s", arg))
		}
		result[currentKey] = append(result[currentKey], arg)
	}

	return result, nil
}
