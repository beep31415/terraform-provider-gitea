package provider

import (
	"context"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &teamDataSource{}
	_ datasource.DataSourceWithConfigure = &teamDataSource{}
)

type teamDataSource struct {
	client      *api.APIClient
	teamAdapter *adapters.TeamAdapter
}

type teamUnitDataSource struct {
	Name       types.String `tfsdk:"name"`
	Permission types.String `tfsdk:"permission"`
}

func NewTeamDataSource() datasource.DataSource {
	return &teamDataSource{}
}

func (d *teamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (d *teamDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a Gitea team.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The team name.",
				Required:    true,
			},
			"organization": schema.StringAttribute{
				Description: "The organization the team belongs to.",
				Required:    true,
			},
			"permission": schema.StringAttribute{
				Description: "Team permission within the organization.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The team description",
				Computed:    true,
			},
			"can_create_org_repo": schema.BoolAttribute{
				Description: "Flag that indicates whether team can create organization repos.",
				Computed:    true,
			},
			"includes_all_repositories": schema.BoolAttribute{
				Description: "Flag indicating whether team has access to all repos.",
				Computed:    true,
			},
			"members": schema.ListAttribute{
				Description: "Team members.",
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}

func (d *teamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.APIClient)
	d.teamAdapter = adapters.NewTeamAdapter(d.client)
}

func (d *teamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.TeamDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.teamAdapter.GetTeamByOrgAndName(ctx, state.Organization.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Gitea team.",
			adapters.GetAPIErrorMessage(err),
		)

		return
	}

	var diags diag.Diagnostics
	state, diags = models.NewTeamDataSource(ctx, res)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
