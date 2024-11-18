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

// checks if the EntitymanagerUpdateAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerUpdateAttributeRequest{}

// EntitymanagerUpdateAttributeRequest struct for EntitymanagerUpdateAttributeRequest
type EntitymanagerUpdateAttributeRequest struct {
	TenantId             *string                                     `json:"tenantId,omitempty"`
	Code                 *string                                     `json:"code,omitempty"`
	EntityData           *EntitymanagerEntityIdentifier              `json:"entityData,omitempty"`
	EntityId             *string                                     `json:"entityId,omitempty"`
	Payload              *EntitymanagerUpdateAttributeRequestPayload `json:"payload,omitempty"`
	FieldMask            *string                                     `json:"fieldMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerUpdateAttributeRequest EntitymanagerUpdateAttributeRequest

// NewEntitymanagerUpdateAttributeRequest instantiates a new EntitymanagerUpdateAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerUpdateAttributeRequest() *EntitymanagerUpdateAttributeRequest {
	this := EntitymanagerUpdateAttributeRequest{}
	return &this
}

// NewEntitymanagerUpdateAttributeRequestWithDefaults instantiates a new EntitymanagerUpdateAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerUpdateAttributeRequestWithDefaults() *EntitymanagerUpdateAttributeRequest {
	this := EntitymanagerUpdateAttributeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerUpdateAttributeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerUpdateAttributeRequest) SetCode(v string) {
	o.Code = &v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetEntityData() EntitymanagerEntityIdentifier {
	if o == nil || IsNil(o.EntityData) {
		var ret EntitymanagerEntityIdentifier
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetEntityDataOk() (*EntitymanagerEntityIdentifier, bool) {
	if o == nil || IsNil(o.EntityData) {
		return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasEntityData() bool {
	if o != nil && !IsNil(o.EntityData) {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given EntitymanagerEntityIdentifier and assigns it to the EntityData field.
func (o *EntitymanagerUpdateAttributeRequest) SetEntityData(v EntitymanagerEntityIdentifier) {
	o.EntityData = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntitymanagerUpdateAttributeRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetPayload() EntitymanagerUpdateAttributeRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret EntitymanagerUpdateAttributeRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetPayloadOk() (*EntitymanagerUpdateAttributeRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given EntitymanagerUpdateAttributeRequestPayload and assigns it to the Payload field.
func (o *EntitymanagerUpdateAttributeRequest) SetPayload(v EntitymanagerUpdateAttributeRequestPayload) {
	o.Payload = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequest) GetFieldMask() string {
	if o == nil || IsNil(o.FieldMask) {
		var ret string
		return ret
	}
	return *o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequest) GetFieldMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMask) {
		return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequest) HasFieldMask() bool {
	if o != nil && !IsNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given string and assigns it to the FieldMask field.
func (o *EntitymanagerUpdateAttributeRequest) SetFieldMask(v string) {
	o.FieldMask = &v
}

func (o EntitymanagerUpdateAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerUpdateAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.EntityData) {
		toSerialize["entityData"] = o.EntityData
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerUpdateAttributeRequest) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerUpdateAttributeRequest := _EntitymanagerUpdateAttributeRequest{}

	err = json.Unmarshal(data, &varEntitymanagerUpdateAttributeRequest)

	if err != nil {
		return err
	}

	*o = EntitymanagerUpdateAttributeRequest(varEntitymanagerUpdateAttributeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "entityData")
		delete(additionalProperties, "entityId")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "fieldMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EntitymanagerUpdateAttributeRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EntitymanagerUpdateAttributeRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEntitymanagerUpdateAttributeRequest struct {
	value *EntitymanagerUpdateAttributeRequest
	isSet bool
}

func (v NullableEntitymanagerUpdateAttributeRequest) Get() *EntitymanagerUpdateAttributeRequest {
	return v.value
}

func (v *NullableEntitymanagerUpdateAttributeRequest) Set(val *EntitymanagerUpdateAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerUpdateAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerUpdateAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerUpdateAttributeRequest(val *EntitymanagerUpdateAttributeRequest) *NullableEntitymanagerUpdateAttributeRequest {
	return &NullableEntitymanagerUpdateAttributeRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerUpdateAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerUpdateAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
