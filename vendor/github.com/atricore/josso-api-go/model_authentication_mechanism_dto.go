/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.4.3-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// AuthenticationMechanismDTO struct for AuthenticationMechanismDTO
type AuthenticationMechanismDTO struct {
	DelegatedAuthentication *DelegatedAuthenticationDTO `json:"delegatedAuthentication,omitempty"`
	DisplayName             *string                     `json:"displayName,omitempty"`
	ElementId               *string                     `json:"elementId,omitempty"`
	Id                      *int64                      `json:"id,omitempty"`
	Name                    *string                     `json:"name,omitempty"`
	Priority                *int32                      `json:"priority,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AuthenticationMechanismDTO AuthenticationMechanismDTO

// NewAuthenticationMechanismDTO instantiates a new AuthenticationMechanismDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMechanismDTO() *AuthenticationMechanismDTO {
	this := AuthenticationMechanismDTO{}
	return &this
}

// NewAuthenticationMechanismDTOWithDefaults instantiates a new AuthenticationMechanismDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMechanismDTOWithDefaults() *AuthenticationMechanismDTO {
	this := AuthenticationMechanismDTO{}
	return &this
}

// GetDelegatedAuthentication returns the DelegatedAuthentication field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetDelegatedAuthentication() DelegatedAuthenticationDTO {
	if o == nil || o.DelegatedAuthentication == nil {
		var ret DelegatedAuthenticationDTO
		return ret
	}
	return *o.DelegatedAuthentication
}

// GetDelegatedAuthenticationOk returns a tuple with the DelegatedAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetDelegatedAuthenticationOk() (*DelegatedAuthenticationDTO, bool) {
	if o == nil || o.DelegatedAuthentication == nil {
		return nil, false
	}
	return o.DelegatedAuthentication, true
}

// HasDelegatedAuthentication returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasDelegatedAuthentication() bool {
	if o != nil && o.DelegatedAuthentication != nil {
		return true
	}

	return false
}

// SetDelegatedAuthentication gets a reference to the given DelegatedAuthenticationDTO and assigns it to the DelegatedAuthentication field.
func (o *AuthenticationMechanismDTO) SetDelegatedAuthentication(v DelegatedAuthenticationDTO) {
	o.DelegatedAuthentication = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthenticationMechanismDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *AuthenticationMechanismDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AuthenticationMechanismDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationMechanismDTO) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AuthenticationMechanismDTO) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMechanismDTO) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AuthenticationMechanismDTO) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AuthenticationMechanismDTO) SetPriority(v int32) {
	o.Priority = &v
}

func (o AuthenticationMechanismDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegatedAuthentication != nil {
		toSerialize["delegatedAuthentication"] = o.DelegatedAuthentication
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationMechanismDTO) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationMechanismDTO := _AuthenticationMechanismDTO{}

	if err = json.Unmarshal(bytes, &varAuthenticationMechanismDTO); err == nil {
		*o = AuthenticationMechanismDTO(varAuthenticationMechanismDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delegatedAuthentication")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticationMechanismDTO struct {
	value *AuthenticationMechanismDTO
	isSet bool
}

func (v NullableAuthenticationMechanismDTO) Get() *AuthenticationMechanismDTO {
	return v.value
}

func (v *NullableAuthenticationMechanismDTO) Set(val *AuthenticationMechanismDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMechanismDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMechanismDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMechanismDTO(val *AuthenticationMechanismDTO) *NullableAuthenticationMechanismDTO {
	return &NullableAuthenticationMechanismDTO{value: val, isSet: true}
}

func (v NullableAuthenticationMechanismDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMechanismDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
