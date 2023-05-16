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

// checks if the EditRepoOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditRepoOption{}

// EditRepoOption EditRepoOption options when editing a repository's properties
type EditRepoOption struct {
	// either `true` to allow mark pr as merged manually, or `false` to prevent it.
	AllowManualMerge *bool `json:"allow_manual_merge,omitempty"`
	// either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
	AllowMergeCommits *bool `json:"allow_merge_commits,omitempty"`
	// either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
	AllowRebase *bool `json:"allow_rebase,omitempty"`
	// either `true` to allow rebase with explicit merge commits (--no-ff), or `false` to prevent rebase with explicit merge commits.
	AllowRebaseExplicit *bool `json:"allow_rebase_explicit,omitempty"`
	// either `true` to allow updating pull request branch by rebase, or `false` to prevent it.
	AllowRebaseUpdate *bool `json:"allow_rebase_update,omitempty"`
	// either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
	AllowSquashMerge *bool `json:"allow_squash_merge,omitempty"`
	// set to `true` to archive this repository.
	Archived *bool `json:"archived,omitempty"`
	// either `true` to enable AutodetectManualMerge, or `false` to prevent it. Note: In some special cases, misjudgments can occur.
	AutodetectManualMerge *bool `json:"autodetect_manual_merge,omitempty"`
	// set to `true` to allow edits from maintainers by default
	DefaultAllowMaintainerEdit *bool `json:"default_allow_maintainer_edit,omitempty"`
	// sets the default branch for this repository.
	DefaultBranch *string `json:"default_branch,omitempty"`
	// set to `true` to delete pr branch after merge by default
	DefaultDeleteBranchAfterMerge *bool `json:"default_delete_branch_after_merge,omitempty"`
	// set to a merge style to be used by this repository: \"merge\", \"rebase\", \"rebase-merge\", or \"squash\".
	DefaultMergeStyle *string `json:"default_merge_style,omitempty"`
	// a short description of the repository.
	Description *string `json:"description,omitempty"`
	// enable prune - remove obsolete remote-tracking references
	EnablePrune     *bool            `json:"enable_prune,omitempty"`
	ExternalTracker *ExternalTracker `json:"external_tracker,omitempty"`
	ExternalWiki    *ExternalWiki    `json:"external_wiki,omitempty"`
	// either `true` to enable issues for this repository or `false` to disable them.
	HasIssues *bool `json:"has_issues,omitempty"`
	// either `true` to enable project unit, or `false` to disable them.
	HasProjects *bool `json:"has_projects,omitempty"`
	// either `true` to allow pull requests, or `false` to prevent pull request.
	HasPullRequests *bool `json:"has_pull_requests,omitempty"`
	// either `true` to enable the wiki for this repository or `false` to disable it.
	HasWiki *bool `json:"has_wiki,omitempty"`
	// either `true` to ignore whitespace for conflicts, or `false` to not ignore whitespace.
	IgnoreWhitespaceConflicts *bool            `json:"ignore_whitespace_conflicts,omitempty"`
	InternalTracker           *InternalTracker `json:"internal_tracker,omitempty"`
	// set to a string like `8h30m0s` to set the mirror interval time
	MirrorInterval *string `json:"mirror_interval,omitempty"`
	// name of the repository
	Name *string `json:"name,omitempty"`
	// either `true` to make the repository private or `false` to make it public. Note: you will get a 422 error if the organization restricts changing repository visibility to organization owners and a non-owner tries to change the value of private.
	Private *bool `json:"private,omitempty"`
	// either `true` to make this repository a template or `false` to make it a normal repository
	Template *bool `json:"template,omitempty"`
	// a URL with more information about the repository.
	Website *string `json:"website,omitempty"`
}

// NewEditRepoOption instantiates a new EditRepoOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditRepoOption() *EditRepoOption {
	this := EditRepoOption{}
	return &this
}

// NewEditRepoOptionWithDefaults instantiates a new EditRepoOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditRepoOptionWithDefaults() *EditRepoOption {
	this := EditRepoOption{}
	return &this
}

// GetAllowManualMerge returns the AllowManualMerge field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowManualMerge() bool {
	if o == nil || IsNil(o.AllowManualMerge) {
		var ret bool
		return ret
	}
	return *o.AllowManualMerge
}

// GetAllowManualMergeOk returns a tuple with the AllowManualMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowManualMergeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowManualMerge) {
		return nil, false
	}
	return o.AllowManualMerge, true
}

