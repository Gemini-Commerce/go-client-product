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

// checks if the ProductListProductsBySkuResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductListProductsBySkuResponse{}

// ProductListProductsBySkuResponse struct for ProductListProductsBySkuResponse
type ProductListProductsBySkuResponse struct {
	Products []ProductProductEntity `json:"products,omitempty"`
}

// NewProductListProductsBySkuResponse instantiates a new ProductListProductsBySkuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListProductsBySkuResponse() *ProductListProductsBySkuResponse {
	this := ProductListProductsBySkuResponse{}
	return &this
}

// NewProductListProductsBySkuResponseWithDefaults instantiates a new ProductListProductsBySkuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListProductsBySkuResponseWithDefaults() *ProductListProductsBySkuResponse {
	this := ProductListProductsBySkuResponse{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ProductListProductsBySkuResponse) GetProducts() []ProductProductEntity {
	if o == nil || IsNil(o.Products) {
		var ret []ProductProductEntity
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsBySkuResponse) GetProductsOk() ([]ProductProductEntity, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ProductListProductsBySkuResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductProductEntity and assigns it to the Products field.
func (o *ProductListProductsBySkuResponse) SetProducts(v []ProductProductEntity) {
	o.Products = v
}

func (o ProductListProductsBySkuResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductListProductsBySkuResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableProductListProductsBySkuResponse struct {
	value *ProductListProductsBySkuResponse
	isSet bool
}

func (v NullableProductListProductsBySkuResponse) Get() *ProductListProductsBySkuResponse {
	return v.value
}

func (v *NullableProductListProductsBySkuResponse) Set(val *ProductListProductsBySkuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListProductsBySkuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListProductsBySkuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListProductsBySkuResponse(val *ProductListProductsBySkuResponse) *NullableProductListProductsBySkuResponse {
	return &NullableProductListProductsBySkuResponse{value: val, isSet: true}
}

func (v NullableProductListProductsBySkuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListProductsBySkuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


