package proxy

import (
	"strings"

	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func GetApiInitError(err error) string {
	return getAPIErrorMessage(err)
}

func toDiagnosticError(err error, title string) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		title,
		getAPIErrorMessage(err),
	)
}

func toDiagnosticArrayError(err error, title string) diag.Diagnostics {
	return diag.Diagnostics{
		diag.NewErrorDiagnostic(
			title,
			getAPIErrorMessage(err),
		),
	}
}

func isErrorNotFound(err error) bool {
	if err == nil {
		return false
	}

	apiError, ok := err.(*api.GenericOpenAPIError)
	if !ok {
		return false
	}

	return strings.Contains(apiError.Error(), "404")
}

func getAPIErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	apiError, ok := err.(*api.GenericOpenAPIError)
	if !ok {
		return err.Error()
	}

	return string(apiError.Body())
}
