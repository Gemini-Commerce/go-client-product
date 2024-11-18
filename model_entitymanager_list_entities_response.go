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

// checks if the EntitymanagerListEntitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerListEntitiesResponse{}

// EntitymanagerListEntitiesResponse struct for EntitymanagerListEntitiesResponse
type EntitymanagerListEntitiesResponse struct {
	Entities []EntitymanagerEntity `json:"entities,omitempty"`
	NextPage *int32                `json:"nextPage,omitempty"`
}

// NewEntitymanagerListEntitiesResponse instantiates a new EntitymanagerListEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListEntitiesResponse() *EntitymanagerListEntitiesResponse {
	this := EntitymanagerListEntitiesResponse{}
	return &this
}

// NewEntitymanagerListEntitiesResponseWithDefaults instantiates a new EntitymanagerListEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListEntitiesResponseWithDefaults() *EntitymanagerListEntitiesResponse {
	this := EntitymanagerListEntitiesResponse{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesResponse) GetEntities() []EntitymanagerEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []EntitymanagerEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesResponse) GetEntitiesOk() ([]EntitymanagerEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesResponse) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntitymanagerEntity and assigns it to the Entities field.
func (o *EntitymanagerListEntitiesResponse) SetEntities(v []EntitymanagerEntity) {
	o.Entities = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesResponse) GetNextPage() int32 {
	if o == nil || IsNil(o.NextPage) {
		var ret int32
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesResponse) GetNextPageOk() (*int32, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given int32 and assigns it to the NextPage field.
func (o *EntitymanagerListEntitiesResponse) SetNextPage(v int32) {
	o.NextPage = &v
}

func (o EntitymanagerListEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerListEntitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	return toSerialize, nil
}

type NullableEntitymanagerListEntitiesResponse struct {
	value *EntitymanagerListEntitiesResponse
	isSet bool
}

func (v NullableEntitymanagerListEntitiesResponse) Get() *EntitymanagerListEntitiesResponse {
	return v.value
}

func (v *NullableEntitymanagerListEntitiesResponse) Set(val *EntitymanagerListEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListEntitiesResponse(val *EntitymanagerListEntitiesResponse) *NullableEntitymanagerListEntitiesResponse {
	return &NullableEntitymanagerListEntitiesResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerListEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
