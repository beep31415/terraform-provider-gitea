package models

import (
	"context"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamDataSourceModel struct {
	Id                      types.Int64  `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Organization            types.String `tfsdk:"organization"`
	Permission              types.String `tfsdk:"permission"`
	Description             types.String `tfsdk:"description"`
	CanCreateOrgRepo        types.Bool   `tfsdk:"can_create_org_repo"`
	IncludesAllRepositories types.Bool   `tfsdk:"includes_all_repositories"`
	Units                   types.Map    `tfsdk:"units"`
	Members                 types.List   `tfsdk:"members"`
}

func NewTeamDataSource(ctx context.Context, team *adapters.Team) (TeamDataSourceModel, diag.Diagnostics) {
	members, diags := types.ListValueFrom(ctx, types.StringType, team.Members)
	if diags.HasError() {
		return TeamDataSourceModel{}, diags
	}

	return TeamDataSourceModel{
		Id:                      types.Int64Value(team.Team.GetId()),
		Name:                    types.StringValue(team.Team.GetName()),
		Organization:            types.StringValue(*team.Team.GetOrganization().Name),
		Permission:              types.StringValue(team.Team.GetPermission()),
		Description:             types.StringValue(team.Team.GetDescription()),
		CanCreateOrgRepo:        types.BoolValue(team.Team.GetCanCreateOrgRepo()),
		IncludesAllRepositories: types.BoolValue(team.Team.GetIncludesAllRepositories()),
		Members:                 members,
	}, nil
}
