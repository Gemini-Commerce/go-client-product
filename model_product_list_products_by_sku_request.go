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

// checks if the ProductListProductsBySkuRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductListProductsBySkuRequest{}

// ProductListProductsBySkuRequest struct for ProductListProductsBySkuRequest
type ProductListProductsBySkuRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Sku *string `json:"sku,omitempty"`
}

// NewProductListProductsBySkuRequest instantiates a new ProductListProductsBySkuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListProductsBySkuRequest() *ProductListProductsBySkuRequest {
	this := ProductListProductsBySkuRequest{}
	return &this
}

// NewProductListProductsBySkuRequestWithDefaults instantiates a new ProductListProductsBySkuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListProductsBySkuRequestWithDefaults() *ProductListProductsBySkuRequest {
	this := ProductListProductsBySkuRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductListProductsBySkuRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsBySkuRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductListProductsBySkuRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductListProductsBySkuRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductListProductsBySkuRequest) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsBySkuRequest) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductListProductsBySkuRequest) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductListProductsBySkuRequest) SetSku(v string) {
	o.Sku = &v
}

func (o ProductListProductsBySkuRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductListProductsBySkuRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableProductListProductsBySkuRequest struct {
	value *ProductListProductsBySkuRequest
	isSet bool
}

func (v NullableProductListProductsBySkuRequest) Get() *ProductListProductsBySkuRequest {
	return v.value
}

func (v *NullableProductListProductsBySkuRequest) Set(val *ProductListProductsBySkuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListProductsBySkuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListProductsBySkuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListProductsBySkuRequest(val *ProductListProductsBySkuRequest) *NullableProductListProductsBySkuRequest {
	return &NullableProductListProductsBySkuRequest{value: val, isSet: true}
}

func (v NullableProductListProductsBySkuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListProductsBySkuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


