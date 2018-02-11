package parser

import (
	"fmt"
)

type Dependency struct {
	Group    string
	Artifact string
	Version  string
}

func (d *Dependency) Empty() bool {
	return d.Group == ""
}

func (d *Dependency) String() string {
	return fmt.Sprintf("group: %s artifact: %s version: %s", d.Group, d.Artifact, d.Version)
}

type Parser interface {
	match(line string) bool
	build(line string) (Dependency, error)
}

type NoParserFoundError struct {
	Line string
}

func (n *NoParserFoundError) Error() string {
	return fmt.Sprintf("line did not match any parsers: %s", n.Line)
}

var availableParser []Parser = make([]Parser, 0)

func init() {
	availableParser = append(availableParser, &TypeOne{})
	availableParser = append(availableParser, &TypeTwo{})
	availableParser = append(availableParser, &TypeThree{})
	availableParser = append(availableParser, &TypeFour{})
	availableParser = append(availableParser, &TypeFive{})
	availableParser = append(availableParser, &TypeSix{})
}

func ParseDeps(line string) (Dependency, error) {
	for _, parser := range availableParser {
		if parser.match(line) {
			return parser.build(line)
		}
	}

	return Dependency{}, &NoParserFoundError{Line: line}
}
