package datasources

import (
	"context"
	"strings"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/provider/errors"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

type userDataSourceModel struct {
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

type userDataSource struct {
	client *api.APIClient
}

func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a Gitea organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "The user name.",
				Required:    true,
			},
			"login_name": schema.StringAttribute{
				Description: "The user login name.",
				Computed:    true,
			},
			"email": schema.StringAttribute{
				Description: "The user email.",
				Computed:    true,
			},
			"is_admin": schema.BoolAttribute{
				Description: "Flag that indicates if user is an admin.",
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Flag that indicates if user is active.",
				Computed:    true,
			},
			"prohibit_login": schema.BoolAttribute{
				Description: "Flag that indicates if user can login.",
				Computed:    true,
			},
			"restricted": schema.BoolAttribute{
				Description: "Flag that indicates if user is restricted.",
				Computed:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "The user visibility.",
				Computed:    true,
			},
			"last_login": schema.StringAttribute{
				Description: "Last user login.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Description: "User creation date.",
				Computed:    true,
			},
		},
	}
}

func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.APIClient)
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state userDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, _, err := d.client.UserAPI.
		UserGet(ctx, strings.ToLower(state.Name.ValueString())).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Gitea user.",
			errors.GetAPIError(err),
		)
		return
	}

	state = userDataSourceModel{
		ID:            types.Int64Value(res.GetId()),
		Name:          types.StringValue(res.GetLogin()),
		LoginName:     types.StringValue(res.GetLoginName()),
		Email:         types.StringValue(res.GetEmail()),
		IsAdmin:       types.BoolValue(res.GetIsAdmin()),
		IsActive:      types.BoolValue(res.GetActive()),
		ProhibitLogin: types.BoolValue(res.GetProhibitLogin()),
		Restricted:    types.BoolValue(res.GetRestricted()),
		Visibility:    types.StringValue(res.GetVisibility()),
		LastLogin:     types.StringValue(res.GetLastLogin().String()),
		Created:       types.StringValue(res.GetCreated().String()),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
