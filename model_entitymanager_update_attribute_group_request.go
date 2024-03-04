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

// checks if the EntitymanagerUpdateAttributeGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerUpdateAttributeGroupRequest{}

// EntitymanagerUpdateAttributeGroupRequest struct for EntitymanagerUpdateAttributeGroupRequest
type EntitymanagerUpdateAttributeGroupRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
	Payload *EntitymanagerUpdateAttributeGroupRequestPayload `json:"payload,omitempty"`
	FieldMask *string `json:"fieldMask,omitempty"`
}

// NewEntitymanagerUpdateAttributeGroupRequest instantiates a new EntitymanagerUpdateAttributeGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerUpdateAttributeGroupRequest() *EntitymanagerUpdateAttributeGroupRequest {
	this := EntitymanagerUpdateAttributeGroupRequest{}
	return &this
}

// NewEntitymanagerUpdateAttributeGroupRequestWithDefaults instantiates a new EntitymanagerUpdateAttributeGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerUpdateAttributeGroupRequestWithDefaults() *EntitymanagerUpdateAttributeGroupRequest {
	this := EntitymanagerUpdateAttributeGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerUpdateAttributeGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerUpdateAttributeGroupRequest) SetCode(v string) {
	o.Code = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetPayload() EntitymanagerUpdateAttributeGroupRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret EntitymanagerUpdateAttributeGroupRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetPayloadOk() (*EntitymanagerUpdateAttributeGroupRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given EntitymanagerUpdateAttributeGroupRequestPayload and assigns it to the Payload field.
func (o *EntitymanagerUpdateAttributeGroupRequest) SetPayload(v EntitymanagerUpdateAttributeGroupRequestPayload) {
	o.Payload = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetFieldMask() string {
	if o == nil || IsNil(o.FieldMask) {
		var ret string
		return ret
	}
	return *o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) GetFieldMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMask) {
		return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeGroupRequest) HasFieldMask() bool {
	if o != nil && !IsNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given string and assigns it to the FieldMask field.
func (o *EntitymanagerUpdateAttributeGroupRequest) SetFieldMask(v string) {
	o.FieldMask = &v
}

func (o EntitymanagerUpdateAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerUpdateAttributeGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}
	return toSerialize, nil
}

type NullableEntitymanagerUpdateAttributeGroupRequest struct {
	value *EntitymanagerUpdateAttributeGroupRequest
	isSet bool
}

func (v NullableEntitymanagerUpdateAttributeGroupRequest) Get() *EntitymanagerUpdateAttributeGroupRequest {
	return v.value
}

func (v *NullableEntitymanagerUpdateAttributeGroupRequest) Set(val *EntitymanagerUpdateAttributeGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerUpdateAttributeGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerUpdateAttributeGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerUpdateAttributeGroupRequest(val *EntitymanagerUpdateAttributeGroupRequest) *NullableEntitymanagerUpdateAttributeGroupRequest {
	return &NullableEntitymanagerUpdateAttributeGroupRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerUpdateAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerUpdateAttributeGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


