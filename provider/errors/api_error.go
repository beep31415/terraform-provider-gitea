package errors

import (
	"strings"

	"terraform-provider-gitea/api"
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	apiError, ok := err.(*api.GenericOpenAPIError)
	if !ok {
		return false
	}

	return strings.Contains(apiError.Error(), "404")
}
