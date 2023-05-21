package converters

import (
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserConverter struct{}

func (UserConverter) ReadToDataSource(model models.UserDataSourceModel, user *api.User) {
	model.ID = types.Int64Value(user.GetId())
	model.Name = types.StringValue(user.GetLogin())
	model.LoginName = types.StringValue(user.GetLoginName())
	model.Email = types.StringValue(user.GetEmail())
	model.IsAdmin = types.BoolValue(user.GetIsAdmin())
	model.IsActive = types.BoolValue(user.GetActive())
	model.ProhibitLogin = types.BoolValue(user.GetProhibitLogin())
	model.Restricted = types.BoolValue(user.GetRestricted())
	model.Visibility = types.StringValue(user.GetVisibility())
	model.LastLogin = types.StringValue(user.GetLastLogin().String())
	model.Created = types.StringValue(user.GetCreated().String())
}
