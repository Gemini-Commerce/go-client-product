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

// checks if the EntitymanagerCreateAttributeOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerCreateAttributeOptionsResponse{}

// EntitymanagerCreateAttributeOptionsResponse struct for EntitymanagerCreateAttributeOptionsResponse
type EntitymanagerCreateAttributeOptionsResponse struct {
	Options              []EntitymanagerAttributeOption       `json:"options,omitempty"`
	Errors               []EntitymanagerAttributeOptionErrors `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerCreateAttributeOptionsResponse EntitymanagerCreateAttributeOptionsResponse

// NewEntitymanagerCreateAttributeOptionsResponse instantiates a new EntitymanagerCreateAttributeOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateAttributeOptionsResponse() *EntitymanagerCreateAttributeOptionsResponse {
	this := EntitymanagerCreateAttributeOptionsResponse{}
	return &this
}

// NewEntitymanagerCreateAttributeOptionsResponseWithDefaults instantiates a new EntitymanagerCreateAttributeOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateAttributeOptionsResponseWithDefaults() *EntitymanagerCreateAttributeOptionsResponse {
	this := EntitymanagerCreateAttributeOptionsResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeOptionsResponse) GetOptions() []EntitymanagerAttributeOption {
	if o == nil || IsNil(o.Options) {
		var ret []EntitymanagerAttributeOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeOptionsResponse) GetOptionsOk() ([]EntitymanagerAttributeOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeOptionsResponse) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerAttributeOption and assigns it to the Options field.
func (o *EntitymanagerCreateAttributeOptionsResponse) SetOptions(v []EntitymanagerAttributeOption) {
	o.Options = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeOptionsResponse) GetErrors() []EntitymanagerAttributeOptionErrors {
	if o == nil || IsNil(o.Errors) {
		var ret []EntitymanagerAttributeOptionErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeOptionsResponse) GetErrorsOk() ([]EntitymanagerAttributeOptionErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeOptionsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []EntitymanagerAttributeOptionErrors and assigns it to the Errors field.
func (o *EntitymanagerCreateAttributeOptionsResponse) SetErrors(v []EntitymanagerAttributeOptionErrors) {
	o.Errors = v
}

func (o EntitymanagerCreateAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerCreateAttributeOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerCreateAttributeOptionsResponse) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerCreateAttributeOptionsResponse := _EntitymanagerCreateAttributeOptionsResponse{}

	err = json.Unmarshal(data, &varEntitymanagerCreateAttributeOptionsResponse)

	if err != nil {
		return err
	}

	*o = EntitymanagerCreateAttributeOptionsResponse(varEntitymanagerCreateAttributeOptionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "options")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitymanagerCreateAttributeOptionsResponse struct {
	value *EntitymanagerCreateAttributeOptionsResponse
	isSet bool
}

func (v NullableEntitymanagerCreateAttributeOptionsResponse) Get() *EntitymanagerCreateAttributeOptionsResponse {
	return v.value
}

func (v *NullableEntitymanagerCreateAttributeOptionsResponse) Set(val *EntitymanagerCreateAttributeOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateAttributeOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateAttributeOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateAttributeOptionsResponse(val *EntitymanagerCreateAttributeOptionsResponse) *NullableEntitymanagerCreateAttributeOptionsResponse {
	return &NullableEntitymanagerCreateAttributeOptionsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateAttributeOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
