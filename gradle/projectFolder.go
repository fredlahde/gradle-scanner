package gradle

import (
	"os"
	"sync"
	"github.com/MichaelTJones/walk"
	"strings"
)

const (
	GRADLEFILE = "build.gradle"
)

type ProjectsFolder struct {
	Path     string
	Projects map[string]Project
	mu       sync.RWMutex // guards projects
}

func NewProjectsFolder(basePath string) *ProjectsFolder {
	f := &ProjectsFolder{Path: basePath}
	f.Projects = make(map[string]Project)
	return f
}

func (g *ProjectsFolder) FindGradleProjects() {
	walk.Walk(g.Path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || info.Name() != GRADLEFILE || strings.Contains(path, ".gradle/wrapper") {
			return nil
		}
		g.mu.Lock()
		if _, ok := g.Projects[path]; !ok {
			g.Projects[path] = newProject(path)
		}
		g.mu.Unlock()
		return nil
	})
}
