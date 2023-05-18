package adapter

import (
	"context"
	"fmt"

	"terraform-provider-gitea/api"
)

func GetTeamByOrgAndName(ctx context.Context, client *api.APIClient, org, name string) (*api.Team, error) {
	res, _, err := client.OrganizationAPI.
		TeamSearch(ctx, org).
		Q(name).
		Execute()
	if err != nil {
		return nil, err
	}

	if len(res.Data) == 0 {
		return nil, fmt.Errorf("Team with name '%s' was not found.", name)
	}

	if len(res.Data) > 1 {
		return nil, fmt.Errorf("Multiple teams with name '%s' were found.", name)
	}

	return &res.Data[0], nil
}
