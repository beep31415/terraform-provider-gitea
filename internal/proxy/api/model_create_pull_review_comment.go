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

// checks if the CreatePullReviewComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePullReviewComment{}

// CreatePullReviewComment CreatePullReviewComment represent a review comment for creation api
type CreatePullReviewComment struct {
	Body *string `json:"body,omitempty"`
	// if comment to new file line or 0
	NewPosition *int64 `json:"new_position,omitempty"`
	// if comment to old file line or 0
	OldPosition *int64 `json:"old_position,omitempty"`
	// the tree path
	Path *string `json:"path,omitempty"`
}

// NewCreatePullReviewComment instantiates a new CreatePullReviewComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePullReviewComment() *CreatePullReviewComment {
	this := CreatePullReviewComment{}
	return &this
}

// NewCreatePullReviewCommentWithDefaults instantiates a new CreatePullReviewComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePullReviewCommentWithDefaults() *CreatePullReviewComment {
	this := CreatePullReviewComment{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreatePullReviewComment) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePullReviewComment) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreatePullReviewComment) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CreatePullReviewComment) SetBody(v string) {
	o.Body = &v
}

// GetNewPosition returns the NewPosition field value if set, zero value otherwise.
func (o *CreatePullReviewComment) GetNewPosition() int64 {
	if o == nil || IsNil(o.NewPosition) {
		var ret int64
		return ret
	}
	return *o.NewPosition
}

// GetNewPositionOk returns a tuple with the NewPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePullReviewComment) GetNewPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.NewPosition) {
		return nil, false
	}
	return o.NewPosition, true
}

// HasNewPosition returns a boolean if a field has been set.
func (o *CreatePullReviewComment) HasNewPosition() bool {
	if o != nil && !IsNil(o.NewPosition) {
		return true
	}

	return false
}

// SetNewPosition gets a reference to the given int64 and assigns it to the NewPosition field.
func (o *CreatePullReviewComment) SetNewPosition(v int64) {
	o.NewPosition = &v
}

// GetOldPosition returns the OldPosition field value if set, zero value otherwise.
func (o *CreatePullReviewComment) GetOldPosition() int64 {
	if o == nil || IsNil(o.OldPosition) {
		var ret int64
		return ret
	}
	return *o.OldPosition
}

// GetOldPositionOk returns a tuple with the OldPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePullReviewComment) GetOldPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.OldPosition) {
		return nil, false
	}
	return o.OldPosition, true
}

// HasOldPosition returns a boolean if a field has been set.
func (o *CreatePullReviewComment) HasOldPosition() bool {
	if o != nil && !IsNil(o.OldPosition) {
		return true
	}

	return false
}

// SetOldPosition gets a reference to the given int64 and assigns it to the OldPosition field.
func (o *CreatePullReviewComment) SetOldPosition(v int64) {
	o.OldPosition = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CreatePullReviewComment) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePullReviewComment) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CreatePullReviewComment) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CreatePullReviewComment) SetPath(v string) {
	o.Path = &v
}

func (o CreatePullReviewComment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePullReviewComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.NewPosition) {
		toSerialize["new_position"] = o.NewPosition
	}
	if !IsNil(o.OldPosition) {
		toSerialize["old_position"] = o.OldPosition
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableCreatePullReviewComment struct {
	value *CreatePullReviewComment
	isSet bool
}

func (v NullableCreatePullReviewComment) Get() *CreatePullReviewComment {
	return v.value
}

func (v *NullableCreatePullReviewComment) Set(val *CreatePullReviewComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePullReviewComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePullReviewComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePullReviewComment(val *CreatePullReviewComment) *NullableCreatePullReviewComment {
	return &NullableCreatePullReviewComment{value: val, isSet: true}
}

func (v NullableCreatePullReviewComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePullReviewComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
