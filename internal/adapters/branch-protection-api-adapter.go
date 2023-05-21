package adapters

import (
	"context"
	"terraform-provider-gitea/api"
)

type BranchProtectionAdapter struct {
	client *api.APIClient
}

type BranchProtection struct {
	Owner            string
	Repo             string
	BranchProtection *api.BranchProtection
}

type AddBranchProtectionOptions struct {
	Owner     string
	Repo      string
	AddOption api.CreateBranchProtectionOption
}

type EditBranchProtectionOptions struct {
	Owner      string
	Repo       string
	BranchName string
	EditOption api.EditBranchProtectionOption
}

func NewBranchProtectionAdapter(client *api.APIClient) *BranchProtectionAdapter {
	return &BranchProtectionAdapter{
		client: client,
	}
}

func (b *BranchProtectionAdapter) GetByOwnerRepoAndBranchName(ctx context.Context, owner, repo, branchName string) (*BranchProtection, error) {
	res, _, err := b.client.RepositoryAPI.
		RepoGetBranchProtection(ctx, owner, repo, branchName).
		Execute()

	return &BranchProtection{
		Owner:            owner,
		Repo:             repo,
		BranchProtection: res,
	}, err
}

func (b *BranchProtectionAdapter) Create(ctx context.Context, options AddBranchProtectionOptions) (*BranchProtection, error) {
	res, _, err := b.client.RepositoryAPI.
		RepoCreateBranchProtection(ctx, options.Owner, options.Repo).
		Body(options.AddOption).
		Execute()

	return &BranchProtection{
		Owner:            options.Owner,
		Repo:             options.Repo,
		BranchProtection: res,
	}, err
}

func (b *BranchProtectionAdapter) Edit(ctx context.Context, options EditBranchProtectionOptions) (*BranchProtection, error) {
	res, _, err := b.client.RepositoryAPI.
		RepoEditBranchProtection(ctx, options.Owner, options.Repo, options.BranchName).
		Body(options.EditOption).
		Execute()

	return &BranchProtection{
		Owner:            options.Owner,
		Repo:             options.Repo,
		BranchProtection: res,
	}, err
}

func (b *BranchProtectionAdapter) Delete(ctx context.Context, owner, repo, branchName string) error {
	_, err := b.GetByOwnerRepoAndBranchName(ctx, owner, repo, branchName)
	if err != nil {
		return err
	}

	_, err = b.client.RepositoryAPI.
		RepoDeleteBranchProtection(ctx, owner, repo, branchName).
		Execute()

	return err
}
