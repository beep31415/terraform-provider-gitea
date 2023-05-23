package provider

import (
	"context"
	"strings"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &branchProtectionResource{}
	_ resource.ResourceWithConfigure   = &branchProtectionResource{}
	_ resource.ResourceWithImportState = &branchProtectionResource{}
)

type branchProtectionResource struct {
	proxy proxy.ProxyResource[models.BranchProtectionResourceModel]
}

func NewBranchProtectionResource() resource.Resource {
	return &branchProtectionResource{}
}

func (r *branchProtectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_branch_protection"
}

func (r *branchProtectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.proxy = req.ProviderData.(*proxy.Factory).NewBranchProtectionResourceProxy()
}

func (r *branchProtectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Create(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.FillResource(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Update(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Delete(ctx, &state)...)
}

func (r *branchProtectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ids := strings.Split(req.ID, "/")
	if len(ids) != 3 {
		resp.Diagnostics.AddError(
			"Error Importing Gitea branch protection.",
			"Must provide import id in the form of owner_name/repo_name/branch_name.",
		)

		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("owner"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("repo"), ids[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("branch_name"), ids[2])...)
}

func (d *branchProtectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gitea branch protection.",
		Attributes: map[string]schema.Attribute{
			"branch_name": schema.StringAttribute{
				Description: "The branch name targeted by the rule.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"owner": schema.StringAttribute{
				Description: "The owner of the repo the rule belongs to.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"repo": schema.StringAttribute{
				Description: "The repo that the rule belongs to.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"approvals_whitelist_usernames": schema.ListAttribute{
				Description: "Whitelist of users allowed for approval.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"approvals_whitelist_teams": schema.ListAttribute{
				Description: "Whitelist of teams allowed for approval.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"block_on_official_review_requests": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks on official review requests. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"block_on_outdated_branch": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when branch is outdated. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"block_on_rejected_reviews": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when pull request is rejected. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"dismiss_stale_approvals": schema.BoolAttribute{
				Description: "Flag indicating whether to ignore stale approvals on pull requests. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_approvals_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed for approvals. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_merge_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed to merge. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_push": schema.BoolAttribute{
				Description: "Flag indicating whether to allow push to protected branch. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_push_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable whitelist of users allowed to push. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"merge_whitelist_usernames": schema.ListAttribute{
				Description: "Whitelist of users allowed to merge.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"merge_whitelist_teams": schema.ListAttribute{
				Description: "Whitelist of teams allowed to merge.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"protected_file_patterns": schema.StringAttribute{
				Description: "File pattern of protected files. Defaults to `\"\"`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"push_whitelist_deploy_keys": schema.BoolAttribute{
				Description: "Flag indicating whether to push deploy keys on whitelisted users. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"push_whitelist_usernames": schema.ListAttribute{
				Description: "Flag indicating whether to push usernames on whitelisted users.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"push_whitelist_teams": schema.ListAttribute{
				Description: "Whitelist of teams allowed to push.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
			"require_signed_commits": schema.BoolAttribute{
				Description: "Flag indicating whether to require a signed commit on push. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"required_approvals": schema.Int64Attribute{
				Description: "Number of required approvals. Defaults to `0`.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(0),
			},
			"unprotected_file_patterns": schema.StringAttribute{
				Description: "File pattern of unprotected files. Defaults to `\"\"`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
		},
	}
}
