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

// checks if the ProductListProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductListProductsResponse{}

// ProductListProductsResponse struct for ProductListProductsResponse
type ProductListProductsResponse struct {
	Products []ProductProductEntity `json:"products,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewProductListProductsResponse instantiates a new ProductListProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListProductsResponse() *ProductListProductsResponse {
	this := ProductListProductsResponse{}
	return &this
}

// NewProductListProductsResponseWithDefaults instantiates a new ProductListProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListProductsResponseWithDefaults() *ProductListProductsResponse {
	this := ProductListProductsResponse{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ProductListProductsResponse) GetProducts() []ProductProductEntity {
	if o == nil || IsNil(o.Products) {
		var ret []ProductProductEntity
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsResponse) GetProductsOk() ([]ProductProductEntity, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ProductListProductsResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductProductEntity and assigns it to the Products field.
func (o *ProductListProductsResponse) SetProducts(v []ProductProductEntity) {
	o.Products = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ProductListProductsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListProductsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ProductListProductsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ProductListProductsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ProductListProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductListProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableProductListProductsResponse struct {
	value *ProductListProductsResponse
	isSet bool
}

func (v NullableProductListProductsResponse) Get() *ProductListProductsResponse {
	return v.value
}

func (v *NullableProductListProductsResponse) Set(val *ProductListProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListProductsResponse(val *ProductListProductsResponse) *NullableProductListProductsResponse {
	return &NullableProductListProductsResponse{value: val, isSet: true}
}

func (v NullableProductListProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


