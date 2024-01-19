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

// checks if the ProductUpdateAssetEntryPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateAssetEntryPayload{}

// ProductUpdateAssetEntryPayload struct for ProductUpdateAssetEntryPayload
type ProductUpdateAssetEntryPayload struct {
	AssetGrn *string `json:"assetGrn,omitempty"`
	LocalizedAssetGrn *ProductLocalizedAsset `json:"localizedAssetGrn,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Metadata []ProductAssetsEntryMetadata `json:"metadata,omitempty"`
}

// NewProductUpdateAssetEntryPayload instantiates a new ProductUpdateAssetEntryPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateAssetEntryPayload() *ProductUpdateAssetEntryPayload {
	this := ProductUpdateAssetEntryPayload{}
	return &this
}

// NewProductUpdateAssetEntryPayloadWithDefaults instantiates a new ProductUpdateAssetEntryPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateAssetEntryPayloadWithDefaults() *ProductUpdateAssetEntryPayload {
	this := ProductUpdateAssetEntryPayload{}
	return &this
}

// GetAssetGrn returns the AssetGrn field value if set, zero value otherwise.
func (o *ProductUpdateAssetEntryPayload) GetAssetGrn() string {
	if o == nil || IsNil(o.AssetGrn) {
		var ret string
		return ret
	}
	return *o.AssetGrn
}

// GetAssetGrnOk returns a tuple with the AssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateAssetEntryPayload) GetAssetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AssetGrn) {
		return nil, false
	}
	return o.AssetGrn, true
}

// HasAssetGrn returns a boolean if a field has been set.
func (o *ProductUpdateAssetEntryPayload) HasAssetGrn() bool {
	if o != nil && !IsNil(o.AssetGrn) {
		return true
	}

	return false
}

// SetAssetGrn gets a reference to the given string and assigns it to the AssetGrn field.
func (o *ProductUpdateAssetEntryPayload) SetAssetGrn(v string) {
	o.AssetGrn = &v
}

// GetLocalizedAssetGrn returns the LocalizedAssetGrn field value if set, zero value otherwise.
func (o *ProductUpdateAssetEntryPayload) GetLocalizedAssetGrn() ProductLocalizedAsset {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		var ret ProductLocalizedAsset
		return ret
	}
	return *o.LocalizedAssetGrn
}

// GetLocalizedAssetGrnOk returns a tuple with the LocalizedAssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateAssetEntryPayload) GetLocalizedAssetGrnOk() (*ProductLocalizedAsset, bool) {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		return nil, false
	}
	return o.LocalizedAssetGrn, true
}

// HasLocalizedAssetGrn returns a boolean if a field has been set.
func (o *ProductUpdateAssetEntryPayload) HasLocalizedAssetGrn() bool {
	if o != nil && !IsNil(o.LocalizedAssetGrn) {
		return true
	}

	return false
}

// SetLocalizedAssetGrn gets a reference to the given ProductLocalizedAsset and assigns it to the LocalizedAssetGrn field.
func (o *ProductUpdateAssetEntryPayload) SetLocalizedAssetGrn(v ProductLocalizedAsset) {
	o.LocalizedAssetGrn = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductUpdateAssetEntryPayload) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateAssetEntryPayload) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductUpdateAssetEntryPayload) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ProductUpdateAssetEntryPayload) SetPosition(v int64) {
	o.Position = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProductUpdateAssetEntryPayload) GetMetadata() []ProductAssetsEntryMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret []ProductAssetsEntryMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateAssetEntryPayload) GetMetadataOk() ([]ProductAssetsEntryMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProductUpdateAssetEntryPayload) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ProductAssetsEntryMetadata and assigns it to the Metadata field.
func (o *ProductUpdateAssetEntryPayload) SetMetadata(v []ProductAssetsEntryMetadata) {
	o.Metadata = v
}

func (o ProductUpdateAssetEntryPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateAssetEntryPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetGrn) {
		toSerialize["assetGrn"] = o.AssetGrn
	}
	if !IsNil(o.LocalizedAssetGrn) {
		toSerialize["localizedAssetGrn"] = o.LocalizedAssetGrn
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableProductUpdateAssetEntryPayload struct {
	value *ProductUpdateAssetEntryPayload
	isSet bool
}

func (v NullableProductUpdateAssetEntryPayload) Get() *ProductUpdateAssetEntryPayload {
	return v.value
}

func (v *NullableProductUpdateAssetEntryPayload) Set(val *ProductUpdateAssetEntryPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateAssetEntryPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateAssetEntryPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateAssetEntryPayload(val *ProductUpdateAssetEntryPayload) *NullableProductUpdateAssetEntryPayload {
	return &NullableProductUpdateAssetEntryPayload{value: val, isSet: true}
}

func (v NullableProductUpdateAssetEntryPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateAssetEntryPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

