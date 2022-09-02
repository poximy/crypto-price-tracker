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

// FIXME panic if err
func LoadTemplate() (*template.Template, error) {
	html, _ := loadHTML()
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return &template.Template{}, errors.New("unable to load page")
	}
	return tmpl, nil
}
