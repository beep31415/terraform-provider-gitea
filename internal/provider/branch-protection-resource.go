package provider

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"
	"terraform-provider-gitea/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &branchProtectionResource{}
	_ resource.ResourceWithConfigure   = &branchProtectionResource{}
	_ resource.ResourceWithImportState = &branchProtectionResource{}
)

type branchProtectionResource struct {
	client *api.APIClient
}

func NewBranchProtectionResource() resource.Resource {
	return &branchProtectionResource{}
}

func (r *branchProtectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_branch_protection"
}

func (d *branchProtectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gitea organization.",
		Attributes: map[string]schema.Attribute{
			"branch_name": schema.StringAttribute{
				Description: "The branch name targeted by the rule.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"owner": schema.StringAttribute{
				Description: "The owner of the repo the rule belongs to.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"repo": schema.StringAttribute{
				Description: "The repo that the rule belongs to.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"approvals_whitelist_username": schema.ListAttribute{
				Description: "Whitelist of users allowed for approval.",
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
			"block_on_official_review_requests": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks on official review requests. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"block_on_outdated_branch": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when branch is outdated. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"block_on_rejected_reviews": schema.BoolAttribute{
				Description: "Flag indicating whether pull request blocks when pull request is rejected. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"dismiss_stale_approvals": schema.BoolAttribute{
				Description: "Flag indicating whether to ignore stale approvals on pull requests. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_approvals_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed for approvals. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_merge_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable list of users allowed to merge. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_push": schema.BoolAttribute{
				Description: "Flag indicating whether to allow push to protected branch. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_push_whitelist": schema.BoolAttribute{
				Description: "Flag indicating whether to enable whitelist of users allowed to push. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"merge_whitelist_usernames": schema.ListAttribute{
				Description: "Whitelist of users allowed to merge.",
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
			"protected_file_patterns": schema.StringAttribute{
				Description: "File pattern of protected files. Defaults to `\"\"`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"push_whitelist_deploy_keys": schema.BoolAttribute{
				Description: "Flag indicating whether to push deploy keys on whitelisted users. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"push_whitelist_usernames": schema.ListAttribute{
				Description: "Flag indicating whether to push usernames on whitelisted users.",
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
			"require_signed_commits": schema.BoolAttribute{
				Description: "Flag indicating whether to require a signed commit on push. Defaults to `false`.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"required_approvals": schema.Int64Attribute{
				Description: "Number of required approvals. Defaults to `0`.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(0),
			},
			"unprotected_file_patterns": schema.StringAttribute{
				Description: "File pattern of unprotected files. Defaults to `\"\"`.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
		},
	}
}

func (r *branchProtectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.APIClient)
}

func (r *branchProtectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var approvalWhiteList []string
	var mergeWhiteList []string
	var pushWhiteList []string
	resp.Diagnostics.Append(plan.ApprovalsWhitelistUsername.ElementsAs(ctx, &approvalWhiteList, true)...)
	resp.Diagnostics.Append(plan.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, true)...)
	resp.Diagnostics.Append(plan.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, true)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, _, err := r.client.RepositoryAPI.
		RepoCreateBranchProtection(ctx, plan.Owner.ValueString(), plan.Repo.ValueString()).
		Body(api.CreateBranchProtectionOption{
			RuleName:                      plan.BranchName.ValueStringPointer(),
			BranchName:                    plan.BranchName.ValueStringPointer(),
			ApprovalsWhitelistUsername:    approvalWhiteList,
			BlockOnOfficialReviewRequests: plan.BlockOnOfficialReviewRequests.ValueBoolPointer(),
			BlockOnOutdatedBranch:         plan.BlockOnOutdatedBranch.ValueBoolPointer(),
			BlockOnRejectedReviews:        plan.BlockOnRejectedReviews.ValueBoolPointer(),
			DismissStaleApprovals:         plan.DismissStaleApprovals.ValueBoolPointer(),
			EnableApprovalsWhitelist:      plan.EnableApprovalsWhitelist.ValueBoolPointer(),
			EnableMergeWhitelist:          plan.EnableMergeWhitelist.ValueBoolPointer(),
			EnablePush:                    plan.EnablePush.ValueBoolPointer(),
			EnablePushWhitelist:           plan.EnablePushWhitelist.ValueBoolPointer(),
			MergeWhitelistUsernames:       mergeWhiteList,
			ProtectedFilePatterns:         plan.ProtectedFilePatterns.ValueStringPointer(),
			PushWhitelistDeployKeys:       plan.PushWhitelistDeployKeys.ValueBoolPointer(),
			PushWhitelistUsernames:        pushWhiteList,
			RequireSignedCommits:          plan.RequireSignedCommits.ValueBoolPointer(),
			RequiredApprovals:             plan.RequiredApprovals.ValueInt64Pointer(),
			UnprotectedFilePatterns:       plan.UnprotectedFilePatterns.ValueStringPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Gitea branch protection.",
			fmt.Sprintf("Could not Create Gitea branch protection for '%s/%s/%s', unexpected error: %s",
				plan.Owner.ValueString(), plan.Repo.ValueString(), plan.BranchName.ValueString(), adapters.GetAPIErrorMessage(err)),
		)

		return
	}

	plan.BranchName = types.StringValue(res.GetBranchName())
	plan.BlockOnOfficialReviewRequests = types.BoolValue(res.GetBlockOnOfficialReviewRequests())
	plan.BlockOnOutdatedBranch = types.BoolValue(res.GetBlockOnOutdatedBranch())
	plan.BlockOnRejectedReviews = types.BoolValue(res.GetBlockOnRejectedReviews())
	plan.DismissStaleApprovals = types.BoolValue(res.GetDismissStaleApprovals())
	plan.EnableApprovalsWhitelist = types.BoolValue(res.GetEnableApprovalsWhitelist())
	plan.EnableMergeWhitelist = types.BoolValue(res.GetEnableMergeWhitelist())
	plan.EnablePush = types.BoolValue(res.GetEnablePush())
	plan.EnablePushWhitelist = types.BoolValue(res.GetEnablePushWhitelist())
	plan.ProtectedFilePatterns = types.StringValue(res.GetProtectedFilePatterns())
	plan.PushWhitelistDeployKeys = types.BoolValue(res.GetPushWhitelistDeployKeys())
	plan.RequireSignedCommits = types.BoolValue(res.GetRequireSignedCommits())
	plan.RequiredApprovals = types.Int64Value(res.GetRequiredApprovals())
	plan.UnprotectedFilePatterns = types.StringValue(res.GetUnprotectedFilePatterns())

	tfApprovalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetApprovalsWhitelistUsername())
	resp.Diagnostics.Append(diags...)
	tfMergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetMergeWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	tfPushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetPushWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ApprovalsWhitelistUsername = tfApprovalWhiteList
	plan.MergeWhitelistUsernames = tfMergeWhiteList
	plan.PushWhitelistUsernames = tfPushWhiteList

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.getBranchProtection(ctx, state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Gitea branch protection.",
			fmt.Sprintf("Could not read Gitea branch protection for '%s/%s/%s', unexpected error: %s",
				state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString(), adapters.GetAPIErrorMessage(err)),
		)

		return
	}

	state.BranchName = types.StringValue(res.GetBranchName())
	state.BlockOnOfficialReviewRequests = types.BoolValue(res.GetBlockOnOfficialReviewRequests())
	state.BlockOnOutdatedBranch = types.BoolValue(res.GetBlockOnOutdatedBranch())
	state.BlockOnRejectedReviews = types.BoolValue(res.GetBlockOnRejectedReviews())
	state.DismissStaleApprovals = types.BoolValue(res.GetDismissStaleApprovals())
	state.EnableApprovalsWhitelist = types.BoolValue(res.GetEnableApprovalsWhitelist())
	state.EnableMergeWhitelist = types.BoolValue(res.GetEnableMergeWhitelist())
	state.EnablePush = types.BoolValue(res.GetEnablePush())
	state.EnablePushWhitelist = types.BoolValue(res.GetEnablePushWhitelist())
	state.ProtectedFilePatterns = types.StringValue(res.GetProtectedFilePatterns())
	state.PushWhitelistDeployKeys = types.BoolValue(res.GetPushWhitelistDeployKeys())
	state.RequireSignedCommits = types.BoolValue(res.GetRequireSignedCommits())
	state.RequiredApprovals = types.Int64Value(res.GetRequiredApprovals())
	state.UnprotectedFilePatterns = types.StringValue(res.GetUnprotectedFilePatterns())

	tfApprovalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetApprovalsWhitelistUsername())
	resp.Diagnostics.Append(diags...)
	tfMergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetMergeWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	tfPushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetPushWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.ApprovalsWhitelistUsername = tfApprovalWhiteList
	state.MergeWhitelistUsernames = tfMergeWhiteList
	state.PushWhitelistUsernames = tfPushWhiteList

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var approvalWhiteList []string
	var mergeWhiteList []string
	var pushWhiteList []string
	resp.Diagnostics.Append(plan.ApprovalsWhitelistUsername.ElementsAs(ctx, &approvalWhiteList, false)...)
	resp.Diagnostics.Append(plan.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, false)...)
	resp.Diagnostics.Append(plan.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, _, err := r.client.RepositoryAPI.
		RepoEditBranchProtection(ctx, plan.Owner.ValueString(), plan.Repo.ValueString(), plan.BranchName.ValueString()).
		Body(api.EditBranchProtectionOption{
			ApprovalsWhitelistUsername:    approvalWhiteList,
			BlockOnOfficialReviewRequests: plan.BlockOnOfficialReviewRequests.ValueBoolPointer(),
			BlockOnOutdatedBranch:         plan.BlockOnOutdatedBranch.ValueBoolPointer(),
			BlockOnRejectedReviews:        plan.BlockOnRejectedReviews.ValueBoolPointer(),
			DismissStaleApprovals:         plan.DismissStaleApprovals.ValueBoolPointer(),
			EnableApprovalsWhitelist:      plan.EnableApprovalsWhitelist.ValueBoolPointer(),
			EnableMergeWhitelist:          plan.EnableMergeWhitelist.ValueBoolPointer(),
			EnablePush:                    plan.EnablePush.ValueBoolPointer(),
			EnablePushWhitelist:           plan.EnablePushWhitelist.ValueBoolPointer(),
			MergeWhitelistUsernames:       mergeWhiteList,
			ProtectedFilePatterns:         plan.ProtectedFilePatterns.ValueStringPointer(),
			PushWhitelistDeployKeys:       plan.PushWhitelistDeployKeys.ValueBoolPointer(),
			PushWhitelistUsernames:        pushWhiteList,
			RequireSignedCommits:          plan.RequireSignedCommits.ValueBoolPointer(),
			RequiredApprovals:             plan.RequiredApprovals.ValueInt64Pointer(),
			UnprotectedFilePatterns:       plan.UnprotectedFilePatterns.ValueStringPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Gitea branch protection.",
			fmt.Sprintf("Could not Update Gitea branch protection for '%s/%s/%s', unexpected error: %s",
				plan.Owner.ValueString(), plan.Repo.ValueString(), plan.BranchName.ValueString(), adapters.GetAPIErrorMessage(err)),
		)

		return
	}

	plan.BranchName = types.StringValue(res.GetBranchName())
	plan.BlockOnOfficialReviewRequests = types.BoolValue(res.GetBlockOnOfficialReviewRequests())
	plan.BlockOnOutdatedBranch = types.BoolValue(res.GetBlockOnOutdatedBranch())
	plan.BlockOnRejectedReviews = types.BoolValue(res.GetBlockOnRejectedReviews())
	plan.DismissStaleApprovals = types.BoolValue(res.GetDismissStaleApprovals())
	plan.EnableApprovalsWhitelist = types.BoolValue(res.GetEnableApprovalsWhitelist())
	plan.EnableMergeWhitelist = types.BoolValue(res.GetEnableMergeWhitelist())
	plan.EnablePush = types.BoolValue(res.GetEnablePush())
	plan.EnablePushWhitelist = types.BoolValue(res.GetEnablePushWhitelist())
	plan.ProtectedFilePatterns = types.StringValue(res.GetProtectedFilePatterns())
	plan.PushWhitelistDeployKeys = types.BoolValue(res.GetPushWhitelistDeployKeys())
	plan.RequireSignedCommits = types.BoolValue(res.GetRequireSignedCommits())
	plan.RequiredApprovals = types.Int64Value(res.GetRequiredApprovals())
	plan.UnprotectedFilePatterns = types.StringValue(res.GetUnprotectedFilePatterns())

	tfApprovalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetApprovalsWhitelistUsername())
	resp.Diagnostics.Append(diags...)
	tfMergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetMergeWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	tfPushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, res.GetPushWhitelistUsernames())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ApprovalsWhitelistUsername = tfApprovalWhiteList
	plan.MergeWhitelistUsernames = tfMergeWhiteList
	plan.PushWhitelistUsernames = tfPushWhiteList

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *branchProtectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.BranchProtectionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.getBranchProtection(ctx, state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString())
	if err != nil {
		if adapters.IsErrorNotFound(err) {
			return
		}

		resp.Diagnostics.AddError(
			"Error Delete Gitea branch protection.",
			fmt.Sprintf("Could not check if branch protection exists for '%s/%s/%s', unexpected error: %s",
				state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString(), adapters.GetAPIErrorMessage(err)),
		)

		return
	}

	_, err = r.client.RepositoryAPI.
		RepoDeleteBranchProtection(ctx, state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Gitea branch protection.",
			fmt.Sprintf("Could not Delete Gitea branch protection for '%s/%s/%s', unexpected error: %s",
				state.Owner.ValueString(), state.Repo.ValueString(), state.BranchName.ValueString(), adapters.GetAPIErrorMessage(err)),
		)

		return
	}
}

func (r *branchProtectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ids := strings.Split(req.ID, "/")
	if len(ids) != 3 {
		resp.Diagnostics.AddError(
			"Error Importing Gitea branch protection.",
			"Must provide import id in the form of owner_name/repo_name/branch_name.",
		)

		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("owner"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("repo"), ids[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("branch_name"), ids[2])...)
}

func (r *branchProtectionResource) getBranchProtection(ctx context.Context, owner, repo, name string) (*api.BranchProtection, error) {
	res, _, err := r.client.RepositoryAPI.
		RepoGetBranchProtection(ctx, owner, repo, name).
		Execute()

	return res, err
}
