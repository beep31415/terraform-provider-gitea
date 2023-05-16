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

// checks if the MigrateRepoOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateRepoOptions{}

// MigrateRepoOptions MigrateRepoOptions options for migrating repository's this is used to interact with api v1
type MigrateRepoOptions struct {
	AuthPassword   *string `json:"auth_password,omitempty"`
	AuthToken      *string `json:"auth_token,omitempty"`
	AuthUsername   *string `json:"auth_username,omitempty"`
	CloneAddr      string  `json:"clone_addr"`
	Description    *string `json:"description,omitempty"`
	Issues         *bool   `json:"issues,omitempty"`
	Labels         *bool   `json:"labels,omitempty"`
	Lfs            *bool   `json:"lfs,omitempty"`
	LfsEndpoint    *string `json:"lfs_endpoint,omitempty"`
	Milestones     *bool   `json:"milestones,omitempty"`
	Mirror         *bool   `json:"mirror,omitempty"`
	MirrorInterval *string `json:"mirror_interval,omitempty"`
	Private        *bool   `json:"private,omitempty"`
	PullRequests   *bool   `json:"pull_requests,omitempty"`
	Releases       *bool   `json:"releases,omitempty"`
	RepoName       string  `json:"repo_name"`
	// Name of User or Organisation who will own Repo after migration
	RepoOwner *string `json:"repo_owner,omitempty"`
	Service   *string `json:"service,omitempty"`
	// deprecated (only for backwards compatibility)
	Uid  *int64 `json:"uid,omitempty"`
	Wiki *bool  `json:"wiki,omitempty"`
}

// NewMigrateRepoOptions instantiates a new MigrateRepoOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateRepoOptions(cloneAddr string, repoName string) *MigrateRepoOptions {
	this := MigrateRepoOptions{}
	this.CloneAddr = cloneAddr
	this.RepoName = repoName
	return &this
}

