package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	cache map[string]*template.Template
	mutex sync.RWMutex
	devMode bool
	templateDir string
}

func NewTemplateRenderer(templateDir string, isDevMode bool) *TemplateRenderer {
	return &TemplateRenderer{
		cache: make(map[string]*template.Template),
		devMode: isDevMode,
		templateDir: templateDir,
	}
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateRenderer) getTemplate(templateName string) (*template.Template, error) {
	if !t.devMode {
		t.mutex.RLock()
		if tmpl, ok := t.cache[templateName]; ok {
			t.mutex.RUnlock()
			return tmpl, nil
		}
		t.mutex.RUnlock()
	}

	tmpl, err := t.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	if !t.devMode {
		t.mutex.Lock()
		t.cache[templateName] = tmpl
		t.mutex.Unlock()
	}
	
	return tmpl, nil
}

func (t *TemplateRenderer) parseTemplate(templateName string) (*template.Template, error) {
	templatePath := filepath.Join(t.templateDir, templateName)

	files := []string{templatePath}

	layoutPath := filepath.Join(t.templateDir, "layout.html")
	if _, err := os.Stat(layoutPath); err == nil {
		files = append(files, layoutPath)
	}

	partialPath := filepath.Join(t.templateDir, "partial.html")
	if _, err := os.Stat(partialPath); err == nil {
		files = append(files, partialPath)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}