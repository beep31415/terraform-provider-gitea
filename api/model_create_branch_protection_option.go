/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.19.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateBranchProtectionOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBranchProtectionOption{}

// CreateBranchProtectionOption CreateBranchProtectionOption options for creating a branch protection
type CreateBranchProtectionOption struct {
	ApprovalsWhitelistTeams       []string `json:"approvals_whitelist_teams,omitempty"`
	ApprovalsWhitelistUsername    []string `json:"approvals_whitelist_username,omitempty"`
	BlockOnOfficialReviewRequests *bool    `json:"block_on_official_review_requests,omitempty"`
	BlockOnOutdatedBranch         *bool    `json:"block_on_outdated_branch,omitempty"`
	BlockOnRejectedReviews        *bool    `json:"block_on_rejected_reviews,omitempty"`
	// Deprecated: true
	BranchName               *string  `json:"branch_name,omitempty"`
	DismissStaleApprovals    *bool    `json:"dismiss_stale_approvals,omitempty"`
	EnableApprovalsWhitelist *bool    `json:"enable_approvals_whitelist,omitempty"`
	EnableMergeWhitelist     *bool    `json:"enable_merge_whitelist,omitempty"`
	EnablePush               *bool    `json:"enable_push,omitempty"`
	EnablePushWhitelist      *bool    `json:"enable_push_whitelist,omitempty"`
	EnableStatusCheck        *bool    `json:"enable_status_check,omitempty"`
	MergeWhitelistTeams      []string `json:"merge_whitelist_teams,omitempty"`
	MergeWhitelistUsernames  []string `json:"merge_whitelist_usernames,omitempty"`
	ProtectedFilePatterns    *string  `json:"protected_file_patterns,omitempty"`
	PushWhitelistDeployKeys  *bool    `json:"push_whitelist_deploy_keys,omitempty"`
	PushWhitelistTeams       []string `json:"push_whitelist_teams,omitempty"`
	PushWhitelistUsernames   []string `json:"push_whitelist_usernames,omitempty"`
	RequireSignedCommits     *bool    `json:"require_signed_commits,omitempty"`
	RequiredApprovals        *int64   `json:"required_approvals,omitempty"`
	RuleName                 *string  `json:"rule_name,omitempty"`
	StatusCheckContexts      []string `json:"status_check_contexts,omitempty"`
	UnprotectedFilePatterns  *string  `json:"unprotected_file_patterns,omitempty"`
}

// NewCreateBranchProtectionOption instantiates a new CreateBranchProtectionOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBranchProtectionOption() *CreateBranchProtectionOption {
	this := CreateBranchProtectionOption{}
	return &this
}

// NewCreateBranchProtectionOptionWithDefaults instantiates a new CreateBranchProtectionOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBranchProtectionOptionWithDefaults() *CreateBranchProtectionOption {
	this := CreateBranchProtectionOption{}
	return &this
}

// GetApprovalsWhitelistTeams returns the ApprovalsWhitelistTeams field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetApprovalsWhitelistTeams() []string {
	if o == nil || IsNil(o.ApprovalsWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.ApprovalsWhitelistTeams
}

// GetApprovalsWhitelistTeamsOk returns a tuple with the ApprovalsWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetApprovalsWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApprovalsWhitelistTeams) {
		return nil, false
	}
	return o.ApprovalsWhitelistTeams, true
}

// HasApprovalsWhitelistTeams returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasApprovalsWhitelistTeams() bool {
	if o != nil && !IsNil(o.ApprovalsWhitelistTeams) {
		return true
	}

	return false
}

// SetApprovalsWhitelistTeams gets a reference to the given []string and assigns it to the ApprovalsWhitelistTeams field.
func (o *CreateBranchProtectionOption) SetApprovalsWhitelistTeams(v []string) {
	o.ApprovalsWhitelistTeams = v
}

// GetApprovalsWhitelistUsername returns the ApprovalsWhitelistUsername field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetApprovalsWhitelistUsername() []string {
	if o == nil || IsNil(o.ApprovalsWhitelistUsername) {
		var ret []string
		return ret
	}
	return o.ApprovalsWhitelistUsername
}

// GetApprovalsWhitelistUsernameOk returns a tuple with the ApprovalsWhitelistUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetApprovalsWhitelistUsernameOk() ([]string, bool) {
	if o == nil || IsNil(o.ApprovalsWhitelistUsername) {
		return nil, false
	}
	return o.ApprovalsWhitelistUsername, true
}

