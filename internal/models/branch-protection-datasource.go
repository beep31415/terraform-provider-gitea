package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BranchProtectionDataSourceModel struct {
	Owner                         types.String `tfsdk:"owner"`
	Repo                          types.String `tfsdk:"repo"`
	BranchName                    types.String `tfsdk:"branch_name"`
	ApprovalsWhitelistUsernames   types.List   `tfsdk:"approvals_whitelist_usernames"`
	ApprovalsWhitelistTeams       types.List   `tfsdk:"approvals_whitelist_teams"`
	BlockOnOfficialReviewRequests types.Bool   `tfsdk:"block_on_official_review_requests"`
	BlockOnOutdatedBranch         types.Bool   `tfsdk:"block_on_outdated_branch"`
	BlockOnRejectedReviews        types.Bool   `tfsdk:"block_on_rejected_reviews"`
	DismissStaleApprovals         types.Bool   `tfsdk:"dismiss_stale_approvals"`
	EnableApprovalsWhitelist      types.Bool   `tfsdk:"enable_approvals_whitelist"`
	EnableMergeWhitelist          types.Bool   `tfsdk:"enable_merge_whitelist"`
	EnablePush                    types.Bool   `tfsdk:"enable_push"`
	EnablePushWhitelist           types.Bool   `tfsdk:"enable_push_whitelist"`
	MergeWhitelistUsernames       types.List   `tfsdk:"merge_whitelist_usernames"`
	MergeWhitelistTeams           types.List   `tfsdk:"merge_whitelist_teams"`
	ProtectedFilePatterns         types.String `tfsdk:"protected_file_patterns"`
	PushWhitelistDeployKeys       types.Bool   `tfsdk:"push_whitelist_deploy_keys"`
	PushWhitelistUsernames        types.List   `tfsdk:"push_whitelist_usernames"`
	PushWhitelistTeams            types.List   `tfsdk:"push_whitelist_teams"`
	RequireSignedCommits          types.Bool   `tfsdk:"require_signed_commits"`
	RequiredApprovals             types.Int64  `tfsdk:"required_approvals"`
	UnprotectedFilePatterns       types.String `tfsdk:"unprotected_file_patterns"`
}
