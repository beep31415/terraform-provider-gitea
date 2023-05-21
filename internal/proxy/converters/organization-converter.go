package converters

import (
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationConverter struct{}

func (OrganizationConverter) ReadToDataSource(model models.OrganizationDataSourceModel, organization *api.Organization) {
	model.ID = types.Int64Value(organization.GetId())
	model.Name = types.StringValue(organization.GetName())
	model.FullName = types.StringValue(organization.GetFullName())
	model.Description = types.StringValue(organization.GetDescription())
	model.Website = types.StringValue(organization.GetWebsite())
	model.Location = types.StringValue(organization.GetLocation())
	model.Visibility = types.StringValue(organization.GetVisibility())
	model.AdminChangeTeamAccess = types.BoolValue(organization.GetRepoAdminChangeTeamAccess())
	model.AvatarURL = types.StringValue(organization.GetAvatarUrl())
}

func (OrganizationConverter) ReadToResource(model models.OrganizationResourceModel, organization *api.Organization) {
	model.ID = types.Int64Value(organization.GetId())
	model.Name = types.StringValue(organization.GetName())
	model.FullName = types.StringValue(organization.GetFullName())
	model.Description = types.StringValue(organization.GetDescription())
	model.Website = types.StringValue(organization.GetWebsite())
	model.Location = types.StringValue(organization.GetLocation())
	model.Visibility = types.StringValue(organization.GetVisibility())
	model.AdminChangeTeamAccess = types.BoolValue(organization.GetRepoAdminChangeTeamAccess())
}

func (OrganizationConverter) ToApiAddOrganizationOptions(model models.OrganizationResourceModel) api.CreateOrgOption {
	return api.CreateOrgOption{
		Username:                  model.Name.ValueString(),
		FullName:                  model.FullName.ValueStringPointer(),
		Description:               model.Description.ValueStringPointer(),
		Website:                   model.Website.ValueStringPointer(),
		Location:                  model.Location.ValueStringPointer(),
		Visibility:                model.Visibility.ValueStringPointer(),
		RepoAdminChangeTeamAccess: model.AdminChangeTeamAccess.ValueBoolPointer(),
	}
}

func (OrganizationConverter) ToApiEditOrganizationOptions(model models.OrganizationResourceModel) api.EditOrgOption {
	return api.EditOrgOption{
		FullName:                  model.FullName.ValueStringPointer(),
		Description:               model.Description.ValueStringPointer(),
		Website:                   model.Website.ValueStringPointer(),
		Location:                  model.Location.ValueStringPointer(),
		Visibility:                model.Visibility.ValueStringPointer(),
		RepoAdminChangeTeamAccess: model.AdminChangeTeamAccess.ValueBoolPointer(),
	}
}
