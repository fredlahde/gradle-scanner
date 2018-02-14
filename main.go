package main

import (
	"path/filepath"
	"k8s.io/client-go/util/homedir"
	"fredlahde.com/gradle-version-checker/gradle"
	"log"
	"fmt"
	"sync"
)

var (
	basePath = filepath.Join(homedir.HomeDir() + "/repos/java")
)

func main() {
	ph := gradle.NewProjectsFolder(basePath)
	ph.FindGradleProjects();
	var wg sync.WaitGroup
	wg.Add(len(ph.Projects))
	for _, p := range ph.Projects {
		go func(prj gradle.Project, wg *sync.WaitGroup) {
			if err := prj.GradleFile.FindDeps(); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(p, &wg)

	}
	wg.Wait()
	printDeps(ph)
}
func printDeps(ph *gradle.ProjectsFolder) {
	counter := 0
	for _, p := range ph.Projects {
		fmt.Printf("Project: %s\n", p.Name)
		for _, d := range p.GradleFile.Deps {
			if !d.Empty() {
				counter++
				fmt.Printf("%s\n", d.String())
			}
		}
		fmt.Println()
	}

	log.Printf("Found %d deps.", counter)
}
