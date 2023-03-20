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

// ProductGetProductByUrlKeyRequest struct for ProductGetProductByUrlKeyRequest
type ProductGetProductByUrlKeyRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	UrlKey   *string `json:"urlKey,omitempty"`
	Locale   *string `json:"locale,omitempty"`
}

// NewProductGetProductByUrlKeyRequest instantiates a new ProductGetProductByUrlKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGetProductByUrlKeyRequest() *ProductGetProductByUrlKeyRequest {
	this := ProductGetProductByUrlKeyRequest{}
	return &this
}

// NewProductGetProductByUrlKeyRequestWithDefaults instantiates a new ProductGetProductByUrlKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGetProductByUrlKeyRequestWithDefaults() *ProductGetProductByUrlKeyRequest {
	this := ProductGetProductByUrlKeyRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductGetProductByUrlKeyRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByUrlKeyRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductGetProductByUrlKeyRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductGetProductByUrlKeyRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductGetProductByUrlKeyRequest) GetUrlKey() string {
	if o == nil || isNil(o.UrlKey) {
		var ret string
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByUrlKeyRequest) GetUrlKeyOk() (*string, bool) {
	if o == nil || isNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductGetProductByUrlKeyRequest) HasUrlKey() bool {
	if o != nil && !isNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given string and assigns it to the UrlKey field.
func (o *ProductGetProductByUrlKeyRequest) SetUrlKey(v string) {
	o.UrlKey = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ProductGetProductByUrlKeyRequest) GetLocale() string {
	if o == nil || isNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGetProductByUrlKeyRequest) GetLocaleOk() (*string, bool) {
	if o == nil || isNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ProductGetProductByUrlKeyRequest) HasLocale() bool {
	if o != nil && !isNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ProductGetProductByUrlKeyRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o ProductGetProductByUrlKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !isNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return json.Marshal(toSerialize)
}

type NullableProductGetProductByUrlKeyRequest struct {
	value *ProductGetProductByUrlKeyRequest
	isSet bool
}

func (v NullableProductGetProductByUrlKeyRequest) Get() *ProductGetProductByUrlKeyRequest {
	return v.value
}

func (v *NullableProductGetProductByUrlKeyRequest) Set(val *ProductGetProductByUrlKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGetProductByUrlKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGetProductByUrlKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGetProductByUrlKeyRequest(val *ProductGetProductByUrlKeyRequest) *NullableProductGetProductByUrlKeyRequest {
	return &NullableProductGetProductByUrlKeyRequest{value: val, isSet: true}
}

func (v NullableProductGetProductByUrlKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGetProductByUrlKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
