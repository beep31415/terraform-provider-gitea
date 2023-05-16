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

// checks if the RepoCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepoCommit{}

// RepoCommit struct for RepoCommit
type RepoCommit struct {
	Author       *CommitUser                `json:"author,omitempty"`
	Committer    *CommitUser                `json:"committer,omitempty"`
	Message      *string                    `json:"message,omitempty"`
	Tree         *CommitMeta                `json:"tree,omitempty"`
	Url          *string                    `json:"url,omitempty"`
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// NewRepoCommit instantiates a new RepoCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepoCommit() *RepoCommit {
	this := RepoCommit{}
	return &this
}

// NewRepoCommitWithDefaults instantiates a new RepoCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepoCommitWithDefaults() *RepoCommit {
	this := RepoCommit{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *RepoCommit) GetAuthor() CommitUser {
	if o == nil || IsNil(o.Author) {
		var ret CommitUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetAuthorOk() (*CommitUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *RepoCommit) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommitUser and assigns it to the Author field.
func (o *RepoCommit) SetAuthor(v CommitUser) {
	o.Author = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *RepoCommit) GetCommitter() CommitUser {
	if o == nil || IsNil(o.Committer) {
		var ret CommitUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetCommitterOk() (*CommitUser, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *RepoCommit) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given CommitUser and assigns it to the Committer field.
func (o *RepoCommit) SetCommitter(v CommitUser) {
	o.Committer = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RepoCommit) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RepoCommit) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RepoCommit) SetMessage(v string) {
	o.Message = &v
}

// GetTree returns the Tree field value if set, zero value otherwise.
func (o *RepoCommit) GetTree() CommitMeta {
	if o == nil || IsNil(o.Tree) {
		var ret CommitMeta
		return ret
	}
	return *o.Tree
}

// GetTreeOk returns a tuple with the Tree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetTreeOk() (*CommitMeta, bool) {
	if o == nil || IsNil(o.Tree) {
		return nil, false
	}
	return o.Tree, true
}

// HasTree returns a boolean if a field has been set.
func (o *RepoCommit) HasTree() bool {
	if o != nil && !IsNil(o.Tree) {
		return true
	}

	return false
}

// SetTree gets a reference to the given CommitMeta and assigns it to the Tree field.
func (o *RepoCommit) SetTree(v CommitMeta) {
	o.Tree = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RepoCommit) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RepoCommit) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RepoCommit) SetUrl(v string) {
	o.Url = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *RepoCommit) GetVerification() PayloadCommitVerification {
	if o == nil || IsNil(o.Verification) {
		var ret PayloadCommitVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepoCommit) GetVerificationOk() (*PayloadCommitVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *RepoCommit) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given PayloadCommitVerification and assigns it to the Verification field.
func (o *RepoCommit) SetVerification(v PayloadCommitVerification) {
	o.Verification = &v
}

func (o RepoCommit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepoCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Tree) {
		toSerialize["tree"] = o.Tree
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	return toSerialize, nil
}

type NullableRepoCommit struct {
	value *RepoCommit
	isSet bool
}

func (v NullableRepoCommit) Get() *RepoCommit {
	return v.value
}

func (v *NullableRepoCommit) Set(val *RepoCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableRepoCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableRepoCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepoCommit(val *RepoCommit) *NullableRepoCommit {
	return &NullableRepoCommit{value: val, isSet: true}
}

func (v NullableRepoCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepoCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
