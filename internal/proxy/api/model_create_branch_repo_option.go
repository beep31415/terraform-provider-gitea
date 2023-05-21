/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-833-g040970c32
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateBranchRepoOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBranchRepoOption{}

// CreateBranchRepoOption CreateBranchRepoOption options when creating a branch in a repository
type CreateBranchRepoOption struct {
	// Name of the branch to create
	NewBranchName string `json:"new_branch_name"`
	// Deprecated: true Name of the old branch to create from
	OldBranchName *string `json:"old_branch_name,omitempty"`
	// Name of the old branch/tag/commit to create from
	OldRefName *string `json:"old_ref_name,omitempty"`
}

// NewCreateBranchRepoOption instantiates a new CreateBranchRepoOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBranchRepoOption(newBranchName string) *CreateBranchRepoOption {
	this := CreateBranchRepoOption{}
	this.NewBranchName = newBranchName
	return &this
}

// NewCreateBranchRepoOptionWithDefaults instantiates a new CreateBranchRepoOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBranchRepoOptionWithDefaults() *CreateBranchRepoOption {
	this := CreateBranchRepoOption{}
	return &this
}

// GetNewBranchName returns the NewBranchName field value
func (o *CreateBranchRepoOption) GetNewBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewBranchName
}

// GetNewBranchNameOk returns a tuple with the NewBranchName field value
// and a boolean to check if the value has been set.
func (o *CreateBranchRepoOption) GetNewBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewBranchName, true
}

// SetNewBranchName sets field value
func (o *CreateBranchRepoOption) SetNewBranchName(v string) {
	o.NewBranchName = v
}

// GetOldBranchName returns the OldBranchName field value if set, zero value otherwise.
func (o *CreateBranchRepoOption) GetOldBranchName() string {
	if o == nil || IsNil(o.OldBranchName) {
		var ret string
		return ret
	}
	return *o.OldBranchName
}

// GetOldBranchNameOk returns a tuple with the OldBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchRepoOption) GetOldBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.OldBranchName) {
		return nil, false
	}
	return o.OldBranchName, true
}

// HasOldBranchName returns a boolean if a field has been set.
func (o *CreateBranchRepoOption) HasOldBranchName() bool {
	if o != nil && !IsNil(o.OldBranchName) {
		return true
	}

	return false
}

// SetOldBranchName gets a reference to the given string and assigns it to the OldBranchName field.
func (o *CreateBranchRepoOption) SetOldBranchName(v string) {
	o.OldBranchName = &v
}

// GetOldRefName returns the OldRefName field value if set, zero value otherwise.
func (o *CreateBranchRepoOption) GetOldRefName() string {
	if o == nil || IsNil(o.OldRefName) {
		var ret string
		return ret
	}
	return *o.OldRefName
}

// GetOldRefNameOk returns a tuple with the OldRefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchRepoOption) GetOldRefNameOk() (*string, bool) {
	if o == nil || IsNil(o.OldRefName) {
		return nil, false
	}
	return o.OldRefName, true
}

// HasOldRefName returns a boolean if a field has been set.
func (o *CreateBranchRepoOption) HasOldRefName() bool {
	if o != nil && !IsNil(o.OldRefName) {
		return true
	}

	return false
}

// SetOldRefName gets a reference to the given string and assigns it to the OldRefName field.
func (o *CreateBranchRepoOption) SetOldRefName(v string) {
	o.OldRefName = &v
}

func (o CreateBranchRepoOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBranchRepoOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_branch_name"] = o.NewBranchName
	if !IsNil(o.OldBranchName) {
		toSerialize["old_branch_name"] = o.OldBranchName
	}
	if !IsNil(o.OldRefName) {
		toSerialize["old_ref_name"] = o.OldRefName
	}
	return toSerialize, nil
}

type NullableCreateBranchRepoOption struct {
	value *CreateBranchRepoOption
	isSet bool
}

func (v NullableCreateBranchRepoOption) Get() *CreateBranchRepoOption {
	return v.value
}

func (v *NullableCreateBranchRepoOption) Set(val *CreateBranchRepoOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBranchRepoOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBranchRepoOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBranchRepoOption(val *CreateBranchRepoOption) *NullableCreateBranchRepoOption {
	return &NullableCreateBranchRepoOption{value: val, isSet: true}
}

func (v NullableCreateBranchRepoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBranchRepoOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}