// HasAllowManualMerge returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowManualMerge() bool {
	if o != nil && !IsNil(o.AllowManualMerge) {
		return true
	}

	return false
}

// SetAllowManualMerge gets a reference to the given bool and assigns it to the AllowManualMerge field.
func (o *EditRepoOption) SetAllowManualMerge(v bool) {
	o.AllowManualMerge = &v
}

// GetAllowMergeCommits returns the AllowMergeCommits field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowMergeCommits() bool {
	if o == nil || IsNil(o.AllowMergeCommits) {
		var ret bool
		return ret
	}
	return *o.AllowMergeCommits
}

// GetAllowMergeCommitsOk returns a tuple with the AllowMergeCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowMergeCommitsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMergeCommits) {
		return nil, false
	}
	return o.AllowMergeCommits, true
}

// HasAllowMergeCommits returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowMergeCommits() bool {
	if o != nil && !IsNil(o.AllowMergeCommits) {
		return true
	}

	return false
}

// SetAllowMergeCommits gets a reference to the given bool and assigns it to the AllowMergeCommits field.
func (o *EditRepoOption) SetAllowMergeCommits(v bool) {
	o.AllowMergeCommits = &v
}

// GetAllowRebase returns the AllowRebase field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowRebase() bool {
	if o == nil || IsNil(o.AllowRebase) {
		var ret bool
		return ret
	}
	return *o.AllowRebase
}

// GetAllowRebaseOk returns a tuple with the AllowRebase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowRebaseOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRebase) {
		return nil, false
	}
	return o.AllowRebase, true
}

// HasAllowRebase returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowRebase() bool {
	if o != nil && !IsNil(o.AllowRebase) {
		return true
	}

	return false
}

// SetAllowRebase gets a reference to the given bool and assigns it to the AllowRebase field.
func (o *EditRepoOption) SetAllowRebase(v bool) {
	o.AllowRebase = &v
}

// GetAllowRebaseExplicit returns the AllowRebaseExplicit field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowRebaseExplicit() bool {
	if o == nil || IsNil(o.AllowRebaseExplicit) {
		var ret bool
		return ret
	}
	return *o.AllowRebaseExplicit
}

// GetAllowRebaseExplicitOk returns a tuple with the AllowRebaseExplicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowRebaseExplicitOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRebaseExplicit) {
		return nil, false
	}
	return o.AllowRebaseExplicit, true
}

// HasAllowRebaseExplicit returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowRebaseExplicit() bool {
	if o != nil && !IsNil(o.AllowRebaseExplicit) {
		return true
	}

	return false
}

// SetAllowRebaseExplicit gets a reference to the given bool and assigns it to the AllowRebaseExplicit field.
func (o *EditRepoOption) SetAllowRebaseExplicit(v bool) {
	o.AllowRebaseExplicit = &v
}

// GetAllowRebaseUpdate returns the AllowRebaseUpdate field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowRebaseUpdate() bool {
	if o == nil || IsNil(o.AllowRebaseUpdate) {
		var ret bool
		return ret
	}
	return *o.AllowRebaseUpdate
}

// GetAllowRebaseUpdateOk returns a tuple with the AllowRebaseUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowRebaseUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRebaseUpdate) {
		return nil, false
	}
	return o.AllowRebaseUpdate, true
}

// HasAllowRebaseUpdate returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowRebaseUpdate() bool {
	if o != nil && !IsNil(o.AllowRebaseUpdate) {
		return true
	}

	return false
}

// SetAllowRebaseUpdate gets a reference to the given bool and assigns it to the AllowRebaseUpdate field.
func (o *EditRepoOption) SetAllowRebaseUpdate(v bool) {
	o.AllowRebaseUpdate = &v
}

// GetAllowSquashMerge returns the AllowSquashMerge field value if set, zero value otherwise.
func (o *EditRepoOption) GetAllowSquashMerge() bool {
	if o == nil || IsNil(o.AllowSquashMerge) {
		var ret bool
		return ret
	}
	return *o.AllowSquashMerge
}

// GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAllowSquashMergeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSquashMerge) {
		return nil, false
	}
	return o.AllowSquashMerge, true
}

// HasAllowSquashMerge returns a boolean if a field has been set.
func (o *EditRepoOption) HasAllowSquashMerge() bool {
	if o != nil && !IsNil(o.AllowSquashMerge) {
		return true
	}

	return false
}

