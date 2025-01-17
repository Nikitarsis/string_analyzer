package args

import (
	"errors"
	"fmt"
)

type ArgsParser struct {
	varnames             map[string]string
	possibleNArgsChecker map[string]func(uint) bool
	functionMapper       map[string]func(...string)
	checklist            map[string]bool
}

func newParserDefault() *ArgsParser {
	return &ArgsParser{make(map[string]string), make(map[string]func(uint) bool), make(map[string]func(...string)), make(map[string]bool)}
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
		return false, fmt.Sprintf("Critical error. Keyname %s already refer to map function, but somehow doesn't exist and don't have check function", entity.ArgName)
	}
	if _, check := ap.checklist[entity.ArgName]; check {
		return false, fmt.Sprintf("Critical error. Keyname %s already contains in checklist, but somehow doesn't exist", entity.ArgName)
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
	ap.checklist[entity.ArgName] = entity.IsNecessary

	return nil
}

func (ap ArgsParser) Parse(argsMap map[string][]string) error {
	for key, isNecessary := range ap.checklist {
		_, check := argsMap[key]
		if isNecessary && !check {
			return errors.New(fmt.Sprintf("Should be key %s", key))
		}
	}
	for key, values := range argsMap {
		pseudonym, check := ap.varnames[key]
		if !check {
			return errors.New(fmt.Sprintf("Key %s doesn't exist", key))
		}
		checker, check := ap.possibleNArgsChecker[pseudonym]
		if !check {
			return errors.New(fmt.Sprintf("No checker for key %s", pseudonym))
		}
		function, check := ap.functionMapper[pseudonym]
		size := uint(len(values))
		if !checker(size) {
			return errors.New(fmt.Sprintf("Incorrect number of arguments %s", key))
		}
		function(values...)
	}
	return nil
}

func (ap ArgsParser) ParseArgs(args ...string) error {
	argsMap, error := DivideArgs(args...)
	if error != nil {
		return error
	}
	return ap.Parse(argsMap)
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
