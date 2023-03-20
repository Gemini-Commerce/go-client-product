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

// EntitymanagerGetAttributeOptionsResponse struct for EntitymanagerGetAttributeOptionsResponse
type EntitymanagerGetAttributeOptionsResponse struct {
	Options []EntitymanagerGetAttributeOptionsResponseOption `json:"options,omitempty"`
}

// NewEntitymanagerGetAttributeOptionsResponse instantiates a new EntitymanagerGetAttributeOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerGetAttributeOptionsResponse() *EntitymanagerGetAttributeOptionsResponse {
	this := EntitymanagerGetAttributeOptionsResponse{}
	return &this
}

// NewEntitymanagerGetAttributeOptionsResponseWithDefaults instantiates a new EntitymanagerGetAttributeOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerGetAttributeOptionsResponseWithDefaults() *EntitymanagerGetAttributeOptionsResponse {
	this := EntitymanagerGetAttributeOptionsResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerGetAttributeOptionsResponse) GetOptions() []EntitymanagerGetAttributeOptionsResponseOption {
	if o == nil || isNil(o.Options) {
		var ret []EntitymanagerGetAttributeOptionsResponseOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerGetAttributeOptionsResponse) GetOptionsOk() ([]EntitymanagerGetAttributeOptionsResponseOption, bool) {
	if o == nil || isNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerGetAttributeOptionsResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerGetAttributeOptionsResponseOption and assigns it to the Options field.
func (o *EntitymanagerGetAttributeOptionsResponse) SetOptions(v []EntitymanagerGetAttributeOptionsResponseOption) {
	o.Options = v
}

func (o EntitymanagerGetAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerGetAttributeOptionsResponse struct {
	value *EntitymanagerGetAttributeOptionsResponse
	isSet bool
}

func (v NullableEntitymanagerGetAttributeOptionsResponse) Get() *EntitymanagerGetAttributeOptionsResponse {
	return v.value
}

func (v *NullableEntitymanagerGetAttributeOptionsResponse) Set(val *EntitymanagerGetAttributeOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerGetAttributeOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerGetAttributeOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerGetAttributeOptionsResponse(val *EntitymanagerGetAttributeOptionsResponse) *NullableEntitymanagerGetAttributeOptionsResponse {
	return &NullableEntitymanagerGetAttributeOptionsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerGetAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerGetAttributeOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
