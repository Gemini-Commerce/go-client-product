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

// ProductentitymanagerLocalizedText struct for ProductentitymanagerLocalizedText
type ProductentitymanagerLocalizedText struct {
	Value *map[string]string `json:"value,omitempty"`
}

// NewProductentitymanagerLocalizedText instantiates a new ProductentitymanagerLocalizedText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductentitymanagerLocalizedText() *ProductentitymanagerLocalizedText {
	this := ProductentitymanagerLocalizedText{}
	return &this
}

// NewProductentitymanagerLocalizedTextWithDefaults instantiates a new ProductentitymanagerLocalizedText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductentitymanagerLocalizedTextWithDefaults() *ProductentitymanagerLocalizedText {
	this := ProductentitymanagerLocalizedText{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductentitymanagerLocalizedText) GetValue() map[string]string {
	if o == nil || isNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductentitymanagerLocalizedText) GetValueOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductentitymanagerLocalizedText) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *ProductentitymanagerLocalizedText) SetValue(v map[string]string) {
	o.Value = &v
}

func (o ProductentitymanagerLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableProductentitymanagerLocalizedText struct {
	value *ProductentitymanagerLocalizedText
	isSet bool
}

func (v NullableProductentitymanagerLocalizedText) Get() *ProductentitymanagerLocalizedText {
	return v.value
}

func (v *NullableProductentitymanagerLocalizedText) Set(val *ProductentitymanagerLocalizedText) {
	v.value = val
	v.isSet = true
}

func (v NullableProductentitymanagerLocalizedText) IsSet() bool {
	return v.isSet
}

func (v *NullableProductentitymanagerLocalizedText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductentitymanagerLocalizedText(val *ProductentitymanagerLocalizedText) *NullableProductentitymanagerLocalizedText {
	return &NullableProductentitymanagerLocalizedText{value: val, isSet: true}
}

func (v NullableProductentitymanagerLocalizedText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductentitymanagerLocalizedText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


