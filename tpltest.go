package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Expected %s <templateDir> <pageName>", os.Args[0])
		os.Exit(1)
	}
	templateDir := os.Args[1]
	name := os.Args[2]

	templates, err := loadTemplates(templateDir)
	if err != nil {
		panic(err)
	}

	if err = renderTemplate(templates, name); err != nil {
		panic(err)
	}
}

// Load templates on program initialisation
func loadTemplates(templatesDir string) (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}

	base := filepath.Join(templatesDir, "base.tpl")

	includes, err := filepath.Glob(filepath.Join(templatesDir, "pages", "*.tpl"))
	if err != nil {
		return nil, err
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, p := range includes {
		files := []string{p, base}
		templates[filepath.Base(p)] = template.Must(template.ParseFiles(files...))
	}
	return templates, nil
}

func renderTemplate(templates map[string]*template.Template, name string) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.  See %+v.", name, templates)
	}
	return tmpl.ExecuteTemplate(os.Stdout, "base.tpl", nil)
}
