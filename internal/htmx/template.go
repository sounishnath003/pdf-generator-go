package htmx

import (
	"io"

	"text/template"

	"github.com/labstack/echo"
)

// Template represents a template renderer for Echo framework.
type Template struct {
	Templates *template.Template
}

// Render renders a template with the given name and data.
func (t *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

// NewTemplateRender initializes a new template renderer for Echo.
func NewTemplateRender(e *echo.Echo, paths ...string) {
	tmpl := &template.Template{}

	// Process the variadic passes in the arguments paths
	for _, path := range paths {
		template.Must(tmpl.Funcs(funcMap).ParseGlob(path))
	}

	t := newTemplate(tmpl)
	e.Renderer = t
}

// newTemplate creates a new Template instance.
func newTemplate(tmpl *template.Template) echo.Renderer {
	return &Template{
		Templates: tmpl,
	}
}

// ParseAndRetrieveTemplFile parses and retrieves a template file with custom functions.
func ParseAndRetrieveTemplFile(filePath string) *template.Template {
	tmpl := &template.Template{}
	return template.Must(tmpl.Funcs(funcMap).ParseFiles(filePath))
}

var funcMap = template.FuncMap{"mul": Mul, "sum": Sum}

func Mul(val1, val2 int) int {
	return val1 * val2
}

func Sum(vals ...int) int {
	t := 0
	for e := range vals {
		t += e
	}
	return t
}

// Converts following examples
//
// 123109 -> 1,23,109.00
//
// 1209.98 -> 1,209.98
//
// 121 -> 121.00
//
// use currency https://pkg.go.dev/golang.org/x/text/currency
func numberToINRCurrency(val interface{}) string {
	// TODO:
	return ""
}
