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

// checks if the EntitymanagerGetAttributeOptionsResponseOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerGetAttributeOptionsResponseOption{}

// EntitymanagerGetAttributeOptionsResponseOption struct for EntitymanagerGetAttributeOptionsResponseOption
type EntitymanagerGetAttributeOptionsResponseOption struct {
	ListCode             *string                       `json:"listCode,omitempty"`
	Option               *EntitymanagerAttributeOption `json:"option,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerGetAttributeOptionsResponseOption EntitymanagerGetAttributeOptionsResponseOption

// NewEntitymanagerGetAttributeOptionsResponseOption instantiates a new EntitymanagerGetAttributeOptionsResponseOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerGetAttributeOptionsResponseOption() *EntitymanagerGetAttributeOptionsResponseOption {
	this := EntitymanagerGetAttributeOptionsResponseOption{}
	return &this
}

// NewEntitymanagerGetAttributeOptionsResponseOptionWithDefaults instantiates a new EntitymanagerGetAttributeOptionsResponseOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerGetAttributeOptionsResponseOptionWithDefaults() *EntitymanagerGetAttributeOptionsResponseOption {
	this := EntitymanagerGetAttributeOptionsResponseOption{}
	return &this
}

// GetListCode returns the ListCode field value if set, zero value otherwise.
func (o *EntitymanagerGetAttributeOptionsResponseOption) GetListCode() string {
	if o == nil || IsNil(o.ListCode) {
		var ret string
		return ret
	}
	return *o.ListCode
}

// GetListCodeOk returns a tuple with the ListCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerGetAttributeOptionsResponseOption) GetListCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ListCode) {
		return nil, false
	}
	return o.ListCode, true
}

// HasListCode returns a boolean if a field has been set.
func (o *EntitymanagerGetAttributeOptionsResponseOption) HasListCode() bool {
	if o != nil && !IsNil(o.ListCode) {
		return true
	}

	return false
}

// SetListCode gets a reference to the given string and assigns it to the ListCode field.
func (o *EntitymanagerGetAttributeOptionsResponseOption) SetListCode(v string) {
	o.ListCode = &v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *EntitymanagerGetAttributeOptionsResponseOption) GetOption() EntitymanagerAttributeOption {
	if o == nil || IsNil(o.Option) {
		var ret EntitymanagerAttributeOption
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerGetAttributeOptionsResponseOption) GetOptionOk() (*EntitymanagerAttributeOption, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *EntitymanagerGetAttributeOptionsResponseOption) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given EntitymanagerAttributeOption and assigns it to the Option field.
func (o *EntitymanagerGetAttributeOptionsResponseOption) SetOption(v EntitymanagerAttributeOption) {
	o.Option = &v
}

func (o EntitymanagerGetAttributeOptionsResponseOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerGetAttributeOptionsResponseOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListCode) {
		toSerialize["listCode"] = o.ListCode
	}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerGetAttributeOptionsResponseOption) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerGetAttributeOptionsResponseOption := _EntitymanagerGetAttributeOptionsResponseOption{}

	err = json.Unmarshal(data, &varEntitymanagerGetAttributeOptionsResponseOption)

	if err != nil {
		return err
	}

	*o = EntitymanagerGetAttributeOptionsResponseOption(varEntitymanagerGetAttributeOptionsResponseOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listCode")
		delete(additionalProperties, "option")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EntitymanagerGetAttributeOptionsResponseOption) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EntitymanagerGetAttributeOptionsResponseOption) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEntitymanagerGetAttributeOptionsResponseOption struct {
	value *EntitymanagerGetAttributeOptionsResponseOption
	isSet bool
}

func (v NullableEntitymanagerGetAttributeOptionsResponseOption) Get() *EntitymanagerGetAttributeOptionsResponseOption {
	return v.value
}

func (v *NullableEntitymanagerGetAttributeOptionsResponseOption) Set(val *EntitymanagerGetAttributeOptionsResponseOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerGetAttributeOptionsResponseOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerGetAttributeOptionsResponseOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerGetAttributeOptionsResponseOption(val *EntitymanagerGetAttributeOptionsResponseOption) *NullableEntitymanagerGetAttributeOptionsResponseOption {
	return &NullableEntitymanagerGetAttributeOptionsResponseOption{value: val, isSet: true}
}

func (v NullableEntitymanagerGetAttributeOptionsResponseOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerGetAttributeOptionsResponseOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
