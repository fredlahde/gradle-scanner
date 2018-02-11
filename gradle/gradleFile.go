package gradle

import (
	"strings"
	"io/ioutil"
	"fredlahde.com/gradle-version-checker/gradle/parser"
	"fmt"
	"regexp"
)

type File struct {
	Path   string
	Deps   []parser.Dependency
	errBag []error
}

var depBlockRegex = regexp.MustCompile("^dependencies {$")

func (f *File) FindDeps() error {
	file, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(file), "\n")
	depthInDepBlock := -1;
	isInDepBlock := func() bool {
		return depthInDepBlock >= 0
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if depBlockRegex.MatchString(line) && !isInDepBlock() {
			depthInDepBlock = 0
			continue
		}
		if isInDepBlock() && strings.Contains(line, "{") {
			depthInDepBlock++
			continue
		} else if isInDepBlock() && strings.Contains(line, "}") {
			depthInDepBlock--
			continue
		}

		if !isInDepBlock() {
			continue
		}

		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		dep, err := parser.ParseDeps(line)
		if err != nil {
			if ne, ok := err.(*parser.NoParserFoundError); ok {
				f.errBag = append(f.errBag, fmt.Errorf("file %s has invalid dep: %s", f.Path, ne.Line))
			} else {
				f.errBag = append(f.errBag, err)
			}
		}
		f.Deps = append(f.Deps, dep)
	}

	if len(f.errBag) > 0 {
		return f.errBag[len(f.errBag)-1]
	}

	return nil
}
