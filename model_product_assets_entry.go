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

// checks if the ProductAssetsEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAssetsEntry{}

// ProductAssetsEntry struct for ProductAssetsEntry
type ProductAssetsEntry struct {
	Id *string `json:"id,omitempty"`
	AssetGrn *string `json:"assetGrn,omitempty"`
	LocalizedAssetGrn *ProductLocalizedAsset `json:"localizedAssetGrn,omitempty"`
	Locales []string `json:"locales,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Metadata []ProductAssetsEntryMetadata `json:"metadata,omitempty"`
}

// NewProductAssetsEntry instantiates a new ProductAssetsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAssetsEntry() *ProductAssetsEntry {
	this := ProductAssetsEntry{}
	return &this
}

// NewProductAssetsEntryWithDefaults instantiates a new ProductAssetsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAssetsEntryWithDefaults() *ProductAssetsEntry {
	this := ProductAssetsEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductAssetsEntry) SetId(v string) {
	o.Id = &v
}

// GetAssetGrn returns the AssetGrn field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetAssetGrn() string {
	if o == nil || IsNil(o.AssetGrn) {
		var ret string
		return ret
	}
	return *o.AssetGrn
}

// GetAssetGrnOk returns a tuple with the AssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetAssetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AssetGrn) {
		return nil, false
	}
	return o.AssetGrn, true
}

// HasAssetGrn returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasAssetGrn() bool {
	if o != nil && !IsNil(o.AssetGrn) {
		return true
	}

	return false
}

// SetAssetGrn gets a reference to the given string and assigns it to the AssetGrn field.
func (o *ProductAssetsEntry) SetAssetGrn(v string) {
	o.AssetGrn = &v
}

// GetLocalizedAssetGrn returns the LocalizedAssetGrn field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetLocalizedAssetGrn() ProductLocalizedAsset {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		var ret ProductLocalizedAsset
		return ret
	}
	return *o.LocalizedAssetGrn
}

// GetLocalizedAssetGrnOk returns a tuple with the LocalizedAssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetLocalizedAssetGrnOk() (*ProductLocalizedAsset, bool) {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		return nil, false
	}
	return o.LocalizedAssetGrn, true
}

// HasLocalizedAssetGrn returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasLocalizedAssetGrn() bool {
	if o != nil && !IsNil(o.LocalizedAssetGrn) {
		return true
	}

	return false
}

// SetLocalizedAssetGrn gets a reference to the given ProductLocalizedAsset and assigns it to the LocalizedAssetGrn field.
func (o *ProductAssetsEntry) SetLocalizedAssetGrn(v ProductLocalizedAsset) {
	o.LocalizedAssetGrn = &v
}

// GetLocales returns the Locales field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetLocales() []string {
	if o == nil || IsNil(o.Locales) {
		var ret []string
		return ret
	}
	return o.Locales
}

// GetLocalesOk returns a tuple with the Locales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.Locales) {
		return nil, false
	}
	return o.Locales, true
}

// HasLocales returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasLocales() bool {
	if o != nil && !IsNil(o.Locales) {
		return true
	}

	return false
}

// SetLocales gets a reference to the given []string and assigns it to the Locales field.
func (o *ProductAssetsEntry) SetLocales(v []string) {
	o.Locales = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ProductAssetsEntry) SetPosition(v int64) {
	o.Position = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProductAssetsEntry) GetMetadata() []ProductAssetsEntryMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret []ProductAssetsEntryMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetsEntry) GetMetadataOk() ([]ProductAssetsEntryMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProductAssetsEntry) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ProductAssetsEntryMetadata and assigns it to the Metadata field.
func (o *ProductAssetsEntry) SetMetadata(v []ProductAssetsEntryMetadata) {
	o.Metadata = v
}

func (o ProductAssetsEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAssetsEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AssetGrn) {
		toSerialize["assetGrn"] = o.AssetGrn
	}
	if !IsNil(o.LocalizedAssetGrn) {
		toSerialize["localizedAssetGrn"] = o.LocalizedAssetGrn
	}
	if !IsNil(o.Locales) {
		toSerialize["locales"] = o.Locales
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableProductAssetsEntry struct {
	value *ProductAssetsEntry
	isSet bool
}

func (v NullableProductAssetsEntry) Get() *ProductAssetsEntry {
	return v.value
}

func (v *NullableProductAssetsEntry) Set(val *ProductAssetsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAssetsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAssetsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAssetsEntry(val *ProductAssetsEntry) *NullableProductAssetsEntry {
	return &NullableProductAssetsEntry{value: val, isSet: true}
}

func (v NullableProductAssetsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAssetsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

