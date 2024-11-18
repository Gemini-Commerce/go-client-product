/*
Product Service

Introducing our revolutionary Product Management Service! Designed to streamline your product inventory and elevate customer experiences, our cutting-edge protobuf service is a game-changer in the world of efficient product management.  With our service, you can effortlessly create new products, allowing you to quickly bring your ideas to life and expand your catalog. Retrieve product information in a snap, providing accurate and personalized details to your customers based on their specific needs and preferences.  Stay ahead of the competition by easily updating product information, ensuring your catalog is always up-to-date and optimized. Seamlessly remove products from your inventory when needed, maintaining a clean and relevant product selection.  Enhance the visual appeal of your products with our advanced media gallery functionalities. Effortlessly add and update captivating images and videos to showcase your products in the best possible light, engaging your customers and driving conversions.  Personalization is key in today's market, and our service enables you to offer unique options to your customers. Easily create and manage lists of customizable options for your products, providing flexibility and tailoring to individual preferences.  Attributes play a vital role in defining products, and our service empowers you to effectively manage them. From bulk attribute creation to listing and retrieving attribute options, our service ensures your product information is rich and comprehensive.  Our service extends its capabilities to entity management, allowing you to effortlessly handle different entities and create customized options lists associated with them. This provides further flexibility and customization options for your product offerings.  When it comes to bulk updates, our service has you covered. Effortlessly update multiple products simultaneously, saving you time and streamlining your operations.  Finding specific products and variants is a breeze with our service. Quickly locate products based on their unique stock keeping unit (SKU) values, ensuring efficient inventory management and smooth order fulfillment.  Experience a new level of efficiency and productivity with our Product Management Service. Unlock the full potential of streamlined product management and empower your business to thrive in today's competitive market. Try our service today and elevate your product management to new heights!

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
	Success              *bool                           `json:"success,omitempty"`
	ProductErrors        []ProductProductResponseError   `json:"productErrors,omitempty"`
	AttributeErrors      []ProductAttributeResponseError `json:"attributeErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductUpdateProductResponse ProductUpdateProductResponse

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
	toSerialize, err := o.ToMap()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductUpdateProductResponse) UnmarshalJSON(data []byte) (err error) {
	varProductUpdateProductResponse := _ProductUpdateProductResponse{}

	err = json.Unmarshal(data, &varProductUpdateProductResponse)

	if err != nil {
		return err
	}

	*o = ProductUpdateProductResponse(varProductUpdateProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "productErrors")
		delete(additionalProperties, "attributeErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductUpdateProductResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductUpdateProductResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
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
