package parser

import (
	"fmt"
)

var (
	availableParser = make([]Parser, 0)
)

func init() {
	availableParser = append(availableParser, &jsonConfiguredParser{})
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

func Parse(line string) (Dependency, error) {
	for _, parser := range availableParser {
		if parser.match(line) {
			return parser.build(line)
		}
	}

	return Dependency{}, &NoParserFoundError{Line: line}
}
