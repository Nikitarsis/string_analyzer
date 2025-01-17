package args

import (
	"errors"
	"fmt"
)

type ArgsParser struct {
	varnames             map[string]string
	possibleNArgsChecker map[string]func(uint) bool
	functionMapper       map[string]func(...string)
}

func newParserDefault() *ArgsParser {
	return &ArgsParser{make(map[string]string), make(map[string]func(uint) bool), make(map[string]func(...string))}
}

func (ap ArgsParser) checkEntity(entity argEntity) (bool, string) {
	for _, pseudonym := range entity.Pseudonyms {
		if refer, check := ap.varnames[pseudonym]; check {
			return false, fmt.Sprintf("Pseudonym %s already exists and refer to %s", pseudonym, refer)
		}
	}
	if _, check := ap.varnames[entity.ArgName]; check {
		return false, fmt.Sprintf("Keyname %s already exist", entity.ArgName)
	}
	if _, check := ap.possibleNArgsChecker[entity.ArgName]; check {
		return false, fmt.Sprintf("Critical error. Keyname %s already refer to check function, but somehow doesn't exist.", entity.ArgName)
	}
	if _, check := ap.functionMapper[entity.ArgName]; check {
		return false, fmt.Sprintf("Critical error. Keuname %s already refer to map function, but somehow doesn't exist and don't have check function", entity.ArgName)
	}
	return true, "Ok"
}

func (ap *ArgsParser) addEntity(entity argEntity) error {
	if check, msg := ap.checkEntity(entity); check {
		return errors.New("Entity collision due to: " + msg)
	}

	ap.possibleNArgsChecker[entity.ArgName] = entity.NargsChecker
	ap.functionMapper[entity.ArgName] = entity.AssociatedFunction
	ap.varnames[entity.ArgName] = entity.ArgName
	for _, element := range entity.Pseudonyms {
		ap.varnames[element] = entity.ArgName
	}

	return nil
}

func constructParser(entities ...argEntity) (*ArgsParser, error) {
	var ret = newParserDefault()
	for _, entity := range entities {
		err := ret.addEntity(entity)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
