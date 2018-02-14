package parser

import "fmt"

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
