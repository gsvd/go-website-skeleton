package handlers

import (
	"github.com/Gsvd/website/internal/template_engine"
)

type ViewHandler struct {
	TemplateEngine *template_engine.TemplateEngine
}

func NewViewHandler(te *template_engine.TemplateEngine) *ViewHandler {
	return &ViewHandler{TemplateEngine: te}
}
