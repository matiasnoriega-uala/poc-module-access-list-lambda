package handler

import (
	"context"

	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/internal/processor"
	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/pkg/models"
)

type ModuleAccessListHandler struct {
	Processor processor.Processor
}

func (h *ModuleAccessListHandler) SetProcessor() {
	if h.Processor == nil {
		h.Processor = processor.ModuleAccessListProcessor{}
	}
}

func (h ModuleAccessListHandler) HandleRequest(ctx context.Context, request models.Request) (models.Response, error) {
	h.SetProcessor()
	output, err := h.Processor.Process(request)
	if err != nil {
		return models.Response{}, err
	}
	return output, nil
}
