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

// checks if the PRBranchInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PRBranchInfo{}

// PRBranchInfo PRBranchInfo information about a branch
type PRBranchInfo struct {
	Label  *string     `json:"label,omitempty"`
	Ref    *string     `json:"ref,omitempty"`
	Repo   *Repository `json:"repo,omitempty"`
	RepoId *int64      `json:"repo_id,omitempty"`
	Sha    *string     `json:"sha,omitempty"`
}

// NewPRBranchInfo instantiates a new PRBranchInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPRBranchInfo() *PRBranchInfo {
	this := PRBranchInfo{}
	return &this
}

// NewPRBranchInfoWithDefaults instantiates a new PRBranchInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPRBranchInfoWithDefaults() *PRBranchInfo {
	this := PRBranchInfo{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PRBranchInfo) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRBranchInfo) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PRBranchInfo) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PRBranchInfo) SetLabel(v string) {
	o.Label = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *PRBranchInfo) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRBranchInfo) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *PRBranchInfo) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *PRBranchInfo) SetRef(v string) {
	o.Ref = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *PRBranchInfo) GetRepo() Repository {
	if o == nil || IsNil(o.Repo) {
		var ret Repository
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRBranchInfo) GetRepoOk() (*Repository, bool) {
	if o == nil || IsNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *PRBranchInfo) HasRepo() bool {
	if o != nil && !IsNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given Repository and assigns it to the Repo field.
func (o *PRBranchInfo) SetRepo(v Repository) {
	o.Repo = &v
}

// GetRepoId returns the RepoId field value if set, zero value otherwise.
func (o *PRBranchInfo) GetRepoId() int64 {
	if o == nil || IsNil(o.RepoId) {
		var ret int64
		return ret
	}
	return *o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRBranchInfo) GetRepoIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RepoId) {
		return nil, false
	}
	return o.RepoId, true
}

// HasRepoId returns a boolean if a field has been set.
func (o *PRBranchInfo) HasRepoId() bool {
	if o != nil && !IsNil(o.RepoId) {
		return true
	}

	return false
}

// SetRepoId gets a reference to the given int64 and assigns it to the RepoId field.
func (o *PRBranchInfo) SetRepoId(v int64) {
	o.RepoId = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *PRBranchInfo) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRBranchInfo) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *PRBranchInfo) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *PRBranchInfo) SetSha(v string) {
	o.Sha = &v
}

func (o PRBranchInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PRBranchInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !IsNil(o.RepoId) {
		toSerialize["repo_id"] = o.RepoId
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	return toSerialize, nil
}

type NullablePRBranchInfo struct {
	value *PRBranchInfo
	isSet bool
}

func (v NullablePRBranchInfo) Get() *PRBranchInfo {
	return v.value
}

func (v *NullablePRBranchInfo) Set(val *PRBranchInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePRBranchInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePRBranchInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePRBranchInfo(val *PRBranchInfo) *NullablePRBranchInfo {
	return &NullablePRBranchInfo{value: val, isSet: true}
}

func (v NullablePRBranchInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePRBranchInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}