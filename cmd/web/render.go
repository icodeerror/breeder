package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, data *templateData) {
	var tmpl *template.Template
	if app.config.useCache {
		if templateForMap, ok := app.templateMap[t]; ok {
			tmpl = templateForMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Panicln("Error building template:", err)
			return
		}

		log.Panicln("Building template from disk")

		tmpl = newTemplate
	}

	// if data == nil {
	// 	data = &templateData{}
	// }

	if err := tmpl.ExecuteTemplate(w, t, data); err != nil {
		log.Panicln("Executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.html",
		"./templates/partials/header.partial.html",
		"./templates/partials/footer.partial.html",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}
