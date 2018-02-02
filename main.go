package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jazaret/go-web/controller"
	"github.com/jazaret/go-web/middleware"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", new(middleware.GzipMiddleware))
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	//Content Templates
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to read contents of blocks directiory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directiory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template. Error - " + err.Error())
		}
		result[fi.Name()] = tmpl
	}

	return result
}
