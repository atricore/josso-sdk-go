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
)

// SessionManagerFactoryDTO struct for SessionManagerFactoryDTO
type SessionManagerFactoryDTO struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionManagerFactoryDTO SessionManagerFactoryDTO

// NewSessionManagerFactoryDTO instantiates a new SessionManagerFactoryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionManagerFactoryDTO() *SessionManagerFactoryDTO {
	this := SessionManagerFactoryDTO{}
	return &this
}

// NewSessionManagerFactoryDTOWithDefaults instantiates a new SessionManagerFactoryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionManagerFactoryDTOWithDefaults() *SessionManagerFactoryDTO {
	this := SessionManagerFactoryDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SessionManagerFactoryDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagerFactoryDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SessionManagerFactoryDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SessionManagerFactoryDTO) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SessionManagerFactoryDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagerFactoryDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SessionManagerFactoryDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SessionManagerFactoryDTO) SetName(v string) {
	o.Name = &v
}

func (o SessionManagerFactoryDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionManagerFactoryDTO) UnmarshalJSON(bytes []byte) (err error) {
	varSessionManagerFactoryDTO := _SessionManagerFactoryDTO{}

	if err = json.Unmarshal(bytes, &varSessionManagerFactoryDTO); err == nil {
		*o = SessionManagerFactoryDTO(varSessionManagerFactoryDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionManagerFactoryDTO struct {
	value *SessionManagerFactoryDTO
	isSet bool
}

func (v NullableSessionManagerFactoryDTO) Get() *SessionManagerFactoryDTO {
	return v.value
}

func (v *NullableSessionManagerFactoryDTO) Set(val *SessionManagerFactoryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionManagerFactoryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionManagerFactoryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionManagerFactoryDTO(val *SessionManagerFactoryDTO) *NullableSessionManagerFactoryDTO {
	return &NullableSessionManagerFactoryDTO{value: val, isSet: true}
}

func (v NullableSessionManagerFactoryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionManagerFactoryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


