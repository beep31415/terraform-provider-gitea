/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-833-g040970c32
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the CreateUserOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserOption{}

// CreateUserOption CreateUserOption create user options
type CreateUserOption struct {
	// For explicitly setting the user creation timestamp. Useful when users are migrated from other systems. When omitted, the user's creation timestamp will be set to \"now\".
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	Email              string     `json:"email"`
	FullName           *string    `json:"full_name,omitempty"`
	LoginName          *string    `json:"login_name,omitempty"`
	MustChangePassword *bool      `json:"must_change_password,omitempty"`
	Password           string     `json:"password"`
	Restricted         *bool      `json:"restricted,omitempty"`
	SendNotify         *bool      `json:"send_notify,omitempty"`
	SourceId           *int64     `json:"source_id,omitempty"`
	Username           string     `json:"username"`
	Visibility         *string    `json:"visibility,omitempty"`
}

// NewCreateUserOption instantiates a new CreateUserOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserOption(email string, password string, username string) *CreateUserOption {
	this := CreateUserOption{}
	this.Email = email
	this.Password = password
	this.Username = username
	return &this
}

// NewCreateUserOptionWithDefaults instantiates a new CreateUserOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserOptionWithDefaults() *CreateUserOption {
	this := CreateUserOption{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateUserOption) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateUserOption) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateUserOption) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value
func (o *CreateUserOption) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateUserOption) SetEmail(v string) {
	o.Email = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *CreateUserOption) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *CreateUserOption) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *CreateUserOption) SetFullName(v string) {
	o.FullName = &v
}

// GetLoginName returns the LoginName field value if set, zero value otherwise.
func (o *CreateUserOption) GetLoginName() string {
	if o == nil || IsNil(o.LoginName) {
		var ret string
		return ret
	}
	return *o.LoginName
}

// GetLoginNameOk returns a tuple with the LoginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetLoginNameOk() (*string, bool) {
	if o == nil || IsNil(o.LoginName) {
		return nil, false
	}
	return o.LoginName, true
}

// HasLoginName returns a boolean if a field has been set.
func (o *CreateUserOption) HasLoginName() bool {
	if o != nil && !IsNil(o.LoginName) {
		return true
	}

	return false
}

// SetLoginName gets a reference to the given string and assigns it to the LoginName field.
func (o *CreateUserOption) SetLoginName(v string) {
	o.LoginName = &v
}

// GetMustChangePassword returns the MustChangePassword field value if set, zero value otherwise.
func (o *CreateUserOption) GetMustChangePassword() bool {
	if o == nil || IsNil(o.MustChangePassword) {
		var ret bool
		return ret
	}
	return *o.MustChangePassword
}

// GetMustChangePasswordOk returns a tuple with the MustChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetMustChangePasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.MustChangePassword) {
		return nil, false
	}
	return o.MustChangePassword, true
}

// HasMustChangePassword returns a boolean if a field has been set.
func (o *CreateUserOption) HasMustChangePassword() bool {
	if o != nil && !IsNil(o.MustChangePassword) {
		return true
	}

	return false
}

// SetMustChangePassword gets a reference to the given bool and assigns it to the MustChangePassword field.
func (o *CreateUserOption) SetMustChangePassword(v bool) {
	o.MustChangePassword = &v
}

// GetPassword returns the Password field value
func (o *CreateUserOption) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateUserOption) SetPassword(v string) {
	o.Password = v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *CreateUserOption) GetRestricted() bool {
	if o == nil || IsNil(o.Restricted) {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.Restricted) {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *CreateUserOption) HasRestricted() bool {
	if o != nil && !IsNil(o.Restricted) {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *CreateUserOption) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetSendNotify returns the SendNotify field value if set, zero value otherwise.
func (o *CreateUserOption) GetSendNotify() bool {
	if o == nil || IsNil(o.SendNotify) {
		var ret bool
		return ret
	}
	return *o.SendNotify
}

// GetSendNotifyOk returns a tuple with the SendNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetSendNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNotify) {
		return nil, false
	}
	return o.SendNotify, true
}

// HasSendNotify returns a boolean if a field has been set.
func (o *CreateUserOption) HasSendNotify() bool {
	if o != nil && !IsNil(o.SendNotify) {
		return true
	}

	return false
}

// SetSendNotify gets a reference to the given bool and assigns it to the SendNotify field.
func (o *CreateUserOption) SetSendNotify(v bool) {
	o.SendNotify = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *CreateUserOption) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId) {
		var ret int64
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetSourceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *CreateUserOption) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given int64 and assigns it to the SourceId field.
func (o *CreateUserOption) SetSourceId(v int64) {
	o.SourceId = &v
}

// GetUsername returns the Username field value
func (o *CreateUserOption) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateUserOption) SetUsername(v string) {
	o.Username = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *CreateUserOption) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserOption) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *CreateUserOption) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *CreateUserOption) SetVisibility(v string) {
	o.Visibility = &v
}

func (o CreateUserOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.LoginName) {
		toSerialize["login_name"] = o.LoginName
	}
	if !IsNil(o.MustChangePassword) {
		toSerialize["must_change_password"] = o.MustChangePassword
	}
	toSerialize["password"] = o.Password
	if !IsNil(o.Restricted) {
		toSerialize["restricted"] = o.Restricted
	}
	if !IsNil(o.SendNotify) {
		toSerialize["send_notify"] = o.SendNotify
	}
	if !IsNil(o.SourceId) {
		toSerialize["source_id"] = o.SourceId
	}
	toSerialize["username"] = o.Username
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableCreateUserOption struct {
	value *CreateUserOption
	isSet bool
}

func (v NullableCreateUserOption) Get() *CreateUserOption {
	return v.value
}

func (v *NullableCreateUserOption) Set(val *CreateUserOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserOption(val *CreateUserOption) *NullableCreateUserOption {
	return &NullableCreateUserOption{value: val, isSet: true}
}

func (v NullableCreateUserOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
