package processor

import (
	"encoding/json"
	"fmt"

	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/internal/okta"
	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/pkg/models"
)

type ModuleAccessListProcessor struct {
}

type Processor interface {
	Process(request models.Request) (models.Response, error)
}

func (m ModuleAccessListProcessor) Process(request models.Request) (models.Response, error) {
	var response models.Response
	o := okta.OktaConnect{}
	user, resp := o.GetUser(request.UserEmail)
	fmt.Printf("User: %+v\n Response: %+v\n\n", user, resp)

	json.Unmarshal([]byte(`{"groups": ["operadorMEP","aprobadorMEP"]}`), &response)

	return response, nil
}
