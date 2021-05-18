/*
 * Atricore Console :: Remote : API
 *
 * # Atricore Console API
 *
 * API version: 1.4.3-SNAPSHOT
 * Contact: sgonzalez@atricore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoapi

import (
	"encoding/json"
)

// DeleteApplianceReq struct for DeleteApplianceReq
type DeleteApplianceReq struct {
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteApplianceReq DeleteApplianceReq

// NewDeleteApplianceReq instantiates a new DeleteApplianceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteApplianceReq() *DeleteApplianceReq {
	this := DeleteApplianceReq{}
	return &this
}

// NewDeleteApplianceReqWithDefaults instantiates a new DeleteApplianceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteApplianceReqWithDefaults() *DeleteApplianceReq {
	this := DeleteApplianceReq{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteApplianceReq) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteApplianceReq) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteApplianceReq) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteApplianceReq) SetId(v string) {
	o.Id = &v
}

func (o DeleteApplianceReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteApplianceReq) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteApplianceReq := _DeleteApplianceReq{}

	if err = json.Unmarshal(bytes, &varDeleteApplianceReq); err == nil {
		*o = DeleteApplianceReq(varDeleteApplianceReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteApplianceReq struct {
	value *DeleteApplianceReq
	isSet bool
}

func (v NullableDeleteApplianceReq) Get() *DeleteApplianceReq {
	return v.value
}

func (v *NullableDeleteApplianceReq) Set(val *DeleteApplianceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteApplianceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteApplianceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteApplianceReq(val *DeleteApplianceReq) *NullableDeleteApplianceReq {
	return &NullableDeleteApplianceReq{value: val, isSet: true}
}

func (v NullableDeleteApplianceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteApplianceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


