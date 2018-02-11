package gradle

import "path/filepath"

type Project struct {
	Name       string
	Path       string
	GradleFile *File
}

func newProject(pathToGradleFile string) Project {
	dir := filepath.Dir(pathToGradleFile)
	name := filepath.Base(dir)
	return Project{
		Name:       name,
		Path:       dir,
		GradleFile: &File{Path: pathToGradleFile},
	}
}
