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

// checks if the CreateLabelOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLabelOption{}

// CreateLabelOption CreateLabelOption options for creating a label
type CreateLabelOption struct {
	Color       string  `json:"color"`
	Description *string `json:"description,omitempty"`
	Exclusive   *bool   `json:"exclusive,omitempty"`
	Name        string  `json:"name"`
}

// NewCreateLabelOption instantiates a new CreateLabelOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelOption(color string, name string) *CreateLabelOption {
	this := CreateLabelOption{}
	this.Color = color
	this.Name = name
	return &this
}

// NewCreateLabelOptionWithDefaults instantiates a new CreateLabelOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelOptionWithDefaults() *CreateLabelOption {
	this := CreateLabelOption{}
	return &this
}

// GetColor returns the Color field value
func (o *CreateLabelOption) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *CreateLabelOption) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *CreateLabelOption) SetColor(v string) {
	o.Color = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateLabelOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabelOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateLabelOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateLabelOption) SetDescription(v string) {
	o.Description = &v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise.
func (o *CreateLabelOption) GetExclusive() bool {
	if o == nil || IsNil(o.Exclusive) {
		var ret bool
		return ret
	}
	return *o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabelOption) GetExclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *CreateLabelOption) HasExclusive() bool {
	if o != nil && !IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given bool and assigns it to the Exclusive field.
func (o *CreateLabelOption) SetExclusive(v bool) {
	o.Exclusive = &v
}

// GetName returns the Name field value
func (o *CreateLabelOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateLabelOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateLabelOption) SetName(v string) {
	o.Name = v
}

func (o CreateLabelOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLabelOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["color"] = o.Color
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Exclusive) {
		toSerialize["exclusive"] = o.Exclusive
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCreateLabelOption struct {
	value *CreateLabelOption
	isSet bool
}

func (v NullableCreateLabelOption) Get() *CreateLabelOption {
	return v.value
}

func (v *NullableCreateLabelOption) Set(val *CreateLabelOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelOption(val *CreateLabelOption) *NullableCreateLabelOption {
	return &NullableCreateLabelOption{value: val, isSet: true}
}

func (v NullableCreateLabelOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}