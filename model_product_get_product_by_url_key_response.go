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

// ProductGetProductByUrlKeyResponse struct for ProductGetProductByUrlKeyResponse
type ProductGetProductByUrlKeyResponse struct {
	Product *ProductProductEntity `json:"product,omitempty"`
}

// NewProductGetProductByUrlKeyResponse instantiates a new ProductGetProductByUrlKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGetProductByUrlKeyResponse() *ProductGetProductByUrlKeyResponse {
	this := ProductGetProductByUrlKeyResponse{}
	return &this
}

// NewProductGetProductByUrlKeyResponseWithDefaults instantiates a new ProductGetProductByUrlKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGetProductByUrlKeyResponseWithDefaults() *ProductGetProductByUrlKeyResponse {
	this := ProductGetProductByUrlKeyResponse{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ProductGetProductByUrlKeyResponse) GetProduct() ProductProductEntity {
	if o == nil || isNil(o.Product) {
		var ret ProductProductEntity
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByUrlKeyResponse) GetProductOk() (*ProductProductEntity, bool) {
	if o == nil || isNil(o.Product) {
    return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ProductGetProductByUrlKeyResponse) HasProduct() bool {
	if o != nil && !isNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductProductEntity and assigns it to the Product field.
func (o *ProductGetProductByUrlKeyResponse) SetProduct(v ProductProductEntity) {
	o.Product = &v
}

func (o ProductGetProductByUrlKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return json.Marshal(toSerialize)
}

type NullableProductGetProductByUrlKeyResponse struct {
	value *ProductGetProductByUrlKeyResponse
	isSet bool
}

func (v NullableProductGetProductByUrlKeyResponse) Get() *ProductGetProductByUrlKeyResponse {
	return v.value
}

func (v *NullableProductGetProductByUrlKeyResponse) Set(val *ProductGetProductByUrlKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGetProductByUrlKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGetProductByUrlKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGetProductByUrlKeyResponse(val *ProductGetProductByUrlKeyResponse) *NullableProductGetProductByUrlKeyResponse {
	return &NullableProductGetProductByUrlKeyResponse{value: val, isSet: true}
}

func (v NullableProductGetProductByUrlKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGetProductByUrlKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


