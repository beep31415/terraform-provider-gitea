package models

import (
	"context"
	"terraform-provider-gitea/internal/adapters"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BranchProtectionDataSourceModel struct {
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

func NewBranchProtectionDataSource(ctx context.Context, branchProtection *adapters.BranchProtection) (BranchProtectionDataSourceModel, diag.Diagnostics) {
	approvalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetApprovalsWhitelistUsername())
	if diags.HasError() {
		return BranchProtectionDataSourceModel{}, diags
	}

	mergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetMergeWhitelistUsernames())
	if diags.HasError() {
		return BranchProtectionDataSourceModel{}, diags
	}

	pushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.BranchProtection.GetPushWhitelistUsernames())
	if diags.HasError() {
		return BranchProtectionDataSourceModel{}, diags
	}

	return BranchProtectionDataSourceModel{
		Owner:                         types.StringValue(branchProtection.Owner),
		Repo:                          types.StringValue(branchProtection.Repo),
		BranchName:                    types.StringValue(branchProtection.BranchProtection.GetBranchName()),
		ApprovalsWhitelistUsername:    approvalWhiteList,
		BlockOnOfficialReviewRequests: types.BoolValue(branchProtection.BranchProtection.GetBlockOnOfficialReviewRequests()),
		BlockOnOutdatedBranch:         types.BoolValue(branchProtection.BranchProtection.GetBlockOnOutdatedBranch()),
		BlockOnRejectedReviews:        types.BoolValue(branchProtection.BranchProtection.GetBlockOnRejectedReviews()),
		DismissStaleApprovals:         types.BoolValue(branchProtection.BranchProtection.GetDismissStaleApprovals()),
		EnableApprovalsWhitelist:      types.BoolValue(branchProtection.BranchProtection.GetEnableApprovalsWhitelist()),
		EnableMergeWhitelist:          types.BoolValue(branchProtection.BranchProtection.GetEnableMergeWhitelist()),
		EnablePush:                    types.BoolValue(branchProtection.BranchProtection.GetEnablePush()),
		EnablePushWhitelist:           types.BoolValue(branchProtection.BranchProtection.GetEnablePushWhitelist()),
		MergeWhitelistUsernames:       mergeWhiteList,
		ProtectedFilePatterns:         types.StringValue(branchProtection.BranchProtection.GetProtectedFilePatterns()),
		PushWhitelistDeployKeys:       types.BoolValue(branchProtection.BranchProtection.GetPushWhitelistDeployKeys()),
		PushWhitelistUsernames:        pushWhiteList,
		RequireSignedCommits:          types.BoolValue(branchProtection.BranchProtection.GetRequireSignedCommits()),
		RequiredApprovals:             types.Int64Value(branchProtection.BranchProtection.GetRequiredApprovals()),
		UnprotectedFilePatterns:       types.StringValue(branchProtection.BranchProtection.GetUnprotectedFilePatterns()),
	}, diags
}
