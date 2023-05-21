package adapters

import (
	"context"
	"fmt"

	"terraform-provider-gitea/api"
)

type TeamAdapter struct {
	client *api.APIClient
}

func NewTeamAdapter(client *api.APIClient) *TeamAdapter {
	return &TeamAdapter{
		client: client,
	}
}

func (t *TeamAdapter) GetTeamById(ctx context.Context, id int64) (*api.Team, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgGetTeam(ctx, id).
		Execute()

	return res, err
}

func (t *TeamAdapter) GetTeamByOrgAndName(ctx context.Context, org, name string) (*api.Team, error) {
	res, _, err := t.client.OrganizationAPI.
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

func (t *TeamAdapter) DeleteTeam(ctx context.Context, id int64) error {
	_, err := t.client.OrganizationAPI.
		OrgDeleteTeam(ctx, id).
		Execute()

	return err
}

func (t *TeamAdapter) GetTeamMembers(ctx context.Context, id int64) ([]string, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgListTeamMembers(ctx, id).
		Execute()
	if err != nil {
		return []string{}, err
	}

	users := []string{}
	for _, resUser := range res {
		users = append(users, resUser.GetLogin())
	}

	return users, err
}

func (t *TeamAdapter) AddTeamMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgAddTeamMember(ctx, id, user).
		Execute()

	return err
}

func (t *TeamAdapter) RemoveTeamMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgRemoveTeamMember(ctx, id, user).
		Execute()

	return err
}
