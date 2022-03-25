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

// GetApplianceRes struct for GetApplianceRes
type GetApplianceRes struct {
	Appliance *IdentityApplianceDefinitionDTO `json:"appliance,omitempty"`
	Error *string `json:"error,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetApplianceRes GetApplianceRes

// NewGetApplianceRes instantiates a new GetApplianceRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplianceRes() *GetApplianceRes {
	this := GetApplianceRes{}
	return &this
}

// NewGetApplianceResWithDefaults instantiates a new GetApplianceRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplianceResWithDefaults() *GetApplianceRes {
	this := GetApplianceRes{}
	return &this
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *GetApplianceRes) GetAppliance() IdentityApplianceDefinitionDTO {
	if o == nil || o.Appliance == nil {
		var ret IdentityApplianceDefinitionDTO
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceRes) GetApplianceOk() (*IdentityApplianceDefinitionDTO, bool) {
	if o == nil || o.Appliance == nil {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *GetApplianceRes) HasAppliance() bool {
	if o != nil && o.Appliance != nil {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given IdentityApplianceDefinitionDTO and assigns it to the Appliance field.
func (o *GetApplianceRes) SetAppliance(v IdentityApplianceDefinitionDTO) {
	o.Appliance = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetApplianceRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetApplianceRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetApplianceRes) SetError(v string) {
	o.Error = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetApplianceRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetApplianceRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetApplianceRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetApplianceRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appliance != nil {
		toSerialize["appliance"] = o.Appliance
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetApplianceRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetApplianceRes := _GetApplianceRes{}

	if err = json.Unmarshal(bytes, &varGetApplianceRes); err == nil {
		*o = GetApplianceRes(varGetApplianceRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appliance")
		delete(additionalProperties, "error")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetApplianceRes struct {
	value *GetApplianceRes
	isSet bool
}

func (v NullableGetApplianceRes) Get() *GetApplianceRes {
	return v.value
}

func (v *NullableGetApplianceRes) Set(val *GetApplianceRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplianceRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplianceRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplianceRes(val *GetApplianceRes) *NullableGetApplianceRes {
	return &NullableGetApplianceRes{value: val, isSet: true}
}

func (v NullableGetApplianceRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplianceRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


