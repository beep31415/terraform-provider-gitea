package models

import (
	"context"

	"terraform-provider-gitea/api"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BranchProtectionResourceModel struct {
	Owner                         types.String `tfsdk:"owner"`
	Repo                          types.String `tfsdk:"repo"`
	BranchName                    types.String `tfsdk:"branch_name"`
	ApprovalsWhitelistUsername    types.List   `tfsdk:"approvals_whitelist_username"`
	BlockOnOfficialReviewRequests types.Bool   `tfsdk:"block_on_official_review_requests"`
	BlockOnOutdatedBranch         types.Bool   `tfsdk:"block_on_outdated_branch"`
	BlockOnRejectedReviews        types.Bool   `tfsdk:"block_on_rejected_reviews"`
	DismissStaleApprovals         types.Bool   `tfsdk:"dismiss_stale_approvals"`
	EnableApprovalsWhitelist      types.Bool   `tfsdk:"enable_approvals_whitelist"`
	EnableMergeWhitelist          types.Bool   `tfsdk:"enable_merge_whitelist"`
	EnablePush                    types.Bool   `tfsdk:"enable_push"`
	EnablePushWhitelist           types.Bool   `tfsdk:"enable_push_whitelist"`
	MergeWhitelistUsernames       types.List   `tfsdk:"merge_whitelist_usernames"`
	ProtectedFilePatterns         types.String `tfsdk:"protected_file_patterns"`
	PushWhitelistDeployKeys       types.Bool   `tfsdk:"push_whitelist_deploy_keys"`
	PushWhitelistUsernames        types.List   `tfsdk:"push_whitelist_usernames"`
	RequireSignedCommits          types.Bool   `tfsdk:"require_signed_commits"`
	RequiredApprovals             types.Int64  `tfsdk:"required_approvals"`
	UnprotectedFilePatterns       types.String `tfsdk:"unprotected_file_patterns"`
}

func (b *BranchProtectionResourceModel) ReadFromApi(ctx context.Context, branchProtection *adapters.BranchProtection) diag.Diagnostics {
	tfApprovalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetApprovalsWhitelistUsername())
	if diags.HasError() {
		return diags
	}

	tfMergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetMergeWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	tfPushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetPushWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	b.Owner = types.StringValue(branchProtection.Owner)
	b.Repo = types.StringValue(branchProtection.Repo)
	b.BranchName = types.StringValue(branchProtection.BranchProtection.GetBranchName())
	b.BlockOnOfficialReviewRequests = types.BoolValue(branchProtection.BranchProtection.GetBlockOnOfficialReviewRequests())
	b.BlockOnOutdatedBranch = types.BoolValue(branchProtection.BranchProtection.GetBlockOnOutdatedBranch())
	b.BlockOnRejectedReviews = types.BoolValue(branchProtection.BranchProtection.GetBlockOnRejectedReviews())
	b.DismissStaleApprovals = types.BoolValue(branchProtection.BranchProtection.GetDismissStaleApprovals())
	b.EnableApprovalsWhitelist = types.BoolValue(branchProtection.BranchProtection.GetEnableApprovalsWhitelist())
	b.EnableMergeWhitelist = types.BoolValue(branchProtection.BranchProtection.GetEnableMergeWhitelist())
	b.EnablePush = types.BoolValue(branchProtection.BranchProtection.GetEnablePush())
	b.EnablePushWhitelist = types.BoolValue(branchProtection.BranchProtection.GetEnablePushWhitelist())
	b.ProtectedFilePatterns = types.StringValue(branchProtection.BranchProtection.GetProtectedFilePatterns())
	b.PushWhitelistDeployKeys = types.BoolValue(branchProtection.BranchProtection.GetPushWhitelistDeployKeys())
	b.RequireSignedCommits = types.BoolValue(branchProtection.BranchProtection.GetRequireSignedCommits())
	b.RequiredApprovals = types.Int64Value(branchProtection.BranchProtection.GetRequiredApprovals())
	b.UnprotectedFilePatterns = types.StringValue(branchProtection.BranchProtection.GetUnprotectedFilePatterns())
	b.ApprovalsWhitelistUsername = tfApprovalWhiteList
	b.MergeWhitelistUsernames = tfMergeWhiteList
	b.PushWhitelistUsernames = tfPushWhiteList

	return diags
}

