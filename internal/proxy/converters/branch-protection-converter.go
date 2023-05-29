package converters

import (
	"context"
	"terraform-provider-gitea/internal/models"
	"terraform-provider-gitea/internal/proxy/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BranchProtectionConverter struct{}

func (BranchProtectionConverter) ReadToDataSource(ctx context.Context,
	model *models.BranchProtectionDataSourceModel,
	branchProtection *api.BranchProtection) diag.Diagnostics {

	approvalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetApprovalsWhitelistUsername())
	if diags.HasError() {
		return diags
	}

	mergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetMergeWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	pushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetPushWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	approvalWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetApprovalsWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	mergeWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetMergeWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	pushWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetPushWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	model.BranchName = types.StringValue(branchProtection.GetBranchName())
	model.BlockOnOfficialReviewRequests = types.BoolValue(branchProtection.GetBlockOnOfficialReviewRequests())
	model.BlockOnOutdatedBranch = types.BoolValue(branchProtection.GetBlockOnOutdatedBranch())
	model.BlockOnRejectedReviews = types.BoolValue(branchProtection.GetBlockOnRejectedReviews())
	model.DismissStaleApprovals = types.BoolValue(branchProtection.GetDismissStaleApprovals())
	model.EnableApprovalsWhitelist = types.BoolValue(branchProtection.GetEnableApprovalsWhitelist())
	model.EnableMergeWhitelist = types.BoolValue(branchProtection.GetEnableMergeWhitelist())
	model.EnablePush = types.BoolValue(branchProtection.GetEnablePush())
	model.EnablePushWhitelist = types.BoolValue(branchProtection.GetEnablePushWhitelist())
	model.ProtectedFilePatterns = types.StringValue(branchProtection.GetProtectedFilePatterns())
	model.PushWhitelistDeployKeys = types.BoolValue(branchProtection.GetPushWhitelistDeployKeys())
	model.RequireSignedCommits = types.BoolValue(branchProtection.GetRequireSignedCommits())
	model.RequiredApprovals = types.Int64Value(branchProtection.GetRequiredApprovals())
	model.UnprotectedFilePatterns = types.StringValue(branchProtection.GetUnprotectedFilePatterns())
	model.EnableStatusCheck = types.BoolValue(branchProtection.GetEnableStatusCheck())
	model.ApprovalsWhitelistUsernames = approvalWhiteList
	model.MergeWhitelistUsernames = mergeWhiteList
	model.PushWhitelistUsernames = pushWhiteList
	model.ApprovalsWhitelistTeams = approvalWhiteListTeams
	model.MergeWhitelistTeams = mergeWhiteListTeams
	model.PushWhitelistTeams = pushWhiteListTeams

	return nil
}

func (BranchProtectionConverter) ReadToResource(ctx context.Context,
	model *models.BranchProtectionResourceModel,
	branchProtection *api.BranchProtection) diag.Diagnostics {

	tfApprovalWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetApprovalsWhitelistUsername())
	if diags.HasError() {
		return diags
	}

	tfMergeWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetMergeWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	tfPushWhiteList, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetPushWhitelistUsernames())
	if diags.HasError() {
		return diags
	}

	tfApprovalWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetApprovalsWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	tfMergeWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetMergeWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	tfPushWhiteListTeams, diags := types.ListValueFrom(ctx, types.StringType, branchProtection.GetPushWhitelistTeams())
	if diags.HasError() {
		return diags
	}

	model.BranchName = types.StringValue(branchProtection.GetBranchName())
	model.BlockOnOfficialReviewRequests = types.BoolValue(branchProtection.GetBlockOnOfficialReviewRequests())
	model.BlockOnOutdatedBranch = types.BoolValue(branchProtection.GetBlockOnOutdatedBranch())
	model.BlockOnRejectedReviews = types.BoolValue(branchProtection.GetBlockOnRejectedReviews())
	model.DismissStaleApprovals = types.BoolValue(branchProtection.GetDismissStaleApprovals())
	model.EnableApprovalsWhitelist = types.BoolValue(branchProtection.GetEnableApprovalsWhitelist())
	model.EnableMergeWhitelist = types.BoolValue(branchProtection.GetEnableMergeWhitelist())
	model.EnablePush = types.BoolValue(branchProtection.GetEnablePush())
	model.EnablePushWhitelist = types.BoolValue(branchProtection.GetEnablePushWhitelist())
	model.ProtectedFilePatterns = types.StringValue(branchProtection.GetProtectedFilePatterns())
	model.PushWhitelistDeployKeys = types.BoolValue(branchProtection.GetPushWhitelistDeployKeys())
	model.RequireSignedCommits = types.BoolValue(branchProtection.GetRequireSignedCommits())
	model.RequiredApprovals = types.Int64Value(branchProtection.GetRequiredApprovals())
	model.UnprotectedFilePatterns = types.StringValue(branchProtection.GetUnprotectedFilePatterns())
	model.EnableStatusCheck = types.BoolValue(branchProtection.GetEnableStatusCheck())
	model.ApprovalsWhitelistUsernames = tfApprovalWhiteList
	model.MergeWhitelistUsernames = tfMergeWhiteList
	model.PushWhitelistUsernames = tfPushWhiteList
	model.ApprovalsWhitelistTeams = tfApprovalWhiteListTeams
	model.MergeWhitelistTeams = tfMergeWhiteListTeams
	model.PushWhitelistTeams = tfPushWhiteListTeams

	return diags
}

