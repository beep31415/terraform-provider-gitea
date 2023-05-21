package models

import (
	"context"
	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	allUnits = []string{"repo.code", "repo.issues", "repo.ext_issues", "repo.wiki",
		"repo.pulls", "repo.releases", "repo.projects", "repo.ext_wiki"}
)

type TeamResourceModel struct {
	Id                      types.Int64  `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Organization            types.String `tfsdk:"organization"`
	Permission              types.String `tfsdk:"permission"`
	Description             types.String `tfsdk:"description"`
	CanCreateOrgRepo        types.Bool   `tfsdk:"can_create_org_repo"`
	IncludesAllRepositories types.Bool   `tfsdk:"includes_all_repositories"`
	Members                 types.List   `tfsdk:"members"`
}

func (t *TeamResourceModel) ReadFromApi(ctx context.Context, team *adapters.Team) diag.Diagnostics {
	t.Id = types.Int64Value(team.Team.GetId())
	t.Name = types.StringValue(team.Team.GetName())
	t.Permission = types.StringValue(team.Team.GetPermission())
	t.Description = types.StringValue(team.Team.GetDescription())
	t.CanCreateOrgRepo = types.BoolValue(team.Team.GetCanCreateOrgRepo())
	t.IncludesAllRepositories = types.BoolValue(team.Team.GetIncludesAllRepositories())

	tfMembers, diags := types.ListValueFrom(ctx, types.StringType, team.Members)
	if diags.HasError() {
		return diags
	}
	t.Members = tfMembers

	return diags
}

func (t *TeamResourceModel) ToApiAddTeamOptions(ctx context.Context) (adapters.AddTeamOptions, diag.Diagnostics) {
	var memberList []string
	diags := t.Members.ElementsAs(ctx, &memberList, true)
	if diags.HasError() {
		return adapters.AddTeamOptions{}, diags
	}
	if memberList == nil {
		memberList = []string{}
	}

	return adapters.AddTeamOptions{
		Organization: t.Organization.ValueString(),
		AddOption: api.CreateTeamOption{
			Name:                    t.Name.ValueString(),
			Permission:              t.Permission.ValueStringPointer(),
			Description:             t.Description.ValueStringPointer(),
			CanCreateOrgRepo:        t.CanCreateOrgRepo.ValueBoolPointer(),
			IncludesAllRepositories: t.IncludesAllRepositories.ValueBoolPointer(),
			Units:                   allUnits,
		},
		Members: memberList,
	}, diags
}

func (t *TeamResourceModel) ToApiEditTeamOptions(ctx context.Context) (adapters.EditTeamOptions, diag.Diagnostics) {
	var memberList []string
	diags := t.Members.ElementsAs(ctx, &memberList, true)
	if diags.HasError() {
		return adapters.EditTeamOptions{}, diags
	}
	if memberList == nil {
		memberList = []string{}
	}

	return adapters.EditTeamOptions{
		Id: int32(t.Id.ValueInt64()),
		EditOption: api.EditTeamOption{
			Name:                    t.Name.ValueString(),
			Permission:              t.Permission.ValueStringPointer(),
			Description:             t.Description.ValueStringPointer(),
			CanCreateOrgRepo:        t.CanCreateOrgRepo.ValueBoolPointer(),
			IncludesAllRepositories: t.IncludesAllRepositories.ValueBoolPointer(),
			Units:                   allUnits,
		},
		Members: memberList,
	}, diags
}
