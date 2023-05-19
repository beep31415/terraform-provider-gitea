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

// checks if the CreatePushMirrorOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePushMirrorOption{}

// CreatePushMirrorOption struct for CreatePushMirrorOption
type CreatePushMirrorOption struct {
	Interval       *string `json:"interval,omitempty"`
	RemoteAddress  *string `json:"remote_address,omitempty"`
	RemotePassword *string `json:"remote_password,omitempty"`
	RemoteUsername *string `json:"remote_username,omitempty"`
	SyncOnCommit   *bool   `json:"sync_on_commit,omitempty"`
}

// NewCreatePushMirrorOption instantiates a new CreatePushMirrorOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePushMirrorOption() *CreatePushMirrorOption {
	this := CreatePushMirrorOption{}
	return &this
}

// NewCreatePushMirrorOptionWithDefaults instantiates a new CreatePushMirrorOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePushMirrorOptionWithDefaults() *CreatePushMirrorOption {
	this := CreatePushMirrorOption{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CreatePushMirrorOption) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushMirrorOption) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CreatePushMirrorOption) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *CreatePushMirrorOption) SetInterval(v string) {
	o.Interval = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *CreatePushMirrorOption) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushMirrorOption) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *CreatePushMirrorOption) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *CreatePushMirrorOption) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetRemotePassword returns the RemotePassword field value if set, zero value otherwise.
func (o *CreatePushMirrorOption) GetRemotePassword() string {
	if o == nil || IsNil(o.RemotePassword) {
		var ret string
		return ret
	}
	return *o.RemotePassword
}

// GetRemotePasswordOk returns a tuple with the RemotePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushMirrorOption) GetRemotePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RemotePassword) {
		return nil, false
	}
	return o.RemotePassword, true
}

// HasRemotePassword returns a boolean if a field has been set.
func (o *CreatePushMirrorOption) HasRemotePassword() bool {
	if o != nil && !IsNil(o.RemotePassword) {
		return true
	}

	return false
}

// SetRemotePassword gets a reference to the given string and assigns it to the RemotePassword field.
func (o *CreatePushMirrorOption) SetRemotePassword(v string) {
	o.RemotePassword = &v
}

// GetRemoteUsername returns the RemoteUsername field value if set, zero value otherwise.
func (o *CreatePushMirrorOption) GetRemoteUsername() string {
	if o == nil || IsNil(o.RemoteUsername) {
		var ret string
		return ret
	}
	return *o.RemoteUsername
}

// GetRemoteUsernameOk returns a tuple with the RemoteUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushMirrorOption) GetRemoteUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteUsername) {
		return nil, false
	}
	return o.RemoteUsername, true
}

// HasRemoteUsername returns a boolean if a field has been set.
func (o *CreatePushMirrorOption) HasRemoteUsername() bool {
	if o != nil && !IsNil(o.RemoteUsername) {
		return true
	}

	return false
}

// SetRemoteUsername gets a reference to the given string and assigns it to the RemoteUsername field.
func (o *CreatePushMirrorOption) SetRemoteUsername(v string) {
	o.RemoteUsername = &v
}

// GetSyncOnCommit returns the SyncOnCommit field value if set, zero value otherwise.
func (o *CreatePushMirrorOption) GetSyncOnCommit() bool {
	if o == nil || IsNil(o.SyncOnCommit) {
		var ret bool
		return ret
	}
	return *o.SyncOnCommit
}

// GetSyncOnCommitOk returns a tuple with the SyncOnCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushMirrorOption) GetSyncOnCommitOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncOnCommit) {
		return nil, false
	}
	return o.SyncOnCommit, true
}

// HasSyncOnCommit returns a boolean if a field has been set.
func (o *CreatePushMirrorOption) HasSyncOnCommit() bool {
	if o != nil && !IsNil(o.SyncOnCommit) {
		return true
	}

	return false
}

// SetSyncOnCommit gets a reference to the given bool and assigns it to the SyncOnCommit field.
func (o *CreatePushMirrorOption) SetSyncOnCommit(v bool) {
	o.SyncOnCommit = &v
}

func (o CreatePushMirrorOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePushMirrorOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.RemoteAddress) {
		toSerialize["remote_address"] = o.RemoteAddress
	}
	if !IsNil(o.RemotePassword) {
		toSerialize["remote_password"] = o.RemotePassword
	}
	if !IsNil(o.RemoteUsername) {
		toSerialize["remote_username"] = o.RemoteUsername
	}
	if !IsNil(o.SyncOnCommit) {
		toSerialize["sync_on_commit"] = o.SyncOnCommit
	}
	return toSerialize, nil
}

type NullableCreatePushMirrorOption struct {
	value *CreatePushMirrorOption
	isSet bool
}

func (v NullableCreatePushMirrorOption) Get() *CreatePushMirrorOption {
	return v.value
}

func (v *NullableCreatePushMirrorOption) Set(val *CreatePushMirrorOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePushMirrorOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePushMirrorOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePushMirrorOption(val *CreatePushMirrorOption) *NullableCreatePushMirrorOption {
	return &NullableCreatePushMirrorOption{value: val, isSet: true}
}

func (v NullableCreatePushMirrorOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePushMirrorOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}