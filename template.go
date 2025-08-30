//go:build ignore
// +build ignore

//go:generate echo Setting up your files...
//go:generate go mod init {{ .GitHubRepository }}
//go:generate go run template.go --repo={{ .GitHubRepository }}

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	modulePath := flag.String("repo", "default", "Repo to use in generation")
	flag.Parse()

	projectName := filepath.Base(*modulePath)

	// read template for README.md
	t, err := template.ParseFiles("README.md.tpl")
	if err != nil {
		log.Fatalf("Failed to parse template file: %v", err)
	}

	// open target output file
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// execute template and write to file
	fmt.Printf("Generating README.md")
	t.Execute(f, struct {
		Project    string
		Repository string
	}{
		Project:    projectName,
		Repository: *modulePath,
	})
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
}
