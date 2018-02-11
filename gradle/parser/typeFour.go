package parser

import (
	"regexp"
	"errors"
)

var (
	typeFourRegex = regexp.MustCompile("^providedCompile group:\\s[\"']([a-z.\\d]*)[\"'],\\sname:\\s[\"']([a-z-\\d]*)[\"'],\\sversion:\\s[\"']([\\da-zA-Z.]*)[\"']")
)

type TypeFour struct {
}

func (t *TypeFour) match(line string) bool {
	return typeFourRegex.MatchString(line)
}

func (t *TypeFour) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeFourRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type Four does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
