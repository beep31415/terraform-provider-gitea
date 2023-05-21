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

// checks if the GitignoreTemplateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitignoreTemplateInfo{}

// GitignoreTemplateInfo GitignoreTemplateInfo name and text of a gitignore template
type GitignoreTemplateInfo struct {
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewGitignoreTemplateInfo instantiates a new GitignoreTemplateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitignoreTemplateInfo() *GitignoreTemplateInfo {
	this := GitignoreTemplateInfo{}
	return &this
}

// NewGitignoreTemplateInfoWithDefaults instantiates a new GitignoreTemplateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitignoreTemplateInfoWithDefaults() *GitignoreTemplateInfo {
	this := GitignoreTemplateInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GitignoreTemplateInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitignoreTemplateInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GitignoreTemplateInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GitignoreTemplateInfo) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GitignoreTemplateInfo) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitignoreTemplateInfo) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GitignoreTemplateInfo) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GitignoreTemplateInfo) SetSource(v string) {
	o.Source = &v
}

func (o GitignoreTemplateInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitignoreTemplateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableGitignoreTemplateInfo struct {
	value *GitignoreTemplateInfo
	isSet bool
}

func (v NullableGitignoreTemplateInfo) Get() *GitignoreTemplateInfo {
	return v.value
}

func (v *NullableGitignoreTemplateInfo) Set(val *GitignoreTemplateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGitignoreTemplateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGitignoreTemplateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitignoreTemplateInfo(val *GitignoreTemplateInfo) *NullableGitignoreTemplateInfo {
	return &NullableGitignoreTemplateInfo{value: val, isSet: true}
}

func (v NullableGitignoreTemplateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitignoreTemplateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
