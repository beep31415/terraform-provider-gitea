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

// checks if the FileCommitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileCommitResponse{}

// FileCommitResponse struct for FileCommitResponse
type FileCommitResponse struct {
	Author    *CommitUser  `json:"author,omitempty"`
	Committer *CommitUser  `json:"committer,omitempty"`
	Created   *time.Time   `json:"created,omitempty"`
	HtmlUrl   *string      `json:"html_url,omitempty"`
	Message   *string      `json:"message,omitempty"`
	Parents   []CommitMeta `json:"parents,omitempty"`
	Sha       *string      `json:"sha,omitempty"`
	Tree      *CommitMeta  `json:"tree,omitempty"`
	Url       *string      `json:"url,omitempty"`
}

// NewFileCommitResponse instantiates a new FileCommitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCommitResponse() *FileCommitResponse {
	this := FileCommitResponse{}
	return &this
}

// NewFileCommitResponseWithDefaults instantiates a new FileCommitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCommitResponseWithDefaults() *FileCommitResponse {
	this := FileCommitResponse{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *FileCommitResponse) GetAuthor() CommitUser {
	if o == nil || IsNil(o.Author) {
		var ret CommitUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetAuthorOk() (*CommitUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *FileCommitResponse) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommitUser and assigns it to the Author field.
func (o *FileCommitResponse) SetAuthor(v CommitUser) {
	o.Author = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *FileCommitResponse) GetCommitter() CommitUser {
	if o == nil || IsNil(o.Committer) {
		var ret CommitUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetCommitterOk() (*CommitUser, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *FileCommitResponse) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given CommitUser and assigns it to the Committer field.
func (o *FileCommitResponse) SetCommitter(v CommitUser) {
	o.Committer = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FileCommitResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FileCommitResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FileCommitResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *FileCommitResponse) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *FileCommitResponse) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *FileCommitResponse) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FileCommitResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FileCommitResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FileCommitResponse) SetMessage(v string) {
	o.Message = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *FileCommitResponse) GetParents() []CommitMeta {
	if o == nil || IsNil(o.Parents) {
		var ret []CommitMeta
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetParentsOk() ([]CommitMeta, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *FileCommitResponse) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []CommitMeta and assigns it to the Parents field.
func (o *FileCommitResponse) SetParents(v []CommitMeta) {
	o.Parents = v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *FileCommitResponse) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *FileCommitResponse) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *FileCommitResponse) SetSha(v string) {
	o.Sha = &v
}

// GetTree returns the Tree field value if set, zero value otherwise.
func (o *FileCommitResponse) GetTree() CommitMeta {
	if o == nil || IsNil(o.Tree) {
		var ret CommitMeta
		return ret
	}
	return *o.Tree
}

// GetTreeOk returns a tuple with the Tree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetTreeOk() (*CommitMeta, bool) {
	if o == nil || IsNil(o.Tree) {
		return nil, false
	}
	return o.Tree, true
}

// HasTree returns a boolean if a field has been set.
func (o *FileCommitResponse) HasTree() bool {
	if o != nil && !IsNil(o.Tree) {
		return true
	}

	return false
}

// SetTree gets a reference to the given CommitMeta and assigns it to the Tree field.
func (o *FileCommitResponse) SetTree(v CommitMeta) {
	o.Tree = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FileCommitResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FileCommitResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FileCommitResponse) SetUrl(v string) {
	o.Url = &v
}

func (o FileCommitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileCommitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.Tree) {
		toSerialize["tree"] = o.Tree
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableFileCommitResponse struct {
	value *FileCommitResponse
	isSet bool
}

func (v NullableFileCommitResponse) Get() *FileCommitResponse {
	return v.value
}

func (v *NullableFileCommitResponse) Set(val *FileCommitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCommitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCommitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCommitResponse(val *FileCommitResponse) *NullableFileCommitResponse {
	return &NullableFileCommitResponse{value: val, isSet: true}
}

func (v NullableFileCommitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCommitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