// HasApprovalsWhitelistUsername returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasApprovalsWhitelistUsername() bool {
	if o != nil && !IsNil(o.ApprovalsWhitelistUsername) {
		return true
	}

	return false
}

// SetApprovalsWhitelistUsername gets a reference to the given []string and assigns it to the ApprovalsWhitelistUsername field.
func (o *CreateBranchProtectionOption) SetApprovalsWhitelistUsername(v []string) {
	o.ApprovalsWhitelistUsername = v
}

// GetBlockOnOfficialReviewRequests returns the BlockOnOfficialReviewRequests field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetBlockOnOfficialReviewRequests() bool {
	if o == nil || IsNil(o.BlockOnOfficialReviewRequests) {
		var ret bool
		return ret
	}
	return *o.BlockOnOfficialReviewRequests
}

// GetBlockOnOfficialReviewRequestsOk returns a tuple with the BlockOnOfficialReviewRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetBlockOnOfficialReviewRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnOfficialReviewRequests) {
		return nil, false
	}
	return o.BlockOnOfficialReviewRequests, true
}

// HasBlockOnOfficialReviewRequests returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasBlockOnOfficialReviewRequests() bool {
	if o != nil && !IsNil(o.BlockOnOfficialReviewRequests) {
		return true
	}

	return false
}

// SetBlockOnOfficialReviewRequests gets a reference to the given bool and assigns it to the BlockOnOfficialReviewRequests field.
func (o *CreateBranchProtectionOption) SetBlockOnOfficialReviewRequests(v bool) {
	o.BlockOnOfficialReviewRequests = &v
}

// GetBlockOnOutdatedBranch returns the BlockOnOutdatedBranch field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetBlockOnOutdatedBranch() bool {
	if o == nil || IsNil(o.BlockOnOutdatedBranch) {
		var ret bool
		return ret
	}
	return *o.BlockOnOutdatedBranch
}

// GetBlockOnOutdatedBranchOk returns a tuple with the BlockOnOutdatedBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetBlockOnOutdatedBranchOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnOutdatedBranch) {
		return nil, false
	}
	return o.BlockOnOutdatedBranch, true
}

// HasBlockOnOutdatedBranch returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasBlockOnOutdatedBranch() bool {
	if o != nil && !IsNil(o.BlockOnOutdatedBranch) {
		return true
	}

	return false
}

// SetBlockOnOutdatedBranch gets a reference to the given bool and assigns it to the BlockOnOutdatedBranch field.
func (o *CreateBranchProtectionOption) SetBlockOnOutdatedBranch(v bool) {
	o.BlockOnOutdatedBranch = &v
}

// GetBlockOnRejectedReviews returns the BlockOnRejectedReviews field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetBlockOnRejectedReviews() bool {
	if o == nil || IsNil(o.BlockOnRejectedReviews) {
		var ret bool
		return ret
	}
	return *o.BlockOnRejectedReviews
}

// GetBlockOnRejectedReviewsOk returns a tuple with the BlockOnRejectedReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetBlockOnRejectedReviewsOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnRejectedReviews) {
		return nil, false
	}
	return o.BlockOnRejectedReviews, true
}

// HasBlockOnRejectedReviews returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasBlockOnRejectedReviews() bool {
	if o != nil && !IsNil(o.BlockOnRejectedReviews) {
		return true
	}

	return false
}

// SetBlockOnRejectedReviews gets a reference to the given bool and assigns it to the BlockOnRejectedReviews field.
func (o *CreateBranchProtectionOption) SetBlockOnRejectedReviews(v bool) {
	o.BlockOnRejectedReviews = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *CreateBranchProtectionOption) SetBranchName(v string) {
	o.BranchName = &v
}

// GetDismissStaleApprovals returns the DismissStaleApprovals field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetDismissStaleApprovals() bool {
	if o == nil || IsNil(o.DismissStaleApprovals) {
		var ret bool
		return ret
	}
	return *o.DismissStaleApprovals
}

// GetDismissStaleApprovalsOk returns a tuple with the DismissStaleApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetDismissStaleApprovalsOk() (*bool, bool) {
	if o == nil || IsNil(o.DismissStaleApprovals) {
		return nil, false
	}
	return o.DismissStaleApprovals, true
}

// HasDismissStaleApprovals returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasDismissStaleApprovals() bool {
	if o != nil && !IsNil(o.DismissStaleApprovals) {
		return true
	}

	return false
}

