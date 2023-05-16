/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.19.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the BranchProtection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BranchProtection{}

// BranchProtection BranchProtection represents a branch protection for a repository
type BranchProtection struct {
	ApprovalsWhitelistTeams       []string `json:"approvals_whitelist_teams,omitempty"`
	ApprovalsWhitelistUsername    []string `json:"approvals_whitelist_username,omitempty"`
	BlockOnOfficialReviewRequests *bool    `json:"block_on_official_review_requests,omitempty"`
	BlockOnOutdatedBranch         *bool    `json:"block_on_outdated_branch,omitempty"`
	BlockOnRejectedReviews        *bool    `json:"block_on_rejected_reviews,omitempty"`
	// Deprecated: true
	BranchName               *string    `json:"branch_name,omitempty"`
	CreatedAt                *time.Time `json:"created_at,omitempty"`
	DismissStaleApprovals    *bool      `json:"dismiss_stale_approvals,omitempty"`
	EnableApprovalsWhitelist *bool      `json:"enable_approvals_whitelist,omitempty"`
	EnableMergeWhitelist     *bool      `json:"enable_merge_whitelist,omitempty"`
	EnablePush               *bool      `json:"enable_push,omitempty"`
	EnablePushWhitelist      *bool      `json:"enable_push_whitelist,omitempty"`
	EnableStatusCheck        *bool      `json:"enable_status_check,omitempty"`
	MergeWhitelistTeams      []string   `json:"merge_whitelist_teams,omitempty"`
	MergeWhitelistUsernames  []string   `json:"merge_whitelist_usernames,omitempty"`
	ProtectedFilePatterns    *string    `json:"protected_file_patterns,omitempty"`
	PushWhitelistDeployKeys  *bool      `json:"push_whitelist_deploy_keys,omitempty"`
	PushWhitelistTeams       []string   `json:"push_whitelist_teams,omitempty"`
	PushWhitelistUsernames   []string   `json:"push_whitelist_usernames,omitempty"`
	RequireSignedCommits     *bool      `json:"require_signed_commits,omitempty"`
	RequiredApprovals        *int64     `json:"required_approvals,omitempty"`
	RuleName                 *string    `json:"rule_name,omitempty"`
	StatusCheckContexts      []string   `json:"status_check_contexts,omitempty"`
	UnprotectedFilePatterns  *string    `json:"unprotected_file_patterns,omitempty"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty"`
}

// NewBranchProtection instantiates a new BranchProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchProtection() *BranchProtection {
	this := BranchProtection{}
	return &this
}

// NewBranchProtectionWithDefaults instantiates a new BranchProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchProtectionWithDefaults() *BranchProtection {
	this := BranchProtection{}
	return &this
}

// GetApprovalsWhitelistTeams returns the ApprovalsWhitelistTeams field value if set, zero value otherwise.
func (o *BranchProtection) GetApprovalsWhitelistTeams() []string {
	if o == nil || IsNil(o.ApprovalsWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.ApprovalsWhitelistTeams
}

// GetApprovalsWhitelistTeamsOk returns a tuple with the ApprovalsWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetApprovalsWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApprovalsWhitelistTeams) {
		return nil, false
	}
	return o.ApprovalsWhitelistTeams, true
}

// HasApprovalsWhitelistTeams returns a boolean if a field has been set.
func (o *BranchProtection) HasApprovalsWhitelistTeams() bool {
	if o != nil && !IsNil(o.ApprovalsWhitelistTeams) {
		return true
	}

	return false
}

// SetApprovalsWhitelistTeams gets a reference to the given []string and assigns it to the ApprovalsWhitelistTeams field.
func (o *BranchProtection) SetApprovalsWhitelistTeams(v []string) {
	o.ApprovalsWhitelistTeams = v
}

// GetApprovalsWhitelistUsername returns the ApprovalsWhitelistUsername field value if set, zero value otherwise.
func (o *BranchProtection) GetApprovalsWhitelistUsername() []string {
	if o == nil || IsNil(o.ApprovalsWhitelistUsername) {
		var ret []string
		return ret
	}
	return o.ApprovalsWhitelistUsername
}

