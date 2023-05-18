package adapter

import (
	"strings"

	"terraform-provider-gitea/api"
)

func IsErrorNotFound(err error) bool {
	if err == nil {
		return false
	}

	apiError, ok := err.(*api.GenericOpenAPIError)
	if !ok {
		return false
	}

	return strings.Contains(apiError.Error(), "404")
}

func GetAPIErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	apiError, ok := err.(*api.GenericOpenAPIError)
	if !ok {
		return err.Error()
	}

	return string(apiError.Body())
}