// SetDismissStaleApprovals gets a reference to the given bool and assigns it to the DismissStaleApprovals field.
func (o *CreateBranchProtectionOption) SetDismissStaleApprovals(v bool) {
	o.DismissStaleApprovals = &v
}

// GetEnableApprovalsWhitelist returns the EnableApprovalsWhitelist field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetEnableApprovalsWhitelist() bool {
	if o == nil || IsNil(o.EnableApprovalsWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnableApprovalsWhitelist
}

// GetEnableApprovalsWhitelistOk returns a tuple with the EnableApprovalsWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetEnableApprovalsWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableApprovalsWhitelist) {
		return nil, false
	}
	return o.EnableApprovalsWhitelist, true
}

// HasEnableApprovalsWhitelist returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasEnableApprovalsWhitelist() bool {
	if o != nil && !IsNil(o.EnableApprovalsWhitelist) {
		return true
	}

	return false
}

// SetEnableApprovalsWhitelist gets a reference to the given bool and assigns it to the EnableApprovalsWhitelist field.
func (o *CreateBranchProtectionOption) SetEnableApprovalsWhitelist(v bool) {
	o.EnableApprovalsWhitelist = &v
}

// GetEnableMergeWhitelist returns the EnableMergeWhitelist field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetEnableMergeWhitelist() bool {
	if o == nil || IsNil(o.EnableMergeWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnableMergeWhitelist
}

// GetEnableMergeWhitelistOk returns a tuple with the EnableMergeWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetEnableMergeWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMergeWhitelist) {
		return nil, false
	}
	return o.EnableMergeWhitelist, true
}

// HasEnableMergeWhitelist returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasEnableMergeWhitelist() bool {
	if o != nil && !IsNil(o.EnableMergeWhitelist) {
		return true
	}

	return false
}

// SetEnableMergeWhitelist gets a reference to the given bool and assigns it to the EnableMergeWhitelist field.
func (o *CreateBranchProtectionOption) SetEnableMergeWhitelist(v bool) {
	o.EnableMergeWhitelist = &v
}

// GetEnablePush returns the EnablePush field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetEnablePush() bool {
	if o == nil || IsNil(o.EnablePush) {
		var ret bool
		return ret
	}
	return *o.EnablePush
}

// GetEnablePushOk returns a tuple with the EnablePush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetEnablePushOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePush) {
		return nil, false
	}
	return o.EnablePush, true
}

// HasEnablePush returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasEnablePush() bool {
	if o != nil && !IsNil(o.EnablePush) {
		return true
	}

	return false
}

// SetEnablePush gets a reference to the given bool and assigns it to the EnablePush field.
func (o *CreateBranchProtectionOption) SetEnablePush(v bool) {
	o.EnablePush = &v
}

// GetEnablePushWhitelist returns the EnablePushWhitelist field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetEnablePushWhitelist() bool {
	if o == nil || IsNil(o.EnablePushWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnablePushWhitelist
}

// GetEnablePushWhitelistOk returns a tuple with the EnablePushWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetEnablePushWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePushWhitelist) {
		return nil, false
	}
	return o.EnablePushWhitelist, true
}

// HasEnablePushWhitelist returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasEnablePushWhitelist() bool {
	if o != nil && !IsNil(o.EnablePushWhitelist) {
		return true
	}

	return false
}

// SetEnablePushWhitelist gets a reference to the given bool and assigns it to the EnablePushWhitelist field.
func (o *CreateBranchProtectionOption) SetEnablePushWhitelist(v bool) {
	o.EnablePushWhitelist = &v
}

// GetEnableStatusCheck returns the EnableStatusCheck field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetEnableStatusCheck() bool {
	if o == nil || IsNil(o.EnableStatusCheck) {
		var ret bool
		return ret
	}
	return *o.EnableStatusCheck
}

// GetEnableStatusCheckOk returns a tuple with the EnableStatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetEnableStatusCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableStatusCheck) {
		return nil, false
	}
	return o.EnableStatusCheck, true
}

// HasEnableStatusCheck returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasEnableStatusCheck() bool {
	if o != nil && !IsNil(o.EnableStatusCheck) {
		return true
	}

	return false
}

// SetEnableStatusCheck gets a reference to the given bool and assigns it to the EnableStatusCheck field.
func (o *CreateBranchProtectionOption) SetEnableStatusCheck(v bool) {
	o.EnableStatusCheck = &v
}