// GetApprovalsWhitelistUsernameOk returns a tuple with the ApprovalsWhitelistUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetApprovalsWhitelistUsernameOk() ([]string, bool) {
	if o == nil || IsNil(o.ApprovalsWhitelistUsername) {
		return nil, false
	}
	return o.ApprovalsWhitelistUsername, true
}

// HasApprovalsWhitelistUsername returns a boolean if a field has been set.
func (o *BranchProtection) HasApprovalsWhitelistUsername() bool {
	if o != nil && !IsNil(o.ApprovalsWhitelistUsername) {
		return true
	}

	return false
}

// SetApprovalsWhitelistUsername gets a reference to the given []string and assigns it to the ApprovalsWhitelistUsername field.
func (o *BranchProtection) SetApprovalsWhitelistUsername(v []string) {
	o.ApprovalsWhitelistUsername = v
}

// GetBlockOnOfficialReviewRequests returns the BlockOnOfficialReviewRequests field value if set, zero value otherwise.
func (o *BranchProtection) GetBlockOnOfficialReviewRequests() bool {
	if o == nil || IsNil(o.BlockOnOfficialReviewRequests) {
		var ret bool
		return ret
	}
	return *o.BlockOnOfficialReviewRequests
}

// GetBlockOnOfficialReviewRequestsOk returns a tuple with the BlockOnOfficialReviewRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetBlockOnOfficialReviewRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnOfficialReviewRequests) {
		return nil, false
	}
	return o.BlockOnOfficialReviewRequests, true
}

// HasBlockOnOfficialReviewRequests returns a boolean if a field has been set.
func (o *BranchProtection) HasBlockOnOfficialReviewRequests() bool {
	if o != nil && !IsNil(o.BlockOnOfficialReviewRequests) {
		return true
	}

	return false
}

// SetBlockOnOfficialReviewRequests gets a reference to the given bool and assigns it to the BlockOnOfficialReviewRequests field.
func (o *BranchProtection) SetBlockOnOfficialReviewRequests(v bool) {
	o.BlockOnOfficialReviewRequests = &v
}

// GetBlockOnOutdatedBranch returns the BlockOnOutdatedBranch field value if set, zero value otherwise.
func (o *BranchProtection) GetBlockOnOutdatedBranch() bool {
	if o == nil || IsNil(o.BlockOnOutdatedBranch) {
		var ret bool
		return ret
	}
	return *o.BlockOnOutdatedBranch
}

// GetBlockOnOutdatedBranchOk returns a tuple with the BlockOnOutdatedBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetBlockOnOutdatedBranchOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnOutdatedBranch) {
		return nil, false
	}
	return o.BlockOnOutdatedBranch, true
}

// HasBlockOnOutdatedBranch returns a boolean if a field has been set.
func (o *BranchProtection) HasBlockOnOutdatedBranch() bool {
	if o != nil && !IsNil(o.BlockOnOutdatedBranch) {
		return true
	}

	return false
}

// SetBlockOnOutdatedBranch gets a reference to the given bool and assigns it to the BlockOnOutdatedBranch field.
func (o *BranchProtection) SetBlockOnOutdatedBranch(v bool) {
	o.BlockOnOutdatedBranch = &v
}

// GetBlockOnRejectedReviews returns the BlockOnRejectedReviews field value if set, zero value otherwise.
func (o *BranchProtection) GetBlockOnRejectedReviews() bool {
	if o == nil || IsNil(o.BlockOnRejectedReviews) {
		var ret bool
		return ret
	}
	return *o.BlockOnRejectedReviews
}

// GetBlockOnRejectedReviewsOk returns a tuple with the BlockOnRejectedReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetBlockOnRejectedReviewsOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOnRejectedReviews) {
		return nil, false
	}
	return o.BlockOnRejectedReviews, true
}

// HasBlockOnRejectedReviews returns a boolean if a field has been set.
func (o *BranchProtection) HasBlockOnRejectedReviews() bool {
	if o != nil && !IsNil(o.BlockOnRejectedReviews) {
		return true
	}

	return false
}

// SetBlockOnRejectedReviews gets a reference to the given bool and assigns it to the BlockOnRejectedReviews field.
func (o *BranchProtection) SetBlockOnRejectedReviews(v bool) {
	o.BlockOnRejectedReviews = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *BranchProtection) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *BranchProtection) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *BranchProtection) SetBranchName(v string) {
	o.BranchName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BranchProtection) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BranchProtection) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BranchProtection) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDismissStaleApprovals returns the DismissStaleApprovals field value if set, zero value otherwise.
