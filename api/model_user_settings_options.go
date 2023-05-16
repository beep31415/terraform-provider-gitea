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

// checks if the UserSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettingsOptions{}

// UserSettingsOptions UserSettingsOptions represents options to change user settings
type UserSettingsOptions struct {
	Description   *string `json:"description,omitempty"`
	DiffViewStyle *string `json:"diff_view_style,omitempty"`
	FullName      *string `json:"full_name,omitempty"`
	HideActivity  *bool   `json:"hide_activity,omitempty"`
	// Privacy
	HideEmail *bool   `json:"hide_email,omitempty"`
	Language  *string `json:"language,omitempty"`
	Location  *string `json:"location,omitempty"`
	Theme     *string `json:"theme,omitempty"`
	Website   *string `json:"website,omitempty"`
}

// NewUserSettingsOptions instantiates a new UserSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsOptions() *UserSettingsOptions {
	this := UserSettingsOptions{}
	return &this
}

// NewUserSettingsOptionsWithDefaults instantiates a new UserSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsOptionsWithDefaults() *UserSettingsOptions {
	this := UserSettingsOptions{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserSettingsOptions) SetDescription(v string) {
	o.Description = &v
}

// GetDiffViewStyle returns the DiffViewStyle field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetDiffViewStyle() string {
	if o == nil || IsNil(o.DiffViewStyle) {
		var ret string
		return ret
	}
	return *o.DiffViewStyle
}

// GetDiffViewStyleOk returns a tuple with the DiffViewStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetDiffViewStyleOk() (*string, bool) {
	if o == nil || IsNil(o.DiffViewStyle) {
		return nil, false
	}
	return o.DiffViewStyle, true
}

// HasDiffViewStyle returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasDiffViewStyle() bool {
	if o != nil && !IsNil(o.DiffViewStyle) {
		return true
	}

	return false
}

// SetDiffViewStyle gets a reference to the given string and assigns it to the DiffViewStyle field.
func (o *UserSettingsOptions) SetDiffViewStyle(v string) {
	o.DiffViewStyle = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UserSettingsOptions) SetFullName(v string) {
	o.FullName = &v
}

// GetHideActivity returns the HideActivity field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetHideActivity() bool {
	if o == nil || IsNil(o.HideActivity) {
		var ret bool
		return ret
	}
	return *o.HideActivity
}

// GetHideActivityOk returns a tuple with the HideActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetHideActivityOk() (*bool, bool) {
	if o == nil || IsNil(o.HideActivity) {
		return nil, false
	}
	return o.HideActivity, true
}

// HasHideActivity returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasHideActivity() bool {
	if o != nil && !IsNil(o.HideActivity) {
		return true
	}

	return false
}

// SetHideActivity gets a reference to the given bool and assigns it to the HideActivity field.
func (o *UserSettingsOptions) SetHideActivity(v bool) {
	o.HideActivity = &v
}

// GetHideEmail returns the HideEmail field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetHideEmail() bool {
	if o == nil || IsNil(o.HideEmail) {
		var ret bool
		return ret
	}
	return *o.HideEmail
}

// GetHideEmailOk returns a tuple with the HideEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetHideEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.HideEmail) {
		return nil, false
	}
	return o.HideEmail, true
}

// HasHideEmail returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasHideEmail() bool {
	if o != nil && !IsNil(o.HideEmail) {
		return true
	}

	return false
}

// SetHideEmail gets a reference to the given bool and assigns it to the HideEmail field.
func (o *UserSettingsOptions) SetHideEmail(v bool) {
	o.HideEmail = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UserSettingsOptions) SetLanguage(v string) {
	o.Language = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UserSettingsOptions) SetLocation(v string) {
	o.Location = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *UserSettingsOptions) SetTheme(v string) {
	o.Theme = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *UserSettingsOptions) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsOptions) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *UserSettingsOptions) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *UserSettingsOptions) SetWebsite(v string) {
	o.Website = &v
}

func (o UserSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DiffViewStyle) {
		toSerialize["diff_view_style"] = o.DiffViewStyle
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.HideActivity) {
		toSerialize["hide_activity"] = o.HideActivity
	}
	if !IsNil(o.HideEmail) {
		toSerialize["hide_email"] = o.HideEmail
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableUserSettingsOptions struct {
	value *UserSettingsOptions
	isSet bool
}

func (v NullableUserSettingsOptions) Get() *UserSettingsOptions {
	return v.value
}

func (v *NullableUserSettingsOptions) Set(val *UserSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsOptions(val *UserSettingsOptions) *NullableUserSettingsOptions {
	return &NullableUserSettingsOptions{value: val, isSet: true}
}

func (v NullableUserSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
