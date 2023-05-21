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

// checks if the EditDeadlineOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditDeadlineOption{}

// EditDeadlineOption EditDeadlineOption options for creating a deadline
type EditDeadlineOption struct {
	DueDate time.Time `json:"due_date"`
}

// NewEditDeadlineOption instantiates a new EditDeadlineOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditDeadlineOption(dueDate time.Time) *EditDeadlineOption {
	this := EditDeadlineOption{}
	this.DueDate = dueDate
	return &this
}

// NewEditDeadlineOptionWithDefaults instantiates a new EditDeadlineOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditDeadlineOptionWithDefaults() *EditDeadlineOption {
	this := EditDeadlineOption{}
	return &this
}

// GetDueDate returns the DueDate field value
func (o *EditDeadlineOption) GetDueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *EditDeadlineOption) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *EditDeadlineOption) SetDueDate(v time.Time) {
	o.DueDate = v
}

func (o EditDeadlineOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditDeadlineOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["due_date"] = o.DueDate
	return toSerialize, nil
}

type NullableEditDeadlineOption struct {
	value *EditDeadlineOption
	isSet bool
}

func (v NullableEditDeadlineOption) Get() *EditDeadlineOption {
	return v.value
}

func (v *NullableEditDeadlineOption) Set(val *EditDeadlineOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEditDeadlineOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEditDeadlineOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditDeadlineOption(val *EditDeadlineOption) *NullableEditDeadlineOption {
	return &NullableEditDeadlineOption{value: val, isSet: true}
}

func (v NullableEditDeadlineOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditDeadlineOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}