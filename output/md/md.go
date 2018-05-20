package md

import (
	"fmt"
	"github.com/k1LoW/tbls/schema"
	"os"
	"path/filepath"
	"text/template"
)

// Output generate markdown files.
func Output(s *schema.Schema, path string, force bool) error {
	if !force && outputExists(s, path) {
		return fmt.Errorf("Error: %s", "output files already exists.")
	}
	// README.md
	file, err := os.Create(filepath.Join(path, "README.md"))
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles("output/md/index.md.tmpl"))
	err = tmpl.Execute(file, map[string]interface{}{
		"Schema": s,
	})
	if err != nil {
		return err
	}
	// tables
	for _, t := range s.Tables {
		file, err := os.Create(filepath.Join(path, fmt.Sprintf("%s.md", t.Name)))
		if err != nil {
			return err
		}
		tmpl := template.Must(template.ParseFiles("output/md/table.md.tmpl"))
		err = tmpl.Execute(file, map[string]interface{}{
			"Table": t,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func outputExists(s *schema.Schema, path string) bool {
	// README.md
	if _, err := os.Lstat(filepath.Join(path, "README.md")); err == nil {
		return true
	}
	// tables
	for _, t := range s.Tables {
		if _, err := os.Lstat(filepath.Join(path, fmt.Sprintf("%s.md", t.Name))); err == nil {
			return true
		}
	}
	return false
}
