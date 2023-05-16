/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-790-g61ad4c607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the EditAttachmentOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAttachmentOptions{}

// EditAttachmentOptions EditAttachmentOptions options for editing attachments
type EditAttachmentOptions struct {
	Name *string `json:"name,omitempty"`
}

// NewEditAttachmentOptions instantiates a new EditAttachmentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAttachmentOptions() *EditAttachmentOptions {
	this := EditAttachmentOptions{}
	return &this
}

// NewEditAttachmentOptionsWithDefaults instantiates a new EditAttachmentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAttachmentOptionsWithDefaults() *EditAttachmentOptions {
	this := EditAttachmentOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditAttachmentOptions) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAttachmentOptions) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditAttachmentOptions) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditAttachmentOptions) SetName(v string) {
	o.Name = &v
}

func (o EditAttachmentOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAttachmentOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEditAttachmentOptions struct {
	value *EditAttachmentOptions
	isSet bool
}

func (v NullableEditAttachmentOptions) Get() *EditAttachmentOptions {
	return v.value
}

func (v *NullableEditAttachmentOptions) Set(val *EditAttachmentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAttachmentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAttachmentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAttachmentOptions(val *EditAttachmentOptions) *NullableEditAttachmentOptions {
	return &NullableEditAttachmentOptions{value: val, isSet: true}
}

func (v NullableEditAttachmentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAttachmentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
