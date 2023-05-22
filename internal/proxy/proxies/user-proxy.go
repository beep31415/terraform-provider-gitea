package proxies

import (
	"context"
	"strings"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"
	"terraform-provider-gitea/internal/proxy/errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type UserProxy struct {
	client    *api.APIClient
	converter converters.UserConverter
}

func NewUserProxy(client *api.APIClient) *UserProxy {
	return &UserProxy{
		client:    client,
		converter: converters.UserConverter{},
	}
}

func (u *UserProxy) FillDataSource(ctx context.Context, model models.UserDataSourceModel) diag.Diagnostics {
	res, _, err := u.client.UserAPI.
		UserGet(ctx, strings.ToLower(model.Name.ValueString())).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea user.")
	}

	u.converter.ReadToDataSource(model, res)

	return nil
}
