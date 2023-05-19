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

// checks if the RepoTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepoTransfer{}

// RepoTransfer RepoTransfer represents a pending repo transfer
type RepoTransfer struct {
	Doer      *User  `json:"doer,omitempty"`
	Recipient *User  `json:"recipient,omitempty"`
	Teams     []Team `json:"teams,omitempty"`
}

// NewRepoTransfer instantiates a new RepoTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepoTransfer() *RepoTransfer {
	this := RepoTransfer{}
	return &this
}

// NewRepoTransferWithDefaults instantiates a new RepoTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepoTransferWithDefaults() *RepoTransfer {
	this := RepoTransfer{}
	return &this
}

// GetDoer returns the Doer field value if set, zero value otherwise.
func (o *RepoTransfer) GetDoer() User {
	if o == nil || IsNil(o.Doer) {
		var ret User
		return ret
	}
	return *o.Doer
}

// GetDoerOk returns a tuple with the Doer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoTransfer) GetDoerOk() (*User, bool) {
	if o == nil || IsNil(o.Doer) {
		return nil, false
	}
	return o.Doer, true
}

// HasDoer returns a boolean if a field has been set.
func (o *RepoTransfer) HasDoer() bool {
	if o != nil && !IsNil(o.Doer) {
		return true
	}

	return false
}

// SetDoer gets a reference to the given User and assigns it to the Doer field.
func (o *RepoTransfer) SetDoer(v User) {
	o.Doer = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *RepoTransfer) GetRecipient() User {
	if o == nil || IsNil(o.Recipient) {
		var ret User
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoTransfer) GetRecipientOk() (*User, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *RepoTransfer) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given User and assigns it to the Recipient field.
func (o *RepoTransfer) SetRecipient(v User) {
	o.Recipient = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *RepoTransfer) GetTeams() []Team {
	if o == nil || IsNil(o.Teams) {
		var ret []Team
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoTransfer) GetTeamsOk() ([]Team, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *RepoTransfer) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []Team and assigns it to the Teams field.
func (o *RepoTransfer) SetTeams(v []Team) {
	o.Teams = v
}

func (o RepoTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepoTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Doer) {
		toSerialize["doer"] = o.Doer
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableRepoTransfer struct {
	value *RepoTransfer
	isSet bool
}

func (v NullableRepoTransfer) Get() *RepoTransfer {
	return v.value
}

func (v *NullableRepoTransfer) Set(val *RepoTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableRepoTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableRepoTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepoTransfer(val *RepoTransfer) *NullableRepoTransfer {
	return &NullableRepoTransfer{value: val, isSet: true}
}

func (v NullableRepoTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepoTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}