// SetAllowSquashMerge gets a reference to the given bool and assigns it to the AllowSquashMerge field.
func (o *EditRepoOption) SetAllowSquashMerge(v bool) {
	o.AllowSquashMerge = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *EditRepoOption) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *EditRepoOption) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *EditRepoOption) SetArchived(v bool) {
	o.Archived = &v
}

// GetAutodetectManualMerge returns the AutodetectManualMerge field value if set, zero value otherwise.
func (o *EditRepoOption) GetAutodetectManualMerge() bool {
	if o == nil || IsNil(o.AutodetectManualMerge) {
		var ret bool
		return ret
	}
	return *o.AutodetectManualMerge
}

// GetAutodetectManualMergeOk returns a tuple with the AutodetectManualMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetAutodetectManualMergeOk() (*bool, bool) {
	if o == nil || IsNil(o.AutodetectManualMerge) {
		return nil, false
	}
	return o.AutodetectManualMerge, true
}

// HasAutodetectManualMerge returns a boolean if a field has been set.
func (o *EditRepoOption) HasAutodetectManualMerge() bool {
	if o != nil && !IsNil(o.AutodetectManualMerge) {
		return true
	}

	return false
}

// SetAutodetectManualMerge gets a reference to the given bool and assigns it to the AutodetectManualMerge field.
func (o *EditRepoOption) SetAutodetectManualMerge(v bool) {
	o.AutodetectManualMerge = &v
}

// GetDefaultAllowMaintainerEdit returns the DefaultAllowMaintainerEdit field value if set, zero value otherwise.
func (o *EditRepoOption) GetDefaultAllowMaintainerEdit() bool {
	if o == nil || IsNil(o.DefaultAllowMaintainerEdit) {
		var ret bool
		return ret
	}
	return *o.DefaultAllowMaintainerEdit
}

// GetDefaultAllowMaintainerEditOk returns a tuple with the DefaultAllowMaintainerEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetDefaultAllowMaintainerEditOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultAllowMaintainerEdit) {
		return nil, false
	}
	return o.DefaultAllowMaintainerEdit, true
}

// HasDefaultAllowMaintainerEdit returns a boolean if a field has been set.
func (o *EditRepoOption) HasDefaultAllowMaintainerEdit() bool {
	if o != nil && !IsNil(o.DefaultAllowMaintainerEdit) {
		return true
	}

	return false
}

// SetDefaultAllowMaintainerEdit gets a reference to the given bool and assigns it to the DefaultAllowMaintainerEdit field.
func (o *EditRepoOption) SetDefaultAllowMaintainerEdit(v bool) {
	o.DefaultAllowMaintainerEdit = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *EditRepoOption) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *EditRepoOption) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *EditRepoOption) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetDefaultDeleteBranchAfterMerge returns the DefaultDeleteBranchAfterMerge field value if set, zero value otherwise.
func (o *EditRepoOption) GetDefaultDeleteBranchAfterMerge() bool {
	if o == nil || IsNil(o.DefaultDeleteBranchAfterMerge) {
		var ret bool
		return ret
	}
	return *o.DefaultDeleteBranchAfterMerge
}

// GetDefaultDeleteBranchAfterMergeOk returns a tuple with the DefaultDeleteBranchAfterMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetDefaultDeleteBranchAfterMergeOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultDeleteBranchAfterMerge) {
		return nil, false
	}
	return o.DefaultDeleteBranchAfterMerge, true
}

// HasDefaultDeleteBranchAfterMerge returns a boolean if a field has been set.
func (o *EditRepoOption) HasDefaultDeleteBranchAfterMerge() bool {
	if o != nil && !IsNil(o.DefaultDeleteBranchAfterMerge) {
		return true
	}

	return false
}

// SetDefaultDeleteBranchAfterMerge gets a reference to the given bool and assigns it to the DefaultDeleteBranchAfterMerge field.
func (o *EditRepoOption) SetDefaultDeleteBranchAfterMerge(v bool) {
	o.DefaultDeleteBranchAfterMerge = &v
}

// GetDefaultMergeStyle returns the DefaultMergeStyle field value if set, zero value otherwise.
func (o *EditRepoOption) GetDefaultMergeStyle() string {
	if o == nil || IsNil(o.DefaultMergeStyle) {
		var ret string
		return ret
	}
	return *o.DefaultMergeStyle
}

// GetDefaultMergeStyleOk returns a tuple with the DefaultMergeStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetDefaultMergeStyleOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultMergeStyle) {
		return nil, false
	}
	return o.DefaultMergeStyle, true
}

