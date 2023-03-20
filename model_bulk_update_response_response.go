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

// BulkUpdateResponseResponse struct for BulkUpdateResponseResponse
type BulkUpdateResponseResponse struct {
	ProductId  *string                 `json:"productId,omitempty"`
	Success    *bool                   `json:"success,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
}

// NewBulkUpdateResponseResponse instantiates a new BulkUpdateResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateResponseResponse() *BulkUpdateResponseResponse {
	this := BulkUpdateResponseResponse{}
	return &this
}

// NewBulkUpdateResponseResponseWithDefaults instantiates a new BulkUpdateResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateResponseResponseWithDefaults() *BulkUpdateResponseResponse {
	this := BulkUpdateResponseResponse{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *BulkUpdateResponseResponse) GetProductId() string {
	if o == nil || isNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateResponseResponse) GetProductIdOk() (*string, bool) {
	if o == nil || isNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *BulkUpdateResponseResponse) HasProductId() bool {
	if o != nil && !isNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *BulkUpdateResponseResponse) SetProductId(v string) {
	o.ProductId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkUpdateResponseResponse) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateResponseResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkUpdateResponseResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *BulkUpdateResponseResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BulkUpdateResponseResponse) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateResponseResponse) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BulkUpdateResponseResponse) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *BulkUpdateResponseResponse) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o BulkUpdateResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableBulkUpdateResponseResponse struct {
	value *BulkUpdateResponseResponse
	isSet bool
}

func (v NullableBulkUpdateResponseResponse) Get() *BulkUpdateResponseResponse {
	return v.value
}

func (v *NullableBulkUpdateResponseResponse) Set(val *BulkUpdateResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateResponseResponse(val *BulkUpdateResponseResponse) *NullableBulkUpdateResponseResponse {
	return &NullableBulkUpdateResponseResponse{value: val, isSet: true}
}

func (v NullableBulkUpdateResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
