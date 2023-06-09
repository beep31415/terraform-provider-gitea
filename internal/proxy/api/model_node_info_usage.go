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

// checks if the NodeInfoUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInfoUsage{}

// NodeInfoUsage NodeInfoUsage contains usage statistics for this server
type NodeInfoUsage struct {
	LocalComments *int64              `json:"localComments,omitempty"`
	LocalPosts    *int64              `json:"localPosts,omitempty"`
	Users         *NodeInfoUsageUsers `json:"users,omitempty"`
}

// NewNodeInfoUsage instantiates a new NodeInfoUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInfoUsage() *NodeInfoUsage {
	this := NodeInfoUsage{}
	return &this
}

// NewNodeInfoUsageWithDefaults instantiates a new NodeInfoUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInfoUsageWithDefaults() *NodeInfoUsage {
	this := NodeInfoUsage{}
	return &this
}

// GetLocalComments returns the LocalComments field value if set, zero value otherwise.
func (o *NodeInfoUsage) GetLocalComments() int64 {
	if o == nil || IsNil(o.LocalComments) {
		var ret int64
		return ret
	}
	return *o.LocalComments
}

// GetLocalCommentsOk returns a tuple with the LocalComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfoUsage) GetLocalCommentsOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalComments) {
		return nil, false
	}
	return o.LocalComments, true
}

// HasLocalComments returns a boolean if a field has been set.
func (o *NodeInfoUsage) HasLocalComments() bool {
	if o != nil && !IsNil(o.LocalComments) {
		return true
	}

	return false
}

// SetLocalComments gets a reference to the given int64 and assigns it to the LocalComments field.
func (o *NodeInfoUsage) SetLocalComments(v int64) {
	o.LocalComments = &v
}

// GetLocalPosts returns the LocalPosts field value if set, zero value otherwise.
func (o *NodeInfoUsage) GetLocalPosts() int64 {
	if o == nil || IsNil(o.LocalPosts) {
		var ret int64
		return ret
	}
	return *o.LocalPosts
}

// GetLocalPostsOk returns a tuple with the LocalPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfoUsage) GetLocalPostsOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalPosts) {
		return nil, false
	}
	return o.LocalPosts, true
}

// HasLocalPosts returns a boolean if a field has been set.
func (o *NodeInfoUsage) HasLocalPosts() bool {
	if o != nil && !IsNil(o.LocalPosts) {
		return true
	}

	return false
}

// SetLocalPosts gets a reference to the given int64 and assigns it to the LocalPosts field.
func (o *NodeInfoUsage) SetLocalPosts(v int64) {
	o.LocalPosts = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NodeInfoUsage) GetUsers() NodeInfoUsageUsers {
	if o == nil || IsNil(o.Users) {
		var ret NodeInfoUsageUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfoUsage) GetUsersOk() (*NodeInfoUsageUsers, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NodeInfoUsage) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given NodeInfoUsageUsers and assigns it to the Users field.
func (o *NodeInfoUsage) SetUsers(v NodeInfoUsageUsers) {
	o.Users = &v
}

func (o NodeInfoUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInfoUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalComments) {
		toSerialize["localComments"] = o.LocalComments
	}
	if !IsNil(o.LocalPosts) {
		toSerialize["localPosts"] = o.LocalPosts
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableNodeInfoUsage struct {
	value *NodeInfoUsage
	isSet bool
}

func (v NullableNodeInfoUsage) Get() *NodeInfoUsage {
	return v.value
}

func (v *NullableNodeInfoUsage) Set(val *NodeInfoUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInfoUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInfoUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInfoUsage(val *NodeInfoUsage) *NullableNodeInfoUsage {
	return &NullableNodeInfoUsage{value: val, isSet: true}
}

func (v NullableNodeInfoUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInfoUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
