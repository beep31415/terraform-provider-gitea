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

// checks if the NodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInfo{}

// NodeInfo NodeInfo contains standardized way of exposing metadata about a server running one of the distributed social networks
type NodeInfo struct {
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	OpenRegistrations *bool                  `json:"openRegistrations,omitempty"`
	Protocols         []string               `json:"protocols,omitempty"`
	Services          *NodeInfoServices      `json:"services,omitempty"`
	Software          *NodeInfoSoftware      `json:"software,omitempty"`
	Usage             *NodeInfoUsage         `json:"usage,omitempty"`
	Version           *string                `json:"version,omitempty"`
}

// NewNodeInfo instantiates a new NodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInfo() *NodeInfo {
	this := NodeInfo{}
	return &this
}

// NewNodeInfoWithDefaults instantiates a new NodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInfoWithDefaults() *NodeInfo {
	this := NodeInfo{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NodeInfo) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NodeInfo) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NodeInfo) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetOpenRegistrations returns the OpenRegistrations field value if set, zero value otherwise.
func (o *NodeInfo) GetOpenRegistrations() bool {
	if o == nil || IsNil(o.OpenRegistrations) {
		var ret bool
		return ret
	}
	return *o.OpenRegistrations
}

// GetOpenRegistrationsOk returns a tuple with the OpenRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetOpenRegistrationsOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenRegistrations) {
		return nil, false
	}
	return o.OpenRegistrations, true
}

// HasOpenRegistrations returns a boolean if a field has been set.
func (o *NodeInfo) HasOpenRegistrations() bool {
	if o != nil && !IsNil(o.OpenRegistrations) {
		return true
	}

	return false
}

// SetOpenRegistrations gets a reference to the given bool and assigns it to the OpenRegistrations field.
func (o *NodeInfo) SetOpenRegistrations(v bool) {
	o.OpenRegistrations = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *NodeInfo) GetProtocols() []string {
	if o == nil || IsNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Protocols) {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *NodeInfo) HasProtocols() bool {
	if o != nil && !IsNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *NodeInfo) SetProtocols(v []string) {
	o.Protocols = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *NodeInfo) GetServices() NodeInfoServices {
	if o == nil || IsNil(o.Services) {
		var ret NodeInfoServices
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetServicesOk() (*NodeInfoServices, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *NodeInfo) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given NodeInfoServices and assigns it to the Services field.
func (o *NodeInfo) SetServices(v NodeInfoServices) {
	o.Services = &v
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *NodeInfo) GetSoftware() NodeInfoSoftware {
	if o == nil || IsNil(o.Software) {
		var ret NodeInfoSoftware
		return ret
	}
	return *o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetSoftwareOk() (*NodeInfoSoftware, bool) {
	if o == nil || IsNil(o.Software) {
		return nil, false
	}
	return o.Software, true
}

// HasSoftware returns a boolean if a field has been set.
func (o *NodeInfo) HasSoftware() bool {
	if o != nil && !IsNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given NodeInfoSoftware and assigns it to the Software field.
func (o *NodeInfo) SetSoftware(v NodeInfoSoftware) {
	o.Software = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *NodeInfo) GetUsage() NodeInfoUsage {
	if o == nil || IsNil(o.Usage) {
		var ret NodeInfoUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetUsageOk() (*NodeInfoUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *NodeInfo) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given NodeInfoUsage and assigns it to the Usage field.
func (o *NodeInfo) SetUsage(v NodeInfoUsage) {
	o.Usage = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NodeInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NodeInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NodeInfo) SetVersion(v string) {
	o.Version = &v
}

func (o NodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.OpenRegistrations) {
		toSerialize["openRegistrations"] = o.OpenRegistrations
	}
	if !IsNil(o.Protocols) {
		toSerialize["protocols"] = o.Protocols
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableNodeInfo struct {
	value *NodeInfo
	isSet bool
}

func (v NullableNodeInfo) Get() *NodeInfo {
	return v.value
}

func (v *NullableNodeInfo) Set(val *NodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInfo(val *NodeInfo) *NullableNodeInfo {
	return &NullableNodeInfo{value: val, isSet: true}
}

func (v NullableNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}