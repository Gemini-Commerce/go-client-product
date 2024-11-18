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

// checks if the ProductAttributeToEnrich type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttributeToEnrich{}

// ProductAttributeToEnrich struct for ProductAttributeToEnrich
type ProductAttributeToEnrich struct {
	Code                 *string                       `json:"code,omitempty"`
	Type                 *ProductAttributeToEnrichType `json:"type,omitempty"`
	CanCreateValue       *bool                         `json:"canCreateValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAttributeToEnrich ProductAttributeToEnrich

// NewProductAttributeToEnrich instantiates a new ProductAttributeToEnrich object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttributeToEnrich() *ProductAttributeToEnrich {
	this := ProductAttributeToEnrich{}
	var type_ ProductAttributeToEnrichType = PRODUCTATTRIBUTETOENRICHTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewProductAttributeToEnrichWithDefaults instantiates a new ProductAttributeToEnrich object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeToEnrichWithDefaults() *ProductAttributeToEnrich {
	this := ProductAttributeToEnrich{}
	var type_ ProductAttributeToEnrichType = PRODUCTATTRIBUTETOENRICHTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductAttributeToEnrich) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeToEnrich) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductAttributeToEnrich) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductAttributeToEnrich) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAttributeToEnrich) GetType() ProductAttributeToEnrichType {
	if o == nil || IsNil(o.Type) {
		var ret ProductAttributeToEnrichType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeToEnrich) GetTypeOk() (*ProductAttributeToEnrichType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAttributeToEnrich) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProductAttributeToEnrichType and assigns it to the Type field.
func (o *ProductAttributeToEnrich) SetType(v ProductAttributeToEnrichType) {
	o.Type = &v
}

// GetCanCreateValue returns the CanCreateValue field value if set, zero value otherwise.
func (o *ProductAttributeToEnrich) GetCanCreateValue() bool {
	if o == nil || IsNil(o.CanCreateValue) {
		var ret bool
		return ret
	}
	return *o.CanCreateValue
}

// GetCanCreateValueOk returns a tuple with the CanCreateValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeToEnrich) GetCanCreateValueOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateValue) {
		return nil, false
	}
	return o.CanCreateValue, true
}

// HasCanCreateValue returns a boolean if a field has been set.
func (o *ProductAttributeToEnrich) HasCanCreateValue() bool {
	if o != nil && !IsNil(o.CanCreateValue) {
		return true
	}

	return false
}

// SetCanCreateValue gets a reference to the given bool and assigns it to the CanCreateValue field.
func (o *ProductAttributeToEnrich) SetCanCreateValue(v bool) {
	o.CanCreateValue = &v
}

func (o ProductAttributeToEnrich) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAttributeToEnrich) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CanCreateValue) {
		toSerialize["canCreateValue"] = o.CanCreateValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductAttributeToEnrich) UnmarshalJSON(data []byte) (err error) {
	varProductAttributeToEnrich := _ProductAttributeToEnrich{}

	err = json.Unmarshal(data, &varProductAttributeToEnrich)

	if err != nil {
		return err
	}

	*o = ProductAttributeToEnrich(varProductAttributeToEnrich)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "type")
		delete(additionalProperties, "canCreateValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductAttributeToEnrich) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populate the value of well-known types
func (o *ProductAttributeToEnrich) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductAttributeToEnrich struct {
	value *ProductAttributeToEnrich
	isSet bool
}

func (v NullableProductAttributeToEnrich) Get() *ProductAttributeToEnrich {
	return v.value
}

func (v *NullableProductAttributeToEnrich) Set(val *ProductAttributeToEnrich) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeToEnrich) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeToEnrich) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeToEnrich(val *ProductAttributeToEnrich) *NullableProductAttributeToEnrich {
	return &NullableProductAttributeToEnrich{value: val, isSet: true}
}

func (v NullableProductAttributeToEnrich) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeToEnrich) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
