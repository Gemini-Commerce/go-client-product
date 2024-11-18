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

// checks if the ProductBulkUpdateRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBulkUpdateRequestPayload{}

// ProductBulkUpdateRequestPayload struct for ProductBulkUpdateRequestPayload
type ProductBulkUpdateRequestPayload struct {
	UpdateMask           *ProductFieldMask       `json:"updateMask,omitempty"`
	Attributes           *map[string]ProtobufAny `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductBulkUpdateRequestPayload ProductBulkUpdateRequestPayload

// NewProductBulkUpdateRequestPayload instantiates a new ProductBulkUpdateRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateRequestPayload() *ProductBulkUpdateRequestPayload {
	this := ProductBulkUpdateRequestPayload{}
	return &this
}

// NewProductBulkUpdateRequestPayloadWithDefaults instantiates a new ProductBulkUpdateRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateRequestPayloadWithDefaults() *ProductBulkUpdateRequestPayload {
	this := ProductBulkUpdateRequestPayload{}
	return &this
}

// GetUpdateMask returns the UpdateMask field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestPayload) GetUpdateMask() ProductFieldMask {
	if o == nil || IsNil(o.UpdateMask) {
		var ret ProductFieldMask
		return ret
	}
	return *o.UpdateMask
}

// GetUpdateMaskOk returns a tuple with the UpdateMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestPayload) GetUpdateMaskOk() (*ProductFieldMask, bool) {
	if o == nil || IsNil(o.UpdateMask) {
		return nil, false
	}
	return o.UpdateMask, true
}

// HasUpdateMask returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestPayload) HasUpdateMask() bool {
	if o != nil && !IsNil(o.UpdateMask) {
		return true
	}

	return false
}

// SetUpdateMask gets a reference to the given ProductFieldMask and assigns it to the UpdateMask field.
func (o *ProductBulkUpdateRequestPayload) SetUpdateMask(v ProductFieldMask) {
	o.UpdateMask = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequestPayload) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequestPayload) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequestPayload) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductBulkUpdateRequestPayload) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o ProductBulkUpdateRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBulkUpdateRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateMask) {
		toSerialize["updateMask"] = o.UpdateMask
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductBulkUpdateRequestPayload) UnmarshalJSON(data []byte) (err error) {
	varProductBulkUpdateRequestPayload := _ProductBulkUpdateRequestPayload{}

	err = json.Unmarshal(data, &varProductBulkUpdateRequestPayload)

	if err != nil {
		return err
	}

	*o = ProductBulkUpdateRequestPayload(varProductBulkUpdateRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "updateMask")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductBulkUpdateRequestPayload struct {
	value *ProductBulkUpdateRequestPayload
	isSet bool
}

func (v NullableProductBulkUpdateRequestPayload) Get() *ProductBulkUpdateRequestPayload {
	return v.value
}

func (v *NullableProductBulkUpdateRequestPayload) Set(val *ProductBulkUpdateRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateRequestPayload(val *ProductBulkUpdateRequestPayload) *NullableProductBulkUpdateRequestPayload {
	return &NullableProductBulkUpdateRequestPayload{value: val, isSet: true}
}

func (v NullableProductBulkUpdateRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
