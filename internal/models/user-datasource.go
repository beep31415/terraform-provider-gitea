package models

import "github.com/hashicorp/terraform-plugin-framework/types"

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
