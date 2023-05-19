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

// checks if the TransferRepoOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferRepoOption{}

// TransferRepoOption TransferRepoOption options when transfer a repository's ownership
type TransferRepoOption struct {
	NewOwner string `json:"new_owner"`
	// ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories.
	TeamIds []int64 `json:"team_ids,omitempty"`
}

// NewTransferRepoOption instantiates a new TransferRepoOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepoOption(newOwner string) *TransferRepoOption {
	this := TransferRepoOption{}
	this.NewOwner = newOwner
	return &this
}

// NewTransferRepoOptionWithDefaults instantiates a new TransferRepoOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepoOptionWithDefaults() *TransferRepoOption {
	this := TransferRepoOption{}
	return &this
}

// GetNewOwner returns the NewOwner field value
func (o *TransferRepoOption) GetNewOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewOwner
}

// GetNewOwnerOk returns a tuple with the NewOwner field value
// and a boolean to check if the value has been set.
func (o *TransferRepoOption) GetNewOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewOwner, true
}

// SetNewOwner sets field value
func (o *TransferRepoOption) SetNewOwner(v string) {
	o.NewOwner = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *TransferRepoOption) GetTeamIds() []int64 {
	if o == nil || IsNil(o.TeamIds) {
		var ret []int64
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRepoOption) GetTeamIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *TransferRepoOption) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []int64 and assigns it to the TeamIds field.
func (o *TransferRepoOption) SetTeamIds(v []int64) {
	o.TeamIds = v
}

func (o TransferRepoOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRepoOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_owner"] = o.NewOwner
	if !IsNil(o.TeamIds) {
		toSerialize["team_ids"] = o.TeamIds
	}
	return toSerialize, nil
}

type NullableTransferRepoOption struct {
	value *TransferRepoOption
	isSet bool
}

func (v NullableTransferRepoOption) Get() *TransferRepoOption {
	return v.value
}

func (v *NullableTransferRepoOption) Set(val *TransferRepoOption) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepoOption) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepoOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepoOption(val *TransferRepoOption) *NullableTransferRepoOption {
	return &NullableTransferRepoOption{value: val, isSet: true}
}

func (v NullableTransferRepoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepoOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}