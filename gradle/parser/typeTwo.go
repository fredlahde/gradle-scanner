package parser

import (
	"regexp"
	"errors"
)

var (
	typeTwoRegex = regexp.MustCompile("^compile\\s[\"']([\\da-z.-]+):([\\da-z-.]+):([\\da-z.]+)[\"']")
)

type TypeTwo struct {
}

func (t *TypeTwo) match(line string) bool {
	return typeTwoRegex.MatchString(line)
}

func (t *TypeTwo) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeTwoRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type Two does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
