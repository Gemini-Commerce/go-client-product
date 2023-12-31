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

// checks if the ProductBulkUpdateRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBulkUpdateRequestPayload{}

// ProductBulkUpdateRequestPayload struct for ProductBulkUpdateRequestPayload
type ProductBulkUpdateRequestPayload struct {
	UpdateMask *ProductFieldMask `json:"updateMask,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
}

// NewProductBulkUpdateRequestPayload instantiates a new ProductBulkUpdateRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateRequestPayload() *ProductBulkUpdateRequestPayload {
	this := ProductBulkUpdateRequestPayload{}
	return &this
}

// NewProductBulkUpdateRequestPayloadWithDefaults instantiates a new ProductBulkUpdateRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateRequestPayloadWithDefaults() *ProductBulkUpdateRequestPayload {
	this := ProductBulkUpdateRequestPayload{}
	return &this
}

// GetUpdateMask returns the UpdateMask field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestPayload) GetUpdateMask() ProductFieldMask {
	if o == nil || IsNil(o.UpdateMask) {
		var ret ProductFieldMask
		return ret
	}
	return *o.UpdateMask
}

// GetUpdateMaskOk returns a tuple with the UpdateMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestPayload) GetUpdateMaskOk() (*ProductFieldMask, bool) {
	if o == nil || IsNil(o.UpdateMask) {
		return nil, false
	}
	return o.UpdateMask, true
}

// HasUpdateMask returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestPayload) HasUpdateMask() bool {
	if o != nil && !IsNil(o.UpdateMask) {
		return true
	}

	return false
}

// SetUpdateMask gets a reference to the given ProductFieldMask and assigns it to the UpdateMask field.
func (o *ProductBulkUpdateRequestPayload) SetUpdateMask(v ProductFieldMask) {
	o.UpdateMask = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestPayload) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestPayload) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestPayload) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductBulkUpdateRequestPayload) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o ProductBulkUpdateRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBulkUpdateRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateMask) {
		toSerialize["updateMask"] = o.UpdateMask
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableProductBulkUpdateRequestPayload struct {
	value *ProductBulkUpdateRequestPayload
	isSet bool
}

func (v NullableProductBulkUpdateRequestPayload) Get() *ProductBulkUpdateRequestPayload {
	return v.value
}

func (v *NullableProductBulkUpdateRequestPayload) Set(val *ProductBulkUpdateRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateRequestPayload(val *ProductBulkUpdateRequestPayload) *NullableProductBulkUpdateRequestPayload {
	return &NullableProductBulkUpdateRequestPayload{value: val, isSet: true}
}

func (v NullableProductBulkUpdateRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


