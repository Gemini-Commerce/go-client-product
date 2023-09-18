/*
product/product.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product

import (
	"encoding/json"
)

// EntitymanagerUpdateAttributeOptionsResponse struct for EntitymanagerUpdateAttributeOptionsResponse
type EntitymanagerUpdateAttributeOptionsResponse struct {
	Options []EntitymanagerAttributeOption `json:"options,omitempty"`
	Errors []EntitymanagerAttributeOptionErrors `json:"errors,omitempty"`
}

// NewEntitymanagerUpdateAttributeOptionsResponse instantiates a new EntitymanagerUpdateAttributeOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerUpdateAttributeOptionsResponse() *EntitymanagerUpdateAttributeOptionsResponse {
	this := EntitymanagerUpdateAttributeOptionsResponse{}
	return &this
}

// NewEntitymanagerUpdateAttributeOptionsResponseWithDefaults instantiates a new EntitymanagerUpdateAttributeOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerUpdateAttributeOptionsResponseWithDefaults() *EntitymanagerUpdateAttributeOptionsResponse {
	this := EntitymanagerUpdateAttributeOptionsResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeOptionsResponse) GetOptions() []EntitymanagerAttributeOption {
	if o == nil || isNil(o.Options) {
		var ret []EntitymanagerAttributeOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeOptionsResponse) GetOptionsOk() ([]EntitymanagerAttributeOption, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeOptionsResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerAttributeOption and assigns it to the Options field.
func (o *EntitymanagerUpdateAttributeOptionsResponse) SetOptions(v []EntitymanagerAttributeOption) {
	o.Options = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeOptionsResponse) GetErrors() []EntitymanagerAttributeOptionErrors {
	if o == nil || isNil(o.Errors) {
		var ret []EntitymanagerAttributeOptionErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeOptionsResponse) GetErrorsOk() ([]EntitymanagerAttributeOptionErrors, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeOptionsResponse) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []EntitymanagerAttributeOptionErrors and assigns it to the Errors field.
func (o *EntitymanagerUpdateAttributeOptionsResponse) SetErrors(v []EntitymanagerAttributeOptionErrors) {
	o.Errors = v
}

func (o EntitymanagerUpdateAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerUpdateAttributeOptionsResponse struct {
	value *EntitymanagerUpdateAttributeOptionsResponse
	isSet bool
}

func (v NullableEntitymanagerUpdateAttributeOptionsResponse) Get() *EntitymanagerUpdateAttributeOptionsResponse {
	return v.value
}

func (v *NullableEntitymanagerUpdateAttributeOptionsResponse) Set(val *EntitymanagerUpdateAttributeOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerUpdateAttributeOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerUpdateAttributeOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerUpdateAttributeOptionsResponse(val *EntitymanagerUpdateAttributeOptionsResponse) *NullableEntitymanagerUpdateAttributeOptionsResponse {
	return &NullableEntitymanagerUpdateAttributeOptionsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerUpdateAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerUpdateAttributeOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


