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

// checks if the BulkUpdateAssetsEntriesRequestUpdateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateAssetsEntriesRequestUpdateEntity{}

// BulkUpdateAssetsEntriesRequestUpdateEntity struct for BulkUpdateAssetsEntriesRequestUpdateEntity
type BulkUpdateAssetsEntriesRequestUpdateEntity struct {
	Id                   *string                         `json:"id,omitempty"`
	Payload              *ProductUpdateAssetEntryPayload `json:"payload,omitempty"`
	PayloadMask          *string                         `json:"payloadMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkUpdateAssetsEntriesRequestUpdateEntity BulkUpdateAssetsEntriesRequestUpdateEntity

// NewBulkUpdateAssetsEntriesRequestUpdateEntity instantiates a new BulkUpdateAssetsEntriesRequestUpdateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateAssetsEntriesRequestUpdateEntity() *BulkUpdateAssetsEntriesRequestUpdateEntity {
	this := BulkUpdateAssetsEntriesRequestUpdateEntity{}
	return &this
}

// NewBulkUpdateAssetsEntriesRequestUpdateEntityWithDefaults instantiates a new BulkUpdateAssetsEntriesRequestUpdateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateAssetsEntriesRequestUpdateEntityWithDefaults() *BulkUpdateAssetsEntriesRequestUpdateEntity {
	this := BulkUpdateAssetsEntriesRequestUpdateEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayload() ProductUpdateAssetEntryPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ProductUpdateAssetEntryPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayloadOk() (*ProductUpdateAssetEntryPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductUpdateAssetEntryPayload and assigns it to the Payload field.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) SetPayload(v ProductUpdateAssetEntryPayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o BulkUpdateAssetsEntriesRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateAssetsEntriesRequestUpdateEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) UnmarshalJSON(data []byte) (err error) {
	varBulkUpdateAssetsEntriesRequestUpdateEntity := _BulkUpdateAssetsEntriesRequestUpdateEntity{}

	err = json.Unmarshal(data, &varBulkUpdateAssetsEntriesRequestUpdateEntity)

	if err != nil {
		return err
	}

	*o = BulkUpdateAssetsEntriesRequestUpdateEntity(varBulkUpdateAssetsEntriesRequestUpdateEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkUpdateAssetsEntriesRequestUpdateEntity struct {
	value *BulkUpdateAssetsEntriesRequestUpdateEntity
	isSet bool
}

func (v NullableBulkUpdateAssetsEntriesRequestUpdateEntity) Get() *BulkUpdateAssetsEntriesRequestUpdateEntity {
	return v.value
}

func (v *NullableBulkUpdateAssetsEntriesRequestUpdateEntity) Set(val *BulkUpdateAssetsEntriesRequestUpdateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateAssetsEntriesRequestUpdateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateAssetsEntriesRequestUpdateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateAssetsEntriesRequestUpdateEntity(val *BulkUpdateAssetsEntriesRequestUpdateEntity) *NullableBulkUpdateAssetsEntriesRequestUpdateEntity {
	return &NullableBulkUpdateAssetsEntriesRequestUpdateEntity{value: val, isSet: true}
}

func (v NullableBulkUpdateAssetsEntriesRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateAssetsEntriesRequestUpdateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
