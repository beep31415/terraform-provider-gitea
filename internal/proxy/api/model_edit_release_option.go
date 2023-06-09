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

// checks if the EditReleaseOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditReleaseOption{}

// EditReleaseOption EditReleaseOption options when editing a release
type EditReleaseOption struct {
	Body            *string `json:"body,omitempty"`
	Draft           *bool   `json:"draft,omitempty"`
	Name            *string `json:"name,omitempty"`
	Prerelease      *bool   `json:"prerelease,omitempty"`
	TagName         *string `json:"tag_name,omitempty"`
	TargetCommitish *string `json:"target_commitish,omitempty"`
}

// NewEditReleaseOption instantiates a new EditReleaseOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditReleaseOption() *EditReleaseOption {
	this := EditReleaseOption{}
	return &this
}

// NewEditReleaseOptionWithDefaults instantiates a new EditReleaseOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditReleaseOptionWithDefaults() *EditReleaseOption {
	this := EditReleaseOption{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *EditReleaseOption) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *EditReleaseOption) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *EditReleaseOption) SetBody(v string) {
	o.Body = &v
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *EditReleaseOption) GetDraft() bool {
	if o == nil || IsNil(o.Draft) {
		var ret bool
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.Draft) {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *EditReleaseOption) HasDraft() bool {
	if o != nil && !IsNil(o.Draft) {
		return true
	}

	return false
}

// SetDraft gets a reference to the given bool and assigns it to the Draft field.
func (o *EditReleaseOption) SetDraft(v bool) {
	o.Draft = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditReleaseOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditReleaseOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditReleaseOption) SetName(v string) {
	o.Name = &v
}

// GetPrerelease returns the Prerelease field value if set, zero value otherwise.
func (o *EditReleaseOption) GetPrerelease() bool {
	if o == nil || IsNil(o.Prerelease) {
		var ret bool
		return ret
	}
	return *o.Prerelease
}

// GetPrereleaseOk returns a tuple with the Prerelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetPrereleaseOk() (*bool, bool) {
	if o == nil || IsNil(o.Prerelease) {
		return nil, false
	}
	return o.Prerelease, true
}

// HasPrerelease returns a boolean if a field has been set.
func (o *EditReleaseOption) HasPrerelease() bool {
	if o != nil && !IsNil(o.Prerelease) {
		return true
	}

	return false
}

// SetPrerelease gets a reference to the given bool and assigns it to the Prerelease field.
func (o *EditReleaseOption) SetPrerelease(v bool) {
	o.Prerelease = &v
}

// GetTagName returns the TagName field value if set, zero value otherwise.
func (o *EditReleaseOption) GetTagName() string {
	if o == nil || IsNil(o.TagName) {
		var ret string
		return ret
	}
	return *o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetTagNameOk() (*string, bool) {
	if o == nil || IsNil(o.TagName) {
		return nil, false
	}
	return o.TagName, true
}

// HasTagName returns a boolean if a field has been set.
func (o *EditReleaseOption) HasTagName() bool {
	if o != nil && !IsNil(o.TagName) {
		return true
	}

	return false
}

// SetTagName gets a reference to the given string and assigns it to the TagName field.
func (o *EditReleaseOption) SetTagName(v string) {
	o.TagName = &v
}

// GetTargetCommitish returns the TargetCommitish field value if set, zero value otherwise.
func (o *EditReleaseOption) GetTargetCommitish() string {
	if o == nil || IsNil(o.TargetCommitish) {
		var ret string
		return ret
	}
	return *o.TargetCommitish
}

// GetTargetCommitishOk returns a tuple with the TargetCommitish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditReleaseOption) GetTargetCommitishOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCommitish) {
		return nil, false
	}
	return o.TargetCommitish, true
}

// HasTargetCommitish returns a boolean if a field has been set.
func (o *EditReleaseOption) HasTargetCommitish() bool {
	if o != nil && !IsNil(o.TargetCommitish) {
		return true
	}

	return false
}

// SetTargetCommitish gets a reference to the given string and assigns it to the TargetCommitish field.
func (o *EditReleaseOption) SetTargetCommitish(v string) {
	o.TargetCommitish = &v
}

func (o EditReleaseOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditReleaseOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Draft) {
		toSerialize["draft"] = o.Draft
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Prerelease) {
		toSerialize["prerelease"] = o.Prerelease
	}
	if !IsNil(o.TagName) {
		toSerialize["tag_name"] = o.TagName
	}
	if !IsNil(o.TargetCommitish) {
		toSerialize["target_commitish"] = o.TargetCommitish
	}
	return toSerialize, nil
}

type NullableEditReleaseOption struct {
	value *EditReleaseOption
	isSet bool
}

func (v NullableEditReleaseOption) Get() *EditReleaseOption {
	return v.value
}

func (v *NullableEditReleaseOption) Set(val *EditReleaseOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEditReleaseOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEditReleaseOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditReleaseOption(val *EditReleaseOption) *NullableEditReleaseOption {
	return &NullableEditReleaseOption{value: val, isSet: true}
}

func (v NullableEditReleaseOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditReleaseOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
