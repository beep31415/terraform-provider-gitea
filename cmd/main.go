//go:build !debug

package main

import (
	"context"

	"terraform-provider-gitea/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name gitea --rendered-website-dir ../docs

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "terraform.local/local/gitea",
	})
}
