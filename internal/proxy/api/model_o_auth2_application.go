/*
Gitea API.

This documentation describes the Gitea API.

API version: 1.20.0+dev-833-g040970c32
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the OAuth2Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Application{}

// OAuth2Application struct for OAuth2Application
type OAuth2Application struct {
	ClientId           *string    `json:"client_id,omitempty"`
	ClientSecret       *string    `json:"client_secret,omitempty"`
	ConfidentialClient *bool      `json:"confidential_client,omitempty"`
	Created            *time.Time `json:"created,omitempty"`
	Id                 *int64     `json:"id,omitempty"`
	Name               *string    `json:"name,omitempty"`
	RedirectUris       []string   `json:"redirect_uris,omitempty"`
}

// NewOAuth2Application instantiates a new OAuth2Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Application() *OAuth2Application {
	this := OAuth2Application{}
	return &this
}

// NewOAuth2ApplicationWithDefaults instantiates a new OAuth2Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ApplicationWithDefaults() *OAuth2Application {
	this := OAuth2Application{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2Application) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2Application) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2Application) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuth2Application) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth2Application) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuth2Application) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetConfidentialClient returns the ConfidentialClient field value if set, zero value otherwise.
func (o *OAuth2Application) GetConfidentialClient() bool {
	if o == nil || IsNil(o.ConfidentialClient) {
		var ret bool
		return ret
	}
	return *o.ConfidentialClient
}

// GetConfidentialClientOk returns a tuple with the ConfidentialClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetConfidentialClientOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfidentialClient) {
		return nil, false
	}
	return o.ConfidentialClient, true
}

// HasConfidentialClient returns a boolean if a field has been set.
func (o *OAuth2Application) HasConfidentialClient() bool {
	if o != nil && !IsNil(o.ConfidentialClient) {
		return true
	}

	return false
}

// SetConfidentialClient gets a reference to the given bool and assigns it to the ConfidentialClient field.
func (o *OAuth2Application) SetConfidentialClient(v bool) {
	o.ConfidentialClient = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OAuth2Application) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OAuth2Application) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OAuth2Application) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2Application) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2Application) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OAuth2Application) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuth2Application) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuth2Application) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuth2Application) SetName(v string) {
	o.Name = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *OAuth2Application) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Application) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OAuth2Application) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *OAuth2Application) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

func (o OAuth2Application) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.ConfidentialClient) {
		toSerialize["confidential_client"] = o.ConfidentialClient
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	return toSerialize, nil
}

type NullableOAuth2Application struct {
	value *OAuth2Application
	isSet bool
}

func (v NullableOAuth2Application) Get() *OAuth2Application {
	return v.value
}

func (v *NullableOAuth2Application) Set(val *OAuth2Application) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Application) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Application) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Application(val *OAuth2Application) *NullableOAuth2Application {
	return &NullableOAuth2Application{value: val, isSet: true}
}

func (v NullableOAuth2Application) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Application) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}