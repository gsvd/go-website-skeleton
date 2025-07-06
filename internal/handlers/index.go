package handlers

import (
	"net/http"

	"github.com/Gsvd/website/internal/template_engine/layout"
	"github.com/Gsvd/website/internal/template_engine/view"
)

func (vh *ViewHandler) ShowIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":   "go-website-skeleton",
		"Message": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Labore perspiciatis dicta autem laboriosam? Incidunt consequuntur nostrum aliquam illum, iste distinctio consequatur quod facere officia ipsa ducimus similique, rerum nemo est.",
	}

	if err := vh.TemplateEngine.RenderView(w, layout.Base, view.Index, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
