package parser

import (
	"regexp"
	"errors"
)

var (
	typeOneRegex = regexp.MustCompile("^compile group:\\s[\"']([a-z.\\d]*)[\"'],\\sname:\\s[\"']([\\da-z-.]*)[\"'],\\sversion:\\s[\"']([\\da-zA-Z.]*)[\"']")
)

type TypeOne struct {
}

func (t *TypeOne) match(line string) bool {
	return typeOneRegex.MatchString(line)
}

func (t *TypeOne) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeOneRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type one does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
