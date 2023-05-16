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

// checks if the Permission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permission{}

// Permission Permission represents a set of permissions
type Permission struct {
	Admin *bool `json:"admin,omitempty"`
	Pull  *bool `json:"pull,omitempty"`
	Push  *bool `json:"push,omitempty"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission() *Permission {
	this := Permission{}
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *Permission) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *Permission) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *Permission) SetAdmin(v bool) {
	o.Admin = &v
}

// GetPull returns the Pull field value if set, zero value otherwise.
func (o *Permission) GetPull() bool {
	if o == nil || IsNil(o.Pull) {
		var ret bool
		return ret
	}
	return *o.Pull
}

// GetPullOk returns a tuple with the Pull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetPullOk() (*bool, bool) {
	if o == nil || IsNil(o.Pull) {
		return nil, false
	}
	return o.Pull, true
}

// HasPull returns a boolean if a field has been set.
func (o *Permission) HasPull() bool {
	if o != nil && !IsNil(o.Pull) {
		return true
	}

	return false
}

// SetPull gets a reference to the given bool and assigns it to the Pull field.
func (o *Permission) SetPull(v bool) {
	o.Pull = &v
}

// GetPush returns the Push field value if set, zero value otherwise.
func (o *Permission) GetPush() bool {
	if o == nil || IsNil(o.Push) {
		var ret bool
		return ret
	}
	return *o.Push
}

// GetPushOk returns a tuple with the Push field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetPushOk() (*bool, bool) {
	if o == nil || IsNil(o.Push) {
		return nil, false
	}
	return o.Push, true
}

// HasPush returns a boolean if a field has been set.
func (o *Permission) HasPush() bool {
	if o != nil && !IsNil(o.Push) {
		return true
	}

	return false
}

// SetPush gets a reference to the given bool and assigns it to the Push field.
func (o *Permission) SetPush(v bool) {
	o.Push = &v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Pull) {
		toSerialize["pull"] = o.Pull
	}
	if !IsNil(o.Push) {
		toSerialize["push"] = o.Push
	}
	return toSerialize, nil
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
