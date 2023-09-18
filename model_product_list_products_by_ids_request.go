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

// ProductListProductsByIdsRequest struct for ProductListProductsByIdsRequest
type ProductListProductsByIdsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Ids []string `json:"ids,omitempty"`
}

// NewProductListProductsByIdsRequest instantiates a new ProductListProductsByIdsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListProductsByIdsRequest() *ProductListProductsByIdsRequest {
	this := ProductListProductsByIdsRequest{}
	return &this
}

// NewProductListProductsByIdsRequestWithDefaults instantiates a new ProductListProductsByIdsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListProductsByIdsRequestWithDefaults() *ProductListProductsByIdsRequest {
	this := ProductListProductsByIdsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductListProductsByIdsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsByIdsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductListProductsByIdsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductListProductsByIdsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ProductListProductsByIdsRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsByIdsRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ProductListProductsByIdsRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ProductListProductsByIdsRequest) SetIds(v []string) {
	o.Ids = v
}

func (o ProductListProductsByIdsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableProductListProductsByIdsRequest struct {
	value *ProductListProductsByIdsRequest
	isSet bool
}

func (v NullableProductListProductsByIdsRequest) Get() *ProductListProductsByIdsRequest {
	return v.value
}

func (v *NullableProductListProductsByIdsRequest) Set(val *ProductListProductsByIdsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListProductsByIdsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListProductsByIdsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListProductsByIdsRequest(val *ProductListProductsByIdsRequest) *NullableProductListProductsByIdsRequest {
	return &NullableProductListProductsByIdsRequest{value: val, isSet: true}
}

func (v NullableProductListProductsByIdsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListProductsByIdsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


