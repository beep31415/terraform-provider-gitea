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

// checks if the CreateTeamOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTeamOption{}

// CreateTeamOption CreateTeamOption options for creating a team
type CreateTeamOption struct {
	CanCreateOrgRepo        *bool              `json:"can_create_org_repo,omitempty"`
	Description             *string            `json:"description,omitempty"`
	IncludesAllRepositories *bool              `json:"includes_all_repositories,omitempty"`
	Name                    string             `json:"name"`
	Permission              *string            `json:"permission,omitempty"`
	Units                   []string           `json:"units,omitempty"`
	UnitsMap                *map[string]string `json:"units_map,omitempty"`
}

// NewCreateTeamOption instantiates a new CreateTeamOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTeamOption(name string) *CreateTeamOption {
	this := CreateTeamOption{}
	this.Name = name
	return &this
}

// NewCreateTeamOptionWithDefaults instantiates a new CreateTeamOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTeamOptionWithDefaults() *CreateTeamOption {
	this := CreateTeamOption{}
	return &this
}

// GetCanCreateOrgRepo returns the CanCreateOrgRepo field value if set, zero value otherwise.
func (o *CreateTeamOption) GetCanCreateOrgRepo() bool {
	if o == nil || IsNil(o.CanCreateOrgRepo) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrgRepo
}

// GetCanCreateOrgRepoOk returns a tuple with the CanCreateOrgRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetCanCreateOrgRepoOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrgRepo) {
		return nil, false
	}
	return o.CanCreateOrgRepo, true
}

// HasCanCreateOrgRepo returns a boolean if a field has been set.
func (o *CreateTeamOption) HasCanCreateOrgRepo() bool {
	if o != nil && !IsNil(o.CanCreateOrgRepo) {
		return true
	}

	return false
}

// SetCanCreateOrgRepo gets a reference to the given bool and assigns it to the CanCreateOrgRepo field.
func (o *CreateTeamOption) SetCanCreateOrgRepo(v bool) {
	o.CanCreateOrgRepo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTeamOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTeamOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTeamOption) SetDescription(v string) {
	o.Description = &v
}

// GetIncludesAllRepositories returns the IncludesAllRepositories field value if set, zero value otherwise.
func (o *CreateTeamOption) GetIncludesAllRepositories() bool {
	if o == nil || IsNil(o.IncludesAllRepositories) {
		var ret bool
		return ret
	}
	return *o.IncludesAllRepositories
}

// GetIncludesAllRepositoriesOk returns a tuple with the IncludesAllRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetIncludesAllRepositoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesAllRepositories) {
		return nil, false
	}
	return o.IncludesAllRepositories, true
}

// HasIncludesAllRepositories returns a boolean if a field has been set.
func (o *CreateTeamOption) HasIncludesAllRepositories() bool {
	if o != nil && !IsNil(o.IncludesAllRepositories) {
		return true
	}

	return false
}

// SetIncludesAllRepositories gets a reference to the given bool and assigns it to the IncludesAllRepositories field.
func (o *CreateTeamOption) SetIncludesAllRepositories(v bool) {
	o.IncludesAllRepositories = &v
}

// GetName returns the Name field value
func (o *CreateTeamOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTeamOption) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *CreateTeamOption) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *CreateTeamOption) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *CreateTeamOption) SetPermission(v string) {
	o.Permission = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CreateTeamOption) GetUnits() []string {
	if o == nil || IsNil(o.Units) {
		var ret []string
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetUnitsOk() ([]string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CreateTeamOption) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []string and assigns it to the Units field.
func (o *CreateTeamOption) SetUnits(v []string) {
	o.Units = v
}

// GetUnitsMap returns the UnitsMap field value if set, zero value otherwise.
func (o *CreateTeamOption) GetUnitsMap() map[string]string {
	if o == nil || IsNil(o.UnitsMap) {
		var ret map[string]string
		return ret
	}
	return *o.UnitsMap
}

// GetUnitsMapOk returns a tuple with the UnitsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamOption) GetUnitsMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UnitsMap) {
		return nil, false
	}
	return o.UnitsMap, true
}

// HasUnitsMap returns a boolean if a field has been set.
func (o *CreateTeamOption) HasUnitsMap() bool {
	if o != nil && !IsNil(o.UnitsMap) {
		return true
	}

	return false
}

// SetUnitsMap gets a reference to the given map[string]string and assigns it to the UnitsMap field.
func (o *CreateTeamOption) SetUnitsMap(v map[string]string) {
	o.UnitsMap = &v
}

func (o CreateTeamOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTeamOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanCreateOrgRepo) {
		toSerialize["can_create_org_repo"] = o.CanCreateOrgRepo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IncludesAllRepositories) {
		toSerialize["includes_all_repositories"] = o.IncludesAllRepositories
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.UnitsMap) {
		toSerialize["units_map"] = o.UnitsMap
	}
	return toSerialize, nil
}

type NullableCreateTeamOption struct {
	value *CreateTeamOption
	isSet bool
}

func (v NullableCreateTeamOption) Get() *CreateTeamOption {
	return v.value
}

func (v *NullableCreateTeamOption) Set(val *CreateTeamOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTeamOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTeamOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTeamOption(val *CreateTeamOption) *NullableCreateTeamOption {
	return &NullableCreateTeamOption{value: val, isSet: true}
}

func (v NullableCreateTeamOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTeamOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
