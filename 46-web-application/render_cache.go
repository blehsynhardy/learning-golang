package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

// RenderCache holds parsed templates in memory and controls
// whether to re-parse on every request (dev) or cache them (prod).
type RenderCache struct {
	cache       map[string]*template.Template
	isDev       bool
	mu          sync.RWMutex
	templateDir string
}

type templateData struct {
	Form            *Form
	IsAuthenticated bool
	Flash           string
}

// NewRenderCache creates a new RenderCache.
// Set isDev = true during development so templates reload on every request.
// Set isDev = false in production to cache parsed templates in memory.
func NewRenderCache(isDev bool, templateDir string) *RenderCache {
	return &RenderCache{
		cache:       make(map[string]*template.Template),
		isDev:       isDev,
		templateDir: templateDir,
	}
}

// Render fetches (or parses) a template and executes it with the given data,
// writing the result directly to the http.ResponseWriter.
func (rc *RenderCache) Render(w http.ResponseWriter, templateName string, data interface{}) error {
	templ, err := rc.getTemplate(templateName)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err = templ.ExecuteTemplate(w, "base.html", data); err != nil {
		// Headers are already sent at this point, so just log it.
		log.Printf("error executing template %q: %v", templateName, err)
		return err
	}

	return nil
}

// getTemplate returns a cached template in production or a freshly parsed
// template in development mode.
func (rc *RenderCache) getTemplate(templateName string) (*template.Template, error) {
	if rc.isDev {
		// Dev mode: always read from disk so changes are reflected immediately.
		return rc.parseTemplate(templateName)
	}

	// Production: try the read lock first (fast path for cache hits).
	rc.mu.RLock()
	if templ, ok := rc.cache[templateName]; ok {
		rc.mu.RUnlock()
		return templ, nil
	}
	rc.mu.RUnlock()

	// Cache miss: parse the template, then store it under a write lock.
	templ, err := rc.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	rc.mu.Lock()
	// Double-check: another goroutine may have stored it while we were parsing.
	if _, ok := rc.cache[templateName]; !ok {
		rc.cache[templateName] = templ
	}
	rc.mu.Unlock()

	return templ, nil
}

// parseTemplate builds a *template.Template from:
//  1. The requested template file
//  2. All files in <templateDir>/layouts/*.html
//  3. All files in <templateDir>/partials/*.html
//
// This lets templates reference shared layouts and partials via
// {{template "name" .}} without any extra wiring.
func (rc *RenderCache) parseTemplate(templateName string) (*template.Template, error) {
	files := []string{path.Join(rc.templateDir, templateName)}

	layouts, err := filepath.Glob(path.Join(rc.templateDir, "layouts/*.html"))
	if err != nil {
		return nil, err
	}
	files = append(files, layouts...)

	partials, err := filepath.Glob(path.Join(rc.templateDir, "partials/*.html"))
	if err != nil {
		return nil, err
	}
	files = append(files, partials...)

	return template.ParseFiles(files...)
}
