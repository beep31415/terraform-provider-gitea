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

// checks if the CreateForkOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateForkOption{}

// CreateForkOption CreateForkOption options for creating a fork
type CreateForkOption struct {
	// name of the forked repository
	Name *string `json:"name,omitempty"`
	// organization name, if forking into an organization
	Organization *string `json:"organization,omitempty"`
}

// NewCreateForkOption instantiates a new CreateForkOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateForkOption() *CreateForkOption {
	this := CreateForkOption{}
	return &this
}

// NewCreateForkOptionWithDefaults instantiates a new CreateForkOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateForkOptionWithDefaults() *CreateForkOption {
	this := CreateForkOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateForkOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateForkOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateForkOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateForkOption) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CreateForkOption) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateForkOption) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateForkOption) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *CreateForkOption) SetOrganization(v string) {
	o.Organization = &v
}

func (o CreateForkOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateForkOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableCreateForkOption struct {
	value *CreateForkOption
	isSet bool
}

func (v NullableCreateForkOption) Get() *CreateForkOption {
	return v.value
}

func (v *NullableCreateForkOption) Set(val *CreateForkOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateForkOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateForkOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateForkOption(val *CreateForkOption) *NullableCreateForkOption {
	return &NullableCreateForkOption{value: val, isSet: true}
}

func (v NullableCreateForkOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateForkOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
