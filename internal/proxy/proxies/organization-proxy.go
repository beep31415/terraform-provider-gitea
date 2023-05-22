package proxies

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"
	"terraform-provider-gitea/internal/proxy/converters"
	"terraform-provider-gitea/internal/proxy/errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type OrganizationProxy struct {
	client    *api.APIClient
	converter converters.OrganizationConverter
}

func NewOrganizationProxy(client *api.APIClient) *OrganizationProxy {
	return &OrganizationProxy{
		client:    client,
		converter: converters.OrganizationConverter{},
	}
}

func (o *OrganizationProxy) FillDataSource(ctx context.Context, model models.OrganizationDataSourceModel) diag.Diagnostics {
	res, err := o.getByName(ctx, model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea organization.")
	}

	o.converter.ReadToDataSource(model, res)

	return nil
}

func (o *OrganizationProxy) FillResource(ctx context.Context, model models.OrganizationResourceModel) diag.Diagnostics {
	res, err := o.getByName(ctx, model.Name.ValueString())
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Reading Gitea organization.")
	}

	o.converter.ReadToResource(model, res)

	return nil
}

func (o *OrganizationProxy) Create(ctx context.Context, model models.OrganizationResourceModel) diag.Diagnostics {
	option := o.converter.ToApiAddOrganizationOptions(model)

	res, _, err := o.client.OrganizationAPI.
		OrgCreate(ctx).
		Organization(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Creating Gitea organization.")
	}

	o.converter.ReadToResource(model, res)

	return nil
}

func (o *OrganizationProxy) Update(ctx context.Context, model models.OrganizationResourceModel) diag.Diagnostics {
	option := o.converter.ToApiEditOrganizationOptions(model)

	res, _, err := o.client.OrganizationAPI.
		OrgEdit(ctx, model.Name.ValueString()).
		Body(option).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Updating Gitea organization.")
	}

	o.converter.ReadToResource(model, res)

	return nil
}

func (o *OrganizationProxy) Delete(ctx context.Context, model models.OrganizationResourceModel) diag.Diagnostics {
	res, err := o.getByName(ctx, model.Name.ValueString())
	if err != nil {
		if errors.IsErrorNotFound(err) {
			return nil
		}

		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea organization.")
	}

	_, err = o.client.OrganizationAPI.
		OrgDelete(ctx, res.GetName()).
		Execute()
	if err != nil {
		return errors.ToDiagnosticArrayError(err, "Error Deleting Gitea organization.")
	}

	return nil
}

func (o *OrganizationProxy) getByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	return res, err
}