func (o *BranchProtection) GetDismissStaleApprovals() bool {
	if o == nil || IsNil(o.DismissStaleApprovals) {
		var ret bool
		return ret
	}
	return *o.DismissStaleApprovals
}

// GetDismissStaleApprovalsOk returns a tuple with the DismissStaleApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetDismissStaleApprovalsOk() (*bool, bool) {
	if o == nil || IsNil(o.DismissStaleApprovals) {
		return nil, false
	}
	return o.DismissStaleApprovals, true
}

// HasDismissStaleApprovals returns a boolean if a field has been set.
func (o *BranchProtection) HasDismissStaleApprovals() bool {
	if o != nil && !IsNil(o.DismissStaleApprovals) {
		return true
	}

	return false
}

// SetDismissStaleApprovals gets a reference to the given bool and assigns it to the DismissStaleApprovals field.
func (o *BranchProtection) SetDismissStaleApprovals(v bool) {
	o.DismissStaleApprovals = &v
}

// GetEnableApprovalsWhitelist returns the EnableApprovalsWhitelist field value if set, zero value otherwise.
func (o *BranchProtection) GetEnableApprovalsWhitelist() bool {
	if o == nil || IsNil(o.EnableApprovalsWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnableApprovalsWhitelist
}

// GetEnableApprovalsWhitelistOk returns a tuple with the EnableApprovalsWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetEnableApprovalsWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableApprovalsWhitelist) {
		return nil, false
	}
	return o.EnableApprovalsWhitelist, true
}

// HasEnableApprovalsWhitelist returns a boolean if a field has been set.
func (o *BranchProtection) HasEnableApprovalsWhitelist() bool {
	if o != nil && !IsNil(o.EnableApprovalsWhitelist) {
		return true
	}

	return false
}

// SetEnableApprovalsWhitelist gets a reference to the given bool and assigns it to the EnableApprovalsWhitelist field.
func (o *BranchProtection) SetEnableApprovalsWhitelist(v bool) {
	o.EnableApprovalsWhitelist = &v
}

// GetEnableMergeWhitelist returns the EnableMergeWhitelist field value if set, zero value otherwise.
func (o *BranchProtection) GetEnableMergeWhitelist() bool {
	if o == nil || IsNil(o.EnableMergeWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnableMergeWhitelist
}

// GetEnableMergeWhitelistOk returns a tuple with the EnableMergeWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetEnableMergeWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMergeWhitelist) {
		return nil, false
	}
	return o.EnableMergeWhitelist, true
}

// HasEnableMergeWhitelist returns a boolean if a field has been set.
func (o *BranchProtection) HasEnableMergeWhitelist() bool {
	if o != nil && !IsNil(o.EnableMergeWhitelist) {
		return true
	}

	return false
}

// SetEnableMergeWhitelist gets a reference to the given bool and assigns it to the EnableMergeWhitelist field.
func (o *BranchProtection) SetEnableMergeWhitelist(v bool) {
	o.EnableMergeWhitelist = &v
}

// GetEnablePush returns the EnablePush field value if set, zero value otherwise.
func (o *BranchProtection) GetEnablePush() bool {
	if o == nil || IsNil(o.EnablePush) {
		var ret bool
		return ret
	}
	return *o.EnablePush
}

// GetEnablePushOk returns a tuple with the EnablePush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetEnablePushOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePush) {
		return nil, false
	}
	return o.EnablePush, true
}

// HasEnablePush returns a boolean if a field has been set.
func (o *BranchProtection) HasEnablePush() bool {
	if o != nil && !IsNil(o.EnablePush) {
		return true
	}

	return false
}

// SetEnablePush gets a reference to the given bool and assigns it to the EnablePush field.
func (o *BranchProtection) SetEnablePush(v bool) {
	o.EnablePush = &v
}

// GetEnablePushWhitelist returns the EnablePushWhitelist field value if set, zero value otherwise.
func (o *BranchProtection) GetEnablePushWhitelist() bool {
	if o == nil || IsNil(o.EnablePushWhitelist) {
		var ret bool
		return ret
	}
	return *o.EnablePushWhitelist
}

