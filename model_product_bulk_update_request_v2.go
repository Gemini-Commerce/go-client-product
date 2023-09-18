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

// ProductBulkUpdateRequestV2 struct for ProductBulkUpdateRequestV2
type ProductBulkUpdateRequestV2 struct {
	TenantId *string `json:"tenantId,omitempty"`
	ProductIds []string `json:"productIds,omitempty"`
	Payload *ProductBulkUpdateRequestV2Payload `json:"payload,omitempty"`
}

// NewProductBulkUpdateRequestV2 instantiates a new ProductBulkUpdateRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateRequestV2() *ProductBulkUpdateRequestV2 {
	this := ProductBulkUpdateRequestV2{}
	return &this
}

// NewProductBulkUpdateRequestV2WithDefaults instantiates a new ProductBulkUpdateRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateRequestV2WithDefaults() *ProductBulkUpdateRequestV2 {
	this := ProductBulkUpdateRequestV2{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestV2) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestV2) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestV2) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductBulkUpdateRequestV2) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestV2) GetProductIds() []string {
	if o == nil || isNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestV2) GetProductIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ProductIds) {
    return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestV2) HasProductIds() bool {
	if o != nil && !isNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *ProductBulkUpdateRequestV2) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestV2) GetPayload() ProductBulkUpdateRequestV2Payload {
	if o == nil || isNil(o.Payload) {
		var ret ProductBulkUpdateRequestV2Payload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestV2) GetPayloadOk() (*ProductBulkUpdateRequestV2Payload, bool) {
	if o == nil || isNil(o.Payload) {
    return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestV2) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductBulkUpdateRequestV2Payload and assigns it to the Payload field.
func (o *ProductBulkUpdateRequestV2) SetPayload(v ProductBulkUpdateRequestV2Payload) {
	o.Payload = &v
}

func (o ProductBulkUpdateRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ProductIds) {
		toSerialize["productIds"] = o.ProductIds
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableProductBulkUpdateRequestV2 struct {
	value *ProductBulkUpdateRequestV2
	isSet bool
}

func (v NullableProductBulkUpdateRequestV2) Get() *ProductBulkUpdateRequestV2 {
	return v.value
}

func (v *NullableProductBulkUpdateRequestV2) Set(val *ProductBulkUpdateRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateRequestV2(val *ProductBulkUpdateRequestV2) *NullableProductBulkUpdateRequestV2 {
	return &NullableProductBulkUpdateRequestV2{value: val, isSet: true}
}

func (v NullableProductBulkUpdateRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


