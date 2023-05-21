//go:build debug

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"terraform-provider-gitea/internal/proxy/api"
)

var base_url string
var token string

func main() {
	ctx := context.Background()
	url, _ := url.Parse(base_url)

	cfg := api.NewConfiguration()
	cfg.Host = url.Host
	cfg.Scheme = url.Scheme
	cfg.AddDefaultHeader("Authorization", "token "+token)

	c := api.NewAPIClient(cfg)

	cfg.HTTPClient.Transport = http.DefaultTransport
	cfg.HTTPClient.Transport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	res, _, err := c.OrganizationAPI.
		TeamSearch(ctx, "nas").
		//Q("nas").
		Execute()
	if err != nil {
		panic(fmt.Sprintf("error: %+v", err))
	}

	fmt.Printf("res: %+v", len(res.Data))
}
