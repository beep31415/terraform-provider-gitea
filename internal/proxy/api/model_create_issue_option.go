/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-833-g040970c32
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the CreateIssueOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIssueOption{}

// CreateIssueOption CreateIssueOption options to create one issue
type CreateIssueOption struct {
	// deprecated
	Assignee  *string    `json:"assignee,omitempty"`
	Assignees []string   `json:"assignees,omitempty"`
	Body      *string    `json:"body,omitempty"`
	Closed    *bool      `json:"closed,omitempty"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	// list of label ids
	Labels []int64 `json:"labels,omitempty"`
	// milestone id
	Milestone *int64  `json:"milestone,omitempty"`
	Ref       *string `json:"ref,omitempty"`
	Title     string  `json:"title"`
}

// NewCreateIssueOption instantiates a new CreateIssueOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueOption(title string) *CreateIssueOption {
	this := CreateIssueOption{}
	this.Title = title
	return &this
}

// NewCreateIssueOptionWithDefaults instantiates a new CreateIssueOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueOptionWithDefaults() *CreateIssueOption {
	this := CreateIssueOption{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *CreateIssueOption) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *CreateIssueOption) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *CreateIssueOption) SetAssignee(v string) {
	o.Assignee = &v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *CreateIssueOption) GetAssignees() []string {
	if o == nil || IsNil(o.Assignees) {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetAssigneesOk() ([]string, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *CreateIssueOption) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *CreateIssueOption) SetAssignees(v []string) {
	o.Assignees = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateIssueOption) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateIssueOption) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CreateIssueOption) SetBody(v string) {
	o.Body = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *CreateIssueOption) GetClosed() bool {
	if o == nil || IsNil(o.Closed) {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.Closed) {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *CreateIssueOption) HasClosed() bool {
	if o != nil && !IsNil(o.Closed) {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *CreateIssueOption) SetClosed(v bool) {
	o.Closed = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *CreateIssueOption) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *CreateIssueOption) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *CreateIssueOption) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateIssueOption) GetLabels() []int64 {
	if o == nil || IsNil(o.Labels) {
		var ret []int64
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetLabelsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateIssueOption) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []int64 and assigns it to the Labels field.
func (o *CreateIssueOption) SetLabels(v []int64) {
	o.Labels = v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *CreateIssueOption) GetMilestone() int64 {
	if o == nil || IsNil(o.Milestone) {
		var ret int64
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetMilestoneOk() (*int64, bool) {
	if o == nil || IsNil(o.Milestone) {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *CreateIssueOption) HasMilestone() bool {
	if o != nil && !IsNil(o.Milestone) {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given int64 and assigns it to the Milestone field.
func (o *CreateIssueOption) SetMilestone(v int64) {
	o.Milestone = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *CreateIssueOption) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *CreateIssueOption) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *CreateIssueOption) SetRef(v string) {
	o.Ref = &v
}

// GetTitle returns the Title field value
func (o *CreateIssueOption) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateIssueOption) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateIssueOption) SetTitle(v string) {
	o.Title = v
}

func (o CreateIssueOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Closed) {
		toSerialize["closed"] = o.Closed
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Milestone) {
		toSerialize["milestone"] = o.Milestone
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

type NullableCreateIssueOption struct {
	value *CreateIssueOption
	isSet bool
}

func (v NullableCreateIssueOption) Get() *CreateIssueOption {
	return v.value
}

func (v *NullableCreateIssueOption) Set(val *CreateIssueOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueOption(val *CreateIssueOption) *NullableCreateIssueOption {
	return &NullableCreateIssueOption{value: val, isSet: true}
}

func (v NullableCreateIssueOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
