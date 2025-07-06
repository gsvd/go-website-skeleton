package handlers

import (
	"net/http"

	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
)

func (vh *ViewHandler) ShowHelloWorld(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "go-website-skeleton: Hello World",
	}

	if err := vh.TemplateEngine.RenderView(w, layout.Base, view.HelloWorld, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
