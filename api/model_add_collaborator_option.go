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

// checks if the AddCollaboratorOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCollaboratorOption{}

// AddCollaboratorOption AddCollaboratorOption options when adding a user as a collaborator of a repository
type AddCollaboratorOption struct {
	Permission *string `json:"permission,omitempty"`
}

// NewAddCollaboratorOption instantiates a new AddCollaboratorOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCollaboratorOption() *AddCollaboratorOption {
	this := AddCollaboratorOption{}
	return &this
}

// NewAddCollaboratorOptionWithDefaults instantiates a new AddCollaboratorOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCollaboratorOptionWithDefaults() *AddCollaboratorOption {
	this := AddCollaboratorOption{}
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *AddCollaboratorOption) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCollaboratorOption) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *AddCollaboratorOption) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *AddCollaboratorOption) SetPermission(v string) {
	o.Permission = &v
}

func (o AddCollaboratorOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCollaboratorOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	return toSerialize, nil
}

type NullableAddCollaboratorOption struct {
	value *AddCollaboratorOption
	isSet bool
}

func (v NullableAddCollaboratorOption) Get() *AddCollaboratorOption {
	return v.value
}

func (v *NullableAddCollaboratorOption) Set(val *AddCollaboratorOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCollaboratorOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCollaboratorOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCollaboratorOption(val *AddCollaboratorOption) *NullableAddCollaboratorOption {
	return &NullableAddCollaboratorOption{value: val, isSet: true}
}

func (v NullableAddCollaboratorOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCollaboratorOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}