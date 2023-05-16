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

// checks if the CommitAffectedFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitAffectedFiles{}

// CommitAffectedFiles CommitAffectedFiles store information about files affected by the commit
type CommitAffectedFiles struct {
	Filename *string `json:"filename,omitempty"`
}

// NewCommitAffectedFiles instantiates a new CommitAffectedFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitAffectedFiles() *CommitAffectedFiles {
	this := CommitAffectedFiles{}
	return &this
}

// NewCommitAffectedFilesWithDefaults instantiates a new CommitAffectedFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitAffectedFilesWithDefaults() *CommitAffectedFiles {
	this := CommitAffectedFiles{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CommitAffectedFiles) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitAffectedFiles) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CommitAffectedFiles) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CommitAffectedFiles) SetFilename(v string) {
	o.Filename = &v
}

func (o CommitAffectedFiles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitAffectedFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	return toSerialize, nil
}

type NullableCommitAffectedFiles struct {
	value *CommitAffectedFiles
	isSet bool
}

func (v NullableCommitAffectedFiles) Get() *CommitAffectedFiles {
	return v.value
}

func (v *NullableCommitAffectedFiles) Set(val *CommitAffectedFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitAffectedFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitAffectedFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitAffectedFiles(val *CommitAffectedFiles) *NullableCommitAffectedFiles {
	return &NullableCommitAffectedFiles{value: val, isSet: true}
}

func (v NullableCommitAffectedFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitAffectedFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
