package models

import "github.com/hashicorp/terraform-plugin-framework/types"

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
