package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type RepositoryDataSourceModel struct {
	ID                            types.Int64  `tfsdk:"id"`
	Owner                         types.String `tfsdk:"owner"`
	Name                          types.String `tfsdk:"name"`
	AllowMerge                    types.Bool   `tfsdk:"allow_merge_commits"`
	AllowRebase                   types.Bool   `tfsdk:"allow_rebase"`
	AllowRebaseMerge              types.Bool   `tfsdk:"allow_rebase_explicit"`
	AllowRebaseUpdate             types.Bool   `tfsdk:"allow_rebase_update"`
	AllowSquash                   types.Bool   `tfsdk:"allow_squash_merge"`
	Archived                      types.Bool   `tfsdk:"archived"`
	AvatarURL                     types.String `tfsdk:"avatar_url"`
	CloneURL                      types.String `tfsdk:"clone_url"`
	Created                       types.String `tfsdk:"created_at"`
	DefaultAllowMaintainerEdit    types.Bool   `tfsdk:"default_allow_maintainer_edit"`
	DefaultBranch                 types.String `tfsdk:"default_branch"`
	DefaultDeleteBranchAfterMerge types.Bool   `tfsdk:"default_delete_branch_after_merge"`
	DefaultMergeStyle             types.String `tfsdk:"default_merge_style"`
	Description                   types.String `tfsdk:"description"`
	Empty                         types.Bool   `tfsdk:"empty"`
	Fork                          types.Bool   `tfsdk:"fork"`
	Forks                         types.Int64  `tfsdk:"forks_count"`
	FullName                      types.String `tfsdk:"full_name"`
	HTMLURL                       types.String `tfsdk:"html_url"`
	HasIssues                     types.Bool   `tfsdk:"has_issues"`
	HasProjects                   types.Bool   `tfsdk:"has_projects"`
	HasPullRequests               types.Bool   `tfsdk:"has_pull_requests"`
	HasWiki                       types.Bool   `tfsdk:"has_wiki"`
	IgnoreWhitespaceConflicts     types.Bool   `tfsdk:"ignore_whitespace_conflicts"`
	Internal                      types.Bool   `tfsdk:"internal"`
	Language                      types.String `tfsdk:"language"`
	LanguagesURL                  types.String `tfsdk:"languages_url"`
	Link                          types.String `tfsdk:"link"`
	Mirror                        types.Bool   `tfsdk:"mirror"`
	MirrorInterval                types.String `tfsdk:"mirror_interval"`
	MirrorUpdated                 types.String `tfsdk:"mirror_updated"`
	OpenIssues                    types.Int64  `tfsdk:"open_issues_count"`
	OpenPulls                     types.Int64  `tfsdk:"open_pr_counter"`
	OriginalURL                   types.String `tfsdk:"original_url"`
	Private                       types.Bool   `tfsdk:"private"`
	Releases                      types.Int64  `tfsdk:"release_counter"`
	SSHURL                        types.String `tfsdk:"ssh_url"`
	Size                          types.Int64  `tfsdk:"size"`
	Stars                         types.Int64  `tfsdk:"stars_count"`
	Template                      types.Bool   `tfsdk:"template"`
	Updated                       types.String `tfsdk:"updated_at"`
	Watchers                      types.Int64  `tfsdk:"watchers_count"`
	Website                       types.String `tfsdk:"website"`
}
