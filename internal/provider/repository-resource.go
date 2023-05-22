package provider

import (
	"context"
	"strings"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/pkg/plans"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ resource.Resource                = &repoResource{}
	_ resource.ResourceWithConfigure   = &repoResource{}
	_ resource.ResourceWithImportState = &repoResource{}
)

type repoResource struct {
	client *api.APIClient
	proxy  *proxy.RepositoryProxy
}

func NewRepoResource() resource.Resource {
	return &repoResource{}
}

func (r *repoResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repo"
}

func (r *repoResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
	r.proxy = proxy.NewRepositoryProxy(r.client)
}

func (r *repoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RepositoryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Create(ctx, plan))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *repoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.RepositoryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.FillResource(ctx, state))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *repoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.RepositoryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Edit(ctx, plan))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *repoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.RepositoryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Delete(ctx, state))
}

func (r *repoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ids := strings.Split(req.ID, "/")
	if len(ids) != 2 {
		resp.Diagnostics.AddError(
			"Error Importing Gitea repository.",
			"Must provide import id in the form of owner_name/repo_name.",
		)

		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("owner"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), ids[1])...)
}

func (d *repoResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gitea repository.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the repository without spaces.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"owner": schema.StringAttribute{
				Description: "The owner of the repository without spaces. " +
					"`!!IMPORTANT` Note that if owner is not an organization it has to be set as " +
					"the user related to the terrafom credentials for this provider since " +
					"Gitea does not offer an API to create repos owned by other users.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"auto_init": schema.BoolAttribute{
				Description: "Whether the repository should be auto-initialized.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"default_branch": schema.StringAttribute{
				Description: "DefaultBranch of the repository (used when initializes and in template). Defaults to `master`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("master"),
			},
			"description": schema.StringAttribute{
				Description: "Description of the repository to create.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"gitignores": schema.StringAttribute{
				Description: "Gitignores to use.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"issue_labels": schema.StringAttribute{
				Description: "Label-Set to use.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"license": schema.StringAttribute{
				Description: "License to use.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"private": schema.BoolAttribute{
				Description: "Whether the repository is private. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"readme": schema.StringAttribute{
				Description: "Readme of the repository to create.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"template": schema.BoolAttribute{
				Description: "Whether the repository is template. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"trust_model": schema.StringAttribute{
				Description: "TrustModel of the repository. Can be: default, collaborator, committer or collaboratorcommitter. Defaults to `default`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "collaborator", "committer", "collaboratorcommitter"),
				},
			},
			"allow_manual_merge": schema.BoolAttribute{
				Description: "Either `true` to allow mark pr as merged manually, or `false` to prevent it.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"allow_merge_commits": schema.BoolAttribute{
				Description: "Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"allow_rebase": schema.BoolAttribute{
				Description: "Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"allow_rebase_explicit": schema.BoolAttribute{
				Description: "Either `true` to allow rebase with explicit merge commits (--no-ff), or `false` to prevent rebase with explicit merge commits. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"allow_rebase_update": schema.BoolAttribute{
				Description: "Either `true` to allow updating pull request branch by rebase, or `false` to prevent it. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"allow_squash_merge": schema.BoolAttribute{
				Description: "Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"archived": schema.BoolAttribute{
				Description: "Set to `true` to archive this repository. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"autodetect_manual_merge": schema.BoolAttribute{
				Description: "Either `true` to enable AutodetectManualMerge, or `false` to prevent it. Note: In some special cases, misjudgments can occur.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"default_allow_maintainer_edit": schema.BoolAttribute{
				Description: "Set to `true` to allow edits from maintainers by default. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"default_delete_branch_after_merge": schema.BoolAttribute{
				Description: "Set to `true` to delete pr branch after merge by default. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"default_merge_style": schema.StringAttribute{
				Description: "Set to a merge style to be used by this repository. Can be: merge, rebase, rebase-merge or squash. Defaults to `merge`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("merge"),
				Validators: []validator.String{
					stringvalidator.OneOf("merge", "rebase", "rebase-merge", "squash"),
				},
			},
			"enable_prune": schema.BoolAttribute{
				Description: "Enable prune - remove obsolete remote-tracking references.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"has_issues": schema.BoolAttribute{
				Description: "Either `true` to enable issues for this repository or `false` to disable them. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"has_projects": schema.BoolAttribute{
				Description: "Either `true` to enable project unit, or `false` to disable them. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"has_pull_requests": schema.BoolAttribute{
				Description: "Either `true` to allow pull requests, or `false` to prevent pull request. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"has_wiki": schema.BoolAttribute{
				Description: "Either `true` to enable the wiki for this repository or `false` to disable it. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"ignore_whitespace_conflicts": schema.BoolAttribute{
				Description: "Either `true` to ignore whitespace for conflicts, or `false` to not ignore whitespace. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"website": schema.StringAttribute{
				Description: "A URL with more information about the repository.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
		},
	}
}
