package models

import (
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
