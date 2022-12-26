package pages

import (
	"os"
	"text/template"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

const file = "./pages/index.html"

func loadHTML() string {
	b, err := os.ReadFile(file)
	if err != nil {
		panic("unable to load file: " + file)
	}

	b = bundle(b)
	html := string(b)

	return html
}

func bundle(b []byte) []byte {
	const mediaType = "text/html"

	m := minify.New()
	m.AddFunc(mediaType, html.Minify)

	b, err := m.Bytes(mediaType, b)
	if err != nil {
		panic("unable to minify: " + file)
	}

	return b
}

func LoadTemplate() *template.Template {
	html := loadHTML()

	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		panic("HTML cannot be parsed")
	}
	return tmpl
}
