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
	"fmt"
)

// ProductAttributeInReviewJobType the model 'ProductAttributeInReviewJobType'
type ProductAttributeInReviewJobType string

// List of productAttributeInReviewJobType
const (
	PRODUCTATTRIBUTEINREVIEWJOBTYPE_UNKNOWN     ProductAttributeInReviewJobType = "ATTRIBUTE_IN_REVIEW_JOB_TYPE_UNKNOWN"
	PRODUCTATTRIBUTEINREVIEWJOBTYPE_ENRICHMENT  ProductAttributeInReviewJobType = "ATTRIBUTE_IN_REVIEW_JOB_TYPE_ENRICHMENT"
	PRODUCTATTRIBUTEINREVIEWJOBTYPE_TRANSLATION ProductAttributeInReviewJobType = "ATTRIBUTE_IN_REVIEW_JOB_TYPE_TRANSLATION"
)

// All allowed values of ProductAttributeInReviewJobType enum
var AllowedProductAttributeInReviewJobTypeEnumValues = []ProductAttributeInReviewJobType{
	"ATTRIBUTE_IN_REVIEW_JOB_TYPE_UNKNOWN",
	"ATTRIBUTE_IN_REVIEW_JOB_TYPE_ENRICHMENT",
	"ATTRIBUTE_IN_REVIEW_JOB_TYPE_TRANSLATION",
}

func (v *ProductAttributeInReviewJobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductAttributeInReviewJobType(value)
	for _, existing := range AllowedProductAttributeInReviewJobTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductAttributeInReviewJobType", value)
}

// NewProductAttributeInReviewJobTypeFromValue returns a pointer to a valid ProductAttributeInReviewJobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductAttributeInReviewJobTypeFromValue(v string) (*ProductAttributeInReviewJobType, error) {
	ev := ProductAttributeInReviewJobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductAttributeInReviewJobType: valid values are %v", v, AllowedProductAttributeInReviewJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductAttributeInReviewJobType) IsValid() bool {
	for _, existing := range AllowedProductAttributeInReviewJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productAttributeInReviewJobType value
func (v ProductAttributeInReviewJobType) Ptr() *ProductAttributeInReviewJobType {
	return &v
}

type NullableProductAttributeInReviewJobType struct {
	value *ProductAttributeInReviewJobType
	isSet bool
}

func (v NullableProductAttributeInReviewJobType) Get() *ProductAttributeInReviewJobType {
	return v.value
}

func (v *NullableProductAttributeInReviewJobType) Set(val *ProductAttributeInReviewJobType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeInReviewJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeInReviewJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeInReviewJobType(val *ProductAttributeInReviewJobType) *NullableProductAttributeInReviewJobType {
	return &NullableProductAttributeInReviewJobType{value: val, isSet: true}
}

func (v NullableProductAttributeInReviewJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeInReviewJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
