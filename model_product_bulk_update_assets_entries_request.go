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

// checks if the ProductBulkUpdateAssetsEntriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBulkUpdateAssetsEntriesRequest{}

// ProductBulkUpdateAssetsEntriesRequest struct for ProductBulkUpdateAssetsEntriesRequest
type ProductBulkUpdateAssetsEntriesRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	Assets []BulkUpdateAssetsEntriesRequestUpdateEntity `json:"assets,omitempty"`
}

// NewProductBulkUpdateAssetsEntriesRequest instantiates a new ProductBulkUpdateAssetsEntriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateAssetsEntriesRequest() *ProductBulkUpdateAssetsEntriesRequest {
	this := ProductBulkUpdateAssetsEntriesRequest{}
	return &this
}

// NewProductBulkUpdateAssetsEntriesRequestWithDefaults instantiates a new ProductBulkUpdateAssetsEntriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateAssetsEntriesRequestWithDefaults() *ProductBulkUpdateAssetsEntriesRequest {
	this := ProductBulkUpdateAssetsEntriesRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductBulkUpdateAssetsEntriesRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductBulkUpdateAssetsEntriesRequest) SetProductId(v string) {
	o.ProductId = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetAssets() []BulkUpdateAssetsEntriesRequestUpdateEntity {
	if o == nil || IsNil(o.Assets) {
		var ret []BulkUpdateAssetsEntriesRequestUpdateEntity
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) GetAssetsOk() ([]BulkUpdateAssetsEntriesRequestUpdateEntity, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ProductBulkUpdateAssetsEntriesRequest) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []BulkUpdateAssetsEntriesRequestUpdateEntity and assigns it to the Assets field.
func (o *ProductBulkUpdateAssetsEntriesRequest) SetAssets(v []BulkUpdateAssetsEntriesRequestUpdateEntity) {
	o.Assets = v
}

func (o ProductBulkUpdateAssetsEntriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBulkUpdateAssetsEntriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	return toSerialize, nil
}

type NullableProductBulkUpdateAssetsEntriesRequest struct {
	value *ProductBulkUpdateAssetsEntriesRequest
	isSet bool
}

func (v NullableProductBulkUpdateAssetsEntriesRequest) Get() *ProductBulkUpdateAssetsEntriesRequest {
	return v.value
}

func (v *NullableProductBulkUpdateAssetsEntriesRequest) Set(val *ProductBulkUpdateAssetsEntriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateAssetsEntriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateAssetsEntriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateAssetsEntriesRequest(val *ProductBulkUpdateAssetsEntriesRequest) *NullableProductBulkUpdateAssetsEntriesRequest {
	return &NullableProductBulkUpdateAssetsEntriesRequest{value: val, isSet: true}
}

func (v NullableProductBulkUpdateAssetsEntriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateAssetsEntriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

