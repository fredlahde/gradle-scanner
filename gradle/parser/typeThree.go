package parser

import (
	"regexp"
	"errors"
)

var (
	typeThreeRegex = regexp.MustCompile("^testCompile group:\\s[\"']([a-z.\\d]*)[\"'],\\sname:\\s[\"']([a-z-\\d]*)[\"'],\\sversion:\\s[\"']([\\da-zA-Z.]*)[\"']")
)

type TypeThree struct {
}

func (t *TypeThree) match(line string) bool {
	return typeThreeRegex.MatchString(line)
}

func (t *TypeThree) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeThreeRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type Three does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
