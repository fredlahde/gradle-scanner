package parser

import (
	"regexp"
	"errors"
)

var (
	typeFiveRegex = regexp.MustCompile("^classpath\\s[\"']([\\da-z.-]+):([\\da-z-.]+):([\\da-z.]+)[\"']")
)

type TypeFive struct {
}

func (t *TypeFive) match(line string) bool {
	return typeFiveRegex.MatchString(line)
}

func (t *TypeFive) build(line string) (Dependency, error) {
	dep := Dependency{}
	matches := typeFiveRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("type Five does not match")
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil
}
