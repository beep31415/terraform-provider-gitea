package adapters

import (
	"context"
	"terraform-provider-gitea/api"
)

type RepositoryAdapter struct {
	client *api.APIClient
}

type RepositoryApiOptions struct {
	Owner                         string
	Name                          string
	AutoInit                      *bool
	DefaultBranch                 *string
	Description                   *string
	Gitignores                    *string
	IssueLabels                   *string
	License                       *string
	Private                       *bool
	Readme                        *string
	Template                      *bool
	TrustModel                    *string
	AllowManualMerge              *bool
	AllowMergeCommits             *bool
	AllowRebase                   *bool
	AllowRebaseExplicit           *bool
	AllowRebaseUpdate             *bool
	AllowSquashMerge              *bool
	Archived                      *bool
	AutodetectManualMerge         *bool
	DefaultAllowMaintainerEdit    *bool
	DefaultDeleteBranchAfterMerge *bool
	DefaultMergeStyle             *string
	EnablePrune                   *bool
	HasIssues                     *bool
	HasProjects                   *bool
	HasPullRequests               *bool
	HasWiki                       *bool
	IgnoreWhitespaceConflicts     *bool
	Website                       *string
}

func NewRepositorydapter(client *api.APIClient) *RepositoryAdapter {
	return &RepositoryAdapter{
		client: client,
	}
}

func (r *RepositoryAdapter) GetByOwnerAndName(ctx context.Context, owner, name string) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGet(ctx, owner, name).
		Execute()

	return res, err
}

func (r *RepositoryAdapter) Create(ctx context.Context, options RepositoryApiOptions) (*api.Repository, error) {
	org, err := r.getOrgByName(ctx, options.Owner)
	if err != nil {
		return nil, err
	}

	createOptions := api.CreateRepoOption{
		Name:          options.Name,
		AutoInit:      options.AutoInit,
		DefaultBranch: options.DefaultBranch,
		Description:   options.Description,
		Gitignores:    options.Gitignores,
		IssueLabels:   options.IssueLabels,
		License:       options.License,
		Private:       options.Private,
		Readme:        options.Readme,
		Template:      options.Template,
		TrustModel:    options.TrustModel,
	}

	var res *api.Repository
	if org != nil {
		res, err = r.createOrgRepo(ctx, options.Owner, createOptions)
	} else {
		res, err = r.createCurrentUserRepo(ctx, createOptions)
	}
	if err != nil {
		return res, err
	}

	return r.Edit(ctx, options)
}

func (r *RepositoryAdapter) Edit(ctx context.Context, options RepositoryApiOptions) (*api.Repository, error) {
	editOptions := api.EditRepoOption{
		Name:                          &options.Name,
		AllowManualMerge:              options.AllowManualMerge,
		AllowMergeCommits:             options.AllowMergeCommits,
		AllowRebase:                   options.AllowRebase,
		AllowRebaseExplicit:           options.AllowRebaseExplicit,
		AllowRebaseUpdate:             options.AllowRebaseUpdate,
		AllowSquashMerge:              options.AllowSquashMerge,
		Archived:                      options.Archived,
		AutodetectManualMerge:         options.AutodetectManualMerge,
		DefaultAllowMaintainerEdit:    options.DefaultAllowMaintainerEdit,
		DefaultBranch:                 options.DefaultBranch,
		DefaultDeleteBranchAfterMerge: options.DefaultDeleteBranchAfterMerge,
		DefaultMergeStyle:             options.DefaultMergeStyle,
		Description:                   options.Description,
		EnablePrune:                   options.EnablePrune,
		HasIssues:                     options.HasIssues,
		HasProjects:                   options.HasProjects,
		HasPullRequests:               options.HasPullRequests,
		HasWiki:                       options.HasWiki,
		IgnoreWhitespaceConflicts:     options.IgnoreWhitespaceConflicts,
		Private:                       options.Private,
		Template:                      options.Template,
		Website:                       options.Website,
	}

	res, _, err := r.client.RepositoryAPI.
		RepoEdit(ctx, options.Owner, options.Name).
		Body(editOptions).
		Execute()

	return res, err
}

func (r *RepositoryAdapter) Delete(ctx context.Context, id int64) error {
	res, err := r.getById(ctx, id)
	if err != nil {
		if IsErrorNotFound(err) {
			return nil
		}

		return err
	}

	_, err = r.client.RepositoryAPI.
		RepoDelete(ctx, *res.GetOwner().Login, res.GetName()).
		Execute()

	return err
}

func (r *RepositoryAdapter) getById(ctx context.Context, id int64) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGetByID(ctx, id).
		Execute()

	return res, err
}

func (r *RepositoryAdapter) createOrgRepo(ctx context.Context, org string, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.OrganizationAPI.
		CreateOrgRepo(ctx, org).
		Body(opts).
		Execute()

	return res, err
}

func (r *RepositoryAdapter) createCurrentUserRepo(ctx context.Context, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		CreateCurrentUserRepo(ctx).
		Body(opts).
		Execute()

	return res, err
}

func (r *RepositoryAdapter) getOrgByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	if IsErrorNotFound(err) {
		res = nil
		err = nil
	}

	return res, err
}
