package models

import (
	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RepositoryResourceModel struct {
	ID                            types.Int64  `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	Owner                         types.String `tfsdk:"owner"`
	AutoInit                      types.Bool   `tfsdk:"auto_init"`
	DefaultBranch                 types.String `tfsdk:"default_branch"`
	Description                   types.String `tfsdk:"description"`
	Gitignores                    types.String `tfsdk:"gitignores"`
	IssueLabels                   types.String `tfsdk:"issue_labels"`
	License                       types.String `tfsdk:"license"`
	Private                       types.Bool   `tfsdk:"private"`
	Readme                        types.String `tfsdk:"readme"`
	Template                      types.Bool   `tfsdk:"template"`
	TrustModel                    types.String `tfsdk:"trust_model"`
	AllowManualMerge              types.Bool   `tfsdk:"allow_manual_merge"`
	AllowMerge                    types.Bool   `tfsdk:"allow_merge_commits"`
	AllowRebase                   types.Bool   `tfsdk:"allow_rebase"`
	AllowRebaseMerge              types.Bool   `tfsdk:"allow_rebase_explicit"`
	AllowRebaseUpdate             types.Bool   `tfsdk:"allow_rebase_update"`
	AllowSquash                   types.Bool   `tfsdk:"allow_squash_merge"`
	Archived                      types.Bool   `tfsdk:"archived"`
	AutodetectManualMerge         types.Bool   `tfsdk:"autodetect_manual_merge"`
	DefaultAllowMaintainerEdit    types.Bool   `tfsdk:"default_allow_maintainer_edit"`
	DefaultDeleteBranchAfterMerge types.Bool   `tfsdk:"default_delete_branch_after_merge"`
	DefaultMergeStyle             types.String `tfsdk:"default_merge_style"`
	EnablePrune                   types.Bool   `tfsdk:"enable_prune"`
	HasIssues                     types.Bool   `tfsdk:"has_issues"`
	HasProjects                   types.Bool   `tfsdk:"has_projects"`
	HasPullRequests               types.Bool   `tfsdk:"has_pull_requests"`
	HasWiki                       types.Bool   `tfsdk:"has_wiki"`
	IgnoreWhitespaceConflicts     types.Bool   `tfsdk:"ignore_whitespace_conflicts"`
	Website                       types.String `tfsdk:"website"`
}

func (r *RepositoryResourceModel) ReadFromApi(repository *api.Repository) {
	r.ID = types.Int64Value(repository.GetId())
	r.Name = types.StringValue(repository.GetName())
	r.Owner = types.StringValue(*repository.GetOwner().Login)
	r.DefaultBranch = types.StringValue(repository.GetDefaultBranch())
	r.Description = types.StringValue(repository.GetDescription())
	r.Private = types.BoolValue(repository.GetPrivate())
	r.Template = types.BoolValue(repository.GetTemplate())
	r.AllowMerge = types.BoolValue(repository.GetAllowMergeCommits())
	r.AllowRebase = types.BoolValue(repository.GetAllowRebase())
	r.AllowRebaseMerge = types.BoolValue(repository.GetAllowRebaseExplicit())
	r.AllowRebaseUpdate = types.BoolValue(repository.GetAllowRebaseUpdate())
	r.AllowSquash = types.BoolValue(repository.GetAllowSquashMerge())
	r.Archived = types.BoolValue(repository.GetArchived())
	r.DefaultAllowMaintainerEdit = types.BoolValue(repository.GetDefaultAllowMaintainerEdit())
	r.DefaultDeleteBranchAfterMerge = types.BoolValue(repository.GetDefaultDeleteBranchAfterMerge())
	r.DefaultMergeStyle = types.StringValue(repository.GetDefaultMergeStyle())
	r.HasIssues = types.BoolValue(repository.GetHasIssues())
	r.HasProjects = types.BoolValue(repository.GetHasProjects())
	r.HasPullRequests = types.BoolValue(repository.GetHasPullRequests())
	r.HasWiki = types.BoolValue(repository.GetHasWiki())
	r.IgnoreWhitespaceConflicts = types.BoolValue(repository.GetIgnoreWhitespaceConflicts())
	r.Website = types.StringValue(repository.GetWebsite())
	if r.AutoInit.IsNull() {
		r.AutoInit = types.BoolValue(false)
	}
	if r.Gitignores.IsNull() {
		r.Gitignores = types.StringValue("")
	}
	if r.IssueLabels.IsNull() {
		r.IssueLabels = types.StringValue("")
	}
	if r.License.IsNull() {
		r.License = types.StringValue("")
	}
	if r.Readme.IsNull() {
		r.Readme = types.StringValue("")
	}
	if r.AllowManualMerge.IsNull() {
		r.AllowManualMerge = types.BoolValue(false)
	}
	if r.AutodetectManualMerge.IsNull() {
		r.AutodetectManualMerge = types.BoolValue(false)
	}
	if r.EnablePrune.IsNull() {
		r.EnablePrune = types.BoolValue(false)
	}
	if r.TrustModel.IsNull() {
		r.TrustModel = types.StringValue("default")
	}
}

func (r *RepositoryResourceModel) ToApiRepositoryOptions() adapters.RepositoryApiOptions {
	return adapters.RepositoryApiOptions{
		Owner:                         r.Owner.ValueString(),
		Name:                          r.Name.ValueString(),
		AutoInit:                      r.AutoInit.ValueBoolPointer(),
		DefaultBranch:                 r.DefaultBranch.ValueStringPointer(),
		Description:                   r.Description.ValueStringPointer(),
		Gitignores:                    r.Gitignores.ValueStringPointer(),
		IssueLabels:                   r.IssueLabels.ValueStringPointer(),
		License:                       r.License.ValueStringPointer(),
		Private:                       r.Private.ValueBoolPointer(),
		Readme:                        r.Readme.ValueStringPointer(),
		Template:                      r.Template.ValueBoolPointer(),
		TrustModel:                    r.TrustModel.ValueStringPointer(),
		AllowManualMerge:              r.AllowManualMerge.ValueBoolPointer(),
		AllowMergeCommits:             r.AllowMerge.ValueBoolPointer(),
		AllowRebase:                   r.AllowRebase.ValueBoolPointer(),
		AllowRebaseExplicit:           r.AllowRebaseMerge.ValueBoolPointer(),
		AllowRebaseUpdate:             r.AllowRebaseUpdate.ValueBoolPointer(),
		AllowSquashMerge:              r.AllowSquash.ValueBoolPointer(),
		Archived:                      r.Archived.ValueBoolPointer(),
		AutodetectManualMerge:         r.AutodetectManualMerge.ValueBoolPointer(),
		DefaultAllowMaintainerEdit:    r.DefaultAllowMaintainerEdit.ValueBoolPointer(),
		DefaultDeleteBranchAfterMerge: r.DefaultDeleteBranchAfterMerge.ValueBoolPointer(),
		DefaultMergeStyle:             r.DefaultMergeStyle.ValueStringPointer(),
		EnablePrune:                   r.EnablePrune.ValueBoolPointer(),
		HasIssues:                     r.HasIssues.ValueBoolPointer(),
		HasProjects:                   r.HasProjects.ValueBoolPointer(),
		HasPullRequests:               r.HasPullRequests.ValueBoolPointer(),
		HasWiki:                       r.HasWiki.ValueBoolPointer(),
		IgnoreWhitespaceConflicts:     r.IgnoreWhitespaceConflicts.ValueBoolPointer(),
		Website:                       r.Website.ValueStringPointer(),
	}
}
