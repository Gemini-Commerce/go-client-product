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

// ProductBulkUpdateResponse struct for ProductBulkUpdateResponse
type ProductBulkUpdateResponse struct {
	ProductResponse []BulkUpdateResponseResponse `json:"productResponse,omitempty"`
}

// NewProductBulkUpdateResponse instantiates a new ProductBulkUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateResponse() *ProductBulkUpdateResponse {
	this := ProductBulkUpdateResponse{}
	return &this
}

// NewProductBulkUpdateResponseWithDefaults instantiates a new ProductBulkUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateResponseWithDefaults() *ProductBulkUpdateResponse {
	this := ProductBulkUpdateResponse{}
	return &this
}

// GetProductResponse returns the ProductResponse field value if set, zero value otherwise.
func (o *ProductBulkUpdateResponse) GetProductResponse() []BulkUpdateResponseResponse {
	if o == nil || isNil(o.ProductResponse) {
		var ret []BulkUpdateResponseResponse
		return ret
	}
	return o.ProductResponse
}

// GetProductResponseOk returns a tuple with the ProductResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateResponse) GetProductResponseOk() ([]BulkUpdateResponseResponse, bool) {
	if o == nil || isNil(o.ProductResponse) {
    return nil, false
	}
	return o.ProductResponse, true
}

// HasProductResponse returns a boolean if a field has been set.
func (o *ProductBulkUpdateResponse) HasProductResponse() bool {
	if o != nil && !isNil(o.ProductResponse) {
		return true
	}

	return false
}

// SetProductResponse gets a reference to the given []BulkUpdateResponseResponse and assigns it to the ProductResponse field.
func (o *ProductBulkUpdateResponse) SetProductResponse(v []BulkUpdateResponseResponse) {
	o.ProductResponse = v
}

func (o ProductBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProductResponse) {
		toSerialize["productResponse"] = o.ProductResponse
	}
	return json.Marshal(toSerialize)
}

type NullableProductBulkUpdateResponse struct {
	value *ProductBulkUpdateResponse
	isSet bool
}

func (v NullableProductBulkUpdateResponse) Get() *ProductBulkUpdateResponse {
	return v.value
}

func (v *NullableProductBulkUpdateResponse) Set(val *ProductBulkUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateResponse(val *ProductBulkUpdateResponse) *NullableProductBulkUpdateResponse {
	return &NullableProductBulkUpdateResponse{value: val, isSet: true}
}

func (v NullableProductBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


