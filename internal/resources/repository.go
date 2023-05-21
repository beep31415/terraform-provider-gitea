package resources

import (
	"context"
	"strings"
	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/plans"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &repoResource{}
	_ resource.ResourceWithConfigure   = &repoResource{}
	_ resource.ResourceWithImportState = &repoResource{}
)

type repoResource struct {
	client *api.APIClient
}

func NewRepoResource() resource.Resource {
	return &repoResource{}
}

func (r *repoResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repo"
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

func (r *repoResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
}

func (r *repoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RepositoryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	owner := plan.Owner.ValueString()

	org, err := r.getOrgByName(ctx, owner)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea repository.",
			"Could not check if owner is an organization for "+owner+": "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	opts := api.CreateRepoOption{
		Name:          plan.Name.ValueString(),
		AutoInit:      plan.AutoInit.ValueBoolPointer(),
		DefaultBranch: plan.DefaultBranch.ValueStringPointer(),
		Description:   plan.Description.ValueStringPointer(),
		Gitignores:    plan.Gitignores.ValueStringPointer(),
		IssueLabels:   plan.IssueLabels.ValueStringPointer(),
		License:       plan.License.ValueStringPointer(),
		Private:       plan.Private.ValueBoolPointer(),
		Readme:        plan.Readme.ValueStringPointer(),
		Template:      plan.Template.ValueBoolPointer(),
		TrustModel:    plan.TrustModel.ValueStringPointer(),
	}

	var res *api.Repository
	if org != nil {
		res, err = r.createOrgRepo(ctx, owner, opts)
	} else {
		res, err = r.createCurrentUserRepo(ctx, opts)
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea repository.",
			"Could not create repository, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	plan.ID = types.Int64Value(res.GetId())
	plan.Name = types.StringValue(res.GetName())
	plan.DefaultBranch = types.StringValue(res.GetDefaultBranch())
	plan.Description = types.StringValue(res.GetDescription())
	plan.Private = types.BoolValue(res.GetPrivate())
	plan.Template = types.BoolValue(res.GetTemplate())

	resEdit, _, err := r.client.RepositoryAPI.
		RepoEdit(ctx, owner, res.GetName()).
		Body(api.EditRepoOption{
			Name:                          plan.Name.ValueStringPointer(),
			AllowManualMerge:              plan.AllowManualMerge.ValueBoolPointer(),
			AllowMergeCommits:             plan.AllowMerge.ValueBoolPointer(),
			AllowRebase:                   plan.AllowRebase.ValueBoolPointer(),
			AllowRebaseExplicit:           plan.AllowRebaseMerge.ValueBoolPointer(),
			AllowRebaseUpdate:             plan.AllowRebaseUpdate.ValueBoolPointer(),
			AllowSquashMerge:              plan.AllowSquash.ValueBoolPointer(),
			Archived:                      plan.Archived.ValueBoolPointer(),
			AutodetectManualMerge:         plan.AutodetectManualMerge.ValueBoolPointer(),
			DefaultAllowMaintainerEdit:    plan.DefaultAllowMaintainerEdit.ValueBoolPointer(),
			DefaultBranch:                 plan.DefaultBranch.ValueStringPointer(),
			DefaultDeleteBranchAfterMerge: plan.DefaultDeleteBranchAfterMerge.ValueBoolPointer(),
			DefaultMergeStyle:             plan.DefaultMergeStyle.ValueStringPointer(),
			Description:                   plan.Description.ValueStringPointer(),
			EnablePrune:                   plan.EnablePrune.ValueBoolPointer(),
			HasIssues:                     plan.HasIssues.ValueBoolPointer(),
			HasProjects:                   plan.HasProjects.ValueBoolPointer(),
			HasPullRequests:               plan.HasPullRequests.ValueBoolPointer(),
			HasWiki:                       plan.HasWiki.ValueBoolPointer(),
			IgnoreWhitespaceConflicts:     plan.IgnoreWhitespaceConflicts.ValueBoolPointer(),
			Private:                       plan.Private.ValueBoolPointer(),
			Template:                      plan.Template.ValueBoolPointer(),
			Website:                       plan.Website.ValueStringPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea repository.",
			"Could not update all repository values, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	plan.Name = types.StringValue(resEdit.GetName())
	plan.AllowMerge = types.BoolValue(resEdit.GetAllowMergeCommits())
	plan.AllowRebase = types.BoolValue(resEdit.GetAllowRebase())
	plan.AllowRebaseMerge = types.BoolValue(resEdit.GetAllowRebaseExplicit())
	plan.AllowRebaseUpdate = types.BoolValue(resEdit.GetAllowRebaseUpdate())
	plan.AllowSquash = types.BoolValue(resEdit.GetAllowSquashMerge())
	plan.Archived = types.BoolValue(resEdit.GetArchived())
	plan.DefaultAllowMaintainerEdit = types.BoolValue(resEdit.GetDefaultAllowMaintainerEdit())
	plan.DefaultBranch = types.StringValue(resEdit.GetDefaultBranch())
	plan.DefaultDeleteBranchAfterMerge = types.BoolValue(resEdit.GetDefaultDeleteBranchAfterMerge())
	plan.DefaultMergeStyle = types.StringValue(resEdit.GetDefaultMergeStyle())
	plan.Description = types.StringValue(resEdit.GetDescription())
	plan.HasIssues = types.BoolValue(resEdit.GetHasIssues())
	plan.HasProjects = types.BoolValue(resEdit.GetHasProjects())
	plan.HasPullRequests = types.BoolValue(resEdit.GetHasPullRequests())
	plan.HasWiki = types.BoolValue(resEdit.GetHasWiki())
	plan.IgnoreWhitespaceConflicts = types.BoolValue(resEdit.GetIgnoreWhitespaceConflicts())
	plan.Private = types.BoolValue(resEdit.GetPrivate())
	plan.Template = types.BoolValue(resEdit.GetTemplate())
	plan.Website = types.StringValue(resEdit.GetWebsite())

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

	res, err := r.getRepositoryByName(ctx, state.Owner.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea repository.",
			"Could not read Gitea repository ID "+state.ID.String()+": "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	state.ID = types.Int64Value(res.GetId())
	state.Name = types.StringValue(res.GetName())
	state.Owner = types.StringValue(*res.GetOwner().Login)
	state.DefaultBranch = types.StringValue(res.GetDefaultBranch())
	state.Description = types.StringValue(res.GetDescription())
	state.Private = types.BoolValue(res.GetPrivate())
	state.Template = types.BoolValue(res.GetTemplate())
	state.AllowMerge = types.BoolValue(res.GetAllowMergeCommits())
	state.AllowRebase = types.BoolValue(res.GetAllowRebase())
	state.AllowRebaseMerge = types.BoolValue(res.GetAllowRebaseExplicit())
	state.AllowRebaseUpdate = types.BoolValue(res.GetAllowRebaseUpdate())
	state.AllowSquash = types.BoolValue(res.GetAllowSquashMerge())
	state.Archived = types.BoolValue(res.GetArchived())
	state.DefaultAllowMaintainerEdit = types.BoolValue(res.GetDefaultAllowMaintainerEdit())
	state.DefaultDeleteBranchAfterMerge = types.BoolValue(res.GetDefaultDeleteBranchAfterMerge())
	state.DefaultMergeStyle = types.StringValue(res.GetDefaultMergeStyle())
	state.HasIssues = types.BoolValue(res.GetHasIssues())
	state.HasProjects = types.BoolValue(res.GetHasProjects())
	state.HasPullRequests = types.BoolValue(res.GetHasPullRequests())
	state.HasWiki = types.BoolValue(res.GetHasWiki())
	state.IgnoreWhitespaceConflicts = types.BoolValue(res.GetIgnoreWhitespaceConflicts())
	state.Website = types.StringValue(res.GetWebsite())
	if state.AutoInit.IsNull() {
		state.AutoInit = types.BoolValue(false)
	}
	if state.Gitignores.IsNull() {
		state.Gitignores = types.StringValue("")
	}
	if state.IssueLabels.IsNull() {
		state.IssueLabels = types.StringValue("")
	}
	if state.License.IsNull() {
		state.License = types.StringValue("")
	}
	if state.Readme.IsNull() {
		state.Readme = types.StringValue("")
	}
	if state.AllowManualMerge.IsNull() {
		state.AllowManualMerge = types.BoolValue(false)
	}
	if state.AutodetectManualMerge.IsNull() {
		state.AutodetectManualMerge = types.BoolValue(false)
	}
	if state.EnablePrune.IsNull() {
		state.EnablePrune = types.BoolValue(false)
	}
	if state.TrustModel.IsNull() {
		state.TrustModel = types.StringValue("default")
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

	res, _, err := r.client.RepositoryAPI.
		RepoEdit(ctx, plan.Owner.ValueString(), plan.Name.ValueString()).
		Body(api.EditRepoOption{
			Name:                          plan.Name.ValueStringPointer(),
			AllowManualMerge:              plan.AllowManualMerge.ValueBoolPointer(),
			AllowMergeCommits:             plan.AllowMerge.ValueBoolPointer(),
			AllowRebase:                   plan.AllowRebase.ValueBoolPointer(),
			AllowRebaseExplicit:           plan.AllowRebaseMerge.ValueBoolPointer(),
			AllowRebaseUpdate:             plan.AllowRebaseUpdate.ValueBoolPointer(),
			AllowSquashMerge:              plan.AllowSquash.ValueBoolPointer(),
			Archived:                      plan.Archived.ValueBoolPointer(),
			AutodetectManualMerge:         plan.AutodetectManualMerge.ValueBoolPointer(),
			DefaultAllowMaintainerEdit:    plan.DefaultAllowMaintainerEdit.ValueBoolPointer(),
			DefaultBranch:                 plan.DefaultBranch.ValueStringPointer(),
			DefaultDeleteBranchAfterMerge: plan.DefaultDeleteBranchAfterMerge.ValueBoolPointer(),
			DefaultMergeStyle:             plan.DefaultMergeStyle.ValueStringPointer(),
			Description:                   plan.Description.ValueStringPointer(),
			EnablePrune:                   plan.EnablePrune.ValueBoolPointer(),
			HasIssues:                     plan.HasIssues.ValueBoolPointer(),
			HasProjects:                   plan.HasProjects.ValueBoolPointer(),
			HasPullRequests:               plan.HasPullRequests.ValueBoolPointer(),
			HasWiki:                       plan.HasWiki.ValueBoolPointer(),
			IgnoreWhitespaceConflicts:     plan.IgnoreWhitespaceConflicts.ValueBoolPointer(),
			Private:                       plan.Private.ValueBoolPointer(),
			Template:                      plan.Template.ValueBoolPointer(),
			Website:                       plan.Website.ValueStringPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Gitea repository.",
			"Could not update repository, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	plan.ID = types.Int64Value(res.GetId())
	plan.Name = types.StringValue(res.GetName())
	plan.Owner = types.StringValue(*res.GetOwner().Login)
	plan.DefaultBranch = types.StringValue(res.GetDefaultBranch())
	plan.Description = types.StringValue(res.GetDescription())
	plan.Private = types.BoolValue(res.GetPrivate())
	plan.Template = types.BoolValue(res.GetTemplate())
	plan.AllowMerge = types.BoolValue(res.GetAllowMergeCommits())
	plan.AllowRebase = types.BoolValue(res.GetAllowRebase())
	plan.AllowRebaseMerge = types.BoolValue(res.GetAllowRebaseExplicit())
	plan.AllowRebaseUpdate = types.BoolValue(res.GetAllowRebaseUpdate())
	plan.AllowSquash = types.BoolValue(res.GetAllowSquashMerge())
	plan.Archived = types.BoolValue(res.GetArchived())
	plan.DefaultAllowMaintainerEdit = types.BoolValue(res.GetDefaultAllowMaintainerEdit())
	plan.DefaultDeleteBranchAfterMerge = types.BoolValue(res.GetDefaultDeleteBranchAfterMerge())
	plan.DefaultMergeStyle = types.StringValue(res.GetDefaultMergeStyle())
	plan.HasIssues = types.BoolValue(res.GetHasIssues())
	plan.HasProjects = types.BoolValue(res.GetHasProjects())
	plan.HasPullRequests = types.BoolValue(res.GetHasPullRequests())
	plan.HasWiki = types.BoolValue(res.GetHasWiki())
	plan.IgnoreWhitespaceConflicts = types.BoolValue(res.GetIgnoreWhitespaceConflicts())
	plan.Website = types.StringValue(res.GetWebsite())

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

	res, err := r.getRepositoryById(ctx, state.ID.ValueInt64())
	if err != nil {
		if adapters.IsErrorNotFound(err) {
			return
		}

		resp.Diagnostics.AddError(
			"Error Delete Gitea repository.",
			"Could not get repository to delete, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	_, err = r.client.RepositoryAPI.
		RepoDelete(ctx, *res.GetOwner().Login, res.GetName()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Gitea repository.",
			"Could not delete repository, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}
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

func (r *repoResource) getRepositoryById(ctx context.Context, id int64) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGetByID(ctx, id).
		Execute()

	return res, err
}

func (r *repoResource) getRepositoryByName(ctx context.Context, owner, repo string) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGet(ctx, owner, repo).
		Execute()

	return res, err
}

func (r *repoResource) createOrgRepo(ctx context.Context, org string, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.OrganizationAPI.
		CreateOrgRepo(ctx, org).
		Body(opts).
		Execute()

	return res, err
}

func (r *repoResource) createCurrentUserRepo(ctx context.Context, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		CreateCurrentUserRepo(ctx).
		Body(opts).
		Execute()

	return res, err
}

func (r *repoResource) getOrgByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	if adapters.IsErrorNotFound(err) {
		res = nil
		err = nil
	}

	return res, err
}
