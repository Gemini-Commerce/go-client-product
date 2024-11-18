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

// checks if the ProductAttributeInReviewError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttributeInReviewError{}

// ProductAttributeInReviewError struct for ProductAttributeInReviewError
type ProductAttributeInReviewError struct {
	Code                 *string `json:"code,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAttributeInReviewError ProductAttributeInReviewError

// NewProductAttributeInReviewError instantiates a new ProductAttributeInReviewError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttributeInReviewError() *ProductAttributeInReviewError {
	this := ProductAttributeInReviewError{}
	return &this
}

// NewProductAttributeInReviewErrorWithDefaults instantiates a new ProductAttributeInReviewError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeInReviewErrorWithDefaults() *ProductAttributeInReviewError {
	this := ProductAttributeInReviewError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductAttributeInReviewError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReviewError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductAttributeInReviewError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductAttributeInReviewError) SetCode(v string) {
	o.Code = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ProductAttributeInReviewError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReviewError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ProductAttributeInReviewError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ProductAttributeInReviewError) SetReason(v string) {
	o.Reason = &v
}

func (o ProductAttributeInReviewError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAttributeInReviewError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductAttributeInReviewError) UnmarshalJSON(data []byte) (err error) {
	varProductAttributeInReviewError := _ProductAttributeInReviewError{}

	err = json.Unmarshal(data, &varProductAttributeInReviewError)

	if err != nil {
		return err
	}

	*o = ProductAttributeInReviewError(varProductAttributeInReviewError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductAttributeInReviewError struct {
	value *ProductAttributeInReviewError
	isSet bool
}

func (v NullableProductAttributeInReviewError) Get() *ProductAttributeInReviewError {
	return v.value
}

func (v *NullableProductAttributeInReviewError) Set(val *ProductAttributeInReviewError) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeInReviewError) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeInReviewError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeInReviewError(val *ProductAttributeInReviewError) *NullableProductAttributeInReviewError {
	return &NullableProductAttributeInReviewError{value: val, isSet: true}
}

func (v NullableProductAttributeInReviewError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeInReviewError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
