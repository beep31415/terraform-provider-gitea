package provider

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

type userDataSource struct {
	client *api.APIClient
	proxy  proxy.UserProxy
}

func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.APIClient)
	d.proxy = *proxy.NewUserProxy(d.client)
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.UserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(d.proxy.FillDataSource(ctx, state))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
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
