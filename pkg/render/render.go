package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/myqly/bookings/pkg/config"
	"github.com/myqly/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, templateName string, td *models.TemplateData) {
	// create a template cache
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	var err error

	// get requested template from cache
	cachedTemplate, ok := templateCache[templateName]
	if !ok {
		log.Fatal(err)
	}

	// creating this buffer but I do not know why yet.
	// supposedly gives finer grained error checking
	// This allows you to test for an error with the value stored in the map itself
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err = cachedTemplate.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cacheMap := map[string]*template.Template{}

	// for a template, you should start with the template you want to render and then the layouts/partials afterward.
	// should start with *.page.tmpl first

	// get all of the files named *.page.tmpl

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cacheMap, err
	}

	// range through all files ending with .page.tmpl
	for _, page := range pages {
		templateFileName := filepath.Base(page)
		templateSet, err := template.New(templateFileName).ParseFiles(page)
		if err != nil {
			return cacheMap, err
		}

		// find the layouts

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cacheMap, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cacheMap, err
			}
		}

		cacheMap[templateFileName] = templateSet

	}

	return cacheMap, nil

}
