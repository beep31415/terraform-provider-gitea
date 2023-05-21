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

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity struct for Activity
type Activity struct {
	ActUser   *User       `json:"act_user,omitempty"`
	ActUserId *int64      `json:"act_user_id,omitempty"`
	Comment   *Comment    `json:"comment,omitempty"`
	CommentId *int64      `json:"comment_id,omitempty"`
	Content   *string     `json:"content,omitempty"`
	Created   *time.Time  `json:"created,omitempty"`
	Id        *int64      `json:"id,omitempty"`
	IsPrivate *bool       `json:"is_private,omitempty"`
	OpType    *string     `json:"op_type,omitempty"`
	RefName   *string     `json:"ref_name,omitempty"`
	Repo      *Repository `json:"repo,omitempty"`
	RepoId    *int64      `json:"repo_id,omitempty"`
	UserId    *int64      `json:"user_id,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetActUser returns the ActUser field value if set, zero value otherwise.
func (o *Activity) GetActUser() User {
	if o == nil || IsNil(o.ActUser) {
		var ret User
		return ret
	}
	return *o.ActUser
}

// GetActUserOk returns a tuple with the ActUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActUserOk() (*User, bool) {
	if o == nil || IsNil(o.ActUser) {
		return nil, false
	}
	return o.ActUser, true
}

// HasActUser returns a boolean if a field has been set.
func (o *Activity) HasActUser() bool {
	if o != nil && !IsNil(o.ActUser) {
		return true
	}

	return false
}

// SetActUser gets a reference to the given User and assigns it to the ActUser field.
func (o *Activity) SetActUser(v User) {
	o.ActUser = &v
}

// GetActUserId returns the ActUserId field value if set, zero value otherwise.
func (o *Activity) GetActUserId() int64 {
	if o == nil || IsNil(o.ActUserId) {
		var ret int64
		return ret
	}
	return *o.ActUserId
}

// GetActUserIdOk returns a tuple with the ActUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ActUserId) {
		return nil, false
	}
	return o.ActUserId, true
}

// HasActUserId returns a boolean if a field has been set.
func (o *Activity) HasActUserId() bool {
	if o != nil && !IsNil(o.ActUserId) {
		return true
	}

	return false
}

// SetActUserId gets a reference to the given int64 and assigns it to the ActUserId field.
func (o *Activity) SetActUserId(v int64) {
	o.ActUserId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Activity) GetComment() Comment {
	if o == nil || IsNil(o.Comment) {
		var ret Comment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCommentOk() (*Comment, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Activity) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given Comment and assigns it to the Comment field.
func (o *Activity) SetComment(v Comment) {
	o.Comment = &v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *Activity) GetCommentId() int64 {
	if o == nil || IsNil(o.CommentId) {
		var ret int64
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCommentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CommentId) {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *Activity) HasCommentId() bool {
	if o != nil && !IsNil(o.CommentId) {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given int64 and assigns it to the CommentId field.
func (o *Activity) SetCommentId(v int64) {
	o.CommentId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Activity) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Activity) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Activity) SetContent(v string) {
	o.Content = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Activity) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Activity) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Activity) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Activity) SetId(v int64) {
	o.Id = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *Activity) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *Activity) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *Activity) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetOpType returns the OpType field value if set, zero value otherwise.
func (o *Activity) GetOpType() string {
	if o == nil || IsNil(o.OpType) {
		var ret string
		return ret
	}
	return *o.OpType
}

// GetOpTypeOk returns a tuple with the OpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetOpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OpType) {
		return nil, false
	}
	return o.OpType, true
}

// HasOpType returns a boolean if a field has been set.
func (o *Activity) HasOpType() bool {
	if o != nil && !IsNil(o.OpType) {
		return true
	}

	return false
}

// SetOpType gets a reference to the given string and assigns it to the OpType field.
func (o *Activity) SetOpType(v string) {
	o.OpType = &v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *Activity) GetRefName() string {
	if o == nil || IsNil(o.RefName) {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetRefNameOk() (*string, bool) {
	if o == nil || IsNil(o.RefName) {
		return nil, false
	}
	return o.RefName, true
}

// HasRefName returns a boolean if a field has been set.
func (o *Activity) HasRefName() bool {
	if o != nil && !IsNil(o.RefName) {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *Activity) SetRefName(v string) {
	o.RefName = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *Activity) GetRepo() Repository {
	if o == nil || IsNil(o.Repo) {
		var ret Repository
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetRepoOk() (*Repository, bool) {
	if o == nil || IsNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *Activity) HasRepo() bool {
	if o != nil && !IsNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given Repository and assigns it to the Repo field.
func (o *Activity) SetRepo(v Repository) {
	o.Repo = &v
}

// GetRepoId returns the RepoId field value if set, zero value otherwise.
func (o *Activity) GetRepoId() int64 {
	if o == nil || IsNil(o.RepoId) {
		var ret int64
		return ret
	}
	return *o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetRepoIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RepoId) {
		return nil, false
	}
	return o.RepoId, true
}

// HasRepoId returns a boolean if a field has been set.
func (o *Activity) HasRepoId() bool {
	if o != nil && !IsNil(o.RepoId) {
		return true
	}

	return false
}

// SetRepoId gets a reference to the given int64 and assigns it to the RepoId field.
func (o *Activity) SetRepoId(v int64) {
	o.RepoId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Activity) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Activity) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *Activity) SetUserId(v int64) {
	o.UserId = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActUser) {
		toSerialize["act_user"] = o.ActUser
	}
	if !IsNil(o.ActUserId) {
		toSerialize["act_user_id"] = o.ActUserId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CommentId) {
		toSerialize["comment_id"] = o.CommentId
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}
	if !IsNil(o.OpType) {
		toSerialize["op_type"] = o.OpType
	}
	if !IsNil(o.RefName) {
		toSerialize["ref_name"] = o.RefName
	}
	if !IsNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !IsNil(o.RepoId) {
		toSerialize["repo_id"] = o.RepoId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