// GetEnablePushWhitelistOk returns a tuple with the EnablePushWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetEnablePushWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePushWhitelist) {
		return nil, false
	}
	return o.EnablePushWhitelist, true
}

// HasEnablePushWhitelist returns a boolean if a field has been set.
func (o *BranchProtection) HasEnablePushWhitelist() bool {
	if o != nil && !IsNil(o.EnablePushWhitelist) {
		return true
	}

	return false
}

// SetEnablePushWhitelist gets a reference to the given bool and assigns it to the EnablePushWhitelist field.
func (o *BranchProtection) SetEnablePushWhitelist(v bool) {
	o.EnablePushWhitelist = &v
}

// GetEnableStatusCheck returns the EnableStatusCheck field value if set, zero value otherwise.
func (o *BranchProtection) GetEnableStatusCheck() bool {
	if o == nil || IsNil(o.EnableStatusCheck) {
		var ret bool
		return ret
	}
	return *o.EnableStatusCheck
}

// GetEnableStatusCheckOk returns a tuple with the EnableStatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetEnableStatusCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableStatusCheck) {
		return nil, false
	}
	return o.EnableStatusCheck, true
}

// HasEnableStatusCheck returns a boolean if a field has been set.
func (o *BranchProtection) HasEnableStatusCheck() bool {
	if o != nil && !IsNil(o.EnableStatusCheck) {
		return true
	}

	return false
}

// SetEnableStatusCheck gets a reference to the given bool and assigns it to the EnableStatusCheck field.
func (o *BranchProtection) SetEnableStatusCheck(v bool) {
	o.EnableStatusCheck = &v
}

// GetMergeWhitelistTeams returns the MergeWhitelistTeams field value if set, zero value otherwise.
func (o *BranchProtection) GetMergeWhitelistTeams() []string {
	if o == nil || IsNil(o.MergeWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.MergeWhitelistTeams
}

// GetMergeWhitelistTeamsOk returns a tuple with the MergeWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetMergeWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.MergeWhitelistTeams) {
		return nil, false
	}
	return o.MergeWhitelistTeams, true
}

// HasMergeWhitelistTeams returns a boolean if a field has been set.
func (o *BranchProtection) HasMergeWhitelistTeams() bool {
	if o != nil && !IsNil(o.MergeWhitelistTeams) {
		return true
	}

	return false
}

// SetMergeWhitelistTeams gets a reference to the given []string and assigns it to the MergeWhitelistTeams field.
func (o *BranchProtection) SetMergeWhitelistTeams(v []string) {
	o.MergeWhitelistTeams = v
}

// GetMergeWhitelistUsernames returns the MergeWhitelistUsernames field value if set, zero value otherwise.
func (o *BranchProtection) GetMergeWhitelistUsernames() []string {
	if o == nil || IsNil(o.MergeWhitelistUsernames) {
		var ret []string
		return ret
	}
	return o.MergeWhitelistUsernames
}

// GetMergeWhitelistUsernamesOk returns a tuple with the MergeWhitelistUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetMergeWhitelistUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.MergeWhitelistUsernames) {
		return nil, false
	}
	return o.MergeWhitelistUsernames, true
}

// HasMergeWhitelistUsernames returns a boolean if a field has been set.
func (o *BranchProtection) HasMergeWhitelistUsernames() bool {
	if o != nil && !IsNil(o.MergeWhitelistUsernames) {
		return true
	}

	return false
}

// SetMergeWhitelistUsernames gets a reference to the given []string and assigns it to the MergeWhitelistUsernames field.
func (o *BranchProtection) SetMergeWhitelistUsernames(v []string) {
	o.MergeWhitelistUsernames = v
}

// GetProtectedFilePatterns returns the ProtectedFilePatterns field value if set, zero value otherwise.
func (o *BranchProtection) GetProtectedFilePatterns() string {
	if o == nil || IsNil(o.ProtectedFilePatterns) {
		var ret string
		return ret
	}
	return *o.ProtectedFilePatterns
}

