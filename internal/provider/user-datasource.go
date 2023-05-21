package provider

import (
	"context"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

type userDataSource struct {
	client      *api.APIClient
	userAdapter adapters.UserAdapter
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
	d.userAdapter = *adapters.NewUserAdapter(d.client)
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.UserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.userAdapter.GetByName(ctx, state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Gitea user.",
			adapters.GetAPIErrorMessage(err),
		)
		return
	}

	state = models.NewUserDataSource(res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
