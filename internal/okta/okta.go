package okta

import (
	"context"
	"fmt"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

type OktaConnect struct{}

func (o OktaConnect) GetGroups(email string) ([]*okta.Group, *okta.Response) {
	ctx, client, err := okta.NewClient(
		context.TODO(),
		okta.WithOrgUrl("https://dev-75708788.okta.com"),
		okta.WithToken("00gcPFrj49lmSC8xEnrDBK3OgaHno0HqdgbV0kzDQB"),
	)

	if err != nil {
		panic(fmt.Sprintf("Error: %v\n", err))
	}

	groups, resp, err := client.User.ListUserGroups(ctx, email)

	if err != nil {
		panic(fmt.Sprintf("Error Getting User: %v\n", err))
	}

	return groups, resp
}
