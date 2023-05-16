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

// checks if the PayloadUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayloadUser{}

// PayloadUser PayloadUser represents the author or committer of a commit
type PayloadUser struct {
	Email *string `json:"email,omitempty"`
	// Full name of the commit author
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewPayloadUser instantiates a new PayloadUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadUser() *PayloadUser {
	this := PayloadUser{}
	return &this
}

// NewPayloadUserWithDefaults instantiates a new PayloadUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadUserWithDefaults() *PayloadUser {
	this := PayloadUser{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PayloadUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PayloadUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PayloadUser) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PayloadUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PayloadUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PayloadUser) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PayloadUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PayloadUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PayloadUser) SetUsername(v string) {
	o.Username = &v
}

func (o PayloadUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayloadUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullablePayloadUser struct {
	value *PayloadUser
	isSet bool
}

func (v NullablePayloadUser) Get() *PayloadUser {
	return v.value
}

func (v *NullablePayloadUser) Set(val *PayloadUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadUser(val *PayloadUser) *NullablePayloadUser {
	return &NullablePayloadUser{value: val, isSet: true}
}

func (v NullablePayloadUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