// HasDefaultMergeStyle returns a boolean if a field has been set.
func (o *EditRepoOption) HasDefaultMergeStyle() bool {
	if o != nil && !IsNil(o.DefaultMergeStyle) {
		return true
	}

	return false
}

// SetDefaultMergeStyle gets a reference to the given string and assigns it to the DefaultMergeStyle field.
func (o *EditRepoOption) SetDefaultMergeStyle(v string) {
	o.DefaultMergeStyle = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditRepoOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditRepoOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditRepoOption) SetDescription(v string) {
	o.Description = &v
}

// GetEnablePrune returns the EnablePrune field value if set, zero value otherwise.
func (o *EditRepoOption) GetEnablePrune() bool {
	if o == nil || IsNil(o.EnablePrune) {
		var ret bool
		return ret
	}
	return *o.EnablePrune
}

// GetEnablePruneOk returns a tuple with the EnablePrune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetEnablePruneOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePrune) {
		return nil, false
	}
	return o.EnablePrune, true
}

// HasEnablePrune returns a boolean if a field has been set.
func (o *EditRepoOption) HasEnablePrune() bool {
	if o != nil && !IsNil(o.EnablePrune) {
		return true
	}

	return false
}

// SetEnablePrune gets a reference to the given bool and assigns it to the EnablePrune field.
func (o *EditRepoOption) SetEnablePrune(v bool) {
	o.EnablePrune = &v
}

// GetExternalTracker returns the ExternalTracker field value if set, zero value otherwise.
func (o *EditRepoOption) GetExternalTracker() ExternalTracker {
	if o == nil || IsNil(o.ExternalTracker) {
		var ret ExternalTracker
		return ret
	}
	return *o.ExternalTracker
}

// GetExternalTrackerOk returns a tuple with the ExternalTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetExternalTrackerOk() (*ExternalTracker, bool) {
	if o == nil || IsNil(o.ExternalTracker) {
		return nil, false
	}
	return o.ExternalTracker, true
}

// HasExternalTracker returns a boolean if a field has been set.
func (o *EditRepoOption) HasExternalTracker() bool {
	if o != nil && !IsNil(o.ExternalTracker) {
		return true
	}

	return false
}

// SetExternalTracker gets a reference to the given ExternalTracker and assigns it to the ExternalTracker field.
func (o *EditRepoOption) SetExternalTracker(v ExternalTracker) {
	o.ExternalTracker = &v
}

// GetExternalWiki returns the ExternalWiki field value if set, zero value otherwise.
func (o *EditRepoOption) GetExternalWiki() ExternalWiki {
	if o == nil || IsNil(o.ExternalWiki) {
		var ret ExternalWiki
		return ret
	}
	return *o.ExternalWiki
}

// GetExternalWikiOk returns a tuple with the ExternalWiki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetExternalWikiOk() (*ExternalWiki, bool) {
	if o == nil || IsNil(o.ExternalWiki) {
		return nil, false
	}
	return o.ExternalWiki, true
}

// HasExternalWiki returns a boolean if a field has been set.
func (o *EditRepoOption) HasExternalWiki() bool {
	if o != nil && !IsNil(o.ExternalWiki) {
		return true
	}

	return false
}

// SetExternalWiki gets a reference to the given ExternalWiki and assigns it to the ExternalWiki field.
func (o *EditRepoOption) SetExternalWiki(v ExternalWiki) {
	o.ExternalWiki = &v
}

// GetHasIssues returns the HasIssues field value if set, zero value otherwise.
func (o *EditRepoOption) GetHasIssues() bool {
	if o == nil || IsNil(o.HasIssues) {
		var ret bool
		return ret
	}
	return *o.HasIssues
}

// GetHasIssuesOk returns a tuple with the HasIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetHasIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasIssues) {
		return nil, false
	}
	return o.HasIssues, true
}

// HasHasIssues returns a boolean if a field has been set.
func (o *EditRepoOption) HasHasIssues() bool {
	if o != nil && !IsNil(o.HasIssues) {
		return true
	}

	return false
}

// SetHasIssues gets a reference to the given bool and assigns it to the HasIssues field.
func (o *EditRepoOption) SetHasIssues(v bool) {
	o.HasIssues = &v
}

// GetHasProjects returns the HasProjects field value if set, zero value otherwise.
func (o *EditRepoOption) GetHasProjects() bool {
	if o == nil || IsNil(o.HasProjects) {
		var ret bool
		return ret
	}
	return *o.HasProjects
}

// GetHasProjectsOk returns a tuple with the HasProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetHasProjectsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasProjects) {
		return nil, false
	}
	return o.HasProjects, true
}