// NewMigrateRepoOptionsWithDefaults instantiates a new MigrateRepoOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateRepoOptionsWithDefaults() *MigrateRepoOptions {
	this := MigrateRepoOptions{}
	return &this
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetAuthPassword() string {
	if o == nil || IsNil(o.AuthPassword) {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPassword) {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasAuthPassword() bool {
	if o != nil && !IsNil(o.AuthPassword) {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *MigrateRepoOptions) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetAuthToken() string {
	if o == nil || IsNil(o.AuthToken) {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AuthToken) {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasAuthToken() bool {
	if o != nil && !IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *MigrateRepoOptions) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername) {
		var ret string
		return ret
	}
	return *o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUsername) {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasAuthUsername() bool {
	if o != nil && !IsNil(o.AuthUsername) {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *MigrateRepoOptions) SetAuthUsername(v string) {
	o.AuthUsername = &v
}

// GetCloneAddr returns the CloneAddr field value
func (o *MigrateRepoOptions) GetCloneAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloneAddr
}

// GetCloneAddrOk returns a tuple with the CloneAddr field value
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetCloneAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloneAddr, true
}

// SetCloneAddr sets field value
func (o *MigrateRepoOptions) SetCloneAddr(v string) {
	o.CloneAddr = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MigrateRepoOptions) SetDescription(v string) {
	o.Description = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetIssues() bool {
	if o == nil || IsNil(o.Issues) {
		var ret bool
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given bool and assigns it to the Issues field.
func (o *MigrateRepoOptions) SetIssues(v bool) {
	o.Issues = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetLabels() bool {
	if o == nil || IsNil(o.Labels) {
		var ret bool
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetLabelsOk() (*bool, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given bool and assigns it to the Labels field.
func (o *MigrateRepoOptions) SetLabels(v bool) {
	o.Labels = &v
}

// GetLfs returns the Lfs field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetLfs() bool {
	if o == nil || IsNil(o.Lfs) {
		var ret bool
		return ret
	}
	return *o.Lfs
}

// GetLfsOk returns a tuple with the Lfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetLfsOk() (*bool, bool) {
	if o == nil || IsNil(o.Lfs) {
		return nil, false
	}
	return o.Lfs, true
}

// HasLfs returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasLfs() bool {
	if o != nil && !IsNil(o.Lfs) {
		return true
	}

	return false
}

// SetLfs gets a reference to the given bool and assigns it to the Lfs field.
func (o *MigrateRepoOptions) SetLfs(v bool) {
	o.Lfs = &v
}

// GetLfsEndpoint returns the LfsEndpoint field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetLfsEndpoint() string {
	if o == nil || IsNil(o.LfsEndpoint) {
		var ret string
		return ret
	}
	return *o.LfsEndpoint
}

// GetLfsEndpointOk returns a tuple with the LfsEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetLfsEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.LfsEndpoint) {
		return nil, false
	}
	return o.LfsEndpoint, true
}

// HasLfsEndpoint returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasLfsEndpoint() bool {
	if o != nil && !IsNil(o.LfsEndpoint) {
		return true
	}

	return false
}

// SetLfsEndpoint gets a reference to the given string and assigns it to the LfsEndpoint field.
func (o *MigrateRepoOptions) SetLfsEndpoint(v string) {
	o.LfsEndpoint = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetMilestones() bool {
	if o == nil || IsNil(o.Milestones) {
		var ret bool
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetMilestonesOk() (*bool, bool) {
	if o == nil || IsNil(o.Milestones) {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasMilestones() bool {
	if o != nil && !IsNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given bool and assigns it to the Milestones field.
func (o *MigrateRepoOptions) SetMilestones(v bool) {
	o.Milestones = &v
}

// GetMirror returns the Mirror field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetMirror() bool {
	if o == nil || IsNil(o.Mirror) {
		var ret bool
		return ret
	}
	return *o.Mirror
}

// GetMirrorOk returns a tuple with the Mirror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetMirrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Mirror) {
		return nil, false
	}
	return o.Mirror, true
}

// HasMirror returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasMirror() bool {
	if o != nil && !IsNil(o.Mirror) {
		return true
	}

	return false
}

// SetMirror gets a reference to the given bool and assigns it to the Mirror field.
func (o *MigrateRepoOptions) SetMirror(v bool) {
	o.Mirror = &v
}

// GetMirrorInterval returns the MirrorInterval field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetMirrorInterval() string {
	if o == nil || IsNil(o.MirrorInterval) {
		var ret string
		return ret
	}
	return *o.MirrorInterval
}

// GetMirrorIntervalOk returns a tuple with the MirrorInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetMirrorIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MirrorInterval) {
		return nil, false
	}
	return o.MirrorInterval, true
}

// HasMirrorInterval returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasMirrorInterval() bool {
	if o != nil && !IsNil(o.MirrorInterval) {
		return true
	}

	return false
}

// SetMirrorInterval gets a reference to the given string and assigns it to the MirrorInterval field.
func (o *MigrateRepoOptions) SetMirrorInterval(v string) {
	o.MirrorInterval = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *MigrateRepoOptions) SetPrivate(v bool) {
	o.Private = &v
}

// GetPullRequests returns the PullRequests field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetPullRequests() bool {
	if o == nil || IsNil(o.PullRequests) {
		var ret bool
		return ret
	}
	return *o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetPullRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.PullRequests) {
		return nil, false
	}
	return o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasPullRequests() bool {
	if o != nil && !IsNil(o.PullRequests) {
		return true
	}

	return false
}

