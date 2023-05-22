package proxy

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/proxies"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type ProxyDataSource[T any] interface {
	FillDataSource(ctx context.Context, model *T) diag.Diagnostics
}

type ProxyResource[T any] interface {
	FillResource(ctx context.Context, model *T) diag.Diagnostics
	Create(ctx context.Context, model *T) diag.Diagnostics
	Update(ctx context.Context, model *T) diag.Diagnostics
	Delete(ctx context.Context, model *T) diag.Diagnostics
}

type Factory struct {
	apiClient *api.APIClient
}

type ProxyModel struct {
	Scheme       string
	Host         string
	Token        string
	Username     string
	Password     string
	CACertFile   string
	Insecure     bool
	Debug        bool
	DebugLogPath string
}

func NewFactory(model ProxyModel) (*Factory, error) {
	apiClient, err := newAPIClient(model)
	if err != nil {
		return nil, err
	}

	return &Factory{
		apiClient: apiClient,
	}, nil
}

func (f *Factory) NewBranchProtectionDataSourceProxy() ProxyDataSource[models.BranchProtectionDataSourceModel] {
	return proxies.NewBranchProtectionProxy(f.apiClient)
}

func (f *Factory) NewBranchProtectionResourceProxy() ProxyResource[models.BranchProtectionResourceModel] {
	return proxies.NewBranchProtectionProxy(f.apiClient)
}

func (f *Factory) NewOrganizationDataSourceProxy() ProxyDataSource[models.OrganizationDataSourceModel] {
	return proxies.NewOrganizationProxy(f.apiClient)
}

func (f *Factory) NewOrganizationResourceProxy() ProxyResource[models.OrganizationResourceModel] {
	return proxies.NewOrganizationProxy(f.apiClient)
}

func (f *Factory) NewRepositoryDataSourceProxy() ProxyDataSource[models.RepositoryDataSourceModel] {
	return proxies.NewRepositoryProxy(f.apiClient)
}

func (f *Factory) NewRepositoryResourceProxy() ProxyResource[models.RepositoryResourceModel] {
	return proxies.NewRepositoryProxy(f.apiClient)
}

func (f *Factory) NewTeamDataSourceProxy() ProxyDataSource[models.TeamDataSourceModel] {
	return proxies.NewTeamProxy(f.apiClient)
}

func (f *Factory) NewTeamResourceProxy() ProxyResource[models.TeamResourceModel] {
	return proxies.NewTeamProxy(f.apiClient)
}

func (f *Factory) NewUserDataSourceProxy() ProxyDataSource[models.UserDataSourceModel] {
	return proxies.NewUserProxy(f.apiClient)
}

func newAPIClient(model ProxyModel) (*api.APIClient, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: model.Insecure,
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig

	apiClientConfig := api.NewConfiguration()
	apiClientConfig.Host = model.Host
	apiClientConfig.Scheme = model.Scheme
	apiClientConfig.HTTPClient = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	apiClientConfig.Debug = model.Debug
	if apiClientConfig.Debug {
		file, err := os.OpenFile(model.DebugLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}

		log.SetOutput(file)
	}

	apiClient := api.NewAPIClient(apiClientConfig)

	if model.Username != "" {
		basicAuthToken := model.Username + ":" + model.Password
		basicAuthToken = base64.RawStdEncoding.EncodeToString([]byte(basicAuthToken))
		apiClientConfig.AddDefaultHeader("Authorization", "Basic "+basicAuthToken)
	} else {
		apiClientConfig.AddDefaultHeader("Authorization", "token "+model.Token)
	}

	if model.CACertFile != "" {
		cac, err := os.ReadFile(model.CACertFile)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cac)
		tlsConfig.RootCAs = caCertPool
	}

	return apiClient, nil
}
