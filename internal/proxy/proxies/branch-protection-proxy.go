package proxies

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"
	"terraform-provider-gitea/internal/proxy/errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type branchProtectionProxy struct {
	client    *api.APIClient
	converter converters.BranchProtectionConverter
}

func NewBranchProtectionProxy(client *api.APIClient) *branchProtectionProxy {
	return &branchProtectionProxy{
		client:    client,
		converter: converters.BranchProtectionConverter{},
	}
}

func (b *branchProtectionProxy) FillDataSource(ctx context.Context, model models.BranchProtectionDataSourceModel) diag.Diagnostics {
	res, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea branch protection.")
	}

	return b.converter.ReadToDataSource(ctx, model, res)
}

func (b *branchProtectionProxy) FillResource(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	res, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *branchProtectionProxy) Create(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	option, diags := b.converter.ToApiAddBranchProtectionOptions(ctx, model)
	if diags.HasError() {
		return diags
	}

	res, _, err := b.client.RepositoryAPI.
		RepoCreateBranchProtection(ctx, model.Owner.ValueString(), model.Repo.ValueString()).
		Body(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Creating Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *branchProtectionProxy) Update(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	option, diags := b.converter.ToApiEditBranchProtectionOptions(ctx, model)
	if diags.HasError() {
		return diags
	}

	res, _, err := b.client.RepositoryAPI.
		RepoEditBranchProtection(ctx,
			model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString()).
		Body(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Updating Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *branchProtectionProxy) Delete(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	_, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		if errors.IsErrorNotFound(err) {
			return nil
		}

		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea branch protection.")
	}

	_, err = b.client.RepositoryAPI.
		RepoDeleteBranchProtection(ctx,
			model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString()).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea branch protection.")
	}

	return nil
}

func (b *branchProtectionProxy) getByOwnerRepoAndBranchName(ctx context.Context, owner, repo, branchName string) (*api.BranchProtection, error) {
	res, _, err := b.client.RepositoryAPI.
		RepoGetBranchProtection(ctx, owner, repo, branchName).
		Execute()

	return res, err
}
