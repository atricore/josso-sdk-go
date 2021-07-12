/*
 * Atricore Console :: Remote : API
 *
 * # Atricore Console API
 *
 * API version: 1.4.3-SNAPSHOT
 * Contact: sgonzalez@atricore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
	"time"
)

// IdentityApplianceDefinitionDTO struct for IdentityApplianceDefinitionDTO
type IdentityApplianceDefinitionDTO struct {
	ActiveFeatures *[]string `json:"activeFeatures,omitempty"`
	AuthenticationServices *[]AuthenticationServiceDTO `json:"authenticationServices,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	ExecutionEnvironments *[]ExecutionEnvironmentDTO `json:"executionEnvironments,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IdentitySources *[]IdentitySourceDTO `json:"identitySources,omitempty"`
	IdpSelector *EntitySelectionStrategyDTO `json:"idpSelector,omitempty"`
	Keystore *KeystoreDTO `json:"keystore,omitempty"`
	LastModification *time.Time `json:"lastModification,omitempty"`
	Location *LocationDTO `json:"location,omitempty"`
	ModelVersion *string `json:"modelVersion,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Providers *[]ProviderDTO `json:"providers,omitempty"`
	RequiredBundles *[]string `json:"requiredBundles,omitempty"`
	Revision *int32 `json:"revision,omitempty"`
	SecurityConfig *IdentityApplianceSecurityConfigDTO `json:"securityConfig,omitempty"`
	ServiceResources *[]ServiceResourceDTO `json:"serviceResources,omitempty"`
	SupportedRoles *[]string `json:"supportedRoles,omitempty"`
	UserDashboardBranding *UserDashboardBrandingDTO `json:"userDashboardBranding,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityApplianceDefinitionDTO IdentityApplianceDefinitionDTO

// NewIdentityApplianceDefinitionDTO instantiates a new IdentityApplianceDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityApplianceDefinitionDTO() *IdentityApplianceDefinitionDTO {
	this := IdentityApplianceDefinitionDTO{}
	return &this
}

// NewIdentityApplianceDefinitionDTOWithDefaults instantiates a new IdentityApplianceDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityApplianceDefinitionDTOWithDefaults() *IdentityApplianceDefinitionDTO {
	this := IdentityApplianceDefinitionDTO{}
	return &this
}

// GetActiveFeatures returns the ActiveFeatures field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetActiveFeatures() []string {
	if o == nil || o.ActiveFeatures == nil {
		var ret []string
		return ret
	}
	return *o.ActiveFeatures
}

// GetActiveFeaturesOk returns a tuple with the ActiveFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetActiveFeaturesOk() (*[]string, bool) {
	if o == nil || o.ActiveFeatures == nil {
		return nil, false
	}
	return o.ActiveFeatures, true
}

// HasActiveFeatures returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasActiveFeatures() bool {
	if o != nil && o.ActiveFeatures != nil {
		return true
	}

	return false
}

// SetActiveFeatures gets a reference to the given []string and assigns it to the ActiveFeatures field.
func (o *IdentityApplianceDefinitionDTO) SetActiveFeatures(v []string) {
	o.ActiveFeatures = &v
}

// GetAuthenticationServices returns the AuthenticationServices field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetAuthenticationServices() []AuthenticationServiceDTO {
	if o == nil || o.AuthenticationServices == nil {
		var ret []AuthenticationServiceDTO
		return ret
	}
	return *o.AuthenticationServices
}

// GetAuthenticationServicesOk returns a tuple with the AuthenticationServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetAuthenticationServicesOk() (*[]AuthenticationServiceDTO, bool) {
	if o == nil || o.AuthenticationServices == nil {
		return nil, false
	}
	return o.AuthenticationServices, true
}

// HasAuthenticationServices returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasAuthenticationServices() bool {
	if o != nil && o.AuthenticationServices != nil {
		return true
	}

	return false
}

// SetAuthenticationServices gets a reference to the given []AuthenticationServiceDTO and assigns it to the AuthenticationServices field.
func (o *IdentityApplianceDefinitionDTO) SetAuthenticationServices(v []AuthenticationServiceDTO) {
	o.AuthenticationServices = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityApplianceDefinitionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityApplianceDefinitionDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *IdentityApplianceDefinitionDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetExecutionEnvironments returns the ExecutionEnvironments field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetExecutionEnvironments() []ExecutionEnvironmentDTO {
	if o == nil || o.ExecutionEnvironments == nil {
		var ret []ExecutionEnvironmentDTO
		return ret
	}
	return *o.ExecutionEnvironments
}

// GetExecutionEnvironmentsOk returns a tuple with the ExecutionEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetExecutionEnvironmentsOk() (*[]ExecutionEnvironmentDTO, bool) {
	if o == nil || o.ExecutionEnvironments == nil {
		return nil, false
	}
	return o.ExecutionEnvironments, true
}

// HasExecutionEnvironments returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasExecutionEnvironments() bool {
	if o != nil && o.ExecutionEnvironments != nil {
		return true
	}

	return false
}

// SetExecutionEnvironments gets a reference to the given []ExecutionEnvironmentDTO and assigns it to the ExecutionEnvironments field.
func (o *IdentityApplianceDefinitionDTO) SetExecutionEnvironments(v []ExecutionEnvironmentDTO) {
	o.ExecutionEnvironments = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentityApplianceDefinitionDTO) SetId(v int64) {
	o.Id = &v
}

// GetIdentitySources returns the IdentitySources field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetIdentitySources() []IdentitySourceDTO {
	if o == nil || o.IdentitySources == nil {
		var ret []IdentitySourceDTO
		return ret
	}
	return *o.IdentitySources
}

// GetIdentitySourcesOk returns a tuple with the IdentitySources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetIdentitySourcesOk() (*[]IdentitySourceDTO, bool) {
	if o == nil || o.IdentitySources == nil {
		return nil, false
	}
	return o.IdentitySources, true
}

// HasIdentitySources returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasIdentitySources() bool {
	if o != nil && o.IdentitySources != nil {
		return true
	}

	return false
}

// SetIdentitySources gets a reference to the given []IdentitySourceDTO and assigns it to the IdentitySources field.
func (o *IdentityApplianceDefinitionDTO) SetIdentitySources(v []IdentitySourceDTO) {
	o.IdentitySources = &v
}

// GetIdpSelector returns the IdpSelector field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetIdpSelector() EntitySelectionStrategyDTO {
	if o == nil || o.IdpSelector == nil {
		var ret EntitySelectionStrategyDTO
		return ret
	}
	return *o.IdpSelector
}

// GetIdpSelectorOk returns a tuple with the IdpSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetIdpSelectorOk() (*EntitySelectionStrategyDTO, bool) {
	if o == nil || o.IdpSelector == nil {
		return nil, false
	}
	return o.IdpSelector, true
}

// HasIdpSelector returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasIdpSelector() bool {
	if o != nil && o.IdpSelector != nil {
		return true
	}

	return false
}

// SetIdpSelector gets a reference to the given EntitySelectionStrategyDTO and assigns it to the IdpSelector field.
func (o *IdentityApplianceDefinitionDTO) SetIdpSelector(v EntitySelectionStrategyDTO) {
	o.IdpSelector = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetKeystore() KeystoreDTO {
	if o == nil || o.Keystore == nil {
		var ret KeystoreDTO
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetKeystoreOk() (*KeystoreDTO, bool) {
	if o == nil || o.Keystore == nil {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasKeystore() bool {
	if o != nil && o.Keystore != nil {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given KeystoreDTO and assigns it to the Keystore field.
func (o *IdentityApplianceDefinitionDTO) SetKeystore(v KeystoreDTO) {
	o.Keystore = &v
}

// GetLastModification returns the LastModification field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetLastModification() time.Time {
	if o == nil || o.LastModification == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModification
}

// GetLastModificationOk returns a tuple with the LastModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetLastModificationOk() (*time.Time, bool) {
	if o == nil || o.LastModification == nil {
		return nil, false
	}
	return o.LastModification, true
}

// HasLastModification returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasLastModification() bool {
	if o != nil && o.LastModification != nil {
		return true
	}

	return false
}

// SetLastModification gets a reference to the given time.Time and assigns it to the LastModification field.
func (o *IdentityApplianceDefinitionDTO) SetLastModification(v time.Time) {
	o.LastModification = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetLocation() LocationDTO {
	if o == nil || o.Location == nil {
		var ret LocationDTO
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetLocationOk() (*LocationDTO, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationDTO and assigns it to the Location field.
func (o *IdentityApplianceDefinitionDTO) SetLocation(v LocationDTO) {
	o.Location = &v
}

// GetModelVersion returns the ModelVersion field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetModelVersion() string {
	if o == nil || o.ModelVersion == nil {
		var ret string
		return ret
	}
	return *o.ModelVersion
}

// GetModelVersionOk returns a tuple with the ModelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetModelVersionOk() (*string, bool) {
	if o == nil || o.ModelVersion == nil {
		return nil, false
	}
	return o.ModelVersion, true
}

// HasModelVersion returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasModelVersion() bool {
	if o != nil && o.ModelVersion != nil {
		return true
	}

	return false
}

// SetModelVersion gets a reference to the given string and assigns it to the ModelVersion field.
func (o *IdentityApplianceDefinitionDTO) SetModelVersion(v string) {
	o.ModelVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityApplianceDefinitionDTO) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IdentityApplianceDefinitionDTO) SetNamespace(v string) {
	o.Namespace = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetProviders() []ProviderDTO {
	if o == nil || o.Providers == nil {
		var ret []ProviderDTO
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetProvidersOk() (*[]ProviderDTO, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []ProviderDTO and assigns it to the Providers field.
func (o *IdentityApplianceDefinitionDTO) SetProviders(v []ProviderDTO) {
	o.Providers = &v
}

// GetRequiredBundles returns the RequiredBundles field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetRequiredBundles() []string {
	if o == nil || o.RequiredBundles == nil {
		var ret []string
		return ret
	}
	return *o.RequiredBundles
}

// GetRequiredBundlesOk returns a tuple with the RequiredBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetRequiredBundlesOk() (*[]string, bool) {
	if o == nil || o.RequiredBundles == nil {
		return nil, false
	}
	return o.RequiredBundles, true
}

// HasRequiredBundles returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasRequiredBundles() bool {
	if o != nil && o.RequiredBundles != nil {
		return true
	}

	return false
}

// SetRequiredBundles gets a reference to the given []string and assigns it to the RequiredBundles field.
func (o *IdentityApplianceDefinitionDTO) SetRequiredBundles(v []string) {
	o.RequiredBundles = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetRevision() int32 {
	if o == nil || o.Revision == nil {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetRevisionOk() (*int32, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *IdentityApplianceDefinitionDTO) SetRevision(v int32) {
	o.Revision = &v
}

// GetSecurityConfig returns the SecurityConfig field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetSecurityConfig() IdentityApplianceSecurityConfigDTO {
	if o == nil || o.SecurityConfig == nil {
		var ret IdentityApplianceSecurityConfigDTO
		return ret
	}
	return *o.SecurityConfig
}

// GetSecurityConfigOk returns a tuple with the SecurityConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetSecurityConfigOk() (*IdentityApplianceSecurityConfigDTO, bool) {
	if o == nil || o.SecurityConfig == nil {
		return nil, false
	}
	return o.SecurityConfig, true
}

// HasSecurityConfig returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasSecurityConfig() bool {
	if o != nil && o.SecurityConfig != nil {
		return true
	}

	return false
}

// SetSecurityConfig gets a reference to the given IdentityApplianceSecurityConfigDTO and assigns it to the SecurityConfig field.
func (o *IdentityApplianceDefinitionDTO) SetSecurityConfig(v IdentityApplianceSecurityConfigDTO) {
	o.SecurityConfig = &v
}

// GetServiceResources returns the ServiceResources field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetServiceResources() []ServiceResourceDTO {
	if o == nil || o.ServiceResources == nil {
		var ret []ServiceResourceDTO
		return ret
	}
	return *o.ServiceResources
}

// GetServiceResourcesOk returns a tuple with the ServiceResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetServiceResourcesOk() (*[]ServiceResourceDTO, bool) {
	if o == nil || o.ServiceResources == nil {
		return nil, false
	}
	return o.ServiceResources, true
}

// HasServiceResources returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasServiceResources() bool {
	if o != nil && o.ServiceResources != nil {
		return true
	}

	return false
}

// SetServiceResources gets a reference to the given []ServiceResourceDTO and assigns it to the ServiceResources field.
func (o *IdentityApplianceDefinitionDTO) SetServiceResources(v []ServiceResourceDTO) {
	o.ServiceResources = &v
}

// GetSupportedRoles returns the SupportedRoles field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetSupportedRoles() []string {
	if o == nil || o.SupportedRoles == nil {
		var ret []string
		return ret
	}
	return *o.SupportedRoles
}

// GetSupportedRolesOk returns a tuple with the SupportedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetSupportedRolesOk() (*[]string, bool) {
	if o == nil || o.SupportedRoles == nil {
		return nil, false
	}
	return o.SupportedRoles, true
}

// HasSupportedRoles returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasSupportedRoles() bool {
	if o != nil && o.SupportedRoles != nil {
		return true
	}

	return false
}

// SetSupportedRoles gets a reference to the given []string and assigns it to the SupportedRoles field.
func (o *IdentityApplianceDefinitionDTO) SetSupportedRoles(v []string) {
	o.SupportedRoles = &v
}

// GetUserDashboardBranding returns the UserDashboardBranding field value if set, zero value otherwise.
func (o *IdentityApplianceDefinitionDTO) GetUserDashboardBranding() UserDashboardBrandingDTO {
	if o == nil || o.UserDashboardBranding == nil {
		var ret UserDashboardBrandingDTO
		return ret
	}
	return *o.UserDashboardBranding
}

// GetUserDashboardBrandingOk returns a tuple with the UserDashboardBranding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDefinitionDTO) GetUserDashboardBrandingOk() (*UserDashboardBrandingDTO, bool) {
	if o == nil || o.UserDashboardBranding == nil {
		return nil, false
	}
	return o.UserDashboardBranding, true
}

// HasUserDashboardBranding returns a boolean if a field has been set.
func (o *IdentityApplianceDefinitionDTO) HasUserDashboardBranding() bool {
	if o != nil && o.UserDashboardBranding != nil {
		return true
	}

	return false
}

// SetUserDashboardBranding gets a reference to the given UserDashboardBrandingDTO and assigns it to the UserDashboardBranding field.
func (o *IdentityApplianceDefinitionDTO) SetUserDashboardBranding(v UserDashboardBrandingDTO) {
	o.UserDashboardBranding = &v
}

func (o IdentityApplianceDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveFeatures != nil {
		toSerialize["activeFeatures"] = o.ActiveFeatures
	}
	if o.AuthenticationServices != nil {
		toSerialize["authenticationServices"] = o.AuthenticationServices
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ExecutionEnvironments != nil {
		toSerialize["executionEnvironments"] = o.ExecutionEnvironments
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentitySources != nil {
		toSerialize["identitySources"] = o.IdentitySources
	}
	if o.IdpSelector != nil {
		toSerialize["idpSelector"] = o.IdpSelector
	}
	if o.Keystore != nil {
		toSerialize["keystore"] = o.Keystore
	}
	if o.LastModification != nil {
		toSerialize["lastModification"] = o.LastModification
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ModelVersion != nil {
		toSerialize["modelVersion"] = o.ModelVersion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.RequiredBundles != nil {
		toSerialize["requiredBundles"] = o.RequiredBundles
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.SecurityConfig != nil {
		toSerialize["securityConfig"] = o.SecurityConfig
	}
	if o.ServiceResources != nil {
		toSerialize["serviceResources"] = o.ServiceResources
	}
	if o.SupportedRoles != nil {
		toSerialize["supportedRoles"] = o.SupportedRoles
	}
	if o.UserDashboardBranding != nil {
		toSerialize["userDashboardBranding"] = o.UserDashboardBranding
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityApplianceDefinitionDTO) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityApplianceDefinitionDTO := _IdentityApplianceDefinitionDTO{}

	if err = json.Unmarshal(bytes, &varIdentityApplianceDefinitionDTO); err == nil {
		*o = IdentityApplianceDefinitionDTO(varIdentityApplianceDefinitionDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activeFeatures")
		delete(additionalProperties, "authenticationServices")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "executionEnvironments")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identitySources")
		delete(additionalProperties, "idpSelector")
		delete(additionalProperties, "keystore")
		delete(additionalProperties, "lastModification")
		delete(additionalProperties, "location")
		delete(additionalProperties, "modelVersion")
		delete(additionalProperties, "name")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "providers")
		delete(additionalProperties, "requiredBundles")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "securityConfig")
		delete(additionalProperties, "serviceResources")
		delete(additionalProperties, "supportedRoles")
		delete(additionalProperties, "userDashboardBranding")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityApplianceDefinitionDTO struct {
	value *IdentityApplianceDefinitionDTO
	isSet bool
}

func (v NullableIdentityApplianceDefinitionDTO) Get() *IdentityApplianceDefinitionDTO {
	return v.value
}

func (v *NullableIdentityApplianceDefinitionDTO) Set(val *IdentityApplianceDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityApplianceDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityApplianceDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityApplianceDefinitionDTO(val *IdentityApplianceDefinitionDTO) *NullableIdentityApplianceDefinitionDTO {
	return &NullableIdentityApplianceDefinitionDTO{value: val, isSet: true}
}

func (v NullableIdentityApplianceDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityApplianceDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

