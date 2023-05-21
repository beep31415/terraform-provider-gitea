package proxy

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type BranchProtectionProxy struct {
	client    *api.APIClient
	converter converters.BranchProtectionConverter
}

func NewBranchProtectionProxy(client *api.APIClient) *BranchProtectionProxy {
	return &BranchProtectionProxy{
		client:    client,
		converter: converters.BranchProtectionConverter{},
	}
}

func (b *BranchProtectionProxy) FillDataSource(ctx context.Context, model models.BranchProtectionDataSourceModel) diag.Diagnostics {
	res, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		return toDiagnosticArrayError(err, "Error Reading Gitea branch protection.")
	}

	return b.converter.ReadToDataSource(ctx, model, res)
}

func (b *BranchProtectionProxy) FillResource(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	res, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		return toDiagnosticArrayError(err, "Error Reading Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *BranchProtectionProxy) Create(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
	option, diags := b.converter.ToApiAddBranchProtectionOptions(ctx, model)
	if diags.HasError() {
		return diags
	}

	res, _, err := b.client.RepositoryAPI.
		RepoCreateBranchProtection(ctx, model.Owner.ValueString(), model.Repo.ValueString()).
		Body(option).
		Execute()
	if err != nil {
		return toDiagnosticArrayError(err, "Error Creating Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *BranchProtectionProxy) Edit(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostics {
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
		return toDiagnosticArrayError(err, "Error Updating Gitea branch protection.")
	}

	return b.converter.ReadToResource(ctx, model, res)
}

func (b *BranchProtectionProxy) Delete(ctx context.Context, model models.BranchProtectionResourceModel) diag.Diagnostic {
	_, err := b.getByOwnerRepoAndBranchName(ctx,
		model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString())
	if err != nil {
		if isErrorNotFound(err) {
			return nil
		}

		return toDiagnosticError(err, "Error Deleting Gitea branch protection.")
	}

	_, err = b.client.RepositoryAPI.
		RepoDeleteBranchProtection(ctx,
			model.Owner.ValueString(), model.Repo.ValueString(), model.BranchName.ValueString()).
		Execute()
	if err != nil {
		return toDiagnosticError(err, "Error Deleting Gitea branch protection.")
	}

	return nil
}

func (b *BranchProtectionProxy) getByOwnerRepoAndBranchName(ctx context.Context, owner, repo, branchName string) (*api.BranchProtection, error) {
	res, _, err := b.client.RepositoryAPI.
		RepoGetBranchProtection(ctx, owner, repo, branchName).
		Execute()

	return res, err
}
