package plans

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type apiUpdateNotSupported struct{}

func NewApiUpdateNotSupported() apiUpdateNotSupported {
	return apiUpdateNotSupported{}
}

func (apiUpdateNotSupported) Description(ctx context.Context) string {
	return "Plan that will prevent client from changing attributes that are not updatable by the underlining API."
}

func (apiUpdateNotSupported) MarkdownDescription(ctx context.Context) string {
	return "Plan that will prevent client from changing attributes that are not updatable by the underlining API."
}

func (p apiUpdateNotSupported) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		// allow changing state on first read
		return
	}

	if req.StateValue != req.PlanValue {
		resp.Diagnostics.AddAttributeError(req.Path,
			"Attribute change not supported.",
			fmt.Sprintf("API does not support attribute update for \"%s\" from %v to %v. Revert the change to fix this error.",
				req.Path, req.StateValue, req.ConfigValue))
	}
}

func (apiUpdateNotSupported) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		// allow changing state on first read
		return
	}

	if req.StateValue != req.PlanValue {
		resp.Diagnostics.AddAttributeError(req.Path,
			"Attribute change not supported.",
			fmt.Sprintf("API does not support attribute update for \"%s\" from %v to %v. Revert the change to fix this error.",
				req.Path, req.StateValue, req.ConfigValue))
	}
}

func (apiUpdateNotSupported) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.StateValue.IsNull() {
		// allow changing state on first read
		return
	}

	if req.StateValue != req.PlanValue {
		resp.Diagnostics.AddAttributeError(req.Path,
			"Attribute change not supported.",
			fmt.Sprintf("API does not support attribute update for \"%s\" from %v to %v. Revert the change to fix this error.",
				req.Path, req.StateValue, req.ConfigValue))
	}
}
