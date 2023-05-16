/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-790-g61ad4c607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attachment{}

// Attachment Attachment a generic attachment
type Attachment struct {
	BrowserDownloadUrl *string    `json:"browser_download_url,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	DownloadCount      *int64     `json:"download_count,omitempty"`
	Id                 *int64     `json:"id,omitempty"`
	Name               *string    `json:"name,omitempty"`
	Size               *int64     `json:"size,omitempty"`
	Uuid               *string    `json:"uuid,omitempty"`
}

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment() *Attachment {
	this := Attachment{}
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	return &this
}

// GetBrowserDownloadUrl returns the BrowserDownloadUrl field value if set, zero value otherwise.
func (o *Attachment) GetBrowserDownloadUrl() string {
	if o == nil || IsNil(o.BrowserDownloadUrl) {
		var ret string
		return ret
	}
	return *o.BrowserDownloadUrl
}

// GetBrowserDownloadUrlOk returns a tuple with the BrowserDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetBrowserDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserDownloadUrl) {
		return nil, false
	}
	return o.BrowserDownloadUrl, true
}

// HasBrowserDownloadUrl returns a boolean if a field has been set.
func (o *Attachment) HasBrowserDownloadUrl() bool {
	if o != nil && !IsNil(o.BrowserDownloadUrl) {
		return true
	}

	return false
}

// SetBrowserDownloadUrl gets a reference to the given string and assigns it to the BrowserDownloadUrl field.
func (o *Attachment) SetBrowserDownloadUrl(v string) {
	o.BrowserDownloadUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Attachment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Attachment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Attachment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise.
func (o *Attachment) GetDownloadCount() int64 {
	if o == nil || IsNil(o.DownloadCount) {
		var ret int64
		return ret
	}
	return *o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetDownloadCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DownloadCount) {
		return nil, false
	}
	return o.DownloadCount, true
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *Attachment) HasDownloadCount() bool {
	if o != nil && !IsNil(o.DownloadCount) {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given int64 and assigns it to the DownloadCount field.
func (o *Attachment) SetDownloadCount(v int64) {
	o.DownloadCount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Attachment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Attachment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Attachment) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Attachment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Attachment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Attachment) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Attachment) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Attachment) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Attachment) SetSize(v int64) {
	o.Size = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Attachment) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Attachment) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Attachment) SetUuid(v string) {
	o.Uuid = &v
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrowserDownloadUrl) {
		toSerialize["browser_download_url"] = o.BrowserDownloadUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DownloadCount) {
		toSerialize["download_count"] = o.DownloadCount
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
