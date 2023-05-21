package models

import (
	"terraform-provider-gitea/api"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationDataSourceModel struct {
	ID                    types.Int64  `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	FullName              types.String `tfsdk:"full_name"`
	Description           types.String `tfsdk:"description"`
	Website               types.String `tfsdk:"website"`
	Location              types.String `tfsdk:"location"`
	Visibility            types.String `tfsdk:"visibility"`
	AdminChangeTeamAccess types.Bool   `tfsdk:"repo_admin_change_team_access"`
	AvatarURL             types.String `tfsdk:"avatar_url"`
}

func NewOrganizationDataSource(organization *api.Organization) OrganizationDataSourceModel {
	return OrganizationDataSourceModel{
		ID:                    types.Int64Value(organization.GetId()),
		Name:                  types.StringValue(organization.GetName()),
		FullName:              types.StringValue(organization.GetFullName()),
		Description:           types.StringValue(organization.GetDescription()),
		Website:               types.StringValue(organization.GetWebsite()),
		Location:              types.StringValue(organization.GetLocation()),
		Visibility:            types.StringValue(organization.GetVisibility()),
		AdminChangeTeamAccess: types.BoolValue(organization.GetRepoAdminChangeTeamAccess()),
		AvatarURL:             types.StringValue(organization.GetAvatarUrl()),
	}
}
