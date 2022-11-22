package pages

import (
	"errors"
	"os"
	"text/template"
)

func loadHTML() (string, error) {
	file := "./public/index.html"

	b, err := os.ReadFile(file)
	if err != nil {
		return "", errors.New("unable to load page")
	}

	html := string(b)
	return html, nil
}

func LoadTemplate() (*template.Template, error) {
	html, _ := loadHTML()
	tmpl, _ := template.New("index").Parse(html)
	return tmpl, nil
}
