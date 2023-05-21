package provider

import (
	"context"
	"strings"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &repoDataSource{}
	_ datasource.DataSourceWithConfigure = &repoDataSource{}
)

type repoDataSource struct {
	client *api.APIClient
}

func NewRepoDataSource() datasource.DataSource {
	return &repoDataSource{}
}

func (d *repoDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repo"
}

func (d *repoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a Gitea organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
			},
			"owner": schema.StringAttribute{
				Description: "The owner of the repository without spaces. " +
					"`!!IMPORTANT` Note that if owner is not an organization it has to be set as " +
					"the user related to the terrafom credentials for this provider since " +
					"Gitea does not offer an API to create repos owned by other users.",
				Required: true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the repository without spaces.",
				Required:    true,
			},
			"allow_merge_commits": schema.BoolAttribute{
				Description: "Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits. Defaults to `true`.",
				Computed:    true,
			},
			"allow_rebase": schema.BoolAttribute{
				Description: "Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging. Defaults to `true`.",
				Computed:    true,
			},
			"allow_rebase_explicit": schema.BoolAttribute{
				Description: "Either `true` to allow rebase with explicit merge commits (--no-ff), or `false` to prevent rebase with explicit merge commits. Defaults to `true`.",
				Computed:    true,
			},
			"allow_rebase_update": schema.BoolAttribute{
				Description: "Either `true` to allow updating pull request branch by rebase, or `false` to prevent it. Defaults to `true`.",
				Computed:    true,
			},
			"allow_squash_merge": schema.BoolAttribute{
				Description: "Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging. Defaults to `true`.",
				Computed:    true,
			},
			"archived": schema.BoolAttribute{
				Description: "Set to `true` to archive this repository. Defaults to `false`.",
				Computed:    true,
			},
			"avatar_url": schema.StringAttribute{
				Description: "The avatar URL.",
				Computed:    true,
			},
			"clone_url": schema.StringAttribute{
				Description: "The clone URL.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Date for the creation of the repository.",
				Computed:    true,
			},
			"default_allow_maintainer_edit": schema.BoolAttribute{
				Description: "Set to `true` to allow edits from maintainers by default. Defaults to `false`.",
				Computed:    true,
			},
			"default_branch": schema.StringAttribute{
				Description: "DefaultBranch of the repository (used when initializes and in template). Defaults to `master`.",
				Computed:    true,
			},
			"default_delete_branch_after_merge": schema.BoolAttribute{
				Description: "Set to `true` to delete pr branch after merge by default. Defaults to `false`.",
				Computed:    true,
			},
			"default_merge_style": schema.StringAttribute{
				Description: "Set to a merge style to be used by this repository. Can be: merge, rebase, rebase-merge or squash. Defaults to `merge`.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the repository to create.",
				Computed:    true,
			},
			"empty": schema.BoolAttribute{
				Description: "Empty. (???)",
				Computed:    true,
			},
			"fork": schema.BoolAttribute{
				Description: "Flag indicating if repository was forked.",
				Computed:    true,
			},
			"forks_count": schema.Int64Attribute{
				Description: "Fork count.",
				Computed:    true,
			},
			"full_name": schema.StringAttribute{
				Description: "The full name of the repository.",
				Computed:    true,
			},
			"html_url": schema.StringAttribute{
				Description: "The HTML clone url.",
				Computed:    true,
			},
			"has_issues": schema.BoolAttribute{
				Description: "Either `true` to enable issues for this repository or `false` to disable them. Defaults to `true`.",
				Computed:    true,
			},
			"has_projects": schema.BoolAttribute{
				Description: "Either `true` to enable project unit, or `false` to disable them. Defaults to `true`.",
				Computed:    true,
			},
			"has_pull_requests": schema.BoolAttribute{
				Description: "Either `true` to allow pull requests, or `false` to prevent pull request. Defaults to `true`.",
				Computed:    true,
			},
			"has_wiki": schema.BoolAttribute{
				Description: "Either `true` to enable the wiki for this repository or `false` to disable it. Defaults to `true`.",
				Computed:    true,
			},
			"ignore_whitespace_conflicts": schema.BoolAttribute{
				Description: "Either `true` to ignore whitespace for conflicts, or `false` to not ignore whitespace. Defaults to `false`.",
				Computed:    true,
			},
			"internal": schema.BoolAttribute{
				Description: "Internal. (???)",
				Computed:    true,
			},
			"language": schema.StringAttribute{
				Description: "Detected language of repository code.",
				Computed:    true,
			},
			"languages_url": schema.StringAttribute{
				Description: "Detected language URL of repository code.",
				Computed:    true,
			},
			"link": schema.StringAttribute{
				Description: "Link. (???)",
				Computed:    true,
			},
			"mirror": schema.BoolAttribute{
				Description: "Repository mirror.",
				Computed:    true,
			},
			"mirror_interval": schema.StringAttribute{
				Description: "Repository mirror update interval.",
				Computed:    true,
			},
			"mirror_updated": schema.StringAttribute{
				Description: "Repository mirror last updated.",
				Computed:    true,
			},
			"open_issues_count": schema.Int64Attribute{
				Description: "Number of open issues in repository.",
				Computed:    true,
			},
			"open_pr_counter": schema.Int64Attribute{
				Description: "Number of open pull-requests in the repository.",
				Computed:    true,
			},
			"original_url": schema.StringAttribute{
				Description: "Original repository URL.",
				Computed:    true,
			},
			"private": schema.BoolAttribute{
				Description: "Whether the repository is private. Defaults to `false`.",
				Computed:    true,
			},
			"release_counter": schema.Int64Attribute{
				Description: "Number of releases.",
				Computed:    true,
			},
			"ssh_url": schema.StringAttribute{
				Description: "SSH link for repository clone.",
				Computed:    true,
			},
			"size": schema.Int64Attribute{
				Description: "Repository size.",
				Computed:    true,
			},
			"stars_count": schema.Int64Attribute{
				Description: "Awarded stars for current repostiry.",
				Computed:    true,
			},
			"template": schema.BoolAttribute{
				Description: "Whether the repository is template. Defaults to `false`.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Repository last update date.",
				Computed:    true,
			},
			"watchers_count": schema.Int64Attribute{
				Description: "Number of repository watchers.",
				Computed:    true,
			},
			"website": schema.StringAttribute{
				Description: "A URL with more information about the repository.",
				Computed:    true,
			},
		},
	}
}

