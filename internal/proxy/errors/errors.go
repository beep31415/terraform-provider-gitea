package errors

import (
	"strings"

	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ToDiagnosticError(err error, category string) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		category,
		GetAPIErrorMessage(err),
	)
}

func ToDiagnosticArrayError(err error, category string) diag.Diagnostics {
	return diag.Diagnostics{
		diag.NewErrorDiagnostic(
			category,
			GetAPIErrorMessage(err),
		),
	}
}

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
