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

// checks if the ProductUpdateProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateProductResponse{}

// ProductUpdateProductResponse struct for ProductUpdateProductResponse
type ProductUpdateProductResponse struct {
	Success *bool `json:"success,omitempty"`
	ProductErrors []ProductProductResponseError `json:"productErrors,omitempty"`
	AttributeErrors []ProductAttributeResponseError `json:"attributeErrors,omitempty"`
}

// NewProductUpdateProductResponse instantiates a new ProductUpdateProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateProductResponse() *ProductUpdateProductResponse {
	this := ProductUpdateProductResponse{}
	return &this
}

// NewProductUpdateProductResponseWithDefaults instantiates a new ProductUpdateProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateProductResponseWithDefaults() *ProductUpdateProductResponse {
	this := ProductUpdateProductResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ProductUpdateProductResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ProductUpdateProductResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ProductUpdateProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetProductErrors returns the ProductErrors field value if set, zero value otherwise.
func (o *ProductUpdateProductResponse) GetProductErrors() []ProductProductResponseError {
	if o == nil || IsNil(o.ProductErrors) {
		var ret []ProductProductResponseError
		return ret
	}
	return o.ProductErrors
}

// GetProductErrorsOk returns a tuple with the ProductErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductResponse) GetProductErrorsOk() ([]ProductProductResponseError, bool) {
	if o == nil || IsNil(o.ProductErrors) {
		return nil, false
	}
	return o.ProductErrors, true
}

// HasProductErrors returns a boolean if a field has been set.
func (o *ProductUpdateProductResponse) HasProductErrors() bool {
	if o != nil && !IsNil(o.ProductErrors) {
		return true
	}

	return false
}

// SetProductErrors gets a reference to the given []ProductProductResponseError and assigns it to the ProductErrors field.
func (o *ProductUpdateProductResponse) SetProductErrors(v []ProductProductResponseError) {
	o.ProductErrors = v
}

// GetAttributeErrors returns the AttributeErrors field value if set, zero value otherwise.
func (o *ProductUpdateProductResponse) GetAttributeErrors() []ProductAttributeResponseError {
	if o == nil || IsNil(o.AttributeErrors) {
		var ret []ProductAttributeResponseError
		return ret
	}
	return o.AttributeErrors
}

// GetAttributeErrorsOk returns a tuple with the AttributeErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductResponse) GetAttributeErrorsOk() ([]ProductAttributeResponseError, bool) {
	if o == nil || IsNil(o.AttributeErrors) {
		return nil, false
	}
	return o.AttributeErrors, true
}

// HasAttributeErrors returns a boolean if a field has been set.
func (o *ProductUpdateProductResponse) HasAttributeErrors() bool {
	if o != nil && !IsNil(o.AttributeErrors) {
		return true
	}

	return false
}

// SetAttributeErrors gets a reference to the given []ProductAttributeResponseError and assigns it to the AttributeErrors field.
func (o *ProductUpdateProductResponse) SetAttributeErrors(v []ProductAttributeResponseError) {
	o.AttributeErrors = v
}

func (o ProductUpdateProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.ProductErrors) {
		toSerialize["productErrors"] = o.ProductErrors
	}
	if !IsNil(o.AttributeErrors) {
		toSerialize["attributeErrors"] = o.AttributeErrors
	}
	return toSerialize, nil
}

type NullableProductUpdateProductResponse struct {
	value *ProductUpdateProductResponse
	isSet bool
}

func (v NullableProductUpdateProductResponse) Get() *ProductUpdateProductResponse {
	return v.value
}

func (v *NullableProductUpdateProductResponse) Set(val *ProductUpdateProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateProductResponse(val *ProductUpdateProductResponse) *NullableProductUpdateProductResponse {
	return &NullableProductUpdateProductResponse{value: val, isSet: true}
}

func (v NullableProductUpdateProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


