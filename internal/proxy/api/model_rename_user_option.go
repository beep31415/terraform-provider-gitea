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

// checks if the RenameUserOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameUserOption{}

// RenameUserOption RenameUserOption options when renaming a user
type RenameUserOption struct {
	// New username for this user. This name cannot be in use yet by any other user.
	NewUsername string `json:"new_username"`
}

// NewRenameUserOption instantiates a new RenameUserOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameUserOption(newUsername string) *RenameUserOption {
	this := RenameUserOption{}
	this.NewUsername = newUsername
	return &this
}

// NewRenameUserOptionWithDefaults instantiates a new RenameUserOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameUserOptionWithDefaults() *RenameUserOption {
	this := RenameUserOption{}
	return &this
}

// GetNewUsername returns the NewUsername field value
func (o *RenameUserOption) GetNewUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewUsername
}

// GetNewUsernameOk returns a tuple with the NewUsername field value
// and a boolean to check if the value has been set.
func (o *RenameUserOption) GetNewUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewUsername, true
}

// SetNewUsername sets field value
func (o *RenameUserOption) SetNewUsername(v string) {
	o.NewUsername = v
}

func (o RenameUserOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameUserOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_username"] = o.NewUsername
	return toSerialize, nil
}

type NullableRenameUserOption struct {
	value *RenameUserOption
	isSet bool
}

func (v NullableRenameUserOption) Get() *RenameUserOption {
	return v.value
}

func (v *NullableRenameUserOption) Set(val *RenameUserOption) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameUserOption) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameUserOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameUserOption(val *RenameUserOption) *NullableRenameUserOption {
	return &NullableRenameUserOption{value: val, isSet: true}
}

func (v NullableRenameUserOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameUserOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
