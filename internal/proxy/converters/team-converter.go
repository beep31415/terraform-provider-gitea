package converters

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	allUnits = []string{"repo.code", "repo.issues", "repo.ext_issues", "repo.wiki",
		"repo.pulls", "repo.releases", "repo.projects", "repo.ext_wiki"}
)

type TeamConverter struct{}

func (TeamConverter) ReadToDataSource(model *models.TeamDataSourceModel, team *api.Team) {
	model.Id = types.Int64Value(team.GetId())
	model.Name = types.StringValue(team.GetName())
	model.Permission = types.StringValue(team.GetPermission())
	model.Description = types.StringValue(team.GetDescription())
	model.CanCreateOrgRepo = types.BoolValue(team.GetCanCreateOrgRepo())
	model.IncludesAllRepositories = types.BoolValue(team.GetIncludesAllRepositories())
}

func (TeamConverter) ReadMembersToDataSource(ctx context.Context, model *models.TeamDataSourceModel, members []string) diag.Diagnostics {
	tfMembers, diags := types.ListValueFrom(ctx, types.StringType, members)
	if diags.HasError() {
		return diags
	}

	model.Members = tfMembers

	return nil
}

func (TeamConverter) ReadToResource(model *models.TeamResourceModel, team *api.Team) {
	model.Id = types.Int64Value(team.GetId())
	model.Name = types.StringValue(team.GetName())
	model.Permission = types.StringValue(team.GetPermission())
	model.Description = types.StringValue(team.GetDescription())
	model.CanCreateOrgRepo = types.BoolValue(team.GetCanCreateOrgRepo())
	model.IncludesAllRepositories = types.BoolValue(team.GetIncludesAllRepositories())
}

func (TeamConverter) ReadMembersToResource(ctx context.Context, model *models.TeamResourceModel, members []string) diag.Diagnostics {
	tfMembers, diags := types.ListValueFrom(ctx, types.StringType, members)
	if diags.HasError() {
		return diags
	}

	model.Members = tfMembers

	return nil
}

func (TeamConverter) ToApiAddTeamOptions(model models.TeamResourceModel) api.CreateTeamOption {
	return api.CreateTeamOption{
		Name:                    model.Name.ValueString(),
		Permission:              model.Permission.ValueStringPointer(),
		Description:             model.Description.ValueStringPointer(),
		CanCreateOrgRepo:        model.CanCreateOrgRepo.ValueBoolPointer(),
		IncludesAllRepositories: model.IncludesAllRepositories.ValueBoolPointer(),
		Units:                   allUnits,
	}
}

func (TeamConverter) ToApiEditTeamOptions(model models.TeamResourceModel) api.EditTeamOption {
	return api.EditTeamOption{
		Name:                    model.Name.ValueString(),
		Permission:              model.Permission.ValueStringPointer(),
		Description:             model.Description.ValueStringPointer(),
		CanCreateOrgRepo:        model.CanCreateOrgRepo.ValueBoolPointer(),
		IncludesAllRepositories: model.IncludesAllRepositories.ValueBoolPointer(),
		Units:                   allUnits,
	}
}

func (TeamConverter) ToMembers(ctx context.Context, tfMembers types.List) ([]string, diag.Diagnostics) {
	var memberList []string
	diags := tfMembers.ElementsAs(ctx, &memberList, true)
	if diags.HasError() {
		return []string{}, diags
	}

	if memberList == nil {
		memberList = []string{}
	}

	return memberList, nil
}