// HasHasProjects returns a boolean if a field has been set.
func (o *EditRepoOption) HasHasProjects() bool {
	if o != nil && !IsNil(o.HasProjects) {
		return true
	}

	return false
}

// SetHasProjects gets a reference to the given bool and assigns it to the HasProjects field.
func (o *EditRepoOption) SetHasProjects(v bool) {
	o.HasProjects = &v
}

// GetHasPullRequests returns the HasPullRequests field value if set, zero value otherwise.
func (o *EditRepoOption) GetHasPullRequests() bool {
	if o == nil || IsNil(o.HasPullRequests) {
		var ret bool
		return ret
	}
	return *o.HasPullRequests
}

// GetHasPullRequestsOk returns a tuple with the HasPullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetHasPullRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPullRequests) {
		return nil, false
	}
	return o.HasPullRequests, true
}

// HasHasPullRequests returns a boolean if a field has been set.
func (o *EditRepoOption) HasHasPullRequests() bool {
	if o != nil && !IsNil(o.HasPullRequests) {
		return true
	}

	return false
}

// SetHasPullRequests gets a reference to the given bool and assigns it to the HasPullRequests field.
func (o *EditRepoOption) SetHasPullRequests(v bool) {
	o.HasPullRequests = &v
}

// GetHasWiki returns the HasWiki field value if set, zero value otherwise.
func (o *EditRepoOption) GetHasWiki() bool {
	if o == nil || IsNil(o.HasWiki) {
		var ret bool
		return ret
	}
	return *o.HasWiki
}

// GetHasWikiOk returns a tuple with the HasWiki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetHasWikiOk() (*bool, bool) {
	if o == nil || IsNil(o.HasWiki) {
		return nil, false
	}
	return o.HasWiki, true
}

// HasHasWiki returns a boolean if a field has been set.
func (o *EditRepoOption) HasHasWiki() bool {
	if o != nil && !IsNil(o.HasWiki) {
		return true
	}

	return false
}

// SetHasWiki gets a reference to the given bool and assigns it to the HasWiki field.
func (o *EditRepoOption) SetHasWiki(v bool) {
	o.HasWiki = &v
}

// GetIgnoreWhitespaceConflicts returns the IgnoreWhitespaceConflicts field value if set, zero value otherwise.
func (o *EditRepoOption) GetIgnoreWhitespaceConflicts() bool {
	if o == nil || IsNil(o.IgnoreWhitespaceConflicts) {
		var ret bool
		return ret
	}
	return *o.IgnoreWhitespaceConflicts
}

// GetIgnoreWhitespaceConflictsOk returns a tuple with the IgnoreWhitespaceConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetIgnoreWhitespaceConflictsOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreWhitespaceConflicts) {
		return nil, false
	}
	return o.IgnoreWhitespaceConflicts, true
}

// HasIgnoreWhitespaceConflicts returns a boolean if a field has been set.
func (o *EditRepoOption) HasIgnoreWhitespaceConflicts() bool {
	if o != nil && !IsNil(o.IgnoreWhitespaceConflicts) {
		return true
	}

	return false
}

// SetIgnoreWhitespaceConflicts gets a reference to the given bool and assigns it to the IgnoreWhitespaceConflicts field.
func (o *EditRepoOption) SetIgnoreWhitespaceConflicts(v bool) {
	o.IgnoreWhitespaceConflicts = &v
}

// GetInternalTracker returns the InternalTracker field value if set, zero value otherwise.
func (o *EditRepoOption) GetInternalTracker() InternalTracker {
	if o == nil || IsNil(o.InternalTracker) {
		var ret InternalTracker
		return ret
	}
	return *o.InternalTracker
}

// GetInternalTrackerOk returns a tuple with the InternalTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetInternalTrackerOk() (*InternalTracker, bool) {
	if o == nil || IsNil(o.InternalTracker) {
		return nil, false
	}
	return o.InternalTracker, true
}

// HasInternalTracker returns a boolean if a field has been set.
func (o *EditRepoOption) HasInternalTracker() bool {
	if o != nil && !IsNil(o.InternalTracker) {
		return true
	}

	return false
}

// SetInternalTracker gets a reference to the given InternalTracker and assigns it to the InternalTracker field.
func (o *EditRepoOption) SetInternalTracker(v InternalTracker) {
	o.InternalTracker = &v
}

