package proxies

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"
	"terraform-provider-gitea/internal/proxy/errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type RepositoryProxy struct {
	client    *api.APIClient
	converter converters.RepositoryConverter
}

func NewRepositoryProxy(client *api.APIClient) *RepositoryProxy {
	return &RepositoryProxy{
		client:    client,
		converter: converters.RepositoryConverter{},
	}
}

func (r *RepositoryProxy) FillDataSource(ctx context.Context, model models.RepositoryDataSourceModel) diag.Diagnostics {
	res, err := r.getByOwnerAndName(ctx, model.Owner.ValueString(), model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea repository.")
	}

	r.converter.ReadToDataSource(model, res)

	return nil
}

func (r *RepositoryProxy) FillResource(ctx context.Context, model models.RepositoryResourceModel) diag.Diagnostics {
	res, err := r.getByOwnerAndName(ctx, model.Owner.ValueString(), model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea repository.")
	}

	r.converter.ReadToResource(model, res)

	return nil
}

func (r *RepositoryProxy) Create(ctx context.Context, model models.RepositoryResourceModel) diag.Diagnostics {
	org, err := r.getOrgByName(ctx, model.Owner.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Creating Gitea repository.")
	}

	createOptions := r.converter.ToApiAddRepositoryOptions(model)
	var res *api.Repository

	if org != nil {
		res, err = r.createOrgRepo(ctx, model.Owner.ValueString(), createOptions)
	} else {
		res, err = r.createCurrentUserRepo(ctx, createOptions)
	}
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Creating Gitea repository.")
	}

	r.converter.ReadToResource(model, res)

	return r.Update(ctx, model)
}

func (r *RepositoryProxy) Update(ctx context.Context, model models.RepositoryResourceModel) diag.Diagnostics {
	editOptions := r.converter.ToApiEditRepositoryOptions(model)

	res, _, err := r.client.RepositoryAPI.
		RepoEdit(ctx, model.Owner.ValueString(), model.Name.ValueString()).
		Body(editOptions).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Updating Gitea repository.")
	}

	r.converter.ReadToResource(model, res)

	return nil
}

func (r *RepositoryProxy) Delete(ctx context.Context, model models.RepositoryResourceModel) diag.Diagnostics {
	res, err := r.getById(ctx, model.ID.ValueInt64())
	if err != nil {
		if errors.IsErrorNotFound(err) {
			return nil
		}

		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea repository.")
	}

	_, err = r.client.RepositoryAPI.
		RepoDelete(ctx, *res.GetOwner().Login, res.GetName()).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea repository.")
	}

	return nil
}

func (r *RepositoryProxy) getById(ctx context.Context, id int64) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGetByID(ctx, id).
		Execute()

	return res, err
}

func (r *RepositoryProxy) getByOwnerAndName(ctx context.Context, owner, name string) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGet(ctx, owner, name).
		Execute()

	return res, err
}

func (r *RepositoryProxy) createOrgRepo(ctx context.Context, org string, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.OrganizationAPI.
		CreateOrgRepo(ctx, org).
		Body(opts).
		Execute()

	return res, err
}

func (r *RepositoryProxy) createCurrentUserRepo(ctx context.Context, opts api.CreateRepoOption) (*api.Repository, error) {
	res, _, err := r.client.RepositoryAPI.
		CreateCurrentUserRepo(ctx).
		Body(opts).
		Execute()

	return res, err
}

func (r *RepositoryProxy) getOrgByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	if errors.IsErrorNotFound(err) {
		res = nil
		err = nil
	}

	return res, err
}
