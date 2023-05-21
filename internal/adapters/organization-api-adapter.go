package adapters

import (
	"context"

	"terraform-provider-gitea/api"
)

type OrganizationAdapter struct {
	client *api.APIClient
}

type EditOrganizationOptions struct {
	Name       string
	EditOption api.EditOrgOption
}

func NewOrganizationAdapter(client *api.APIClient) *OrganizationAdapter {
	return &OrganizationAdapter{
		client: client,
	}
}

func (o *OrganizationAdapter) GetByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	return res, err
}

func (o *OrganizationAdapter) Create(ctx context.Context, option api.CreateOrgOption) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgCreate(ctx).
		Organization(option).
		Execute()

	return res, err
}

func (o *OrganizationAdapter) Edit(ctx context.Context, options EditOrganizationOptions) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgEdit(ctx, options.Name).
		Body(options.EditOption).
		Execute()

	return res, err
}

func (o *OrganizationAdapter) Delete(ctx context.Context, name string) error {
	res, err := o.GetByName(ctx, name)
	if err != nil {
		return err
	}

	_, err = o.client.OrganizationAPI.
		OrgDelete(ctx, res.GetName()).
		Execute()

	return err
}
