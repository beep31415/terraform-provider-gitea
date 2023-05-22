package proxies

import (
	"context"
	"fmt"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"
	"terraform-provider-gitea/internal/proxy/errors"
	"terraform-provider-gitea/pkg/collections"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type TeamProxy struct {
	client    *api.APIClient
	converter converters.TeamConverter
}

func NewTeamProxy(client *api.APIClient) *TeamProxy {
	return &TeamProxy{
		client:    client,
		converter: converters.TeamConverter{},
	}
}

func (t *TeamProxy) FillDataSource(ctx context.Context, model models.TeamDataSourceModel) diag.Diagnostics {
	res, err := t.getByOrgAndName(ctx, model.Organization.ValueString(), model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea team.")
	}

	t.converter.ReadToDataSource(model, res)

	var members []string
	members, err = t.getMembers(ctx, res.GetId())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea team.")
	}

	return t.converter.ReadMembersToDataSource(ctx, model, members)
}

func (t *TeamProxy) FillResource(ctx context.Context, model models.TeamResourceModel) diag.Diagnostics {
	res, err := t.getByOrgAndName(ctx, model.Organization.ValueString(), model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea team.")
	}

	t.converter.ReadToResource(model, res)

	var members []string
	members, err = t.getMembers(ctx, res.GetId())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea team.")
	}

	return t.converter.ReadMembersToResource(ctx, model, members)
}

func (t *TeamProxy) Create(ctx context.Context, model models.TeamResourceModel) diag.Diagnostics {
	option := t.converter.ToApiAddTeamOptions(model)

	res, _, err := t.client.OrganizationAPI.
		OrgCreateTeam(ctx, model.Organization.ValueString()).
		Body(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Creating Gitea team.")
	}

	t.converter.ReadToResource(model, res)

	members, diags := t.converter.ToMembers(ctx, model.Members)
	if diags.HasError() {
		return diags
	}

	for _, user := range members {
		err := t.addMember(ctx, res.GetId(), user)
		if err != nil {
			diags.Append(errors.ToDiagnosticError(err, "Error Creating Gitea team."))
		}
	}

	return diags
}

func (t *TeamProxy) Update(ctx context.Context, model models.TeamResourceModel) diag.Diagnostics {
	option := t.converter.ToApiEditTeamOptions(model)

	res, _, err := t.client.OrganizationAPI.
		OrgEditTeam(ctx, int32(model.Id.ValueInt64())).
		Body(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Updating Gitea team.")
	}

	t.converter.ReadToResource(model, res)

	existingMemberList, err := t.getMembers(ctx, res.GetId())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Updating Gitea team.")
	}

	members, diags := t.converter.ToMembers(ctx, model.Members)
	if diags.HasError() {
		return diags
	}

	toAdd, toRemove := collections.GetChanges(existingMemberList, members)

	for _, member := range toRemove {
		err = t.removeMember(ctx, res.GetId(), member)
		if err != nil {
			diags.Append(errors.ToDiagnosticError(err, "Error Updating Gitea team."))
		}
	}

	for _, member := range toAdd {
		err = t.addMember(ctx, res.GetId(), member)
		if err != nil {
			diags.Append(errors.ToDiagnosticError(err, "Error Updating Gitea team."))
		}
	}

	return diags
}

func (t *TeamProxy) Delete(ctx context.Context, model models.TeamResourceModel) diag.Diagnostics {
	res, err := t.getById(ctx, model.Id.ValueInt64())
	if err != nil {
		if errors.IsErrorNotFound(err) {
			return nil
		}

		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea team.")
	}

	members, err := t.getMembers(ctx, res.GetId())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea team.")
	}

	var diags diag.Diagnostics
	for _, member := range members {
		err = t.removeMember(ctx, res.GetId(), member)
		if err != nil {
			diags.Append(errors.ToDiagnosticError(err, "Error Deleting Gitea team."))
		}
	}

	_, err = t.client.OrganizationAPI.
		OrgDeleteTeam(ctx, res.GetId()).
		Execute()
	if err != nil {
		diags.Append(errors.ToDiagnosticError(err, "Error Deleting Gitea team."))
	}

	return diags
}

func (t *TeamProxy) getById(ctx context.Context, id int64) (*api.Team, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgGetTeam(ctx, id).
		Execute()

	return res, err
}

func (t *TeamProxy) getByOrgAndName(ctx context.Context, org, name string) (*api.Team, error) {
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

	return &res.Data[0], err
}

func (t *TeamProxy) getMembers(ctx context.Context, id int64) ([]string, error) {
	res, _, err := t.client.OrganizationAPI.
		OrgListTeamMembers(ctx, id).
		Execute()
	if err != nil {
		return []string{}, err
	}

	members := []string{}
	for _, resMember := range res {
		members = append(members, resMember.GetLogin())
	}

	return members, err
}

func (t *TeamProxy) addMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgAddTeamMember(ctx, id, user).
		Execute()

	return err
}

func (t *TeamProxy) removeMember(ctx context.Context, id int64, user string) error {
	_, err := t.client.OrganizationAPI.
		OrgRemoveTeamMember(ctx, id, user).
		Execute()

	return err
}
