package proxy

import (
	"context"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/proxies"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type ProxyDataSource[T any] interface {
	FillDataSource(ctx context.Context, model T) diag.Diagnostics
}

type ProxyResource[T any] interface {
	FillResource(ctx context.Context, model T) diag.Diagnostics
	Create(ctx context.Context, model T) diag.Diagnostics
	Update(ctx context.Context, model T) diag.Diagnostics
	Delete(ctx context.Context, model T) diag.Diagnostics
}

type Factory struct {
	apiClient *api.APIClient
}

func NewFactory(apiClient *api.APIClient) *Factory {
	return &Factory{
		apiClient: apiClient,
	}
}

func (f *Factory) GetBranchProtectionDataSourceProxy() ProxyDataSource[models.BranchProtectionDataSourceModel] {
	return proxies.NewBranchProtectionProxy(f.apiClient)
}

func (f *Factory) GetBranchProtectionResourceProxy() ProxyResource[models.BranchProtectionResourceModel] {
	return proxies.NewBranchProtectionProxy(f.apiClient)
}

func (f *Factory) GetOrganizationDataSourceProxy() ProxyDataSource[models.OrganizationDataSourceModel] {
	return proxies.NewOrganizationProxy(f.apiClient)
}

func (f *Factory) GetOrganizationResourceProxy() ProxyResource[models.OrganizationResourceModel] {
	return proxies.NewOrganizationProxy(f.apiClient)
}

func (f *Factory) GetRepositoryDataSourceProxy() ProxyDataSource[models.RepositoryDataSourceModel] {
	return proxies.NewRepositoryProxy(f.apiClient)
}

func (f *Factory) GetRepositoryResourceProxy() ProxyResource[models.RepositoryResourceModel] {
	return proxies.NewRepositoryProxy(f.apiClient)
}

func (f *Factory) GetTeamDataSourceProxy() ProxyDataSource[models.TeamDataSourceModel] {
	return proxies.NewTeamProxy(f.apiClient)
}

func (f *Factory) GetTeamResourceProxy() ProxyResource[models.TeamResourceModel] {
	return proxies.NewTeamProxy(f.apiClient)
}

func (f *Factory) GetUserDataSourceProxy() ProxyDataSource[models.UserDataSourceModel] {
	return proxies.NewUserProxy(f.apiClient)
}
