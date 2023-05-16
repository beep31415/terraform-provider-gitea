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

// checks if the IssueFormField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueFormField{}

// IssueFormField IssueFormField represents a form field
type IssueFormField struct {
	Attributes  map[string]map[string]interface{} `json:"attributes,omitempty"`
	Id          *string                           `json:"id,omitempty"`
	Type        *string                           `json:"type,omitempty"`
	Validations map[string]map[string]interface{} `json:"validations,omitempty"`
}

// NewIssueFormField instantiates a new IssueFormField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueFormField() *IssueFormField {
	this := IssueFormField{}
	return &this
}

// NewIssueFormFieldWithDefaults instantiates a new IssueFormField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueFormFieldWithDefaults() *IssueFormField {
	this := IssueFormField{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IssueFormField) GetAttributes() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFormField) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IssueFormField) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *IssueFormField) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssueFormField) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFormField) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssueFormField) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IssueFormField) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueFormField) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFormField) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueFormField) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueFormField) SetType(v string) {
	o.Type = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *IssueFormField) GetValidations() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Validations) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFormField) GetValidationsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Validations) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *IssueFormField) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given map[string]map[string]interface{} and assigns it to the Validations field.
func (o *IssueFormField) SetValidations(v map[string]map[string]interface{}) {
	o.Validations = v
}

func (o IssueFormField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueFormField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

type NullableIssueFormField struct {
	value *IssueFormField
	isSet bool
}

func (v NullableIssueFormField) Get() *IssueFormField {
	return v.value
}

func (v *NullableIssueFormField) Set(val *IssueFormField) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueFormField) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueFormField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueFormField(val *IssueFormField) *NullableIssueFormField {
	return &NullableIssueFormField{value: val, isSet: true}
}

func (v NullableIssueFormField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueFormField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
