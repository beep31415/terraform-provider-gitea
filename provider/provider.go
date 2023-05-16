package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"terraform-provider-gitea/api"
	"terraform-provider-gitea/provider/datasources"
	"terraform-provider-gitea/provider/resources"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = &giteaProvider{}
)

type giteaProvider struct{}

type giteaProviderModel struct {
	BaseURL      types.String `tfsdk:"base_url"`
	Token        types.String `tfsdk:"token"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	CACertFile   types.String `tfsdk:"cacert_file"`
	Insecure     types.Bool   `tfsdk:"insecure"`
	Debug        types.Bool   `tfsdk:"debug"`
	DebugLogPath types.String `tfsdk:"debug_log_path"`
}

func New() provider.Provider {
	return &giteaProvider{}
}

func (p *giteaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gitea"
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

func (p *giteaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Gitea client.")

	var cfg giteaProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if v, ok := os.LookupEnv("GITEA_BASE_URL"); ok {
		cfg.BaseURL = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_TOKEN"); ok {
		cfg.Token = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_USERNAME"); ok {
		cfg.Username = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_PASSWORD"); ok {
		cfg.Password = types.StringValue(v)
	}

	if cfg.BaseURL.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Invalid base url.",
			"Base url cannot be empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	var parsedUrl *url.URL
	var err error
	if parsedUrl, err = url.Parse(cfg.BaseURL.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Invalid base url.",
			fmt.Sprintf("Base url %s is not a valid url: %+v.", cfg.BaseURL.ValueString(), err),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if parsedUrl.Path != "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Invalid base url.",
			fmt.Sprintf("Base url %s should not contain path, got: %s.", cfg.BaseURL.ValueString(), parsedUrl.Path),
		)
	}

	if cfg.Username.ValueString() == "" && cfg.Token.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Invalid username or token.",
			"Username and token are empty. Must provide exacly one of them.",
		)
	}

	if cfg.Username.ValueString() != "" && cfg.Token.ValueString() != "" {
		resp.Diagnostics.AddError(
			"Invalid username or token.",
			"Username and token are both provided. Only one of them can be set.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if cfg.Username.ValueString() != "" && cfg.Password.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Invalid password.",
			"If username is provided password must also be set.",
		)
	}

	if cfg.Username.ValueString() == "" && cfg.Password.ValueString() != "" {
		resp.Diagnostics.AddError(
			"Invalid username.",
			"If password is provided username must also be set.",
		)
	}

	if cfg.Debug.ValueBool() && (cfg.DebugLogPath.ValueString() == "") {
		resp.Diagnostics.AddError(
			"Invalid debug log path.",
			"If debug is enabled a debug log path must be specified.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	apiClient, err := p.NewAPIClient(cfg)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Gitea API Client.",
			"An unexpected error occurred when creating the Gitea API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Gitea Client Error: "+err.Error(),
		)

		return
	}

	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient

	tflog.Info(ctx, "Configured Gitea client.", map[string]any{"success": true})
}

func (p *giteaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewOrgDataSource,
		datasources.NewUserDataSource,
		datasources.NewRepoDataSource,
		datasources.NewBranchProtectionDataSource,
	}
}

func (p *giteaProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewOrgResource,
		resources.NewRepoResource,
		resources.NewBranchProtectionResource,
	}
}

func (p *giteaProvider) NewAPIClient(cfg giteaProviderModel) (*api.APIClient, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: cfg.Insecure.ValueBool(),
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig

	url, _ := url.Parse(cfg.BaseURL.ValueString())
	apiClientConfig := api.NewConfiguration()
	apiClientConfig.Host = url.Host
	apiClientConfig.Scheme = url.Scheme
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