// GetMirrorInterval returns the MirrorInterval field value if set, zero value otherwise.
func (o *EditRepoOption) GetMirrorInterval() string {
	if o == nil || IsNil(o.MirrorInterval) {
		var ret string
		return ret
	}
	return *o.MirrorInterval
}

// GetMirrorIntervalOk returns a tuple with the MirrorInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetMirrorIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MirrorInterval) {
		return nil, false
	}
	return o.MirrorInterval, true
}

// HasMirrorInterval returns a boolean if a field has been set.
func (o *EditRepoOption) HasMirrorInterval() bool {
	if o != nil && !IsNil(o.MirrorInterval) {
		return true
	}

	return false
}

// SetMirrorInterval gets a reference to the given string and assigns it to the MirrorInterval field.
func (o *EditRepoOption) SetMirrorInterval(v string) {
	o.MirrorInterval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditRepoOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditRepoOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditRepoOption) SetName(v string) {
	o.Name = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *EditRepoOption) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *EditRepoOption) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *EditRepoOption) SetPrivate(v bool) {
	o.Private = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EditRepoOption) GetTemplate() bool {
	if o == nil || IsNil(o.Template) {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EditRepoOption) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *EditRepoOption) SetTemplate(v bool) {
	o.Template = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *EditRepoOption) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditRepoOption) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *EditRepoOption) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *EditRepoOption) SetWebsite(v string) {
	o.Website = &v
}

func (o EditRepoOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditRepoOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowManualMerge) {
		toSerialize["allow_manual_merge"] = o.AllowManualMerge
	}
	if !IsNil(o.AllowMergeCommits) {
		toSerialize["allow_merge_commits"] = o.AllowMergeCommits
	}
	if !IsNil(o.AllowRebase) {
		toSerialize["allow_rebase"] = o.AllowRebase
	}
	if !IsNil(o.AllowRebaseExplicit) {
		toSerialize["allow_rebase_explicit"] = o.AllowRebaseExplicit
	}
	if !IsNil(o.AllowRebaseUpdate) {
		toSerialize["allow_rebase_update"] = o.AllowRebaseUpdate
	}
	if !IsNil(o.AllowSquashMerge) {
		toSerialize["allow_squash_merge"] = o.AllowSquashMerge
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.AutodetectManualMerge) {
		toSerialize["autodetect_manual_merge"] = o.AutodetectManualMerge
	}
	if !IsNil(o.DefaultAllowMaintainerEdit) {
		toSerialize["default_allow_maintainer_edit"] = o.DefaultAllowMaintainerEdit
	}
	if !IsNil(o.DefaultBranch) {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if !IsNil(o.DefaultDeleteBranchAfterMerge) {
		toSerialize["default_delete_branch_after_merge"] = o.DefaultDeleteBranchAfterMerge
	}
	if !IsNil(o.DefaultMergeStyle) {
		toSerialize["default_merge_style"] = o.DefaultMergeStyle
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EnablePrune) {
		toSerialize["enable_prune"] = o.EnablePrune
	}
	if !IsNil(o.ExternalTracker) {
		toSerialize["external_tracker"] = o.ExternalTracker
	}
	if !IsNil(o.ExternalWiki) {
		toSerialize["external_wiki"] = o.ExternalWiki
	}
	if !IsNil(o.HasIssues) {
		toSerialize["has_issues"] = o.HasIssues
	}
	if !IsNil(o.HasProjects) {
		toSerialize["has_projects"] = o.HasProjects
	}
	if !IsNil(o.HasPullRequests) {
		toSerialize["has_pull_requests"] = o.HasPullRequests
	}
	if !IsNil(o.HasWiki) {
		toSerialize["has_wiki"] = o.HasWiki
	}
	if !IsNil(o.IgnoreWhitespaceConflicts) {
		toSerialize["ignore_whitespace_conflicts"] = o.IgnoreWhitespaceConflicts
	}
	if !IsNil(o.InternalTracker) {
		toSerialize["internal_tracker"] = o.InternalTracker
	}
	if !IsNil(o.MirrorInterval) {
		toSerialize["mirror_interval"] = o.MirrorInterval
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableEditRepoOption struct {
	value *EditRepoOption
	isSet bool
}

func (v NullableEditRepoOption) Get() *EditRepoOption {
	return v.value
}

func (v *NullableEditRepoOption) Set(val *EditRepoOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEditRepoOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEditRepoOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditRepoOption(val *EditRepoOption) *NullableEditRepoOption {
	return &NullableEditRepoOption{value: val, isSet: true}
}

func (v NullableEditRepoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditRepoOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
