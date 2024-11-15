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

// checks if the ProductEnrichAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductEnrichAction{}

// ProductEnrichAction struct for ProductEnrichAction
type ProductEnrichAction struct {
	AttributesToEnrich []ProductAttributeToEnrich `json:"attributesToEnrich,omitempty"`
	GenerationLanguage *ProductLanguageCode `json:"generationLanguage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductEnrichAction ProductEnrichAction

// NewProductEnrichAction instantiates a new ProductEnrichAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductEnrichAction() *ProductEnrichAction {
	this := ProductEnrichAction{}
	var generationLanguage ProductLanguageCode = PRODUCTLANGUAGECODE_UNKNOWN
	this.GenerationLanguage = &generationLanguage
	return &this
}

// NewProductEnrichActionWithDefaults instantiates a new ProductEnrichAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductEnrichActionWithDefaults() *ProductEnrichAction {
	this := ProductEnrichAction{}
	var generationLanguage ProductLanguageCode = PRODUCTLANGUAGECODE_UNKNOWN
	this.GenerationLanguage = &generationLanguage
	return &this
}

// GetAttributesToEnrich returns the AttributesToEnrich field value if set, zero value otherwise.
func (o *ProductEnrichAction) GetAttributesToEnrich() []ProductAttributeToEnrich {
	if o == nil || IsNil(o.AttributesToEnrich) {
		var ret []ProductAttributeToEnrich
		return ret
	}
	return o.AttributesToEnrich
}

// GetAttributesToEnrichOk returns a tuple with the AttributesToEnrich field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEnrichAction) GetAttributesToEnrichOk() ([]ProductAttributeToEnrich, bool) {
	if o == nil || IsNil(o.AttributesToEnrich) {
		return nil, false
	}
	return o.AttributesToEnrich, true
}

// HasAttributesToEnrich returns a boolean if a field has been set.
func (o *ProductEnrichAction) HasAttributesToEnrich() bool {
	if o != nil && !IsNil(o.AttributesToEnrich) {
		return true
	}

	return false
}

// SetAttributesToEnrich gets a reference to the given []ProductAttributeToEnrich and assigns it to the AttributesToEnrich field.
func (o *ProductEnrichAction) SetAttributesToEnrich(v []ProductAttributeToEnrich) {
	o.AttributesToEnrich = v
}

// GetGenerationLanguage returns the GenerationLanguage field value if set, zero value otherwise.
func (o *ProductEnrichAction) GetGenerationLanguage() ProductLanguageCode {
	if o == nil || IsNil(o.GenerationLanguage) {
		var ret ProductLanguageCode
		return ret
	}
	return *o.GenerationLanguage
}

// GetGenerationLanguageOk returns a tuple with the GenerationLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEnrichAction) GetGenerationLanguageOk() (*ProductLanguageCode, bool) {
	if o == nil || IsNil(o.GenerationLanguage) {
		return nil, false
	}
	return o.GenerationLanguage, true
}

// HasGenerationLanguage returns a boolean if a field has been set.
func (o *ProductEnrichAction) HasGenerationLanguage() bool {
	if o != nil && !IsNil(o.GenerationLanguage) {
		return true
	}

	return false
}

// SetGenerationLanguage gets a reference to the given ProductLanguageCode and assigns it to the GenerationLanguage field.
func (o *ProductEnrichAction) SetGenerationLanguage(v ProductLanguageCode) {
	o.GenerationLanguage = &v
}

func (o ProductEnrichAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductEnrichAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributesToEnrich) {
		toSerialize["attributesToEnrich"] = o.AttributesToEnrich
	}
	if !IsNil(o.GenerationLanguage) {
		toSerialize["generationLanguage"] = o.GenerationLanguage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductEnrichAction) UnmarshalJSON(data []byte) (err error) {
	varProductEnrichAction := _ProductEnrichAction{}

	err = json.Unmarshal(data, &varProductEnrichAction)

	if err != nil {
		return err
	}

	*o = ProductEnrichAction(varProductEnrichAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributesToEnrich")
		delete(additionalProperties, "generationLanguage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductEnrichAction) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductEnrichAction) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductEnrichAction struct {
	value *ProductEnrichAction
	isSet bool
}

func (v NullableProductEnrichAction) Get() *ProductEnrichAction {
	return v.value
}

func (v *NullableProductEnrichAction) Set(val *ProductEnrichAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProductEnrichAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProductEnrichAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductEnrichAction(val *ProductEnrichAction) *NullableProductEnrichAction {
	return &NullableProductEnrichAction{value: val, isSet: true}
}

func (v NullableProductEnrichAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductEnrichAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


