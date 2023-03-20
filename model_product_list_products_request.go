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

// ProductListProductsRequest struct for ProductListProductsRequest
type ProductListProductsRequest struct {
	TenantId  *string `json:"tenantId,omitempty"`
	PageSize  *string `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewProductListProductsRequest instantiates a new ProductListProductsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListProductsRequest() *ProductListProductsRequest {
	this := ProductListProductsRequest{}
	return &this
}

// NewProductListProductsRequestWithDefaults instantiates a new ProductListProductsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListProductsRequestWithDefaults() *ProductListProductsRequest {
	this := ProductListProductsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductListProductsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductListProductsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductListProductsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductListProductsRequest) GetPageSize() string {
	if o == nil || isNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsRequest) GetPageSizeOk() (*string, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProductListProductsRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *ProductListProductsRequest) SetPageSize(v string) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductListProductsRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ProductListProductsRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductListProductsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ProductListProductsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableProductListProductsRequest struct {
	value *ProductListProductsRequest
	isSet bool
}

func (v NullableProductListProductsRequest) Get() *ProductListProductsRequest {
	return v.value
}

func (v *NullableProductListProductsRequest) Set(val *ProductListProductsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListProductsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListProductsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListProductsRequest(val *ProductListProductsRequest) *NullableProductListProductsRequest {
	return &NullableProductListProductsRequest{value: val, isSet: true}
}

func (v NullableProductListProductsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListProductsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
