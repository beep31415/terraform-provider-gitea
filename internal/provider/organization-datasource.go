package provider

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &orgDataSource{}
	_ datasource.DataSourceWithConfigure = &orgDataSource{}
)

type orgDataSource struct {
	proxy proxy.ProxyDataSource[models.OrganizationDataSourceModel]
}

func NewOrgDataSource() datasource.DataSource {
	return &orgDataSource{}
}

func (d *orgDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org"
}

func (d *orgDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.proxy = req.ProviderData.(*proxy.Factory).GetOrganizationDataSourceProxy()
}

func (d *orgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.OrganizationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(d.proxy.FillDataSource(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *orgDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a Gitea organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the organisation without spaces.",
				Required:    true,
			},
			"full_name": schema.StringAttribute{
				Description: "The full name of the organisation.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The organization description.",
				Computed:    true,
			},
			"website": schema.StringAttribute{
				Description: "The organization website.",
				Computed:    true,
			},
			"location": schema.StringAttribute{
				Description: "The organization location.",
				Computed:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "The organization visibility.",
				Computed:    true,
			},
			"repo_admin_change_team_access": schema.BoolAttribute{
				Description: "Flag that indicates whether admin can change organization team access.",
				Computed:    true,
			},
			"avatar_url": schema.StringAttribute{
				Description: "The organization avatar url.",
				Computed:    true,
			},
		},
	}
}