// GetProtectedFilePatternsOk returns a tuple with the ProtectedFilePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetProtectedFilePatternsOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectedFilePatterns) {
		return nil, false
	}
	return o.ProtectedFilePatterns, true
}

// HasProtectedFilePatterns returns a boolean if a field has been set.
func (o *BranchProtection) HasProtectedFilePatterns() bool {
	if o != nil && !IsNil(o.ProtectedFilePatterns) {
		return true
	}

	return false
}

// SetProtectedFilePatterns gets a reference to the given string and assigns it to the ProtectedFilePatterns field.
func (o *BranchProtection) SetProtectedFilePatterns(v string) {
	o.ProtectedFilePatterns = &v
}

// GetPushWhitelistDeployKeys returns the PushWhitelistDeployKeys field value if set, zero value otherwise.
func (o *BranchProtection) GetPushWhitelistDeployKeys() bool {
	if o == nil || IsNil(o.PushWhitelistDeployKeys) {
		var ret bool
		return ret
	}
	return *o.PushWhitelistDeployKeys
}

// GetPushWhitelistDeployKeysOk returns a tuple with the PushWhitelistDeployKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetPushWhitelistDeployKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.PushWhitelistDeployKeys) {
		return nil, false
	}
	return o.PushWhitelistDeployKeys, true
}

// HasPushWhitelistDeployKeys returns a boolean if a field has been set.
func (o *BranchProtection) HasPushWhitelistDeployKeys() bool {
	if o != nil && !IsNil(o.PushWhitelistDeployKeys) {
		return true
	}

	return false
}

// SetPushWhitelistDeployKeys gets a reference to the given bool and assigns it to the PushWhitelistDeployKeys field.
func (o *BranchProtection) SetPushWhitelistDeployKeys(v bool) {
	o.PushWhitelistDeployKeys = &v
}

// GetPushWhitelistTeams returns the PushWhitelistTeams field value if set, zero value otherwise.
func (o *BranchProtection) GetPushWhitelistTeams() []string {
	if o == nil || IsNil(o.PushWhitelistTeams) {
		var ret []string
		return ret
	}
	return o.PushWhitelistTeams
}

// GetPushWhitelistTeamsOk returns a tuple with the PushWhitelistTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetPushWhitelistTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.PushWhitelistTeams) {
		return nil, false
	}
	return o.PushWhitelistTeams, true
}

// HasPushWhitelistTeams returns a boolean if a field has been set.
func (o *BranchProtection) HasPushWhitelistTeams() bool {
	if o != nil && !IsNil(o.PushWhitelistTeams) {
		return true
	}

	return false
}

// SetPushWhitelistTeams gets a reference to the given []string and assigns it to the PushWhitelistTeams field.
func (o *BranchProtection) SetPushWhitelistTeams(v []string) {
	o.PushWhitelistTeams = v
}

// GetPushWhitelistUsernames returns the PushWhitelistUsernames field value if set, zero value otherwise.
func (o *BranchProtection) GetPushWhitelistUsernames() []string {
	if o == nil || IsNil(o.PushWhitelistUsernames) {
		var ret []string
		return ret
	}
	return o.PushWhitelistUsernames
}

// GetPushWhitelistUsernamesOk returns a tuple with the PushWhitelistUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetPushWhitelistUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.PushWhitelistUsernames) {
		return nil, false
	}
	return o.PushWhitelistUsernames, true
}

// HasPushWhitelistUsernames returns a boolean if a field has been set.
func (o *BranchProtection) HasPushWhitelistUsernames() bool {
	if o != nil && !IsNil(o.PushWhitelistUsernames) {
		return true
	}

	return false
}

// SetPushWhitelistUsernames gets a reference to the given []string and assigns it to the PushWhitelistUsernames field.
func (o *BranchProtection) SetPushWhitelistUsernames(v []string) {
	o.PushWhitelistUsernames = v
}

// GetRequireSignedCommits returns the RequireSignedCommits field value if set, zero value otherwise.
func (o *BranchProtection) GetRequireSignedCommits() bool {
	if o == nil || IsNil(o.RequireSignedCommits) {
		var ret bool
		return ret
	}
	return *o.RequireSignedCommits
}

// GetRequireSignedCommitsOk returns a tuple with the RequireSignedCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetRequireSignedCommitsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedCommits) {
		return nil, false
	}
	return o.RequireSignedCommits, true
}

