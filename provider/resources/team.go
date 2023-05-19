package resources

import (
	"context"
	"strings"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/provider/adapter"
	"terraform-provider-gitea/provider/plans"

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

	allUnits = []string{"repo.code", "repo.issues", "repo.ext_issues", "repo.wiki",
		"repo.pulls", "repo.releases", "repo.projects", "repo.ext_wiki"}
)

type teamResource struct {
	client *api.APIClient
}

type teamResourceModel struct {
	Id                      types.Int64  `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Organization            types.String `tfsdk:"organization"`
	Permission              types.String `tfsdk:"permission"`
	Description             types.String `tfsdk:"description"`
	CanCreateOrgRepo        types.Bool   `tfsdk:"can_create_org_repo"`
	IncludesAllRepositories types.Bool   `tfsdk:"includes_all_repositories"`
	Members                 types.List   `tfsdk:"members"`
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
}

func (r *teamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan teamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var memberList []string
	resp.Diagnostics.Append(plan.Members.ElementsAs(ctx, &memberList, true)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if memberList == nil {
		memberList = []string{}
	}

	res, _, err := r.client.OrganizationAPI.
		OrgCreateTeam(ctx, plan.Organization.ValueString()).
		Body(api.CreateTeamOption{
			Name:                    plan.Name.ValueString(),
			Permission:              plan.Permission.ValueStringPointer(),
			Description:             plan.Description.ValueStringPointer(),
			CanCreateOrgRepo:        plan.CanCreateOrgRepo.ValueBoolPointer(),
			IncludesAllRepositories: plan.IncludesAllRepositories.ValueBoolPointer(),
			Units:                   allUnits,
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea team.",
			"Could not create team, unexpected error: "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	plan.Id = types.Int64Value(res.GetId())
	plan.Permission = types.StringValue(res.GetPermission())
	plan.Description = types.StringValue(res.GetDescription())
	plan.CanCreateOrgRepo = types.BoolValue(res.GetCanCreateOrgRepo())
	plan.IncludesAllRepositories = types.BoolValue(res.GetIncludesAllRepositories())

	for _, user := range memberList {
		err := r.addTeamMember(ctx, res.GetId(), user)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating Gitea team.",
				"Could not add team member: "+user,
			)

			return
		}
	}

	tfMembersList, diags := types.ListValueFrom(ctx, types.StringType, memberList)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.Members = tfMembersList

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state teamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := adapter.GetTeamByOrgAndName(ctx, r.client, state.Organization.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea team.",
			"Could not read Gitea team name "+state.Name.ValueString()+": "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	state.Id = types.Int64Value(res.GetId())
	state.Name = types.StringValue(res.GetName())
	state.Permission = types.StringValue(res.GetPermission())
	state.Description = types.StringValue(res.GetDescription())
	state.CanCreateOrgRepo = types.BoolValue(res.GetCanCreateOrgRepo())
	state.IncludesAllRepositories = types.BoolValue(res.GetIncludesAllRepositories())

	memberList, err := r.getTeamMembers(ctx, res.GetId())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea team.",
			"Could not read Gitea team members for team "+state.Name.ValueString()+": "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	tfMembersList, diags := types.ListValueFrom(ctx, types.StringType, memberList)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.Members = tfMembersList

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan teamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var updatedMemberList []string
	resp.Diagnostics.Append(plan.Members.ElementsAs(ctx, &updatedMemberList, true)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if updatedMemberList == nil {
		updatedMemberList = []string{}
	}

	res, _, err := r.client.OrganizationAPI.
		OrgEditTeam(ctx, int32(plan.Id.ValueInt64())).
		Body(api.EditTeamOption{
			Name:                    plan.Name.ValueString(),
			Permission:              plan.Permission.ValueStringPointer(),
			Description:             plan.Description.ValueStringPointer(),
			CanCreateOrgRepo:        plan.CanCreateOrgRepo.ValueBoolPointer(),
			IncludesAllRepositories: plan.IncludesAllRepositories.ValueBoolPointer(),
			Units:                   allUnits,
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitea team.",
			"Could not update team, unexpected error: "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	plan.Id = types.Int64Value(res.GetId())
	plan.Name = types.StringValue(res.GetName())
	plan.Permission = types.StringValue(res.GetPermission())
	plan.Description = types.StringValue(res.GetDescription())
	plan.CanCreateOrgRepo = types.BoolValue(res.GetCanCreateOrgRepo())
	plan.IncludesAllRepositories = types.BoolValue(res.GetIncludesAllRepositories())

	existingMemberList, err := r.getTeamMembers(ctx, res.GetId())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitea team.",
			"Could not read current Gitea team members for team "+plan.Name.ValueString()+": "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	for _, member := range existingMemberList {
		if !r.contains(member, updatedMemberList...) {
			err = r.removeTeamMember(ctx, res.GetId(), member)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Updating Gitea team.",
					"Could not remove team member "+member+", unexpected error: "+adapter.GetAPIErrorMessage(err),
				)

				return
			}
		}
	}

	for _, member := range updatedMemberList {
		if !r.contains(member, existingMemberList...) {
			err = r.addTeamMember(ctx, res.GetId(), member)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Updating Gitea team.",
					"Could not add team member "+member+", unexpected error: "+adapter.GetAPIErrorMessage(err),
				)

				return
			}
		}
	}

	tfMembersList, diags := types.ListValueFrom(ctx, types.StringType, updatedMemberList)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.Members = tfMembersList

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *teamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state teamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.getTeamById(ctx, state.Id.ValueInt64())
	if err != nil {
		if adapter.IsErrorNotFound(err) {
			return
		}

		resp.Diagnostics.AddError(
			"Error Delete Gitea team.",
			"Could not check if team to delete exists, unexpected error: "+adapter.GetAPIErrorMessage(err),
		)

		return
	}

	var memberList []string
	resp.Diagnostics.Append(state.Members.ElementsAs(ctx, &memberList, true)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, member := range memberList {
		err = r.removeTeamMember(ctx, state.Id.ValueInt64(), member)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Deleting Gitea team.",
				"Could not delete team member: "+member+", unexpected error: "+adapter.GetAPIErrorMessage(err),
			)

			return
		}
	}

	_, err = r.client.OrganizationAPI.
		OrgDeleteTeam(ctx, *res.Id).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Gitea team.",
			"Could not delete team, unexpected error: "+adapter.GetAPIErrorMessage(err),
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

func (r *teamResource) getTeamById(ctx context.Context, id int64) (*api.Team, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgGetTeam(ctx, id).
		Execute()

	return res, err
}

func (r *teamResource) getTeamMembers(ctx context.Context, id int64) ([]string, error) {
	res, _, err := r.client.OrganizationAPI.
		OrgListTeamMembers(ctx, id).
		Execute()
	if err != nil {
		return []string{}, err
	}

	users := []string{}
	for _, resUser := range res {
		users = append(users, resUser.GetLogin())
	}

	return users, err
}

func (r *teamResource) addTeamMember(ctx context.Context, id int64, user string) error {
	_, err := r.client.OrganizationAPI.
		OrgAddTeamMember(ctx, id, user).
		Execute()

	return err
}

func (r *teamResource) removeTeamMember(ctx context.Context, id int64, user string) error {
	_, err := r.client.OrganizationAPI.
		OrgRemoveTeamMember(ctx, id, user).
		Execute()

	return err
}

func (r *teamResource) contains(item string, list ...string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}