func (BranchProtectionConverter) ToApiAddBranchProtectionOptions(ctx context.Context,
	model models.BranchProtectionResourceModel) (api.CreateBranchProtectionOption, diag.Diagnostics) {

	var approvalWhiteList []string
	diags := model.ApprovalsWhitelistUsernames.ElementsAs(ctx, &approvalWhiteList, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	var mergeWhiteList []string
	diags = model.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	var pushWhiteList []string
	diags = model.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	var approvalWhiteListTeams []string
	diags = model.ApprovalsWhitelistTeams.ElementsAs(ctx, &approvalWhiteListTeams, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	var mergeWhiteListTeams []string
	diags = model.MergeWhitelistTeams.ElementsAs(ctx, &mergeWhiteListTeams, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	var pushWhiteListTeams []string
	diags = model.PushWhitelistTeams.ElementsAs(ctx, &pushWhiteListTeams, true)
	if diags.HasError() {
		return api.CreateBranchProtectionOption{}, diags
	}

	return api.CreateBranchProtectionOption{
		RuleName:                      model.BranchName.ValueStringPointer(),
		BranchName:                    model.BranchName.ValueStringPointer(),
		BlockOnOfficialReviewRequests: model.BlockOnOfficialReviewRequests.ValueBoolPointer(),
		BlockOnOutdatedBranch:         model.BlockOnOutdatedBranch.ValueBoolPointer(),
		BlockOnRejectedReviews:        model.BlockOnRejectedReviews.ValueBoolPointer(),
		DismissStaleApprovals:         model.DismissStaleApprovals.ValueBoolPointer(),
		EnableApprovalsWhitelist:      model.EnableApprovalsWhitelist.ValueBoolPointer(),
		EnableMergeWhitelist:          model.EnableMergeWhitelist.ValueBoolPointer(),
		EnablePush:                    model.EnablePush.ValueBoolPointer(),
		EnablePushWhitelist:           model.EnablePushWhitelist.ValueBoolPointer(),
		ProtectedFilePatterns:         model.ProtectedFilePatterns.ValueStringPointer(),
		PushWhitelistDeployKeys:       model.PushWhitelistDeployKeys.ValueBoolPointer(),
		RequireSignedCommits:          model.RequireSignedCommits.ValueBoolPointer(),
		RequiredApprovals:             model.RequiredApprovals.ValueInt64Pointer(),
		UnprotectedFilePatterns:       model.UnprotectedFilePatterns.ValueStringPointer(),
		EnableStatusCheck:             model.EnableStatusCheck.ValueBoolPointer(),
		ApprovalsWhitelistUsername:    approvalWhiteList,
		MergeWhitelistUsernames:       mergeWhiteList,
		PushWhitelistUsernames:        pushWhiteList,
		ApprovalsWhitelistTeams:       approvalWhiteListTeams,
		MergeWhitelistTeams:           mergeWhiteListTeams,
		PushWhitelistTeams:            pushWhiteListTeams,
	}, nil
}

func (BranchProtectionConverter) ToApiEditBranchProtectionOptions(ctx context.Context,
	model models.BranchProtectionResourceModel) (api.EditBranchProtectionOption, diag.Diagnostics) {

	var approvalWhiteList []string
	diags := model.ApprovalsWhitelistUsernames.ElementsAs(ctx, &approvalWhiteList, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	var mergeWhiteList []string
	diags = model.MergeWhitelistUsernames.ElementsAs(ctx, &mergeWhiteList, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	var pushWhiteList []string
	diags = model.PushWhitelistUsernames.ElementsAs(ctx, &pushWhiteList, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	var approvalWhiteListTeams []string
	diags = model.ApprovalsWhitelistTeams.ElementsAs(ctx, &approvalWhiteListTeams, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	var mergeWhiteListTeams []string
	diags = model.MergeWhitelistTeams.ElementsAs(ctx, &mergeWhiteListTeams, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	var pushWhiteListTeams []string
	diags = model.PushWhitelistTeams.ElementsAs(ctx, &pushWhiteListTeams, true)
	if diags.HasError() {
		return api.EditBranchProtectionOption{}, diags
	}

	return api.EditBranchProtectionOption{
		BlockOnOfficialReviewRequests: model.BlockOnOfficialReviewRequests.ValueBoolPointer(),
		BlockOnOutdatedBranch:         model.BlockOnOutdatedBranch.ValueBoolPointer(),
		BlockOnRejectedReviews:        model.BlockOnRejectedReviews.ValueBoolPointer(),
		DismissStaleApprovals:         model.DismissStaleApprovals.ValueBoolPointer(),
		EnableApprovalsWhitelist:      model.EnableApprovalsWhitelist.ValueBoolPointer(),
		EnableMergeWhitelist:          model.EnableMergeWhitelist.ValueBoolPointer(),
		EnablePush:                    model.EnablePush.ValueBoolPointer(),
		EnablePushWhitelist:           model.EnablePushWhitelist.ValueBoolPointer(),
		ProtectedFilePatterns:         model.ProtectedFilePatterns.ValueStringPointer(),
		PushWhitelistDeployKeys:       model.PushWhitelistDeployKeys.ValueBoolPointer(),
		RequireSignedCommits:          model.RequireSignedCommits.ValueBoolPointer(),
		RequiredApprovals:             model.RequiredApprovals.ValueInt64Pointer(),
		UnprotectedFilePatterns:       model.UnprotectedFilePatterns.ValueStringPointer(),
		EnableStatusCheck:             model.EnableStatusCheck.ValueBoolPointer(),
		ApprovalsWhitelistUsername:    approvalWhiteList,
		MergeWhitelistUsernames:       mergeWhiteList,
		PushWhitelistUsernames:        pushWhiteList,
		ApprovalsWhitelistTeams:       approvalWhiteListTeams,
		MergeWhitelistTeams:           mergeWhiteListTeams,
		PushWhitelistTeams:            pushWhiteListTeams,
	}, nil
}
