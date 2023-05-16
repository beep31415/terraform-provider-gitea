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

// checks if the ChangedFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangedFile{}

// ChangedFile ChangedFile store information about files affected by the pull request
type ChangedFile struct {
	Additions        *int64  `json:"additions,omitempty"`
	Changes          *int64  `json:"changes,omitempty"`
	ContentsUrl      *string `json:"contents_url,omitempty"`
	Deletions        *int64  `json:"deletions,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	HtmlUrl          *string `json:"html_url,omitempty"`
	PreviousFilename *string `json:"previous_filename,omitempty"`
	RawUrl           *string `json:"raw_url,omitempty"`
	Status           *string `json:"status,omitempty"`
}

// NewChangedFile instantiates a new ChangedFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangedFile() *ChangedFile {
	this := ChangedFile{}
	return &this
}

// NewChangedFileWithDefaults instantiates a new ChangedFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangedFileWithDefaults() *ChangedFile {
	this := ChangedFile{}
	return &this
}

// GetAdditions returns the Additions field value if set, zero value otherwise.
func (o *ChangedFile) GetAdditions() int64 {
	if o == nil || IsNil(o.Additions) {
		var ret int64
		return ret
	}
	return *o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetAdditionsOk() (*int64, bool) {
	if o == nil || IsNil(o.Additions) {
		return nil, false
	}
	return o.Additions, true
}

// HasAdditions returns a boolean if a field has been set.
func (o *ChangedFile) HasAdditions() bool {
	if o != nil && !IsNil(o.Additions) {
		return true
	}

	return false
}

// SetAdditions gets a reference to the given int64 and assigns it to the Additions field.
func (o *ChangedFile) SetAdditions(v int64) {
	o.Additions = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *ChangedFile) GetChanges() int64 {
	if o == nil || IsNil(o.Changes) {
		var ret int64
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetChangesOk() (*int64, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *ChangedFile) HasChanges() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given int64 and assigns it to the Changes field.
func (o *ChangedFile) SetChanges(v int64) {
	o.Changes = &v
}

// GetContentsUrl returns the ContentsUrl field value if set, zero value otherwise.
func (o *ChangedFile) GetContentsUrl() string {
	if o == nil || IsNil(o.ContentsUrl) {
		var ret string
		return ret
	}
	return *o.ContentsUrl
}

// GetContentsUrlOk returns a tuple with the ContentsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetContentsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContentsUrl) {
		return nil, false
	}
	return o.ContentsUrl, true
}

// HasContentsUrl returns a boolean if a field has been set.
func (o *ChangedFile) HasContentsUrl() bool {
	if o != nil && !IsNil(o.ContentsUrl) {
		return true
	}

	return false
}

// SetContentsUrl gets a reference to the given string and assigns it to the ContentsUrl field.
func (o *ChangedFile) SetContentsUrl(v string) {
	o.ContentsUrl = &v
}

// GetDeletions returns the Deletions field value if set, zero value otherwise.
func (o *ChangedFile) GetDeletions() int64 {
	if o == nil || IsNil(o.Deletions) {
		var ret int64
		return ret
	}
	return *o.Deletions
}

// GetDeletionsOk returns a tuple with the Deletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetDeletionsOk() (*int64, bool) {
	if o == nil || IsNil(o.Deletions) {
		return nil, false
	}
	return o.Deletions, true
}

// HasDeletions returns a boolean if a field has been set.
func (o *ChangedFile) HasDeletions() bool {
	if o != nil && !IsNil(o.Deletions) {
		return true
	}

	return false
}

// SetDeletions gets a reference to the given int64 and assigns it to the Deletions field.
func (o *ChangedFile) SetDeletions(v int64) {
	o.Deletions = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ChangedFile) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ChangedFile) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ChangedFile) SetFilename(v string) {
	o.Filename = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *ChangedFile) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *ChangedFile) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *ChangedFile) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetPreviousFilename returns the PreviousFilename field value if set, zero value otherwise.
func (o *ChangedFile) GetPreviousFilename() string {
	if o == nil || IsNil(o.PreviousFilename) {
		var ret string
		return ret
	}
	return *o.PreviousFilename
}

// GetPreviousFilenameOk returns a tuple with the PreviousFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetPreviousFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousFilename) {
		return nil, false
	}
	return o.PreviousFilename, true
}

// HasPreviousFilename returns a boolean if a field has been set.
func (o *ChangedFile) HasPreviousFilename() bool {
	if o != nil && !IsNil(o.PreviousFilename) {
		return true
	}

	return false
}

// SetPreviousFilename gets a reference to the given string and assigns it to the PreviousFilename field.
func (o *ChangedFile) SetPreviousFilename(v string) {
	o.PreviousFilename = &v
}

// GetRawUrl returns the RawUrl field value if set, zero value otherwise.
func (o *ChangedFile) GetRawUrl() string {
	if o == nil || IsNil(o.RawUrl) {
		var ret string
		return ret
	}
	return *o.RawUrl
}

// GetRawUrlOk returns a tuple with the RawUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetRawUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RawUrl) {
		return nil, false
	}
	return o.RawUrl, true
}

// HasRawUrl returns a boolean if a field has been set.
func (o *ChangedFile) HasRawUrl() bool {
	if o != nil && !IsNil(o.RawUrl) {
		return true
	}

	return false
}

// SetRawUrl gets a reference to the given string and assigns it to the RawUrl field.
func (o *ChangedFile) SetRawUrl(v string) {
	o.RawUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChangedFile) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedFile) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChangedFile) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChangedFile) SetStatus(v string) {
	o.Status = &v
}

func (o ChangedFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangedFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Additions) {
		toSerialize["additions"] = o.Additions
	}
	if !IsNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	if !IsNil(o.ContentsUrl) {
		toSerialize["contents_url"] = o.ContentsUrl
	}
	if !IsNil(o.Deletions) {
		toSerialize["deletions"] = o.Deletions
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.PreviousFilename) {
		toSerialize["previous_filename"] = o.PreviousFilename
	}
	if !IsNil(o.RawUrl) {
		toSerialize["raw_url"] = o.RawUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChangedFile struct {
	value *ChangedFile
	isSet bool
}

func (v NullableChangedFile) Get() *ChangedFile {
	return v.value
}

func (v *NullableChangedFile) Set(val *ChangedFile) {
	v.value = val
	v.isSet = true
}

func (v NullableChangedFile) IsSet() bool {
	return v.isSet
}

func (v *NullableChangedFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangedFile(val *ChangedFile) *NullableChangedFile {
	return &NullableChangedFile{value: val, isSet: true}
}

func (v NullableChangedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangedFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
