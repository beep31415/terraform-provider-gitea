package models

import (
	"terraform-provider-gitea/api"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserDataSourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	Name          types.String `tfsdk:"username"`
	LoginName     types.String `tfsdk:"login_name"`
	Email         types.String `tfsdk:"email"`
	IsAdmin       types.Bool   `tfsdk:"is_admin"`
	IsActive      types.Bool   `tfsdk:"is_active"`
	ProhibitLogin types.Bool   `tfsdk:"prohibit_login"`
	Restricted    types.Bool   `tfsdk:"restricted"`
	Visibility    types.String `tfsdk:"visibility"`
	LastLogin     types.String `tfsdk:"last_login"`
	Created       types.String `tfsdk:"created"`
}

func NewUserDataSource(user *api.User) UserDataSourceModel {
	return UserDataSourceModel{
		ID:            types.Int64Value(user.GetId()),
		Name:          types.StringValue(user.GetLogin()),
		LoginName:     types.StringValue(user.GetLoginName()),
		Email:         types.StringValue(user.GetEmail()),
		IsAdmin:       types.BoolValue(user.GetIsAdmin()),
		IsActive:      types.BoolValue(user.GetActive()),
		ProhibitLogin: types.BoolValue(user.GetProhibitLogin()),
		Restricted:    types.BoolValue(user.GetRestricted()),
		Visibility:    types.StringValue(user.GetVisibility()),
		LastLogin:     types.StringValue(user.GetLastLogin().String()),
		Created:       types.StringValue(user.GetCreated().String()),
	}
}
