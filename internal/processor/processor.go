package processor

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

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
	oktaGroups, _ := o.GetGroups(request.UserEmail)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	// Creates Lambda response from the groups got from Okta
	for _, oktaGroup := range oktaGroups {
		// Initializate variables for groupResponse and modules slice
		groupResponse := models.GroupResponse{}

		// First, assign group name.
		groupResponse.Group = oktaGroup.Profile.Name

		// Get Modules associated to that group
		// Marshalling the GroupModules Lambda request
		groupModuleRequest := models.GroupModulesRequest{Group: string(groupResponse.Group)}

		payload, err := json.Marshal(groupModuleRequest)

		if err != nil {
			panic(fmt.Sprintf("Error marshalling GroupModulesRequest request: %v", err))
		}
		// Invoking GroupModulesLambda
		result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("arn:aws:lambda:us-east-1:234719978790:function:poc-group-modules-lambda"), Payload: payload})

		if err != nil {
			panic(fmt.Sprintf("Error calling poc-group-modules-lambda: %v", err))
		}

		// Unmarshal modules in response
		var resp models.SingleModuleResponse
		err = json.Unmarshal(result.Payload, &resp)

		if err != nil {
			panic(fmt.Sprintf("Error unmarshalling SingleModuleResponse request: %v", err))
		}

		// Append results to groupResponse.
		groupResponse.Modules = append(groupResponse.Modules, resp.Modules...)
		response.GroupsResponse = append(response.GroupsResponse, groupResponse)
	}

	return response, nil
}
