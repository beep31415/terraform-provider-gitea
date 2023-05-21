package provider

import (
	"context"

	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/plans"
	"terraform-provider-gitea/internal/proxy"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ resource.Resource                = &orgResource{}
	_ resource.ResourceWithConfigure   = &orgResource{}
	_ resource.ResourceWithImportState = &orgResource{}
)

type orgResource struct {
	client *api.APIClient
	proxy  *proxy.OrganizationProxy
}

func NewOrgResource() resource.Resource {
	return &orgResource{}
}

func (r *orgResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org"
}

func (r *orgResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
	r.proxy = proxy.NewOrganizationProxy(r.client)
}

func (r *orgResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.OrganizationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Create(ctx, plan))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.OrganizationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.FillResource(ctx, state))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.OrganizationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Edit(ctx, plan))
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.OrganizationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.proxy.Delete(ctx, state))
}

func (r *orgResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
}

func (d *orgResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gitea organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the organisation without spaces.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"full_name": schema.StringAttribute{
				Description: "The full name of the organisation.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				Description: "The organization description.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"website": schema.StringAttribute{
				Description: "The organization website.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"location": schema.StringAttribute{
				Description: "The organization location.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"visibility": schema.StringAttribute{
				Description: "The organization visibility. Possible values: public, limited or private. Defaults to `public`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("public"),
				Validators: []validator.String{
					stringvalidator.OneOf("public", "limited", "private"),
				},
			},
			"repo_admin_change_team_access": schema.BoolAttribute{
				Description: "Flag that indicates whether admin can change organization team access. Defaults to `true`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}
