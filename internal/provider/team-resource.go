package provider

import (
	"context"
	"strings"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/plans"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &teamResource{}
	_ resource.ResourceWithConfigure   = &teamResource{}
	_ resource.ResourceWithImportState = &teamResource{}
)

type teamResource struct {
	client      *api.APIClient
	teamAdapter *adapters.TeamAdapter
}

func NewTeamResource() resource.Resource {
	return &teamResource{}
}

func (r *teamResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (d *teamResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gitea team.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Identifier attribute.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The team name.",
				Required:    true,
			},
			"organization": schema.StringAttribute{
				Description: "The organization the team belongs to.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					plans.NewApiUpdateNotSupported(),
				},
			},
			"permission": schema.StringAttribute{
				Description: "Sets permission for team members. Defaults to `none`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Validators: []validator.String{
					stringvalidator.OneOf("none", "read", "write", "admin"),
				},
			},
			"description": schema.StringAttribute{
				Description: "The team description",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"can_create_org_repo": schema.BoolAttribute{
				Description: "Flag that indicates whether team can create organization repos. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"includes_all_repositories": schema.BoolAttribute{
				Description: "Flag indicating whether team has access to all repos. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"members": schema.ListAttribute{
				Description: "Team members.",
				ElementType: types.StringType,
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf(""),
					),
				},
			},
		},
	}
}

func (r *teamResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
	r.teamAdapter = adapters.NewTeamAdapter(r.client)
}

func (r *teamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.TeamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	addTeamOptions, diags := plan.ToAddTeamOptions(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.teamAdapter.CreateTeam(ctx, addTeamOptions)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea team.",
			"Could not create team, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	diagsRead := plan.ReadFrom(ctx, res)
	resp.Diagnostics.Append(diagsRead...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.TeamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.teamAdapter.GetTeamByOrgAndName(ctx, state.Organization.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea team.",
			"Could not read Gitea team name "+state.Name.ValueString()+": "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	diags := state.ReadFrom(ctx, res)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.TeamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	editTeamOptions, diags := plan.ToEditTeamOptions(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.teamAdapter.EditTeam(ctx, editTeamOptions)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitea team.",
			"Could not update team, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}

	diagsRead := plan.ReadFrom(ctx, res)
	resp.Diagnostics.Append(diagsRead...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.TeamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.teamAdapter.DeleteTeam(ctx, state.Id.ValueInt64())
	if err != nil {
		if adapters.IsErrorNotFound(err) {
			return
		}

		resp.Diagnostics.AddError(
			"Error Delete Gitea team.",
			"Could not delete team, unexpected error: "+adapters.GetAPIErrorMessage(err),
		)

		return
	}
}

func (r *teamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ids := strings.Split(req.ID, "/")
	if len(ids) != 2 {
		resp.Diagnostics.AddError(
			"Error Importing Gitea team.",
			"Must provide import id in the form of org_name/team_name.",
		)

		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), ids[1])...)
}
