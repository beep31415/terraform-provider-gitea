package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"time"

	"terraform-provider-gitea/internal/proxy"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/errors"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = &giteaProvider{}
)

type giteaProvider struct{}

func New() provider.Provider {
	return &giteaProvider{}
}

func (p *giteaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gitea"
}

func (p *giteaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Gitea client.")

	var cfg giteaProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg.loadEnvs()
	resp.Diagnostics.Append(cfg.validate())
	if resp.Diagnostics.HasError() {
		return
	}

	apiClient, err := p.newAPIClient(cfg)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Gitea API Client.",
			"An unexpected error occurred when creating the Gitea API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Gitea Client Error: "+errors.GetApiInitError(err),
		)

		return
	}

	factory := proxy.NewFactory(apiClient)
	resp.DataSourceData = factory
	resp.ResourceData = factory

	tflog.Info(ctx, "Configured Gitea client.", map[string]any{"success": true})
}

func (p *giteaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrgDataSource,
		NewUserDataSource,
		NewRepoDataSource,
		NewBranchProtectionDataSource,
		NewTeamDataSource,
	}
}

func (p *giteaProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrgResource,
		NewRepoResource,
		NewBranchProtectionResource,
		NewTeamResource,
	}
}

func (p *giteaProvider) newAPIClient(cfg giteaProviderModel) (*api.APIClient, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: cfg.Insecure.ValueBool(),
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig

	apiClientConfig := api.NewConfiguration()
	apiClientConfig.Host = cfg.parsedUrl.Host
	apiClientConfig.Scheme = cfg.parsedUrl.Scheme
	apiClientConfig.HTTPClient = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	apiClientConfig.Debug = cfg.Debug.ValueBool()
	if apiClientConfig.Debug {
		file, err := os.OpenFile(cfg.DebugLogPath.ValueString(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}

		log.SetOutput(file)
	}

	apiClient := api.NewAPIClient(apiClientConfig)

	if cfg.Username.ValueString() != "" {
		basicAuthToken := cfg.Username.ValueString() + ":" + cfg.Password.ValueString()
		basicAuthToken = base64.RawStdEncoding.EncodeToString([]byte(basicAuthToken))
		apiClientConfig.AddDefaultHeader("Authorization", "Basic "+basicAuthToken)
	} else {
		apiClientConfig.AddDefaultHeader("Authorization", "token "+cfg.Token.ValueString())
	}

	if cfg.CACertFile.ValueString() != "" {
		cac, err := os.ReadFile(cfg.CACertFile.ValueString())
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cac)
		tlsConfig.RootCAs = caCertPool
	}

	return apiClient, nil
}

func (p *giteaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Gitea hosting server.",
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Gitea server base URL. Can also be set from GITEA_BASE_URL env.",
				Optional:    true,
			},
			"token": schema.StringAttribute{
				Description: "Gitea authentication token. Can also be set from GITEA_TOKEN env.",
				Optional:    true,
				Sensitive:   true,
			},
			"username": schema.StringAttribute{
				Description: "Gitea user name. Can also be set from GITEA_USERNAME env.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Gitea password. Can also be set from GITEA_PASSWORD env.",
				Optional:    true,
				Sensitive:   true,
			},
			"cacert_file": schema.StringAttribute{
				Description: "Gitea TLS trusted CA cert.",
				Optional:    true,
				Sensitive:   true,
			},
			"insecure": schema.BoolAttribute{
				Description: "Disable SSL verification of API calls.",
				Optional:    true,
			},
			"debug": schema.BoolAttribute{
				Description: "Indicates debug mode activation.",
				Optional:    true,
			},
			"debug_log_path": schema.StringAttribute{
				Description: "Path of file to write the debug logs to.",
				Optional:    true,
			},
		},
	}
}
