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

// checks if the SubmitPullReviewOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitPullReviewOptions{}

// SubmitPullReviewOptions SubmitPullReviewOptions are options to submit a pending pull review
type SubmitPullReviewOptions struct {
	Body *string `json:"body,omitempty"`
	// ReviewStateType review state type
	Event *string `json:"event,omitempty"`
}

// NewSubmitPullReviewOptions instantiates a new SubmitPullReviewOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitPullReviewOptions() *SubmitPullReviewOptions {
	this := SubmitPullReviewOptions{}
	return &this
}

// NewSubmitPullReviewOptionsWithDefaults instantiates a new SubmitPullReviewOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitPullReviewOptionsWithDefaults() *SubmitPullReviewOptions {
	this := SubmitPullReviewOptions{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SubmitPullReviewOptions) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitPullReviewOptions) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SubmitPullReviewOptions) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SubmitPullReviewOptions) SetBody(v string) {
	o.Body = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *SubmitPullReviewOptions) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitPullReviewOptions) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *SubmitPullReviewOptions) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *SubmitPullReviewOptions) SetEvent(v string) {
	o.Event = &v
}

func (o SubmitPullReviewOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitPullReviewOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	return toSerialize, nil
}

type NullableSubmitPullReviewOptions struct {
	value *SubmitPullReviewOptions
	isSet bool
}

func (v NullableSubmitPullReviewOptions) Get() *SubmitPullReviewOptions {
	return v.value
}

func (v *NullableSubmitPullReviewOptions) Set(val *SubmitPullReviewOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitPullReviewOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitPullReviewOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitPullReviewOptions(val *SubmitPullReviewOptions) *NullableSubmitPullReviewOptions {
	return &NullableSubmitPullReviewOptions{value: val, isSet: true}
}

func (v NullableSubmitPullReviewOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitPullReviewOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
