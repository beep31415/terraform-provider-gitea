package adapters

import (
	"context"
	"strings"
	"terraform-provider-gitea/api"
)

type UserAdapter struct {
	client *api.APIClient
}

func NewUserAdapter(client *api.APIClient) *UserAdapter {
	return &UserAdapter{
		client: client,
	}
}

func (u *UserAdapter) GetByName(ctx context.Context, name string) (*api.User, error) {
	res, _, err := u.client.UserAPI.
		UserGet(ctx, strings.ToLower(name)).
		Execute()

	return res, err
}
