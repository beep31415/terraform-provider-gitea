/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.19.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the PayloadCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayloadCommit{}

// PayloadCommit PayloadCommit represents a commit
type PayloadCommit struct {
	Added     []string     `json:"added,omitempty"`
	Author    *PayloadUser `json:"author,omitempty"`
	Committer *PayloadUser `json:"committer,omitempty"`
	// sha1 hash of the commit
	Id           *string                    `json:"id,omitempty"`
	Message      *string                    `json:"message,omitempty"`
	Modified     []string                   `json:"modified,omitempty"`
	Removed      []string                   `json:"removed,omitempty"`
	Timestamp    *time.Time                 `json:"timestamp,omitempty"`
	Url          *string                    `json:"url,omitempty"`
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// NewPayloadCommit instantiates a new PayloadCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadCommit() *PayloadCommit {
	this := PayloadCommit{}
	return &this
}

// NewPayloadCommitWithDefaults instantiates a new PayloadCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadCommitWithDefaults() *PayloadCommit {
	this := PayloadCommit{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *PayloadCommit) GetAdded() []string {
	if o == nil || IsNil(o.Added) {
		var ret []string
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *PayloadCommit) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []string and assigns it to the Added field.
func (o *PayloadCommit) SetAdded(v []string) {
	o.Added = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *PayloadCommit) GetAuthor() PayloadUser {
	if o == nil || IsNil(o.Author) {
		var ret PayloadUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetAuthorOk() (*PayloadUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *PayloadCommit) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given PayloadUser and assigns it to the Author field.
func (o *PayloadCommit) SetAuthor(v PayloadUser) {
	o.Author = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *PayloadCommit) GetCommitter() PayloadUser {
	if o == nil || IsNil(o.Committer) {
		var ret PayloadUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetCommitterOk() (*PayloadUser, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *PayloadCommit) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given PayloadUser and assigns it to the Committer field.
func (o *PayloadCommit) SetCommitter(v PayloadUser) {
	o.Committer = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PayloadCommit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PayloadCommit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PayloadCommit) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PayloadCommit) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PayloadCommit) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PayloadCommit) SetMessage(v string) {
	o.Message = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PayloadCommit) GetModified() []string {
	if o == nil || IsNil(o.Modified) {
		var ret []string
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PayloadCommit) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given []string and assigns it to the Modified field.
func (o *PayloadCommit) SetModified(v []string) {
	o.Modified = v
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *PayloadCommit) GetRemoved() []string {
	if o == nil || IsNil(o.Removed) {
		var ret []string
		return ret
	}
	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetRemovedOk() ([]string, bool) {
	if o == nil || IsNil(o.Removed) {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *PayloadCommit) HasRemoved() bool {
	if o != nil && !IsNil(o.Removed) {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given []string and assigns it to the Removed field.
func (o *PayloadCommit) SetRemoved(v []string) {
	o.Removed = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PayloadCommit) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PayloadCommit) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *PayloadCommit) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PayloadCommit) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PayloadCommit) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PayloadCommit) SetUrl(v string) {
	o.Url = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *PayloadCommit) GetVerification() PayloadCommitVerification {
	if o == nil || IsNil(o.Verification) {
		var ret PayloadCommitVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadCommit) GetVerificationOk() (*PayloadCommitVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *PayloadCommit) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given PayloadCommitVerification and assigns it to the Verification field.
func (o *PayloadCommit) SetVerification(v PayloadCommitVerification) {
	o.Verification = &v
}

func (o PayloadCommit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayloadCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Removed) {
		toSerialize["removed"] = o.Removed
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	return toSerialize, nil
}

type NullablePayloadCommit struct {
	value *PayloadCommit
	isSet bool
}

func (v NullablePayloadCommit) Get() *PayloadCommit {
	return v.value
}

func (v *NullablePayloadCommit) Set(val *PayloadCommit) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadCommit) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadCommit(val *PayloadCommit) *NullablePayloadCommit {
	return &NullablePayloadCommit{value: val, isSet: true}
}

func (v NullablePayloadCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
