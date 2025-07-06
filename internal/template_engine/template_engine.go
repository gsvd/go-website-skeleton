package template_engine

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
	"github.com/pkg/errors"
)

type (
	TemplateEngine struct{}
	Option         func(*TemplateEngine) error
)

func New() *TemplateEngine {
	return &TemplateEngine{}
}

func (te *TemplateEngine) RenderView(w http.ResponseWriter, layout layout.Layout, view view.View, data any) error {
	layoutFile := fmt.Sprintf("web/src/views/layouts/%s.html", layout)
	viewFile := fmt.Sprintf("web/src/views/templates/%s.html", view)

	partialFiles, err := filepath.Glob("web/src/views/partials/*.html")
	if err != nil {
		return errors.Wrap(err, "filepath.Glob")
	}

	files := append([]string{layoutFile, viewFile}, partialFiles...)

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return errors.Wrap(err, "template.ParseFiles")
	}

	if err := tmpl.ExecuteTemplate(w, string(layout), data); err != nil {
		return errors.Wrap(err, "tmpl.ExecuteTemplate")
	}

	return nil
}
