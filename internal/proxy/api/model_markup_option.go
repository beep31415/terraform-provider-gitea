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

// checks if the MarkupOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkupOption{}

// MarkupOption MarkupOption markup options
type MarkupOption struct {
	// Context to render  in: body
	Context *string `json:"Context,omitempty"`
	// File path for detecting extension in file mode  in: body
	FilePath *string `json:"FilePath,omitempty"`
	// Mode to render (comment, gfm, markdown, file)  in: body
	Mode *string `json:"Mode,omitempty"`
	// Text markup to render  in: body
	Text *string `json:"Text,omitempty"`
	// Is it a wiki page ?  in: body
	Wiki *bool `json:"Wiki,omitempty"`
}

// NewMarkupOption instantiates a new MarkupOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkupOption() *MarkupOption {
	this := MarkupOption{}
	return &this
}

// NewMarkupOptionWithDefaults instantiates a new MarkupOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkupOptionWithDefaults() *MarkupOption {
	this := MarkupOption{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MarkupOption) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkupOption) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MarkupOption) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *MarkupOption) SetContext(v string) {
	o.Context = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *MarkupOption) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkupOption) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *MarkupOption) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *MarkupOption) SetFilePath(v string) {
	o.FilePath = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MarkupOption) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkupOption) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MarkupOption) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *MarkupOption) SetMode(v string) {
	o.Mode = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MarkupOption) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkupOption) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MarkupOption) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *MarkupOption) SetText(v string) {
	o.Text = &v
}

// GetWiki returns the Wiki field value if set, zero value otherwise.
func (o *MarkupOption) GetWiki() bool {
	if o == nil || IsNil(o.Wiki) {
		var ret bool
		return ret
	}
	return *o.Wiki
}

// GetWikiOk returns a tuple with the Wiki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkupOption) GetWikiOk() (*bool, bool) {
	if o == nil || IsNil(o.Wiki) {
		return nil, false
	}
	return o.Wiki, true
}

// HasWiki returns a boolean if a field has been set.
func (o *MarkupOption) HasWiki() bool {
	if o != nil && !IsNil(o.Wiki) {
		return true
	}

	return false
}

// SetWiki gets a reference to the given bool and assigns it to the Wiki field.
func (o *MarkupOption) SetWiki(v bool) {
	o.Wiki = &v
}

func (o MarkupOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkupOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["Context"] = o.Context
	}
	if !IsNil(o.FilePath) {
		toSerialize["FilePath"] = o.FilePath
	}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	if !IsNil(o.Text) {
		toSerialize["Text"] = o.Text
	}
	if !IsNil(o.Wiki) {
		toSerialize["Wiki"] = o.Wiki
	}
	return toSerialize, nil
}

type NullableMarkupOption struct {
	value *MarkupOption
	isSet bool
}

func (v NullableMarkupOption) Get() *MarkupOption {
	return v.value
}

func (v *NullableMarkupOption) Set(val *MarkupOption) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkupOption) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkupOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkupOption(val *MarkupOption) *NullableMarkupOption {
	return &NullableMarkupOption{value: val, isSet: true}
}

func (v NullableMarkupOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkupOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}