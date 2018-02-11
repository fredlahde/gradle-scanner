package parser

import (
	"regexp"
	"errors"
)

var (
	typeSixRegex = regexp.MustCompile("^compileOnly\\s[\"']([\\da-z.-]+):([\\da-z-.]+):([\\da-z.]+)[\"']")
)

type TypeSix struct {
}

func (t *TypeSix) match(line string) bool {
	return typeSixRegex.MatchString(line)
}

func (t *TypeSix) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeSixRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type Six does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
