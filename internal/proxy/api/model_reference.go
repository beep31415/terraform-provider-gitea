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

// checks if the Reference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reference{}

// Reference struct for Reference
type Reference struct {
	Object *GitObject `json:"object,omitempty"`
	Ref    *string    `json:"ref,omitempty"`
	Url    *string    `json:"url,omitempty"`
}

// NewReference instantiates a new Reference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReference() *Reference {
	this := Reference{}
	return &this
}

// NewReferenceWithDefaults instantiates a new Reference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceWithDefaults() *Reference {
	this := Reference{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Reference) GetObject() GitObject {
	if o == nil || IsNil(o.Object) {
		var ret GitObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetObjectOk() (*GitObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Reference) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given GitObject and assigns it to the Object field.
func (o *Reference) SetObject(v GitObject) {
	o.Object = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Reference) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Reference) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Reference) SetRef(v string) {
	o.Ref = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Reference) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Reference) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Reference) SetUrl(v string) {
	o.Url = &v
}

func (o Reference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableReference struct {
	value *Reference
	isSet bool
}

func (v NullableReference) Get() *Reference {
	return v.value
}

func (v *NullableReference) Set(val *Reference) {
	v.value = val
	v.isSet = true
}

func (v NullableReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReference(val *Reference) *NullableReference {
	return &NullableReference{value: val, isSet: true}
}

func (v NullableReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
