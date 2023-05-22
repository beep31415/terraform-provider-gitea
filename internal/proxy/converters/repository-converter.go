package converters

import (
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RepositoryConverter struct{}

func (RepositoryConverter) ReadToDataSource(model *models.RepositoryDataSourceModel, repository *api.Repository) {
	model.ID = types.Int64Value(repository.GetId())
	model.Owner = types.StringValue(*repository.GetOwner().Login)
	model.Name = types.StringValue(repository.GetName())
	model.AllowMerge = types.BoolValue(repository.GetAllowMergeCommits())
	model.AllowRebase = types.BoolValue(repository.GetAllowRebase())
	model.AllowRebaseMerge = types.BoolValue(repository.GetAllowRebaseExplicit())
	model.AllowRebaseUpdate = types.BoolValue(repository.GetAllowRebaseUpdate())
	model.AllowSquash = types.BoolValue(repository.GetAllowSquashMerge())
	model.Archived = types.BoolValue(repository.GetArchived())
	model.AvatarURL = types.StringValue(repository.GetAvatarUrl())
	model.CloneURL = types.StringValue(repository.GetCloneUrl())
	model.Created = types.StringValue(repository.GetCreatedAt().String())
	model.DefaultAllowMaintainerEdit = types.BoolValue(repository.GetDefaultAllowMaintainerEdit())
	model.DefaultBranch = types.StringValue(repository.GetDefaultBranch())
	model.DefaultDeleteBranchAfterMerge = types.BoolValue(repository.GetDefaultDeleteBranchAfterMerge())
	model.DefaultMergeStyle = types.StringValue(repository.GetDefaultMergeStyle())
	model.Description = types.StringValue(repository.GetDescription())
	model.Empty = types.BoolValue(repository.GetEmpty())
	model.Fork = types.BoolValue(repository.GetFork())
	model.Forks = types.Int64Value(repository.GetForksCount())
	model.FullName = types.StringValue(repository.GetFullName())
	model.HTMLURL = types.StringValue(repository.GetHtmlUrl())
	model.HasIssues = types.BoolValue(repository.GetHasIssues())
	model.HasProjects = types.BoolValue(repository.GetHasProjects())
	model.HasPullRequests = types.BoolValue(repository.GetHasPullRequests())
	model.HasWiki = types.BoolValue(repository.GetHasWiki())
	model.IgnoreWhitespaceConflicts = types.BoolValue(repository.GetIgnoreWhitespaceConflicts())
	model.Internal = types.BoolValue(repository.GetInternal())
	model.Language = types.StringValue(repository.GetLanguage())
	model.LanguagesURL = types.StringValue(repository.GetLanguagesUrl())
	model.Link = types.StringValue(repository.GetLink())
	model.Mirror = types.BoolValue(repository.GetMirror())
	model.MirrorInterval = types.StringValue(repository.GetMirrorInterval())
	model.MirrorUpdated = types.StringValue(repository.GetMirrorUpdated().String())
	model.OpenIssues = types.Int64Value(repository.GetOpenIssuesCount())
	model.OpenPulls = types.Int64Value(repository.GetOpenPrCounter())
	model.OriginalURL = types.StringValue(repository.GetOriginalUrl())
	model.Private = types.BoolValue(repository.GetPrivate())
	model.Releases = types.Int64Value(repository.GetReleaseCounter())
	model.SSHURL = types.StringValue(repository.GetSshUrl())
	model.Size = types.Int64Value(repository.GetSize())
	model.Stars = types.Int64Value(repository.GetStarsCount())
	model.Template = types.BoolValue(repository.GetTemplate())
	model.Updated = types.StringValue(repository.GetUpdatedAt().String())
	model.Watchers = types.Int64Value(repository.GetWatchersCount())
	model.Website = types.StringValue(repository.GetWebsite())
}

func (RepositoryConverter) ReadToResource(model *models.RepositoryResourceModel, repository *api.Repository) {
	model.ID = types.Int64Value(repository.GetId())
	model.Name = types.StringValue(repository.GetName())
	model.Owner = types.StringValue(*repository.GetOwner().Login)
	model.DefaultBranch = types.StringValue(repository.GetDefaultBranch())
	model.Description = types.StringValue(repository.GetDescription())
	model.Private = types.BoolValue(repository.GetPrivate())
	model.Template = types.BoolValue(repository.GetTemplate())
	model.AllowMerge = types.BoolValue(repository.GetAllowMergeCommits())
	model.AllowRebase = types.BoolValue(repository.GetAllowRebase())
	model.AllowRebaseMerge = types.BoolValue(repository.GetAllowRebaseExplicit())
	model.AllowRebaseUpdate = types.BoolValue(repository.GetAllowRebaseUpdate())
	model.AllowSquash = types.BoolValue(repository.GetAllowSquashMerge())
	model.Archived = types.BoolValue(repository.GetArchived())
	model.DefaultAllowMaintainerEdit = types.BoolValue(repository.GetDefaultAllowMaintainerEdit())
	model.DefaultDeleteBranchAfterMerge = types.BoolValue(repository.GetDefaultDeleteBranchAfterMerge())
	model.DefaultMergeStyle = types.StringValue(repository.GetDefaultMergeStyle())
	model.HasIssues = types.BoolValue(repository.GetHasIssues())
	model.HasProjects = types.BoolValue(repository.GetHasProjects())
	model.HasPullRequests = types.BoolValue(repository.GetHasPullRequests())
	model.HasWiki = types.BoolValue(repository.GetHasWiki())
	model.IgnoreWhitespaceConflicts = types.BoolValue(repository.GetIgnoreWhitespaceConflicts())
	model.Website = types.StringValue(repository.GetWebsite())
	if model.AutoInit.IsNull() {
		model.AutoInit = types.BoolValue(false)
	}
	if model.Gitignores.IsNull() {
		model.Gitignores = types.StringValue("")
	}
	if model.IssueLabels.IsNull() {
		model.IssueLabels = types.StringValue("")
	}
	if model.License.IsNull() {
		model.License = types.StringValue("")
	}
	if model.Readme.IsNull() {
		model.Readme = types.StringValue("")
	}
	if model.AllowManualMerge.IsNull() {
		model.AllowManualMerge = types.BoolValue(false)
	}
	if model.AutodetectManualMerge.IsNull() {
		model.AutodetectManualMerge = types.BoolValue(false)
	}
	if model.EnablePrune.IsNull() {
		model.EnablePrune = types.BoolValue(false)
	}
	if model.TrustModel.IsNull() {
		model.TrustModel = types.StringValue("default")
	}
}

func (RepositoryConverter) ToApiAddRepositoryOptions(model models.RepositoryResourceModel) api.CreateRepoOption {
	return api.CreateRepoOption{
		Name:          model.Name.ValueString(),
		AutoInit:      model.AutoInit.ValueBoolPointer(),
		DefaultBranch: model.DefaultBranch.ValueStringPointer(),
		Description:   model.Description.ValueStringPointer(),
		Gitignores:    model.Gitignores.ValueStringPointer(),
		IssueLabels:   model.IssueLabels.ValueStringPointer(),
		License:       model.License.ValueStringPointer(),
		Private:       model.Private.ValueBoolPointer(),
		Readme:        model.Readme.ValueStringPointer(),
		Template:      model.Template.ValueBoolPointer(),
		TrustModel:    model.TrustModel.ValueStringPointer(),
	}
}

func (RepositoryConverter) ToApiEditRepositoryOptions(model models.RepositoryResourceModel) api.EditRepoOption {
	return api.EditRepoOption{
		Name:                          model.Name.ValueStringPointer(),
		AllowManualMerge:              model.AllowManualMerge.ValueBoolPointer(),
		AllowMergeCommits:             model.AllowMerge.ValueBoolPointer(),
		AllowRebase:                   model.AllowRebase.ValueBoolPointer(),
		AllowRebaseExplicit:           model.AllowRebaseMerge.ValueBoolPointer(),
		AllowRebaseUpdate:             model.AllowRebaseUpdate.ValueBoolPointer(),
		AllowSquashMerge:              model.AllowSquash.ValueBoolPointer(),
		Archived:                      model.Archived.ValueBoolPointer(),
		AutodetectManualMerge:         model.AutodetectManualMerge.ValueBoolPointer(),
		DefaultAllowMaintainerEdit:    model.DefaultAllowMaintainerEdit.ValueBoolPointer(),
		DefaultBranch:                 model.DefaultBranch.ValueStringPointer(),
		DefaultDeleteBranchAfterMerge: model.DefaultDeleteBranchAfterMerge.ValueBoolPointer(),
		DefaultMergeStyle:             model.DefaultMergeStyle.ValueStringPointer(),
		Description:                   model.Description.ValueStringPointer(),
		EnablePrune:                   model.EnablePrune.ValueBoolPointer(),
		HasIssues:                     model.HasIssues.ValueBoolPointer(),
		HasProjects:                   model.HasProjects.ValueBoolPointer(),
		HasPullRequests:               model.HasPullRequests.ValueBoolPointer(),
		HasWiki:                       model.HasWiki.ValueBoolPointer(),
		IgnoreWhitespaceConflicts:     model.IgnoreWhitespaceConflicts.ValueBoolPointer(),
		Private:                       model.Private.ValueBoolPointer(),
		Template:                      model.Template.ValueBoolPointer(),
		Website:                       model.Website.ValueStringPointer(),
	}
}
