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

// checks if the NotificationSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSubject{}

// NotificationSubject NotificationSubject contains the notification subject (Issue/Pull/Commit)
type NotificationSubject struct {
	HtmlUrl              *string `json:"html_url,omitempty"`
	LatestCommentHtmlUrl *string `json:"latest_comment_html_url,omitempty"`
	LatestCommentUrl     *string `json:"latest_comment_url,omitempty"`
	// StateType issue state type
	State *string `json:"state,omitempty"`
	Title *string `json:"title,omitempty"`
	// NotifySubjectType represent type of notification subject
	Type *string `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// NewNotificationSubject instantiates a new NotificationSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSubject() *NotificationSubject {
	this := NotificationSubject{}
	return &this
}

// NewNotificationSubjectWithDefaults instantiates a new NotificationSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSubjectWithDefaults() *NotificationSubject {
	this := NotificationSubject{}
	return &this
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *NotificationSubject) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *NotificationSubject) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *NotificationSubject) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetLatestCommentHtmlUrl returns the LatestCommentHtmlUrl field value if set, zero value otherwise.
func (o *NotificationSubject) GetLatestCommentHtmlUrl() string {
	if o == nil || IsNil(o.LatestCommentHtmlUrl) {
		var ret string
		return ret
	}
	return *o.LatestCommentHtmlUrl
}

// GetLatestCommentHtmlUrlOk returns a tuple with the LatestCommentHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetLatestCommentHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LatestCommentHtmlUrl) {
		return nil, false
	}
	return o.LatestCommentHtmlUrl, true
}

// HasLatestCommentHtmlUrl returns a boolean if a field has been set.
func (o *NotificationSubject) HasLatestCommentHtmlUrl() bool {
	if o != nil && !IsNil(o.LatestCommentHtmlUrl) {
		return true
	}

	return false
}

// SetLatestCommentHtmlUrl gets a reference to the given string and assigns it to the LatestCommentHtmlUrl field.
func (o *NotificationSubject) SetLatestCommentHtmlUrl(v string) {
	o.LatestCommentHtmlUrl = &v
}

// GetLatestCommentUrl returns the LatestCommentUrl field value if set, zero value otherwise.
func (o *NotificationSubject) GetLatestCommentUrl() string {
	if o == nil || IsNil(o.LatestCommentUrl) {
		var ret string
		return ret
	}
	return *o.LatestCommentUrl
}

// GetLatestCommentUrlOk returns a tuple with the LatestCommentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetLatestCommentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LatestCommentUrl) {
		return nil, false
	}
	return o.LatestCommentUrl, true
}

// HasLatestCommentUrl returns a boolean if a field has been set.
func (o *NotificationSubject) HasLatestCommentUrl() bool {
	if o != nil && !IsNil(o.LatestCommentUrl) {
		return true
	}

	return false
}

// SetLatestCommentUrl gets a reference to the given string and assigns it to the LatestCommentUrl field.
func (o *NotificationSubject) SetLatestCommentUrl(v string) {
	o.LatestCommentUrl = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NotificationSubject) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NotificationSubject) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NotificationSubject) SetState(v string) {
	o.State = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NotificationSubject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NotificationSubject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NotificationSubject) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationSubject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationSubject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationSubject) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationSubject) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubject) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationSubject) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationSubject) SetUrl(v string) {
	o.Url = &v
}

func (o NotificationSubject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.LatestCommentHtmlUrl) {
		toSerialize["latest_comment_html_url"] = o.LatestCommentHtmlUrl
	}
	if !IsNil(o.LatestCommentUrl) {
		toSerialize["latest_comment_url"] = o.LatestCommentUrl
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableNotificationSubject struct {
	value *NotificationSubject
	isSet bool
}

func (v NullableNotificationSubject) Get() *NotificationSubject {
	return v.value
}

func (v *NullableNotificationSubject) Set(val *NotificationSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSubject(val *NotificationSubject) *NullableNotificationSubject {
	return &NullableNotificationSubject{value: val, isSet: true}
}

func (v NullableNotificationSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
