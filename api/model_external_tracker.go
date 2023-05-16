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

// checks if the ExternalTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalTracker{}

// ExternalTracker ExternalTracker represents settings for external tracker
type ExternalTracker struct {
	// External Issue Tracker URL Format. Use the placeholders {user}, {repo} and {index} for the username, repository name and issue index.
	ExternalTrackerFormat *string `json:"external_tracker_format,omitempty"`
	// External Issue Tracker issue regular expression
	ExternalTrackerRegexpPattern *string `json:"external_tracker_regexp_pattern,omitempty"`
	// External Issue Tracker Number Format, either `numeric`, `alphanumeric`, or `regexp`
	ExternalTrackerStyle *string `json:"external_tracker_style,omitempty"`
	// URL of external issue tracker.
	ExternalTrackerUrl *string `json:"external_tracker_url,omitempty"`
}

// NewExternalTracker instantiates a new ExternalTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTracker() *ExternalTracker {
	this := ExternalTracker{}
	return &this
}

// NewExternalTrackerWithDefaults instantiates a new ExternalTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTrackerWithDefaults() *ExternalTracker {
	this := ExternalTracker{}
	return &this
}

// GetExternalTrackerFormat returns the ExternalTrackerFormat field value if set, zero value otherwise.
func (o *ExternalTracker) GetExternalTrackerFormat() string {
	if o == nil || IsNil(o.ExternalTrackerFormat) {
		var ret string
		return ret
	}
	return *o.ExternalTrackerFormat
}

// GetExternalTrackerFormatOk returns a tuple with the ExternalTrackerFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTracker) GetExternalTrackerFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTrackerFormat) {
		return nil, false
	}
	return o.ExternalTrackerFormat, true
}

// HasExternalTrackerFormat returns a boolean if a field has been set.
func (o *ExternalTracker) HasExternalTrackerFormat() bool {
	if o != nil && !IsNil(o.ExternalTrackerFormat) {
		return true
	}

	return false
}

// SetExternalTrackerFormat gets a reference to the given string and assigns it to the ExternalTrackerFormat field.
func (o *ExternalTracker) SetExternalTrackerFormat(v string) {
	o.ExternalTrackerFormat = &v
}

// GetExternalTrackerRegexpPattern returns the ExternalTrackerRegexpPattern field value if set, zero value otherwise.
func (o *ExternalTracker) GetExternalTrackerRegexpPattern() string {
	if o == nil || IsNil(o.ExternalTrackerRegexpPattern) {
		var ret string
		return ret
	}
	return *o.ExternalTrackerRegexpPattern
}

// GetExternalTrackerRegexpPatternOk returns a tuple with the ExternalTrackerRegexpPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTracker) GetExternalTrackerRegexpPatternOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTrackerRegexpPattern) {
		return nil, false
	}
	return o.ExternalTrackerRegexpPattern, true
}

// HasExternalTrackerRegexpPattern returns a boolean if a field has been set.
func (o *ExternalTracker) HasExternalTrackerRegexpPattern() bool {
	if o != nil && !IsNil(o.ExternalTrackerRegexpPattern) {
		return true
	}

	return false
}

// SetExternalTrackerRegexpPattern gets a reference to the given string and assigns it to the ExternalTrackerRegexpPattern field.
func (o *ExternalTracker) SetExternalTrackerRegexpPattern(v string) {
	o.ExternalTrackerRegexpPattern = &v
}

// GetExternalTrackerStyle returns the ExternalTrackerStyle field value if set, zero value otherwise.
func (o *ExternalTracker) GetExternalTrackerStyle() string {
	if o == nil || IsNil(o.ExternalTrackerStyle) {
		var ret string
		return ret
	}
	return *o.ExternalTrackerStyle
}

// GetExternalTrackerStyleOk returns a tuple with the ExternalTrackerStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTracker) GetExternalTrackerStyleOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTrackerStyle) {
		return nil, false
	}
	return o.ExternalTrackerStyle, true
}

// HasExternalTrackerStyle returns a boolean if a field has been set.
func (o *ExternalTracker) HasExternalTrackerStyle() bool {
	if o != nil && !IsNil(o.ExternalTrackerStyle) {
		return true
	}

	return false
}

// SetExternalTrackerStyle gets a reference to the given string and assigns it to the ExternalTrackerStyle field.
func (o *ExternalTracker) SetExternalTrackerStyle(v string) {
	o.ExternalTrackerStyle = &v
}

// GetExternalTrackerUrl returns the ExternalTrackerUrl field value if set, zero value otherwise.
func (o *ExternalTracker) GetExternalTrackerUrl() string {
	if o == nil || IsNil(o.ExternalTrackerUrl) {
		var ret string
		return ret
	}
	return *o.ExternalTrackerUrl
}

// GetExternalTrackerUrlOk returns a tuple with the ExternalTrackerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTracker) GetExternalTrackerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTrackerUrl) {
		return nil, false
	}
	return o.ExternalTrackerUrl, true
}

// HasExternalTrackerUrl returns a boolean if a field has been set.
func (o *ExternalTracker) HasExternalTrackerUrl() bool {
	if o != nil && !IsNil(o.ExternalTrackerUrl) {
		return true
	}

	return false
}

// SetExternalTrackerUrl gets a reference to the given string and assigns it to the ExternalTrackerUrl field.
func (o *ExternalTracker) SetExternalTrackerUrl(v string) {
	o.ExternalTrackerUrl = &v
}

func (o ExternalTracker) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalTrackerFormat) {
		toSerialize["external_tracker_format"] = o.ExternalTrackerFormat
	}
	if !IsNil(o.ExternalTrackerRegexpPattern) {
		toSerialize["external_tracker_regexp_pattern"] = o.ExternalTrackerRegexpPattern
	}
	if !IsNil(o.ExternalTrackerStyle) {
		toSerialize["external_tracker_style"] = o.ExternalTrackerStyle
	}
	if !IsNil(o.ExternalTrackerUrl) {
		toSerialize["external_tracker_url"] = o.ExternalTrackerUrl
	}
	return toSerialize, nil
}

type NullableExternalTracker struct {
	value *ExternalTracker
	isSet bool
}

func (v NullableExternalTracker) Get() *ExternalTracker {
	return v.value
}

func (v *NullableExternalTracker) Set(val *ExternalTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTracker(val *ExternalTracker) *NullableExternalTracker {
	return &NullableExternalTracker{value: val, isSet: true}
}

func (v NullableExternalTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
