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

// checks if the Commit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Commit{}

// Commit struct for Commit
type Commit struct {
	Author    *User                 `json:"author,omitempty"`
	Commit    *RepoCommit           `json:"commit,omitempty"`
	Committer *User                 `json:"committer,omitempty"`
	Created   *time.Time            `json:"created,omitempty"`
	Files     []CommitAffectedFiles `json:"files,omitempty"`
	HtmlUrl   *string               `json:"html_url,omitempty"`
	Parents   []CommitMeta          `json:"parents,omitempty"`
	Sha       *string               `json:"sha,omitempty"`
	Stats     *CommitStats          `json:"stats,omitempty"`
	Url       *string               `json:"url,omitempty"`
}

// NewCommit instantiates a new Commit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommit() *Commit {
	this := Commit{}
	return &this
}

// NewCommitWithDefaults instantiates a new Commit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitWithDefaults() *Commit {
	this := Commit{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Commit) GetAuthor() User {
	if o == nil || IsNil(o.Author) {
		var ret User
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetAuthorOk() (*User, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Commit) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given User and assigns it to the Author field.
func (o *Commit) SetAuthor(v User) {
	o.Author = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *Commit) GetCommit() RepoCommit {
	if o == nil || IsNil(o.Commit) {
		var ret RepoCommit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetCommitOk() (*RepoCommit, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *Commit) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given RepoCommit and assigns it to the Commit field.
func (o *Commit) SetCommit(v RepoCommit) {
	o.Commit = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *Commit) GetCommitter() User {
	if o == nil || IsNil(o.Committer) {
		var ret User
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetCommitterOk() (*User, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *Commit) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given User and assigns it to the Committer field.
func (o *Commit) SetCommitter(v User) {
	o.Committer = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Commit) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Commit) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Commit) SetCreated(v time.Time) {
	o.Created = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Commit) GetFiles() []CommitAffectedFiles {
	if o == nil || IsNil(o.Files) {
		var ret []CommitAffectedFiles
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetFilesOk() ([]CommitAffectedFiles, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Commit) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []CommitAffectedFiles and assigns it to the Files field.
func (o *Commit) SetFiles(v []CommitAffectedFiles) {
	o.Files = v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *Commit) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *Commit) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *Commit) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *Commit) GetParents() []CommitMeta {
	if o == nil || IsNil(o.Parents) {
		var ret []CommitMeta
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetParentsOk() ([]CommitMeta, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *Commit) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []CommitMeta and assigns it to the Parents field.
func (o *Commit) SetParents(v []CommitMeta) {
	o.Parents = v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *Commit) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *Commit) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *Commit) SetSha(v string) {
	o.Sha = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Commit) GetStats() CommitStats {
	if o == nil || IsNil(o.Stats) {
		var ret CommitStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetStatsOk() (*CommitStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Commit) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given CommitStats and assigns it to the Stats field.
func (o *Commit) SetStats(v CommitStats) {
	o.Stats = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Commit) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Commit) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Commit) SetUrl(v string) {
	o.Url = &v
}

func (o Commit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Commit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCommit struct {
	value *Commit
	isSet bool
}

func (v NullableCommit) Get() *Commit {
	return v.value
}

func (v *NullableCommit) Set(val *Commit) {
	v.value = val
	v.isSet = true
}

func (v NullableCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommit(val *Commit) *NullableCommit {
	return &NullableCommit{value: val, isSet: true}
}

func (v NullableCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
