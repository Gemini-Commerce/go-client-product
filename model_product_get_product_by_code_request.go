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

// checks if the ProductGetProductByCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductGetProductByCodeRequest{}

// ProductGetProductByCodeRequest struct for ProductGetProductByCodeRequest
type ProductGetProductByCodeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewProductGetProductByCodeRequest instantiates a new ProductGetProductByCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGetProductByCodeRequest() *ProductGetProductByCodeRequest {
	this := ProductGetProductByCodeRequest{}
	return &this
}

// NewProductGetProductByCodeRequestWithDefaults instantiates a new ProductGetProductByCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGetProductByCodeRequestWithDefaults() *ProductGetProductByCodeRequest {
	this := ProductGetProductByCodeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductGetProductByCodeRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByCodeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductGetProductByCodeRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductGetProductByCodeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductGetProductByCodeRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductGetProductByCodeRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductGetProductByCodeRequest) SetCode(v string) {
	o.Code = &v
}

func (o ProductGetProductByCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductGetProductByCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableProductGetProductByCodeRequest struct {
	value *ProductGetProductByCodeRequest
	isSet bool
}

func (v NullableProductGetProductByCodeRequest) Get() *ProductGetProductByCodeRequest {
	return v.value
}

func (v *NullableProductGetProductByCodeRequest) Set(val *ProductGetProductByCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGetProductByCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGetProductByCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGetProductByCodeRequest(val *ProductGetProductByCodeRequest) *NullableProductGetProductByCodeRequest {
	return &NullableProductGetProductByCodeRequest{value: val, isSet: true}
}

func (v NullableProductGetProductByCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGetProductByCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


