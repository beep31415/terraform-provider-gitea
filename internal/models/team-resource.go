package models

import (
	"context"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (t *TeamResourceModel) ReadFrom(ctx context.Context, team *adapters.Team) diag.Diagnostics {
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

	return nil
}
