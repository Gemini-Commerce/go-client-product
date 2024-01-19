/*
Product Service

API for managing products

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product

import (
	"encoding/json"
)

// checks if the ProductLocalizedAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductLocalizedAsset{}

// ProductLocalizedAsset struct for ProductLocalizedAsset
type ProductLocalizedAsset struct {
	Value *map[string]string `json:"value,omitempty"`
}

// NewProductLocalizedAsset instantiates a new ProductLocalizedAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductLocalizedAsset() *ProductLocalizedAsset {
	this := ProductLocalizedAsset{}
	return &this
}

// NewProductLocalizedAssetWithDefaults instantiates a new ProductLocalizedAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductLocalizedAssetWithDefaults() *ProductLocalizedAsset {
	this := ProductLocalizedAsset{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductLocalizedAsset) GetValue() map[string]string {
	if o == nil || IsNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductLocalizedAsset) GetValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductLocalizedAsset) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *ProductLocalizedAsset) SetValue(v map[string]string) {
	o.Value = &v
}

func (o ProductLocalizedAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductLocalizedAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableProductLocalizedAsset struct {
	value *ProductLocalizedAsset
	isSet bool
}

func (v NullableProductLocalizedAsset) Get() *ProductLocalizedAsset {
	return v.value
}

func (v *NullableProductLocalizedAsset) Set(val *ProductLocalizedAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableProductLocalizedAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableProductLocalizedAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductLocalizedAsset(val *ProductLocalizedAsset) *NullableProductLocalizedAsset {
	return &NullableProductLocalizedAsset{value: val, isSet: true}
}

func (v NullableProductLocalizedAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductLocalizedAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

