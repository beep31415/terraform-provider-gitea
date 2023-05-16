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

// checks if the EditLabelOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditLabelOption{}

// EditLabelOption EditLabelOption options for editing a label
type EditLabelOption struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Exclusive   *bool   `json:"exclusive,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// NewEditLabelOption instantiates a new EditLabelOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditLabelOption() *EditLabelOption {
	this := EditLabelOption{}
	return &this
}

// NewEditLabelOptionWithDefaults instantiates a new EditLabelOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditLabelOptionWithDefaults() *EditLabelOption {
	this := EditLabelOption{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *EditLabelOption) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLabelOption) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *EditLabelOption) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *EditLabelOption) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditLabelOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLabelOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditLabelOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditLabelOption) SetDescription(v string) {
	o.Description = &v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise.
func (o *EditLabelOption) GetExclusive() bool {
	if o == nil || IsNil(o.Exclusive) {
		var ret bool
		return ret
	}
	return *o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLabelOption) GetExclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *EditLabelOption) HasExclusive() bool {
	if o != nil && !IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given bool and assigns it to the Exclusive field.
func (o *EditLabelOption) SetExclusive(v bool) {
	o.Exclusive = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditLabelOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLabelOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditLabelOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditLabelOption) SetName(v string) {
	o.Name = &v
}

func (o EditLabelOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditLabelOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Exclusive) {
		toSerialize["exclusive"] = o.Exclusive
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEditLabelOption struct {
	value *EditLabelOption
	isSet bool
}

func (v NullableEditLabelOption) Get() *EditLabelOption {
	return v.value
}

func (v *NullableEditLabelOption) Set(val *EditLabelOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEditLabelOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEditLabelOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditLabelOption(val *EditLabelOption) *NullableEditLabelOption {
	return &NullableEditLabelOption{value: val, isSet: true}
}

func (v NullableEditLabelOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditLabelOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
