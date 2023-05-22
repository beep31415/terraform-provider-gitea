package provider

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type giteaProviderModel struct {
	BaseURL      types.String `tfsdk:"base_url"`
	Token        types.String `tfsdk:"token"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	CACertFile   types.String `tfsdk:"cacert_file"`
	Insecure     types.Bool   `tfsdk:"insecure"`
	Debug        types.Bool   `tfsdk:"debug"`
	DebugLogPath types.String `tfsdk:"debug_log_path"`

	parsedUrl *url.URL
}

func (m *giteaProviderModel) loadEnvs() {
	if v, ok := os.LookupEnv("GITEA_BASE_URL"); ok {
		m.BaseURL = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_TOKEN"); ok {
		m.Token = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_USERNAME"); ok {
		m.Username = types.StringValue(v)
	}

	if v, ok := os.LookupEnv("GITEA_PASSWORD"); ok {
		m.Password = types.StringValue(v)
	}
}

func (m *giteaProviderModel) validate() diag.Diagnostic {
	if m.BaseURL.ValueString() == "" {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("base_url"),
			"Invalid base url.",
			"Base url cannot be empty.",
		)
	}

	var err error
	if m.parsedUrl, err = url.Parse(m.BaseURL.ValueString()); err != nil {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("base_url"),
			"Invalid base url.",
			fmt.Sprintf("Base url %s is not a valid url: %+v.", m.BaseURL.ValueString(), err),
		)
	}

	if m.parsedUrl.Path != "" {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("base_url"),
			"Invalid base url.",
			fmt.Sprintf("Base url %s should not contain path, got: %s.", m.BaseURL.ValueString(), m.parsedUrl.Path),
		)
	}

	if m.Username.ValueString() == "" && m.Token.ValueString() == "" {
		return diag.NewErrorDiagnostic(
			"Invalid username or token.",
			"Username and token are empty. Must provide exacly one of them.",
		)
	}

	if m.Username.ValueString() != "" && m.Token.ValueString() != "" {
		return diag.NewErrorDiagnostic(
			"Invalid username or token.",
			"Username and token are both provided. Only one of them can be set.",
		)
	}

	if m.Username.ValueString() != "" && m.Password.ValueString() == "" {
		return diag.NewErrorDiagnostic(
			"Invalid password.",
			"If username is provided password must also be set.",
		)
	}

	if m.Username.ValueString() == "" && m.Password.ValueString() != "" {
		return diag.NewErrorDiagnostic(
			"Invalid username.",
			"If password is provided username must also be set.",
		)
	}

	if m.Debug.ValueBool() && (m.DebugLogPath.ValueString() == "") {
		return diag.NewErrorDiagnostic(
			"Invalid debug log path.",
			"If debug is enabled a debug log path must be specified.",
		)
	}

	return nil
}
