package provider

import (
	"context"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/plans"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &orgResource{}
	_ resource.ResourceWithConfigure   = &orgResource{}
	_ resource.ResourceWithImportState = &orgResource{}
)

type orgResource struct {
	client *api.APIClient
}

func NewOrgResource() resource.Resource {
	return &orgResource{}
}

func (r *orgResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org"
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

func (r *orgResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
}

func (r *orgResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.OrganizationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, _, err := r.client.OrganizationAPI.
		OrgCreate(ctx).
		Organization(api.CreateOrgOption{
			Username:                  plan.Name.ValueString(),
			FullName:                  plan.FullName.ValueStringPointer(),
			Description:               plan.Description.ValueStringPointer(),
			Website:                   plan.Website.ValueStringPointer(),
			Location:                  plan.Location.ValueStringPointer(),
			Visibility:                plan.Visibility.ValueStringPointer(),
			RepoAdminChangeTeamAccess: plan.AdminChangeTeamAccess.ValueBoolPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea organization.",
			"Could not create organization, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	plan.ID = types.Int64Value(res.GetId())
	plan.Name = types.StringValue(res.GetName())
	plan.FullName = types.StringValue(res.GetFullName())
	plan.Description = types.StringValue(res.GetDescription())
	plan.Website = types.StringValue(res.GetWebsite())
	plan.Location = types.StringValue(res.GetLocation())
	plan.Visibility = types.StringValue(res.GetVisibility())
	plan.AdminChangeTeamAccess = types.BoolValue(res.GetRepoAdminChangeTeamAccess())

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

	res, err := r.getOrgByName(ctx, state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea organization.",
			"Could not read Gitea organization name "+state.Name.ValueString()+": "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	state.ID = types.Int64Value(res.GetId())
	state.Name = types.StringValue(res.GetName())
	state.FullName = types.StringValue(res.GetFullName())
	state.Description = types.StringValue(res.GetDescription())
	state.Website = types.StringValue(res.GetWebsite())
	state.Location = types.StringValue(res.GetLocation())
	state.Visibility = types.StringValue(res.GetVisibility())
	state.AdminChangeTeamAccess = types.BoolValue(res.GetRepoAdminChangeTeamAccess())

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

	res, _, err := r.client.OrganizationAPI.
		OrgEdit(ctx, plan.Name.ValueString()).
		Body(api.EditOrgOption{
			FullName:                  plan.FullName.ValueStringPointer(),
			Description:               plan.Description.ValueStringPointer(),
			Website:                   plan.Website.ValueStringPointer(),
			Location:                  plan.Location.ValueStringPointer(),
			Visibility:                plan.Visibility.ValueStringPointer(),
			RepoAdminChangeTeamAccess: plan.AdminChangeTeamAccess.ValueBoolPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitea organization.",
			"Could not update organization, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	plan.ID = types.Int64Value(res.GetId())
	plan.Name = types.StringValue(res.GetName())
	plan.FullName = types.StringValue(res.GetFullName())
	plan.Description = types.StringValue(res.GetDescription())
	plan.Website = types.StringValue(res.GetWebsite())
	plan.Location = types.StringValue(res.GetLocation())
	plan.Visibility = types.StringValue(res.GetVisibility())
	plan.AdminChangeTeamAccess = types.BoolValue(res.GetRepoAdminChangeTeamAccess())

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

	res, err := r.getOrgByName(ctx, state.Name.ValueString())
	if err != nil {
		if adapters.IsErrorNotFound(err) {
			return
		}

		resp.Diagnostics.AddError(
			"Error Delete Gitea organization.",
			"Could not check the organization to delete exists, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	_, err = r.client.OrganizationAPI.
		OrgDelete(ctx, res.GetName()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Gitea organization.",
			"Could not delete organizaton, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}
}

func (r *orgResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
}

func (r *orgResource) getOrgByName(ctx context.Context, name string) (*api.Organization, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgGet(ctx, name).
		Execute()

	return res, err
}