// HasRequireSignedCommits returns a boolean if a field has been set.
func (o *BranchProtection) HasRequireSignedCommits() bool {
	if o != nil && !IsNil(o.RequireSignedCommits) {
		return true
	}

	return false
}

// SetRequireSignedCommits gets a reference to the given bool and assigns it to the RequireSignedCommits field.
func (o *BranchProtection) SetRequireSignedCommits(v bool) {
	o.RequireSignedCommits = &v
}

// GetRequiredApprovals returns the RequiredApprovals field value if set, zero value otherwise.
func (o *BranchProtection) GetRequiredApprovals() int64 {
	if o == nil || IsNil(o.RequiredApprovals) {
		var ret int64
		return ret
	}
	return *o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetRequiredApprovalsOk() (*int64, bool) {
	if o == nil || IsNil(o.RequiredApprovals) {
		return nil, false
	}
	return o.RequiredApprovals, true
}

// HasRequiredApprovals returns a boolean if a field has been set.
func (o *BranchProtection) HasRequiredApprovals() bool {
	if o != nil && !IsNil(o.RequiredApprovals) {
		return true
	}

	return false
}

// SetRequiredApprovals gets a reference to the given int64 and assigns it to the RequiredApprovals field.
func (o *BranchProtection) SetRequiredApprovals(v int64) {
	o.RequiredApprovals = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *BranchProtection) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *BranchProtection) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *BranchProtection) SetRuleName(v string) {
	o.RuleName = &v
}

// GetStatusCheckContexts returns the StatusCheckContexts field value if set, zero value otherwise.
func (o *BranchProtection) GetStatusCheckContexts() []string {
	if o == nil || IsNil(o.StatusCheckContexts) {
		var ret []string
		return ret
	}
	return o.StatusCheckContexts
}

// GetStatusCheckContextsOk returns a tuple with the StatusCheckContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetStatusCheckContextsOk() ([]string, bool) {
	if o == nil || IsNil(o.StatusCheckContexts) {
		return nil, false
	}
	return o.StatusCheckContexts, true
}

// HasStatusCheckContexts returns a boolean if a field has been set.
func (o *BranchProtection) HasStatusCheckContexts() bool {
	if o != nil && !IsNil(o.StatusCheckContexts) {
		return true
	}

	return false
}

// SetStatusCheckContexts gets a reference to the given []string and assigns it to the StatusCheckContexts field.
func (o *BranchProtection) SetStatusCheckContexts(v []string) {
	o.StatusCheckContexts = v
}

// GetUnprotectedFilePatterns returns the UnprotectedFilePatterns field value if set, zero value otherwise.
func (o *BranchProtection) GetUnprotectedFilePatterns() string {
	if o == nil || IsNil(o.UnprotectedFilePatterns) {
		var ret string
		return ret
	}
	return *o.UnprotectedFilePatterns
}

// GetUnprotectedFilePatternsOk returns a tuple with the UnprotectedFilePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetUnprotectedFilePatternsOk() (*string, bool) {
	if o == nil || IsNil(o.UnprotectedFilePatterns) {
		return nil, false
	}
	return o.UnprotectedFilePatterns, true
}

// HasUnprotectedFilePatterns returns a boolean if a field has been set.
func (o *BranchProtection) HasUnprotectedFilePatterns() bool {
	if o != nil && !IsNil(o.UnprotectedFilePatterns) {
		return true
	}

	return false
}

// SetUnprotectedFilePatterns gets a reference to the given string and assigns it to the UnprotectedFilePatterns field.
func (o *BranchProtection) SetUnprotectedFilePatterns(v string) {
	o.UnprotectedFilePatterns = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BranchProtection) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchProtection) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BranchProtection) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BranchProtection) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o BranchProtection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BranchProtection) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableBranchProtection struct {
	value *BranchProtection
	isSet bool
}

func (v NullableBranchProtection) Get() *BranchProtection {
	return v.value
}

func (v *NullableBranchProtection) Set(val *BranchProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchProtection(val *BranchProtection) *NullableBranchProtection {
	return &NullableBranchProtection{value: val, isSet: true}
}

func (v NullableBranchProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