func (b *BranchProtectionResourceModel) ToApiAddBranchProtectionOptions(ctx context.Context) (adapters.AddBranchProtectionOptions, diag.Diagnostics) {
	var approvalWhiteList []string
	diags := b.ApprovalsWhitelistUsername.ElementsAs(ctx, &approvalWhiteList, true)
	if diags.HasError() {
		return adapters.AddBranchProtectionOptions{}, diags
	}

	var mergeWhiteList []string
	diags = b.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, true)
	if diags.HasError() {
		return adapters.AddBranchProtectionOptions{}, diags
	}

	var pushWhiteList []string
	diags = b.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, true)
	if diags.HasError() {
		return adapters.AddBranchProtectionOptions{}, diags
	}

	return adapters.AddBranchProtectionOptions{
		Owner: b.Owner.ValueString(),
		Repo:  b.Repo.ValueString(),
		AddOption: api.CreateBranchProtectionOption{
			RuleName:                      b.BranchName.ValueStringPointer(),
			BranchName:                    b.BranchName.ValueStringPointer(),
			ApprovalsWhitelistUsername:    approvalWhiteList,
			BlockOnOfficialReviewRequests: b.BlockOnOfficialReviewRequests.ValueBoolPointer(),
			BlockOnOutdatedBranch:         b.BlockOnOutdatedBranch.ValueBoolPointer(),
			BlockOnRejectedReviews:        b.BlockOnRejectedReviews.ValueBoolPointer(),
			DismissStaleApprovals:         b.DismissStaleApprovals.ValueBoolPointer(),
			EnableApprovalsWhitelist:      b.EnableApprovalsWhitelist.ValueBoolPointer(),
			EnableMergeWhitelist:          b.EnableMergeWhitelist.ValueBoolPointer(),
			EnablePush:                    b.EnablePush.ValueBoolPointer(),
			EnablePushWhitelist:           b.EnablePushWhitelist.ValueBoolPointer(),
			MergeWhitelistUsernames:       mergeWhiteList,
			ProtectedFilePatterns:         b.ProtectedFilePatterns.ValueStringPointer(),
			PushWhitelistDeployKeys:       b.PushWhitelistDeployKeys.ValueBoolPointer(),
			PushWhitelistUsernames:        pushWhiteList,
			RequireSignedCommits:          b.RequireSignedCommits.ValueBoolPointer(),
			RequiredApprovals:             b.RequiredApprovals.ValueInt64Pointer(),
			UnprotectedFilePatterns:       b.UnprotectedFilePatterns.ValueStringPointer(),
		},
	}, diags
}

func (b *BranchProtectionResourceModel) ToApiEditBranchProtectionOptions(ctx context.Context) (adapters.EditBranchProtectionOptions, diag.Diagnostics) {
	var approvalWhiteList []string
	diags := b.ApprovalsWhitelistUsername.ElementsAs(ctx, &approvalWhiteList, true)
	if diags.HasError() {
		return adapters.EditBranchProtectionOptions{}, diags
	}

	var mergeWhiteList []string
	diags = b.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, true)
	if diags.HasError() {
		return adapters.EditBranchProtectionOptions{}, diags
	}

	var pushWhiteList []string
	diags = b.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, true)
	if diags.HasError() {
		return adapters.EditBranchProtectionOptions{}, diags
	}

	return adapters.EditBranchProtectionOptions{
		Owner:      b.Owner.ValueString(),
		Repo:       b.Repo.ValueString(),
		BranchName: b.BranchName.ValueString(),
		EditOption: api.EditBranchProtectionOption{
			ApprovalsWhitelistUsername:    approvalWhiteList,
			BlockOnOfficialReviewRequests: b.BlockOnOfficialReviewRequests.ValueBoolPointer(),
			BlockOnOutdatedBranch:         b.BlockOnOutdatedBranch.ValueBoolPointer(),
			BlockOnRejectedReviews:        b.BlockOnRejectedReviews.ValueBoolPointer(),
			DismissStaleApprovals:         b.DismissStaleApprovals.ValueBoolPointer(),
			EnableApprovalsWhitelist:      b.EnableApprovalsWhitelist.ValueBoolPointer(),
			EnableMergeWhitelist:          b.EnableMergeWhitelist.ValueBoolPointer(),
			EnablePush:                    b.EnablePush.ValueBoolPointer(),
			EnablePushWhitelist:           b.EnablePushWhitelist.ValueBoolPointer(),
			MergeWhitelistUsernames:       mergeWhiteList,
			ProtectedFilePatterns:         b.ProtectedFilePatterns.ValueStringPointer(),
			PushWhitelistDeployKeys:       b.PushWhitelistDeployKeys.ValueBoolPointer(),
			PushWhitelistUsernames:        pushWhiteList,
			RequireSignedCommits:          b.RequireSignedCommits.ValueBoolPointer(),
			RequiredApprovals:             b.RequiredApprovals.ValueInt64Pointer(),
			UnprotectedFilePatterns:       b.UnprotectedFilePatterns.ValueStringPointer(),
		},
	}, diags
}
