package args

import (
	"errors"
	"fmt"
	"strings"
)

// Simple function that divides args to map that associates keys and additional argument
func DivideArgs(args ...string) (map[string][]string, error) {
	result := make(map[string][]string)
	var currentKey string

	//Checks every argument
	for _, arg := range args {
		//If argument has duble minus, argument is count as single key
		if strings.HasPrefix(arg, "--") {
			currentKey = strings.TrimLeft(arg, "--")
			//Checking incorrect keys
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
		//If argument has single minus, every letter interpretes as key
		if strings.HasPrefix(arg, "-") {
			keys := strings.Split(strings.TrimLeft(arg, "-"), "")
			if len(keys) == 0 {
				return nil, errors.New("Invalid argument key")
			}
			for _, currentKey = range keys {
				//Checking incorrect keys
				if currentKey == "" {
					return nil, errors.New("Invalid argument key")
				}

				if _, exists := result[currentKey]; exists {
					return nil, errors.New(fmt.Sprintf("Duplicate argument key: %s", currentKey))
				}

				result[currentKey] = []string{}
			}
			continue
		}
		if currentKey == "" {
			return nil, errors.New(fmt.Sprintf("Value provided without a key: %s", arg))
		}
		result[currentKey] = append(result[currentKey], arg)
	}

	return result, nil
}
