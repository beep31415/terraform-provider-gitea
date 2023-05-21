package models

import (
	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationResourceModel struct {
	ID                    types.Int64  `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	FullName              types.String `tfsdk:"full_name"`
	Description           types.String `tfsdk:"description"`
	Website               types.String `tfsdk:"website"`
	Location              types.String `tfsdk:"location"`
	Visibility            types.String `tfsdk:"visibility"`
	AdminChangeTeamAccess types.Bool   `tfsdk:"repo_admin_change_team_access"`
}

func (o *OrganizationResourceModel) ReadFromApi(organization *api.Organization) {
	o.ID = types.Int64Value(organization.GetId())
	o.Name = types.StringValue(organization.GetName())
	o.FullName = types.StringValue(organization.GetFullName())
	o.Description = types.StringValue(organization.GetDescription())
	o.Website = types.StringValue(organization.GetWebsite())
	o.Location = types.StringValue(organization.GetLocation())
	o.Visibility = types.StringValue(organization.GetVisibility())
	o.AdminChangeTeamAccess = types.BoolValue(organization.GetRepoAdminChangeTeamAccess())
}

func (o *OrganizationResourceModel) ToApiAddOrganizationOptions() api.CreateOrgOption {
	return api.CreateOrgOption{
		Username:                  o.Name.ValueString(),
		FullName:                  o.FullName.ValueStringPointer(),
		Description:               o.Description.ValueStringPointer(),
		Website:                   o.Website.ValueStringPointer(),
		Location:                  o.Location.ValueStringPointer(),
		Visibility:                o.Visibility.ValueStringPointer(),
		RepoAdminChangeTeamAccess: o.AdminChangeTeamAccess.ValueBoolPointer(),
	}
}

func (o *OrganizationResourceModel) ToApiEditOrganizationOptions() adapters.EditOrganizationOptions {
	return adapters.EditOrganizationOptions{
		Name: o.Name.ValueString(),
		EditOption: api.EditOrgOption{
			FullName:                  o.FullName.ValueStringPointer(),
			Description:               o.Description.ValueStringPointer(),
			Website:                   o.Website.ValueStringPointer(),
			Location:                  o.Location.ValueStringPointer(),
			Visibility:                o.Visibility.ValueStringPointer(),
			RepoAdminChangeTeamAccess: o.AdminChangeTeamAccess.ValueBoolPointer(),
		},
	}
}
