/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-790-g61ad4c607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AnnotatedTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotatedTag{}

// AnnotatedTag AnnotatedTag represents an annotated tag
type AnnotatedTag struct {
	Message      *string                    `json:"message,omitempty"`
	Object       *AnnotatedTagObject        `json:"object,omitempty"`
	Sha          *string                    `json:"sha,omitempty"`
	Tag          *string                    `json:"tag,omitempty"`
	Tagger       *CommitUser                `json:"tagger,omitempty"`
	Url          *string                    `json:"url,omitempty"`
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// NewAnnotatedTag instantiates a new AnnotatedTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotatedTag() *AnnotatedTag {
	this := AnnotatedTag{}
	return &this
}

// NewAnnotatedTagWithDefaults instantiates a new AnnotatedTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotatedTagWithDefaults() *AnnotatedTag {
	this := AnnotatedTag{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AnnotatedTag) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AnnotatedTag) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AnnotatedTag) SetMessage(v string) {
	o.Message = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AnnotatedTag) GetObject() AnnotatedTagObject {
	if o == nil || IsNil(o.Object) {
		var ret AnnotatedTagObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetObjectOk() (*AnnotatedTagObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AnnotatedTag) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given AnnotatedTagObject and assigns it to the Object field.
func (o *AnnotatedTag) SetObject(v AnnotatedTagObject) {
	o.Object = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *AnnotatedTag) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *AnnotatedTag) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *AnnotatedTag) SetSha(v string) {
	o.Sha = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AnnotatedTag) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AnnotatedTag) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AnnotatedTag) SetTag(v string) {
	o.Tag = &v
}

// GetTagger returns the Tagger field value if set, zero value otherwise.
func (o *AnnotatedTag) GetTagger() CommitUser {
	if o == nil || IsNil(o.Tagger) {
		var ret CommitUser
		return ret
	}
	return *o.Tagger
}

// GetTaggerOk returns a tuple with the Tagger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetTaggerOk() (*CommitUser, bool) {
	if o == nil || IsNil(o.Tagger) {
		return nil, false
	}
	return o.Tagger, true
}

// HasTagger returns a boolean if a field has been set.
func (o *AnnotatedTag) HasTagger() bool {
	if o != nil && !IsNil(o.Tagger) {
		return true
	}

	return false
}

// SetTagger gets a reference to the given CommitUser and assigns it to the Tagger field.
func (o *AnnotatedTag) SetTagger(v CommitUser) {
	o.Tagger = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AnnotatedTag) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AnnotatedTag) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AnnotatedTag) SetUrl(v string) {
	o.Url = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *AnnotatedTag) GetVerification() PayloadCommitVerification {
	if o == nil || IsNil(o.Verification) {
		var ret PayloadCommitVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedTag) GetVerificationOk() (*PayloadCommitVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *AnnotatedTag) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given PayloadCommitVerification and assigns it to the Verification field.
func (o *AnnotatedTag) SetVerification(v PayloadCommitVerification) {
	o.Verification = &v
}

func (o AnnotatedTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotatedTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Tagger) {
		toSerialize["tagger"] = o.Tagger
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	return toSerialize, nil
}

type NullableAnnotatedTag struct {
	value *AnnotatedTag
	isSet bool
}

func (v NullableAnnotatedTag) Get() *AnnotatedTag {
	return v.value
}

func (v *NullableAnnotatedTag) Set(val *AnnotatedTag) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotatedTag) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotatedTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotatedTag(val *AnnotatedTag) *NullableAnnotatedTag {
	return &NullableAnnotatedTag{value: val, isSet: true}
}

func (v NullableAnnotatedTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotatedTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
