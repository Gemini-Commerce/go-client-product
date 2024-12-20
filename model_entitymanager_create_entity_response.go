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

// checks if the EntitymanagerCreateEntityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerCreateEntityResponse{}

// EntitymanagerCreateEntityResponse struct for EntitymanagerCreateEntityResponse
type EntitymanagerCreateEntityResponse struct {
	AttributeWriteErrors *EntitymanagerAttributeWriteErrors `json:"attributeWriteErrors,omitempty"`
	Entity               *EntitymanagerEntity               `json:"entity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerCreateEntityResponse EntitymanagerCreateEntityResponse

// NewEntitymanagerCreateEntityResponse instantiates a new EntitymanagerCreateEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateEntityResponse() *EntitymanagerCreateEntityResponse {
	this := EntitymanagerCreateEntityResponse{}
	return &this
}

// NewEntitymanagerCreateEntityResponseWithDefaults instantiates a new EntitymanagerCreateEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateEntityResponseWithDefaults() *EntitymanagerCreateEntityResponse {
	this := EntitymanagerCreateEntityResponse{}
	return &this
}

// GetAttributeWriteErrors returns the AttributeWriteErrors field value if set, zero value otherwise.
func (o *EntitymanagerCreateEntityResponse) GetAttributeWriteErrors() EntitymanagerAttributeWriteErrors {
	if o == nil || IsNil(o.AttributeWriteErrors) {
		var ret EntitymanagerAttributeWriteErrors
		return ret
	}
	return *o.AttributeWriteErrors
}

// GetAttributeWriteErrorsOk returns a tuple with the AttributeWriteErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateEntityResponse) GetAttributeWriteErrorsOk() (*EntitymanagerAttributeWriteErrors, bool) {
	if o == nil || IsNil(o.AttributeWriteErrors) {
		return nil, false
	}
	return o.AttributeWriteErrors, true
}

// HasAttributeWriteErrors returns a boolean if a field has been set.
func (o *EntitymanagerCreateEntityResponse) HasAttributeWriteErrors() bool {
	if o != nil && !IsNil(o.AttributeWriteErrors) {
		return true
	}

	return false
}

// SetAttributeWriteErrors gets a reference to the given EntitymanagerAttributeWriteErrors and assigns it to the AttributeWriteErrors field.
func (o *EntitymanagerCreateEntityResponse) SetAttributeWriteErrors(v EntitymanagerAttributeWriteErrors) {
	o.AttributeWriteErrors = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntitymanagerCreateEntityResponse) GetEntity() EntitymanagerEntity {
	if o == nil || IsNil(o.Entity) {
		var ret EntitymanagerEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateEntityResponse) GetEntityOk() (*EntitymanagerEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntitymanagerCreateEntityResponse) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntitymanagerEntity and assigns it to the Entity field.
func (o *EntitymanagerCreateEntityResponse) SetEntity(v EntitymanagerEntity) {
	o.Entity = &v
}

func (o EntitymanagerCreateEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerCreateEntityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeWriteErrors) {
		toSerialize["attributeWriteErrors"] = o.AttributeWriteErrors
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerCreateEntityResponse) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerCreateEntityResponse := _EntitymanagerCreateEntityResponse{}

	err = json.Unmarshal(data, &varEntitymanagerCreateEntityResponse)

	if err != nil {
		return err
	}

	*o = EntitymanagerCreateEntityResponse(varEntitymanagerCreateEntityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeWriteErrors")
		delete(additionalProperties, "entity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EntitymanagerCreateEntityResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EntitymanagerCreateEntityResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEntitymanagerCreateEntityResponse struct {
	value *EntitymanagerCreateEntityResponse
	isSet bool
}

func (v NullableEntitymanagerCreateEntityResponse) Get() *EntitymanagerCreateEntityResponse {
	return v.value
}

func (v *NullableEntitymanagerCreateEntityResponse) Set(val *EntitymanagerCreateEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateEntityResponse(val *EntitymanagerCreateEntityResponse) *NullableEntitymanagerCreateEntityResponse {
	return &NullableEntitymanagerCreateEntityResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