// GetMergeWhitelistTeams returns the MergeWhitelistTeams field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetMergeWhitelistTeams() []string {
	if o == nil || IsNil(o.MergeWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.MergeWhitelistTeams
}

// GetMergeWhitelistTeamsOk returns a tuple with the MergeWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetMergeWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.MergeWhitelistTeams) {
		return nil, false
	}
	return o.MergeWhitelistTeams, true
}

// HasMergeWhitelistTeams returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasMergeWhitelistTeams() bool {
	if o != nil && !IsNil(o.MergeWhitelistTeams) {
		return true
	}

	return false
}

// SetMergeWhitelistTeams gets a reference to the given []string and assigns it to the MergeWhitelistTeams field.
func (o *CreateBranchProtectionOption) SetMergeWhitelistTeams(v []string) {
	o.MergeWhitelistTeams = v
}

// GetMergeWhitelistUsernames returns the MergeWhitelistUsernames field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetMergeWhitelistUsernames() []string {
	if o == nil || IsNil(o.MergeWhitelistUsernames) {
		var ret []string
		return ret
	}
	return o.MergeWhitelistUsernames
}

// GetMergeWhitelistUsernamesOk returns a tuple with the MergeWhitelistUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetMergeWhitelistUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.MergeWhitelistUsernames) {
		return nil, false
	}
	return o.MergeWhitelistUsernames, true
}

// HasMergeWhitelistUsernames returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasMergeWhitelistUsernames() bool {
	if o != nil && !IsNil(o.MergeWhitelistUsernames) {
		return true
	}

	return false
}

// SetMergeWhitelistUsernames gets a reference to the given []string and assigns it to the MergeWhitelistUsernames field.
func (o *CreateBranchProtectionOption) SetMergeWhitelistUsernames(v []string) {
	o.MergeWhitelistUsernames = v
}

// GetProtectedFilePatterns returns the ProtectedFilePatterns field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetProtectedFilePatterns() string {
	if o == nil || IsNil(o.ProtectedFilePatterns) {
		var ret string
		return ret
	}
	return *o.ProtectedFilePatterns
}

// GetProtectedFilePatternsOk returns a tuple with the ProtectedFilePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetProtectedFilePatternsOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectedFilePatterns) {
		return nil, false
	}
	return o.ProtectedFilePatterns, true
}

// HasProtectedFilePatterns returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasProtectedFilePatterns() bool {
	if o != nil && !IsNil(o.ProtectedFilePatterns) {
		return true
	}

	return false
}

// SetProtectedFilePatterns gets a reference to the given string and assigns it to the ProtectedFilePatterns field.
func (o *CreateBranchProtectionOption) SetProtectedFilePatterns(v string) {
	o.ProtectedFilePatterns = &v
}

// GetPushWhitelistDeployKeys returns the PushWhitelistDeployKeys field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetPushWhitelistDeployKeys() bool {
	if o == nil || IsNil(o.PushWhitelistDeployKeys) {
		var ret bool
		return ret
	}
	return *o.PushWhitelistDeployKeys
}

// GetPushWhitelistDeployKeysOk returns a tuple with the PushWhitelistDeployKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetPushWhitelistDeployKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.PushWhitelistDeployKeys) {
		return nil, false
	}
	return o.PushWhitelistDeployKeys, true
}

// HasPushWhitelistDeployKeys returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasPushWhitelistDeployKeys() bool {
	if o != nil && !IsNil(o.PushWhitelistDeployKeys) {
		return true
	}

	return false
}

// SetPushWhitelistDeployKeys gets a reference to the given bool and assigns it to the PushWhitelistDeployKeys field.
func (o *CreateBranchProtectionOption) SetPushWhitelistDeployKeys(v bool) {
	o.PushWhitelistDeployKeys = &v
}

// GetPushWhitelistTeams returns the PushWhitelistTeams field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetPushWhitelistTeams() []string {
	if o == nil || IsNil(o.PushWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.PushWhitelistTeams
}

// GetPushWhitelistTeamsOk returns a tuple with the PushWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetPushWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.PushWhitelistTeams) {
		return nil, false
	}
	return o.PushWhitelistTeams, true
}

// HasPushWhitelistTeams returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasPushWhitelistTeams() bool {
	if o != nil && !IsNil(o.PushWhitelistTeams) {
		return true
	}

	return false
}

// SetPushWhitelistTeams gets a reference to the given []string and assigns it to the PushWhitelistTeams field.
func (o *CreateBranchProtectionOption) SetPushWhitelistTeams(v []string) {
	o.PushWhitelistTeams = v
}