// SetPullRequests gets a reference to the given bool and assigns it to the PullRequests field.
func (o *MigrateRepoOptions) SetPullRequests(v bool) {
	o.PullRequests = &v
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetReleases() bool {
	if o == nil || IsNil(o.Releases) {
		var ret bool
		return ret
	}
	return *o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetReleasesOk() (*bool, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given bool and assigns it to the Releases field.
func (o *MigrateRepoOptions) SetReleases(v bool) {
	o.Releases = &v
}

// GetRepoName returns the RepoName field value
func (o *MigrateRepoOptions) GetRepoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepoName
}

// GetRepoNameOk returns a tuple with the RepoName field value
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoName, true
}

// SetRepoName sets field value
func (o *MigrateRepoOptions) SetRepoName(v string) {
	o.RepoName = v
}

// GetRepoOwner returns the RepoOwner field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetRepoOwner() string {
	if o == nil || IsNil(o.RepoOwner) {
		var ret string
		return ret
	}
	return *o.RepoOwner
}

// GetRepoOwnerOk returns a tuple with the RepoOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetRepoOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.RepoOwner) {
		return nil, false
	}
	return o.RepoOwner, true
}

// HasRepoOwner returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasRepoOwner() bool {
	if o != nil && !IsNil(o.RepoOwner) {
		return true
	}

	return false
}

// SetRepoOwner gets a reference to the given string and assigns it to the RepoOwner field.
func (o *MigrateRepoOptions) SetRepoOwner(v string) {
	o.RepoOwner = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *MigrateRepoOptions) SetService(v string) {
	o.Service = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetUid() int64 {
	if o == nil || IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetUidOk() (*int64, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *MigrateRepoOptions) SetUid(v int64) {
	o.Uid = &v
}

// GetWiki returns the Wiki field value if set, zero value otherwise.
func (o *MigrateRepoOptions) GetWiki() bool {
	if o == nil || IsNil(o.Wiki) {
		var ret bool
		return ret
	}
	return *o.Wiki
}

// GetWikiOk returns a tuple with the Wiki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateRepoOptions) GetWikiOk() (*bool, bool) {
	if o == nil || IsNil(o.Wiki) {
		return nil, false
	}
	return o.Wiki, true
}

// HasWiki returns a boolean if a field has been set.
func (o *MigrateRepoOptions) HasWiki() bool {
	if o != nil && !IsNil(o.Wiki) {
		return true
	}

	return false
}

// SetWiki gets a reference to the given bool and assigns it to the Wiki field.
func (o *MigrateRepoOptions) SetWiki(v bool) {
	o.Wiki = &v
}

func (o MigrateRepoOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateRepoOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthPassword) {
		toSerialize["auth_password"] = o.AuthPassword
	}
	if !IsNil(o.AuthToken) {
		toSerialize["auth_token"] = o.AuthToken
	}
	if !IsNil(o.AuthUsername) {
		toSerialize["auth_username"] = o.AuthUsername
	}
	toSerialize["clone_addr"] = o.CloneAddr
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Lfs) {
		toSerialize["lfs"] = o.Lfs
	}
	if !IsNil(o.LfsEndpoint) {
		toSerialize["lfs_endpoint"] = o.LfsEndpoint
	}
	if !IsNil(o.Milestones) {
		toSerialize["milestones"] = o.Milestones
	}
	if !IsNil(o.Mirror) {
		toSerialize["mirror"] = o.Mirror
	}
	if !IsNil(o.MirrorInterval) {
		toSerialize["mirror_interval"] = o.MirrorInterval
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.PullRequests) {
		toSerialize["pull_requests"] = o.PullRequests
	}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}
	toSerialize["repo_name"] = o.RepoName
	if !IsNil(o.RepoOwner) {
		toSerialize["repo_owner"] = o.RepoOwner
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Wiki) {
		toSerialize["wiki"] = o.Wiki
	}
	return toSerialize, nil
}

type NullableMigrateRepoOptions struct {
	value *MigrateRepoOptions
	isSet bool
}

func (v NullableMigrateRepoOptions) Get() *MigrateRepoOptions {
	return v.value
}

func (v *NullableMigrateRepoOptions) Set(val *MigrateRepoOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateRepoOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateRepoOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateRepoOptions(val *MigrateRepoOptions) *NullableMigrateRepoOptions {
	return &NullableMigrateRepoOptions{value: val, isSet: true}
}

func (v NullableMigrateRepoOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateRepoOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
