package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/dudad/bookings/pkg/config"
	"github.com/dudad/bookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, fileName string, data *models.TemplateData) {
	templateCache := app.TemplateCache
	t, ok := templateCache[fileName]
	if !ok {
		log.Fatal("Could not find template")
	}
	_ = t.Execute(w, data)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