// GetPushWhitelistUsernames returns the PushWhitelistUsernames field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetPushWhitelistUsernames() []string {
	if o == nil || IsNil(o.PushWhitelistUsernames) {
		var ret []string
		return ret
	}
	return o.PushWhitelistUsernames
}

// GetPushWhitelistUsernamesOk returns a tuple with the PushWhitelistUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetPushWhitelistUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.PushWhitelistUsernames) {
		return nil, false
	}
	return o.PushWhitelistUsernames, true
}

// HasPushWhitelistUsernames returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasPushWhitelistUsernames() bool {
	if o != nil && !IsNil(o.PushWhitelistUsernames) {
		return true
	}

	return false
}

// SetPushWhitelistUsernames gets a reference to the given []string and assigns it to the PushWhitelistUsernames field.
func (o *CreateBranchProtectionOption) SetPushWhitelistUsernames(v []string) {
	o.PushWhitelistUsernames = v
}

// GetRequireSignedCommits returns the RequireSignedCommits field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetRequireSignedCommits() bool {
	if o == nil || IsNil(o.RequireSignedCommits) {
		var ret bool
		return ret
	}
	return *o.RequireSignedCommits
}

// GetRequireSignedCommitsOk returns a tuple with the RequireSignedCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetRequireSignedCommitsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedCommits) {
		return nil, false
	}
	return o.RequireSignedCommits, true
}

// HasRequireSignedCommits returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasRequireSignedCommits() bool {
	if o != nil && !IsNil(o.RequireSignedCommits) {
		return true
	}

	return false
}

// SetRequireSignedCommits gets a reference to the given bool and assigns it to the RequireSignedCommits field.
func (o *CreateBranchProtectionOption) SetRequireSignedCommits(v bool) {
	o.RequireSignedCommits = &v
}

// GetRequiredApprovals returns the RequiredApprovals field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetRequiredApprovals() int64 {
	if o == nil || IsNil(o.RequiredApprovals) {
		var ret int64
		return ret
	}
	return *o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetRequiredApprovalsOk() (*int64, bool) {
	if o == nil || IsNil(o.RequiredApprovals) {
		return nil, false
	}
	return o.RequiredApprovals, true
}

// HasRequiredApprovals returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasRequiredApprovals() bool {
	if o != nil && !IsNil(o.RequiredApprovals) {
		return true
	}

	return false
}

// SetRequiredApprovals gets a reference to the given int64 and assigns it to the RequiredApprovals field.
func (o *CreateBranchProtectionOption) SetRequiredApprovals(v int64) {
	o.RequiredApprovals = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *CreateBranchProtectionOption) SetRuleName(v string) {
	o.RuleName = &v
}

// GetStatusCheckContexts returns the StatusCheckContexts field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetStatusCheckContexts() []string {
	if o == nil || IsNil(o.StatusCheckContexts) {
		var ret []string
		return ret
	}
	return o.StatusCheckContexts
}

// GetStatusCheckContextsOk returns a tuple with the StatusCheckContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetStatusCheckContextsOk() ([]string, bool) {
	if o == nil || IsNil(o.StatusCheckContexts) {
		return nil, false
	}
	return o.StatusCheckContexts, true
}

// HasStatusCheckContexts returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasStatusCheckContexts() bool {
	if o != nil && !IsNil(o.StatusCheckContexts) {
		return true
	}

	return false
}

// SetStatusCheckContexts gets a reference to the given []string and assigns it to the StatusCheckContexts field.
func (o *CreateBranchProtectionOption) SetStatusCheckContexts(v []string) {
	o.StatusCheckContexts = v
}

// GetUnprotectedFilePatterns returns the UnprotectedFilePatterns field value if set, zero value otherwise.
func (o *CreateBranchProtectionOption) GetUnprotectedFilePatterns() string {
	if o == nil || IsNil(o.UnprotectedFilePatterns) {
		var ret string
		return ret
	}
	return *o.UnprotectedFilePatterns
}

// GetUnprotectedFilePatternsOk returns a tuple with the UnprotectedFilePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionOption) GetUnprotectedFilePatternsOk() (*string, bool) {
	if o == nil || IsNil(o.UnprotectedFilePatterns) {
		return nil, false
	}
	return o.UnprotectedFilePatterns, true
}

// HasUnprotectedFilePatterns returns a boolean if a field has been set.
func (o *CreateBranchProtectionOption) HasUnprotectedFilePatterns() bool {
	if o != nil && !IsNil(o.UnprotectedFilePatterns) {
		return true
	}

	return false
}

