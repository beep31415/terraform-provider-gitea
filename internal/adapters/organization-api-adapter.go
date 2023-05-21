package adapters

import (
	"context"

	"terraform-provider-gitea/api"
)

type Organizationdapter struct {
	client *api.APIClient
}

type EditOrganizationOptions struct {
	Name       string
	EditOption api.EditOrgOption
}

func NewOrganizationdapter(client *api.APIClient) *Organizationdapter {
	return &Organizationdapter{
		client: client,
	}
}

func (o *Organizationdapter) GetByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	return res, err
}

func (o *Organizationdapter) Create(ctx context.Context, option api.CreateOrgOption) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgCreate(ctx).
		Organization(option).
		Execute()

	return res, err
}

func (o *Organizationdapter) Edit(ctx context.Context, options EditOrganizationOptions) (*api.Organization, error) {
	res, _, err := o.client.OrganizationAPI.
		OrgEdit(ctx, options.Name).
		Body(options.EditOption).
		Execute()

	return res, err
}

func (o *Organizationdapter) Delete(ctx context.Context, name string) error {
	res, err := o.GetByName(ctx, name)
	if err != nil {
		return err
	}

	_, err = o.client.OrganizationAPI.
		OrgDelete(ctx, res.GetName()).
		Execute()

	return err
}
