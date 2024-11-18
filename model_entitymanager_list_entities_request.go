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

// checks if the EntitymanagerListEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerListEntitiesRequest{}

// EntitymanagerListEntitiesRequest struct for EntitymanagerListEntitiesRequest
type EntitymanagerListEntitiesRequest struct {
	TenantId             *string `json:"tenantId,omitempty"`
	PageSize             *int32  `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerListEntitiesRequest EntitymanagerListEntitiesRequest

// NewEntitymanagerListEntitiesRequest instantiates a new EntitymanagerListEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListEntitiesRequest() *EntitymanagerListEntitiesRequest {
	this := EntitymanagerListEntitiesRequest{}
	return &this
}

// NewEntitymanagerListEntitiesRequestWithDefaults instantiates a new EntitymanagerListEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListEntitiesRequestWithDefaults() *EntitymanagerListEntitiesRequest {
	this := EntitymanagerListEntitiesRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerListEntitiesRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *EntitymanagerListEntitiesRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o EntitymanagerListEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerListEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerListEntitiesRequest) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerListEntitiesRequest := _EntitymanagerListEntitiesRequest{}

	err = json.Unmarshal(data, &varEntitymanagerListEntitiesRequest)

	if err != nil {
		return err
	}

	*o = EntitymanagerListEntitiesRequest(varEntitymanagerListEntitiesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EntitymanagerListEntitiesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EntitymanagerListEntitiesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEntitymanagerListEntitiesRequest struct {
	value *EntitymanagerListEntitiesRequest
	isSet bool
}

func (v NullableEntitymanagerListEntitiesRequest) Get() *EntitymanagerListEntitiesRequest {
	return v.value
}

func (v *NullableEntitymanagerListEntitiesRequest) Set(val *EntitymanagerListEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListEntitiesRequest(val *EntitymanagerListEntitiesRequest) *NullableEntitymanagerListEntitiesRequest {
	return &NullableEntitymanagerListEntitiesRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerListEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
