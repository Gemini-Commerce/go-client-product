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

// checks if the ProductAssetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAssetData{}

// ProductAssetData struct for ProductAssetData
type ProductAssetData struct {
	AssetGrn *string `json:"assetGrn,omitempty"`
	LocalizedAssetGrn *ProductLocalizedAsset `json:"localizedAssetGrn,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Metadata []ProductAssetsEntryMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAssetData ProductAssetData

// NewProductAssetData instantiates a new ProductAssetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAssetData() *ProductAssetData {
	this := ProductAssetData{}
	return &this
}

// NewProductAssetDataWithDefaults instantiates a new ProductAssetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAssetDataWithDefaults() *ProductAssetData {
	this := ProductAssetData{}
	return &this
}

// GetAssetGrn returns the AssetGrn field value if set, zero value otherwise.
func (o *ProductAssetData) GetAssetGrn() string {
	if o == nil || IsNil(o.AssetGrn) {
		var ret string
		return ret
	}
	return *o.AssetGrn
}

// GetAssetGrnOk returns a tuple with the AssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetData) GetAssetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AssetGrn) {
		return nil, false
	}
	return o.AssetGrn, true
}

// HasAssetGrn returns a boolean if a field has been set.
func (o *ProductAssetData) HasAssetGrn() bool {
	if o != nil && !IsNil(o.AssetGrn) {
		return true
	}

	return false
}

// SetAssetGrn gets a reference to the given string and assigns it to the AssetGrn field.
func (o *ProductAssetData) SetAssetGrn(v string) {
	o.AssetGrn = &v
}

// GetLocalizedAssetGrn returns the LocalizedAssetGrn field value if set, zero value otherwise.
func (o *ProductAssetData) GetLocalizedAssetGrn() ProductLocalizedAsset {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		var ret ProductLocalizedAsset
		return ret
	}
	return *o.LocalizedAssetGrn
}

// GetLocalizedAssetGrnOk returns a tuple with the LocalizedAssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetData) GetLocalizedAssetGrnOk() (*ProductLocalizedAsset, bool) {
	if o == nil || IsNil(o.LocalizedAssetGrn) {
		return nil, false
	}
	return o.LocalizedAssetGrn, true
}

// HasLocalizedAssetGrn returns a boolean if a field has been set.
func (o *ProductAssetData) HasLocalizedAssetGrn() bool {
	if o != nil && !IsNil(o.LocalizedAssetGrn) {
		return true
	}

	return false
}

// SetLocalizedAssetGrn gets a reference to the given ProductLocalizedAsset and assigns it to the LocalizedAssetGrn field.
func (o *ProductAssetData) SetLocalizedAssetGrn(v ProductLocalizedAsset) {
	o.LocalizedAssetGrn = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductAssetData) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetData) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductAssetData) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ProductAssetData) SetPosition(v int64) {
	o.Position = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProductAssetData) GetMetadata() []ProductAssetsEntryMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret []ProductAssetsEntryMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAssetData) GetMetadataOk() ([]ProductAssetsEntryMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProductAssetData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ProductAssetsEntryMetadata and assigns it to the Metadata field.
func (o *ProductAssetData) SetMetadata(v []ProductAssetsEntryMetadata) {
	o.Metadata = v
}

func (o ProductAssetData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAssetData) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductAssetData) UnmarshalJSON(data []byte) (err error) {
	varProductAssetData := _ProductAssetData{}

	err = json.Unmarshal(data, &varProductAssetData)

	if err != nil {
		return err
	}

	*o = ProductAssetData(varProductAssetData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assetGrn")
		delete(additionalProperties, "localizedAssetGrn")
		delete(additionalProperties, "position")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductAssetData) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductAssetData) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductAssetData struct {
	value *ProductAssetData
	isSet bool
}

func (v NullableProductAssetData) Get() *ProductAssetData {
	return v.value
}

func (v *NullableProductAssetData) Set(val *ProductAssetData) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAssetData) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAssetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAssetData(val *ProductAssetData) *NullableProductAssetData {
	return &NullableProductAssetData{value: val, isSet: true}
}

func (v NullableProductAssetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAssetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