func (d *repoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.APIClient)
}

func (d *repoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.RepositoryDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, _, err := d.client.RepositoryAPI.
		RepoGet(ctx, strings.ToLower(state.Owner.ValueString()), strings.ToLower(state.Name.ValueString())).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Gitea repository.",
			adapters.GetAPIErrorMessage(err),
		)

		return
	}

	state = models.RepositoryDataSourceModel{
		ID:                            types.Int64Value(res.GetId()),
		Owner:                         types.StringValue(*res.GetOwner().Login),
		Name:                          types.StringValue(res.GetName()),
		AllowMerge:                    types.BoolValue(res.GetAllowMergeCommits()),
		AllowRebase:                   types.BoolValue(res.GetAllowRebase()),
		AllowRebaseMerge:              types.BoolValue(res.GetAllowRebaseExplicit()),
		AllowRebaseUpdate:             types.BoolValue(res.GetAllowRebaseUpdate()),
		AllowSquash:                   types.BoolValue(res.GetAllowSquashMerge()),
		Archived:                      types.BoolValue(res.GetArchived()),
		AvatarURL:                     types.StringValue(res.GetAvatarUrl()),
		CloneURL:                      types.StringValue(res.GetCloneUrl()),
		Created:                       types.StringValue(res.GetCreatedAt().String()),
		DefaultAllowMaintainerEdit:    types.BoolValue(res.GetDefaultAllowMaintainerEdit()),
		DefaultBranch:                 types.StringValue(res.GetDefaultBranch()),
		DefaultDeleteBranchAfterMerge: types.BoolValue(res.GetDefaultDeleteBranchAfterMerge()),
		DefaultMergeStyle:             types.StringValue(res.GetDefaultMergeStyle()),
		Description:                   types.StringValue(res.GetDescription()),
		Empty:                         types.BoolValue(res.GetEmpty()),
		Fork:                          types.BoolValue(res.GetFork()),
		Forks:                         types.Int64Value(res.GetForksCount()),
		FullName:                      types.StringValue(res.GetFullName()),
		HTMLURL:                       types.StringValue(res.GetHtmlUrl()),
		HasIssues:                     types.BoolValue(res.GetHasIssues()),
		HasProjects:                   types.BoolValue(res.GetHasProjects()),
		HasPullRequests:               types.BoolValue(res.GetHasPullRequests()),
		HasWiki:                       types.BoolValue(res.GetHasWiki()),
		IgnoreWhitespaceConflicts:     types.BoolValue(res.GetIgnoreWhitespaceConflicts()),
		Internal:                      types.BoolValue(res.GetInternal()),
		Language:                      types.StringValue(res.GetLanguage()),
		LanguagesURL:                  types.StringValue(res.GetLanguagesUrl()),
		Link:                          types.StringValue(res.GetLink()),
		Mirror:                        types.BoolValue(res.GetMirror()),
		MirrorInterval:                types.StringValue(res.GetMirrorInterval()),
		MirrorUpdated:                 types.StringValue(res.GetMirrorUpdated().String()),
		OpenIssues:                    types.Int64Value(res.GetOpenIssuesCount()),
		OpenPulls:                     types.Int64Value(res.GetOpenPrCounter()),
		OriginalURL:                   types.StringValue(res.GetOriginalUrl()),
		Private:                       types.BoolValue(res.GetPrivate()),
		Releases:                      types.Int64Value(res.GetReleaseCounter()),
		SSHURL:                        types.StringValue(res.GetSshUrl()),
		Size:                          types.Int64Value(res.GetSize()),
		Stars:                         types.Int64Value(res.GetStarsCount()),
		Template:                      types.BoolValue(res.GetTemplate()),
		Updated:                       types.StringValue(res.GetUpdatedAt().String()),
		Watchers:                      types.Int64Value(res.GetWatchersCount()),
		Website:                       types.StringValue(res.GetWebsite()),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
