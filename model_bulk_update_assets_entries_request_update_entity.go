/*
Product Service

API for managing products

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
	Id *string `json:"id,omitempty"`
	Payload *ProductUpdateAssetEntryPayload `json:"payload,omitempty"`
	PayloadMask []string `json:"payloadMask,omitempty"`
}

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
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayloadMask() []string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret []string
		return ret
	}
	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) GetPayloadMaskOk() ([]string, bool) {
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

// SetPayloadMask gets a reference to the given []string and assigns it to the PayloadMask field.
func (o *BulkUpdateAssetsEntriesRequestUpdateEntity) SetPayloadMask(v []string) {
	o.PayloadMask = v
}

func (o BulkUpdateAssetsEntriesRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
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
	return toSerialize, nil
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


