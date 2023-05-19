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

// checks if the WatchInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchInfo{}

// WatchInfo WatchInfo represents an API watch status of one repository
type WatchInfo struct {
	CreatedAt     *time.Time             `json:"created_at,omitempty"`
	Ignored       *bool                  `json:"ignored,omitempty"`
	Reason        map[string]interface{} `json:"reason,omitempty"`
	RepositoryUrl *string                `json:"repository_url,omitempty"`
	Subscribed    *bool                  `json:"subscribed,omitempty"`
	Url           *string                `json:"url,omitempty"`
}

// NewWatchInfo instantiates a new WatchInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchInfo() *WatchInfo {
	this := WatchInfo{}
	return &this
}

// NewWatchInfoWithDefaults instantiates a new WatchInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchInfoWithDefaults() *WatchInfo {
	this := WatchInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WatchInfo) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WatchInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WatchInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *WatchInfo) GetIgnored() bool {
	if o == nil || IsNil(o.Ignored) {
		var ret bool
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetIgnoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *WatchInfo) HasIgnored() bool {
	if o != nil && !IsNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given bool and assigns it to the Ignored field.
func (o *WatchInfo) SetIgnored(v bool) {
	o.Ignored = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *WatchInfo) GetReason() map[string]interface{} {
	if o == nil || IsNil(o.Reason) {
		var ret map[string]interface{}
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetReasonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Reason) {
		return map[string]interface{}{}, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *WatchInfo) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given map[string]interface{} and assigns it to the Reason field.
func (o *WatchInfo) SetReason(v map[string]interface{}) {
	o.Reason = v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *WatchInfo) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *WatchInfo) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *WatchInfo) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetSubscribed returns the Subscribed field value if set, zero value otherwise.
func (o *WatchInfo) GetSubscribed() bool {
	if o == nil || IsNil(o.Subscribed) {
		var ret bool
		return ret
	}
	return *o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscribed) {
		return nil, false
	}
	return o.Subscribed, true
}

// HasSubscribed returns a boolean if a field has been set.
func (o *WatchInfo) HasSubscribed() bool {
	if o != nil && !IsNil(o.Subscribed) {
		return true
	}

	return false
}

// SetSubscribed gets a reference to the given bool and assigns it to the Subscribed field.
func (o *WatchInfo) SetSubscribed(v bool) {
	o.Subscribed = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WatchInfo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchInfo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WatchInfo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WatchInfo) SetUrl(v string) {
	o.Url = &v
}

func (o WatchInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Ignored) {
		toSerialize["ignored"] = o.Ignored
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.RepositoryUrl) {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	if !IsNil(o.Subscribed) {
		toSerialize["subscribed"] = o.Subscribed
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableWatchInfo struct {
	value *WatchInfo
	isSet bool
}

func (v NullableWatchInfo) Get() *WatchInfo {
	return v.value
}

func (v *NullableWatchInfo) Set(val *WatchInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchInfo(val *WatchInfo) *NullableWatchInfo {
	return &NullableWatchInfo{value: val, isSet: true}
}

func (v NullableWatchInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}