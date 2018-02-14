package parser

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"sync"
	"regexp"
	"errors"
)

type typeDefinition struct {
	Regex    string `json:"regex"`
	compiled *regexp.Regexp
	mu       sync.RWMutex
}

func (t *typeDefinition) compile() {
	t.mu.Lock()
	t.compiled = regexp.MustCompile(t.Regex)
	t.mu.Unlock()
}

func (t *typeDefinition) build(line string) (Dependency, error) {
	if t.compiled == nil {
		t.compile()
	}

	dep := Dependency{}
	matches := t.compiled.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return dep, errors.New("JsonMatcher did not work") // TODO better error
	}
	dep.Group = matches[0][1]
	dep.Artifact = matches[0][2]
	dep.Version = matches[0][3]
	return dep, nil

}

func (t *typeDefinition) match(line string) bool {
	if t.compiled == nil {
		t.compile()
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.compiled.MatchString(line)
}

type jsonConfiguredParser struct {
	ready bool
	types []typeDefinition
	mu    sync.RWMutex
}

func (j *jsonConfiguredParser) initTypes() {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.ready {
		return // failSafe, if go routine has waited until here, but already fully ready
	}

	f, err := ioutil.ReadFile("types.json")
	if err != nil {
		log.Fatal(err)
	}

	t := make([]typeDefinition, 0)

	if err = json.Unmarshal(f, &t); err != nil {
		log.Fatal(err)
	}

	j.ready = true
	j.types = t
}

func (j *jsonConfiguredParser) match(line string) bool {
	if !j.ready {
		j.initTypes()
	}
	j.mu.RLock()
	defer j.mu.RUnlock()

	for _, t := range j.types {
		if t.match(line) {
			return true
		}
	}

	return false
}

func (j *jsonConfiguredParser) build(line string) (Dependency, error) {
	if !j.ready {
		j.initTypes()
	}
	j.mu.RLock()
	defer j.mu.RUnlock()

	for _, t := range j.types {
		if t.match(line) {
			return t.build(line)
		}
	}

	return Dependency{}, errors.New("JsonMatcher did not work") // TODO better error
}
