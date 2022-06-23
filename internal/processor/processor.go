package processor

import (
	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/pkg/models"
)

type ModuleAccessListProcessor struct {
}

type Processor interface {
	Process(request models.Request) (models.Response, error)
}

func (m ModuleAccessListProcessor) Process(request models.Request) (models.Response, error) {

}
