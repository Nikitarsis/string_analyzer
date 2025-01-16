package args

import (
	"errors"
)

type ArgsParser struct {
	varnames             map[string]string
	possibleNArgsChecker map[string]func(uint) bool
	functionMapper       map[string]func([]string)
}

func (ap ArgsParser) checkEntity(entity ArgEntity) bool {
	for _, pseudonym := range entity.Pseudonyms {
		if _, check := ap.varnames[pseudonym]; check {
			return false
		}
	}
	if _, check := ap.varnames[entity.ArgName]; check {
		return false
	}
	if _, check := ap.possibleNArgsChecker[entity.ArgName]; check {
		return false
	}
	if _, check := ap.functionMapper[entity.ArgName]; check {
		return false
	}
	return true
}

func (ap *ArgsParser) AddEntity(entity ArgEntity) error {
	if !ap.checkEntity(entity) {
		return errors.New("Entity collision")
	}

	ap.possibleNArgsChecker[entity.ArgName] = entity.NargsChecker
	ap.functionMapper[entity.ArgName] = entity.AssociatedFunction
	ap.varnames[entity.ArgName] = entity.ArgName
	for _, element := range entity.Pseudonyms {
		ap.varnames[element] = entity.ArgName
	}

	return nil
}

func ConstructParser(entities ...ArgEntity) ArgsParser {
	var ret = ArgsParser{make(map[string]string), make(map[string]func(uint) bool), make(map[string]func([]string))}
	for _, entity := range entities {
		ret.AddEntity(entity)
	}
	return ret
}
