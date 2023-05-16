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

// checks if the PullRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PullRequest{}

// PullRequest PullRequest represents a pull request
type PullRequest struct {
	AllowMaintainerEdit *bool         `json:"allow_maintainer_edit,omitempty"`
	Assignee            *User         `json:"assignee,omitempty"`
	Assignees           []User        `json:"assignees,omitempty"`
	Base                *PRBranchInfo `json:"base,omitempty"`
	Body                *string       `json:"body,omitempty"`
	ClosedAt            *time.Time    `json:"closed_at,omitempty"`
	Comments            *int64        `json:"comments,omitempty"`
	CreatedAt           *time.Time    `json:"created_at,omitempty"`
	DiffUrl             *string       `json:"diff_url,omitempty"`
	DueDate             *time.Time    `json:"due_date,omitempty"`
	Head                *PRBranchInfo `json:"head,omitempty"`
	HtmlUrl             *string       `json:"html_url,omitempty"`
	Id                  *int64        `json:"id,omitempty"`
	IsLocked            *bool         `json:"is_locked,omitempty"`
	Labels              []Label       `json:"labels,omitempty"`
	MergeBase           *string       `json:"merge_base,omitempty"`
	MergeCommitSha      *string       `json:"merge_commit_sha,omitempty"`
	Mergeable           *bool         `json:"mergeable,omitempty"`
	Merged              *bool         `json:"merged,omitempty"`
	MergedAt            *time.Time    `json:"merged_at,omitempty"`
	MergedBy            *User         `json:"merged_by,omitempty"`
	Milestone           *Milestone    `json:"milestone,omitempty"`
	Number              *int64        `json:"number,omitempty"`
	PatchUrl            *string       `json:"patch_url,omitempty"`
	// StateType issue state type
	State     *string    `json:"state,omitempty"`
	Title     *string    `json:"title,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Url       *string    `json:"url,omitempty"`
	User      *User      `json:"user,omitempty"`
}

// NewPullRequest instantiates a new PullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequest() *PullRequest {
	this := PullRequest{}
	return &this
}

// NewPullRequestWithDefaults instantiates a new PullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestWithDefaults() *PullRequest {
	this := PullRequest{}
	return &this
}

// GetAllowMaintainerEdit returns the AllowMaintainerEdit field value if set, zero value otherwise.
func (o *PullRequest) GetAllowMaintainerEdit() bool {
	if o == nil || IsNil(o.AllowMaintainerEdit) {
		var ret bool
		return ret
	}
	return *o.AllowMaintainerEdit
}

// GetAllowMaintainerEditOk returns a tuple with the AllowMaintainerEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetAllowMaintainerEditOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMaintainerEdit) {
		return nil, false
	}
	return o.AllowMaintainerEdit, true
}

// HasAllowMaintainerEdit returns a boolean if a field has been set.
func (o *PullRequest) HasAllowMaintainerEdit() bool {
	if o != nil && !IsNil(o.AllowMaintainerEdit) {
		return true
	}

	return false
}

// SetAllowMaintainerEdit gets a reference to the given bool and assigns it to the AllowMaintainerEdit field.
func (o *PullRequest) SetAllowMaintainerEdit(v bool) {
	o.AllowMaintainerEdit = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *PullRequest) GetAssignee() User {
	if o == nil || IsNil(o.Assignee) {
		var ret User
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetAssigneeOk() (*User, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *PullRequest) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given User and assigns it to the Assignee field.
func (o *PullRequest) SetAssignee(v User) {
	o.Assignee = &v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *PullRequest) GetAssignees() []User {
	if o == nil || IsNil(o.Assignees) {
		var ret []User
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetAssigneesOk() ([]User, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *PullRequest) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []User and assigns it to the Assignees field.
func (o *PullRequest) SetAssignees(v []User) {
	o.Assignees = v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *PullRequest) GetBase() PRBranchInfo {
	if o == nil || IsNil(o.Base) {
		var ret PRBranchInfo
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBaseOk() (*PRBranchInfo, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *PullRequest) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given PRBranchInfo and assigns it to the Base field.
func (o *PullRequest) SetBase(v PRBranchInfo) {
	o.Base = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PullRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PullRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PullRequest) SetBody(v string) {
	o.Body = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *PullRequest) GetClosedAt() time.Time {
	if o == nil || IsNil(o.ClosedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedAt) {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *PullRequest) HasClosedAt() bool {
	if o != nil && !IsNil(o.ClosedAt) {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *PullRequest) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PullRequest) GetComments() int64 {
	if o == nil || IsNil(o.Comments) {
		var ret int64
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetCommentsOk() (*int64, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PullRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given int64 and assigns it to the Comments field.
func (o *PullRequest) SetComments(v int64) {
	o.Comments = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PullRequest) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PullRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PullRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDiffUrl returns the DiffUrl field value if set, zero value otherwise.
func (o *PullRequest) GetDiffUrl() string {
	if o == nil || IsNil(o.DiffUrl) {
		var ret string
		return ret
	}
	return *o.DiffUrl
}

// GetDiffUrlOk returns a tuple with the DiffUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetDiffUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DiffUrl) {
		return nil, false
	}
	return o.DiffUrl, true
}

// HasDiffUrl returns a boolean if a field has been set.
func (o *PullRequest) HasDiffUrl() bool {
	if o != nil && !IsNil(o.DiffUrl) {
		return true
	}

	return false
}

// SetDiffUrl gets a reference to the given string and assigns it to the DiffUrl field.
func (o *PullRequest) SetDiffUrl(v string) {
	o.DiffUrl = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *PullRequest) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *PullRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *PullRequest) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetHead returns the Head field value if set, zero value otherwise.
func (o *PullRequest) GetHead() PRBranchInfo {
	if o == nil || IsNil(o.Head) {
		var ret PRBranchInfo
		return ret
	}
	return *o.Head
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetHeadOk() (*PRBranchInfo, bool) {
	if o == nil || IsNil(o.Head) {
		return nil, false
	}
	return o.Head, true
}

// HasHead returns a boolean if a field has been set.
func (o *PullRequest) HasHead() bool {
	if o != nil && !IsNil(o.Head) {
		return true
	}

	return false
}

// SetHead gets a reference to the given PRBranchInfo and assigns it to the Head field.
func (o *PullRequest) SetHead(v PRBranchInfo) {
	o.Head = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *PullRequest) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *PullRequest) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *PullRequest) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PullRequest) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PullRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PullRequest) SetId(v int64) {
	o.Id = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *PullRequest) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *PullRequest) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *PullRequest) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PullRequest) GetLabels() []Label {
	if o == nil || IsNil(o.Labels) {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetLabelsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PullRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *PullRequest) SetLabels(v []Label) {
	o.Labels = v
}

// GetMergeBase returns the MergeBase field value if set, zero value otherwise.
func (o *PullRequest) GetMergeBase() string {
	if o == nil || IsNil(o.MergeBase) {
		var ret string
		return ret
	}
	return *o.MergeBase
}

// GetMergeBaseOk returns a tuple with the MergeBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergeBaseOk() (*string, bool) {
	if o == nil || IsNil(o.MergeBase) {
		return nil, false
	}
	return o.MergeBase, true
}

// HasMergeBase returns a boolean if a field has been set.
func (o *PullRequest) HasMergeBase() bool {
	if o != nil && !IsNil(o.MergeBase) {
		return true
	}

	return false
}

// SetMergeBase gets a reference to the given string and assigns it to the MergeBase field.
func (o *PullRequest) SetMergeBase(v string) {
	o.MergeBase = &v
}

// GetMergeCommitSha returns the MergeCommitSha field value if set, zero value otherwise.
func (o *PullRequest) GetMergeCommitSha() string {
	if o == nil || IsNil(o.MergeCommitSha) {
		var ret string
		return ret
	}
	return *o.MergeCommitSha
}

// GetMergeCommitShaOk returns a tuple with the MergeCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergeCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.MergeCommitSha) {
		return nil, false
	}
	return o.MergeCommitSha, true
}

// HasMergeCommitSha returns a boolean if a field has been set.
func (o *PullRequest) HasMergeCommitSha() bool {
	if o != nil && !IsNil(o.MergeCommitSha) {
		return true
	}

	return false
}

// SetMergeCommitSha gets a reference to the given string and assigns it to the MergeCommitSha field.
func (o *PullRequest) SetMergeCommitSha(v string) {
	o.MergeCommitSha = &v
}

// GetMergeable returns the Mergeable field value if set, zero value otherwise.
func (o *PullRequest) GetMergeable() bool {
	if o == nil || IsNil(o.Mergeable) {
		var ret bool
		return ret
	}
	return *o.Mergeable
}

// GetMergeableOk returns a tuple with the Mergeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergeableOk() (*bool, bool) {
	if o == nil || IsNil(o.Mergeable) {
		return nil, false
	}
	return o.Mergeable, true
}

// HasMergeable returns a boolean if a field has been set.
func (o *PullRequest) HasMergeable() bool {
	if o != nil && !IsNil(o.Mergeable) {
		return true
	}

	return false
}

// SetMergeable gets a reference to the given bool and assigns it to the Mergeable field.
func (o *PullRequest) SetMergeable(v bool) {
	o.Mergeable = &v
}

// GetMerged returns the Merged field value if set, zero value otherwise.
func (o *PullRequest) GetMerged() bool {
	if o == nil || IsNil(o.Merged) {
		var ret bool
		return ret
	}
	return *o.Merged
}

// GetMergedOk returns a tuple with the Merged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergedOk() (*bool, bool) {
	if o == nil || IsNil(o.Merged) {
		return nil, false
	}
	return o.Merged, true
}

// HasMerged returns a boolean if a field has been set.
func (o *PullRequest) HasMerged() bool {
	if o != nil && !IsNil(o.Merged) {
		return true
	}

	return false
}

// SetMerged gets a reference to the given bool and assigns it to the Merged field.
func (o *PullRequest) SetMerged(v bool) {
	o.Merged = &v
}

// GetMergedAt returns the MergedAt field value if set, zero value otherwise.
func (o *PullRequest) GetMergedAt() time.Time {
	if o == nil || IsNil(o.MergedAt) {
		var ret time.Time
		return ret
	}
	return *o.MergedAt
}

// GetMergedAtOk returns a tuple with the MergedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MergedAt) {
		return nil, false
	}
	return o.MergedAt, true
}

// HasMergedAt returns a boolean if a field has been set.
func (o *PullRequest) HasMergedAt() bool {
	if o != nil && !IsNil(o.MergedAt) {
		return true
	}

	return false
}

// SetMergedAt gets a reference to the given time.Time and assigns it to the MergedAt field.
func (o *PullRequest) SetMergedAt(v time.Time) {
	o.MergedAt = &v
}

// GetMergedBy returns the MergedBy field value if set, zero value otherwise.
func (o *PullRequest) GetMergedBy() User {
	if o == nil || IsNil(o.MergedBy) {
		var ret User
		return ret
	}
	return *o.MergedBy
}

// GetMergedByOk returns a tuple with the MergedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergedByOk() (*User, bool) {
	if o == nil || IsNil(o.MergedBy) {
		return nil, false
	}
	return o.MergedBy, true
}

// HasMergedBy returns a boolean if a field has been set.
func (o *PullRequest) HasMergedBy() bool {
	if o != nil && !IsNil(o.MergedBy) {
		return true
	}

	return false
}

// SetMergedBy gets a reference to the given User and assigns it to the MergedBy field.
func (o *PullRequest) SetMergedBy(v User) {
	o.MergedBy = &v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *PullRequest) GetMilestone() Milestone {
	if o == nil || IsNil(o.Milestone) {
		var ret Milestone
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMilestoneOk() (*Milestone, bool) {
	if o == nil || IsNil(o.Milestone) {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *PullRequest) HasMilestone() bool {
	if o != nil && !IsNil(o.Milestone) {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given Milestone and assigns it to the Milestone field.
func (o *PullRequest) SetMilestone(v Milestone) {
	o.Milestone = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PullRequest) GetNumber() int64 {
	if o == nil || IsNil(o.Number) {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PullRequest) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *PullRequest) SetNumber(v int64) {
	o.Number = &v
}

// GetPatchUrl returns the PatchUrl field value if set, zero value otherwise.
func (o *PullRequest) GetPatchUrl() string {
	if o == nil || IsNil(o.PatchUrl) {
		var ret string
		return ret
	}
	return *o.PatchUrl
}

// GetPatchUrlOk returns a tuple with the PatchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetPatchUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PatchUrl) {
		return nil, false
	}
	return o.PatchUrl, true
}

// HasPatchUrl returns a boolean if a field has been set.
func (o *PullRequest) HasPatchUrl() bool {
	if o != nil && !IsNil(o.PatchUrl) {
		return true
	}

	return false
}

// SetPatchUrl gets a reference to the given string and assigns it to the PatchUrl field.
func (o *PullRequest) SetPatchUrl(v string) {
	o.PatchUrl = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PullRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PullRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PullRequest) SetState(v string) {
	o.State = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PullRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PullRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PullRequest) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PullRequest) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PullRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PullRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PullRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PullRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PullRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PullRequest) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PullRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *PullRequest) SetUser(v User) {
	o.User = &v
}

func (o PullRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PullRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowMaintainerEdit) {
		toSerialize["allow_maintainer_edit"] = o.AllowMaintainerEdit
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.ClosedAt) {
		toSerialize["closed_at"] = o.ClosedAt
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DiffUrl) {
		toSerialize["diff_url"] = o.DiffUrl
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.Head) {
		toSerialize["head"] = o.Head
	}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsLocked) {
		toSerialize["is_locked"] = o.IsLocked
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.MergeBase) {
		toSerialize["merge_base"] = o.MergeBase
	}
	if !IsNil(o.MergeCommitSha) {
		toSerialize["merge_commit_sha"] = o.MergeCommitSha
	}
	if !IsNil(o.Mergeable) {
		toSerialize["mergeable"] = o.Mergeable
	}
	if !IsNil(o.Merged) {
		toSerialize["merged"] = o.Merged
	}
	if !IsNil(o.MergedAt) {
		toSerialize["merged_at"] = o.MergedAt
	}
	if !IsNil(o.MergedBy) {
		toSerialize["merged_by"] = o.MergedBy
	}
	if !IsNil(o.Milestone) {
		toSerialize["milestone"] = o.Milestone
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.PatchUrl) {
		toSerialize["patch_url"] = o.PatchUrl
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullablePullRequest struct {
	value *PullRequest
	isSet bool
}

func (v NullablePullRequest) Get() *PullRequest {
	return v.value
}

func (v *NullablePullRequest) Set(val *PullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequest(val *PullRequest) *NullablePullRequest {
	return &NullablePullRequest{value: val, isSet: true}
}

func (v NullablePullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
