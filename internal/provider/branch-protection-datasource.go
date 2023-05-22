package provider

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &branchProtectionDataSource{}
	_ datasource.DataSourceWithConfigure = &branchProtectionDataSource{}
)

type branchProtectionDataSource struct {
	proxy proxy.ProxyDataSource[models.BranchProtectionDataSourceModel]
}

func NewBranchProtectionDataSource() datasource.DataSource {
	return &branchProtectionDataSource{}
}

func (d *branchProtectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_branch_protection"
}

func (d *branchProtectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.proxy = req.ProviderData.(*proxy.Factory).NewBranchProtectionDataSourceProxy()
}

func (d *branchProtectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.BranchProtectionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(d.proxy.FillDataSource(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *branchProtectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a Gitea team.",
		Attributes: map[string]schema.Attribute{
			"owner": schema.StringAttribute{
				Description: "The owner of the repo the rule belongs to.",
				Required:    true,
			},
			"repo": schema.StringAttribute{
				Description: "The repo that the rule belongs to.",
				Required:    true,
			},
			"branch_name": schema.StringAttribute{
				Description: "The branch name targeted by the rule.",
				Computed:    true,
			},
			"approvals_whitelist_username": schema.ListAttribute{
				Description: "Whitelist of users allowed for approval.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"block_on_official_review_requests": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks on official review requests.",
				Computed:    true,
			},
			"block_on_outdated_branch": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when branch is outdated.",
				Computed:    true,
			},
			"block_on_rejected_reviews": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when pull request is rejected.",
				Computed:    true,
			},
			"dismiss_stale_approvals": schema.BoolAttribute{
				Description: "Flag indicating whether to ignore stale approvals on pull requests.",
				Computed:    true,
			},
			"enable_approvals_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed for approvals.",
				Computed:    true,
			},
			"enable_merge_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed to merge.",
				Computed:    true,
			},
			"enable_push": schema.BoolAttribute{
				Description: "Flag indicating whether to allow push to protected branch.",
				Computed:    true,
			},
			"enable_push_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable whitelist of users allowed to push.",
				Computed:    true,
			},
			"merge_whitelist_usernames": schema.ListAttribute{
				Description: "Whitelist of users allowed to merge.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"protected_file_patterns": schema.StringAttribute{
				Description: "File pattern of protected files.",
				Computed:    true,
			},
			"push_whitelist_deploy_keys": schema.BoolAttribute{
				Description: "Flag indicating whether to push deploy keys on whitelisted users.",
				Computed:    true,
			},
			"push_whitelist_usernames": schema.ListAttribute{
				Description: "Flag indicating whether to push usernames on whitelisted users.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"require_signed_commits": schema.BoolAttribute{
				Description: "Flag indicating whether to require a signed commit on push.",
				Computed:    true,
			},
			"required_approvals": schema.Int64Attribute{
				Description: "Flag indicating whether to require approvals on pull requests.",
				Computed:    true,
			},
			"unprotected_file_patterns": schema.StringAttribute{
				Description: "File pattern of unprotected files.",
				Computed:    true,
			},
		},
	}
}
