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

// ProductCreateProductResponse struct for ProductCreateProductResponse
type ProductCreateProductResponse struct {
	Success *bool `json:"success,omitempty"`
	Id *string `json:"id,omitempty"`
	ProductErrors []ProductProductResponseError `json:"productErrors,omitempty"`
	AttributeErrors []ProductAttributeResponseError `json:"attributeErrors,omitempty"`
}

// NewProductCreateProductResponse instantiates a new ProductCreateProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateProductResponse() *ProductCreateProductResponse {
	this := ProductCreateProductResponse{}
	return &this
}

// NewProductCreateProductResponseWithDefaults instantiates a new ProductCreateProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateProductResponseWithDefaults() *ProductCreateProductResponse {
	this := ProductCreateProductResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ProductCreateProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductCreateProductResponse) SetId(v string) {
	o.Id = &v
}

// GetProductErrors returns the ProductErrors field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetProductErrors() []ProductProductResponseError {
	if o == nil || isNil(o.ProductErrors) {
		var ret []ProductProductResponseError
		return ret
	}
	return o.ProductErrors
}

// GetProductErrorsOk returns a tuple with the ProductErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetProductErrorsOk() ([]ProductProductResponseError, bool) {
	if o == nil || isNil(o.ProductErrors) {
    return nil, false
	}
	return o.ProductErrors, true
}

// HasProductErrors returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasProductErrors() bool {
	if o != nil && !isNil(o.ProductErrors) {
		return true
	}

	return false
}

// SetProductErrors gets a reference to the given []ProductProductResponseError and assigns it to the ProductErrors field.
func (o *ProductCreateProductResponse) SetProductErrors(v []ProductProductResponseError) {
	o.ProductErrors = v
}

// GetAttributeErrors returns the AttributeErrors field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetAttributeErrors() []ProductAttributeResponseError {
	if o == nil || isNil(o.AttributeErrors) {
		var ret []ProductAttributeResponseError
		return ret
	}
	return o.AttributeErrors
}

// GetAttributeErrorsOk returns a tuple with the AttributeErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetAttributeErrorsOk() ([]ProductAttributeResponseError, bool) {
	if o == nil || isNil(o.AttributeErrors) {
    return nil, false
	}
	return o.AttributeErrors, true
}

// HasAttributeErrors returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasAttributeErrors() bool {
	if o != nil && !isNil(o.AttributeErrors) {
		return true
	}

	return false
}

// SetAttributeErrors gets a reference to the given []ProductAttributeResponseError and assigns it to the AttributeErrors field.
func (o *ProductCreateProductResponse) SetAttributeErrors(v []ProductAttributeResponseError) {
	o.AttributeErrors = v
}

func (o ProductCreateProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ProductErrors) {
		toSerialize["productErrors"] = o.ProductErrors
	}
	if !isNil(o.AttributeErrors) {
		toSerialize["attributeErrors"] = o.AttributeErrors
	}
	return json.Marshal(toSerialize)
}

type NullableProductCreateProductResponse struct {
	value *ProductCreateProductResponse
	isSet bool
}

func (v NullableProductCreateProductResponse) Get() *ProductCreateProductResponse {
	return v.value
}

func (v *NullableProductCreateProductResponse) Set(val *ProductCreateProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateProductResponse(val *ProductCreateProductResponse) *NullableProductCreateProductResponse {
	return &NullableProductCreateProductResponse{value: val, isSet: true}
}

func (v NullableProductCreateProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