// SetUnprotectedFilePatterns gets a reference to the given string and assigns it to the UnprotectedFilePatterns field.
func (o *CreateBranchProtectionOption) SetUnprotectedFilePatterns(v string) {
	o.UnprotectedFilePatterns = &v
}

func (o CreateBranchProtectionOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBranchProtectionOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprovalsWhitelistTeams) {
		toSerialize["approvals_whitelist_teams"] = o.ApprovalsWhitelistTeams
	}
	if !IsNil(o.ApprovalsWhitelistUsername) {
		toSerialize["approvals_whitelist_username"] = o.ApprovalsWhitelistUsername
	}
	if !IsNil(o.BlockOnOfficialReviewRequests) {
		toSerialize["block_on_official_review_requests"] = o.BlockOnOfficialReviewRequests
	}
	if !IsNil(o.BlockOnOutdatedBranch) {
		toSerialize["block_on_outdated_branch"] = o.BlockOnOutdatedBranch
	}
	if !IsNil(o.BlockOnRejectedReviews) {
		toSerialize["block_on_rejected_reviews"] = o.BlockOnRejectedReviews
	}
	if !IsNil(o.BranchName) {
		toSerialize["branch_name"] = o.BranchName
	}
	if !IsNil(o.DismissStaleApprovals) {
		toSerialize["dismiss_stale_approvals"] = o.DismissStaleApprovals
	}
	if !IsNil(o.EnableApprovalsWhitelist) {
		toSerialize["enable_approvals_whitelist"] = o.EnableApprovalsWhitelist
	}
	if !IsNil(o.EnableMergeWhitelist) {
		toSerialize["enable_merge_whitelist"] = o.EnableMergeWhitelist
	}
	if !IsNil(o.EnablePush) {
		toSerialize["enable_push"] = o.EnablePush
	}
	if !IsNil(o.EnablePushWhitelist) {
		toSerialize["enable_push_whitelist"] = o.EnablePushWhitelist
	}
	if !IsNil(o.EnableStatusCheck) {
		toSerialize["enable_status_check"] = o.EnableStatusCheck
	}
	if !IsNil(o.MergeWhitelistTeams) {
		toSerialize["merge_whitelist_teams"] = o.MergeWhitelistTeams
	}
	if !IsNil(o.MergeWhitelistUsernames) {
		toSerialize["merge_whitelist_usernames"] = o.MergeWhitelistUsernames
	}
	if !IsNil(o.ProtectedFilePatterns) {
		toSerialize["protected_file_patterns"] = o.ProtectedFilePatterns
	}
	if !IsNil(o.PushWhitelistDeployKeys) {
		toSerialize["push_whitelist_deploy_keys"] = o.PushWhitelistDeployKeys
	}
	if !IsNil(o.PushWhitelistTeams) {
		toSerialize["push_whitelist_teams"] = o.PushWhitelistTeams
	}
	if !IsNil(o.PushWhitelistUsernames) {
		toSerialize["push_whitelist_usernames"] = o.PushWhitelistUsernames
	}
	if !IsNil(o.RequireSignedCommits) {
		toSerialize["require_signed_commits"] = o.RequireSignedCommits
	}
	if !IsNil(o.RequiredApprovals) {
		toSerialize["required_approvals"] = o.RequiredApprovals
	}
	if !IsNil(o.RuleName) {
		toSerialize["rule_name"] = o.RuleName
	}
	if !IsNil(o.StatusCheckContexts) {
		toSerialize["status_check_contexts"] = o.StatusCheckContexts
	}
	if !IsNil(o.UnprotectedFilePatterns) {
		toSerialize["unprotected_file_patterns"] = o.UnprotectedFilePatterns
	}
	return toSerialize, nil
}

type NullableCreateBranchProtectionOption struct {
	value *CreateBranchProtectionOption
	isSet bool
}

func (v NullableCreateBranchProtectionOption) Get() *CreateBranchProtectionOption {
	return v.value
}

func (v *NullableCreateBranchProtectionOption) Set(val *CreateBranchProtectionOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBranchProtectionOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBranchProtectionOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBranchProtectionOption(val *CreateBranchProtectionOption) *NullableCreateBranchProtectionOption {
	return &NullableCreateBranchProtectionOption{value: val, isSet: true}
}

func (v NullableCreateBranchProtectionOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBranchProtectionOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
