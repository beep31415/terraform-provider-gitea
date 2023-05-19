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

// checks if the FileDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDeleteResponse{}

// FileDeleteResponse FileDeleteResponse contains information about a repo's file that was deleted
type FileDeleteResponse struct {
	Commit       *FileCommitResponse        `json:"commit,omitempty"`
	Content      map[string]interface{}     `json:"content,omitempty"`
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// NewFileDeleteResponse instantiates a new FileDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDeleteResponse() *FileDeleteResponse {
	this := FileDeleteResponse{}
	return &this
}

// NewFileDeleteResponseWithDefaults instantiates a new FileDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDeleteResponseWithDefaults() *FileDeleteResponse {
	this := FileDeleteResponse{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *FileDeleteResponse) GetCommit() FileCommitResponse {
	if o == nil || IsNil(o.Commit) {
		var ret FileCommitResponse
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDeleteResponse) GetCommitOk() (*FileCommitResponse, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *FileDeleteResponse) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given FileCommitResponse and assigns it to the Commit field.
func (o *FileDeleteResponse) SetCommit(v FileCommitResponse) {
	o.Commit = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FileDeleteResponse) GetContent() map[string]interface{} {
	if o == nil || IsNil(o.Content) {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDeleteResponse) GetContentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Content) {
		return map[string]interface{}{}, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FileDeleteResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *FileDeleteResponse) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *FileDeleteResponse) GetVerification() PayloadCommitVerification {
	if o == nil || IsNil(o.Verification) {
		var ret PayloadCommitVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDeleteResponse) GetVerificationOk() (*PayloadCommitVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *FileDeleteResponse) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given PayloadCommitVerification and assigns it to the Verification field.
func (o *FileDeleteResponse) SetVerification(v PayloadCommitVerification) {
	o.Verification = &v
}

func (o FileDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	return toSerialize, nil
}

type NullableFileDeleteResponse struct {
	value *FileDeleteResponse
	isSet bool
}

func (v NullableFileDeleteResponse) Get() *FileDeleteResponse {
	return v.value
}

func (v *NullableFileDeleteResponse) Set(val *FileDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDeleteResponse(val *FileDeleteResponse) *NullableFileDeleteResponse {
	return &NullableFileDeleteResponse{value: val, isSet: true}
}

func (v NullableFileDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}