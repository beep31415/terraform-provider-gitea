package adapters

import (
	"context"
	"fmt"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/pkg/arrays"
)

type TeamAdapter struct {
	client *api.APIClient
}

type AddTeamOptions struct {
	Organization string
	AddOption    api.CreateTeamOption
	Members      []string
}

type EditTeamOptions struct {
	Id         int32
	EditOption api.EditTeamOption
	Members    []string
}

type Team struct {
	Team    *api.Team
	Members []string
}

func NewTeamAdapter(client *api.APIClient) *TeamAdapter {
	return &TeamAdapter{
		client: client,
	}
}

func (t *TeamAdapter) GetById(ctx context.Context, id int64) (*Team, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgGetTeam(ctx, id).
		Execute()
	if err != nil {
		return nil, err
	}

	members, err := t.getMembers(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Team{
		Team:    res,
		Members: members,
	}, err
}

func (t *TeamAdapter) GetByOrgAndName(ctx context.Context, org, name string) (*Team, error) {
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

	members, err := t.getMembers(ctx, res.Data[0].GetId())
	if err != nil {
		return nil, err
	}

	return &Team{
		Team:    &res.Data[0],
		Members: members,
	}, err
}

func (t *TeamAdapter) Create(ctx context.Context, options AddTeamOptions) (*Team, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgCreateTeam(ctx, options.Organization).
		Body(options.AddOption).
		Execute()
	if err != nil {
		return nil, err
	}

	for _, user := range options.Members {
		err := t.addMember(ctx, res.GetId(), user)
		if err != nil {
			return nil, err
		}
	}

	return &Team{
		Team:    res,
		Members: options.Members,
	}, err
}

func (t *TeamAdapter) Edit(ctx context.Context, options EditTeamOptions) (*Team, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgEditTeam(ctx, options.Id).
		Body(options.EditOption).
		Execute()
	if err != nil {
		return nil, err
	}

	existingMemberList, err := t.getMembers(ctx, int64(options.Id))
	if err != nil {
		return nil, err
	}

	for _, member := range existingMemberList {
		if !arrays.Contains(member, options.Members...) {
			err = t.removeMember(ctx, int64(options.Id), member)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, member := range options.Members {
		if !arrays.Contains(member, existingMemberList...) {
			err = t.addMember(ctx, int64(options.Id), member)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Team{
		Team:    res,
		Members: options.Members,
	}, err
}

func (t *TeamAdapter) Delete(ctx context.Context, id int64) error {
	_, err := t.GetById(ctx, id)
	if err != nil {
		return err
	}

	members, err := t.getMembers(ctx, id)
	if err != nil {
		return err
	}

	for _, member := range members {
		err = t.removeMember(ctx, id, member)
		if err != nil {
			return err
		}
	}

	_, err = t.client.OrganizationAPI.
		OrgDeleteTeam(ctx, id).
		Execute()

	return err
}

func (t *TeamAdapter) getMembers(ctx context.Context, id int64) ([]string, error) {
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

func (t *TeamAdapter) addMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgAddTeamMember(ctx, id, user).
		Execute()

	return err
}

func (t *TeamAdapter) removeMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgRemoveTeamMember(ctx, id, user).
		Execute()

	return err
}
