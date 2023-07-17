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

// ProductProductEntity struct for ProductProductEntity
type ProductProductEntity struct {
	TenantId *string `json:"tenantId,omitempty"`
	Grn *string `json:"grn,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	EntityCode *string `json:"entityCode,omitempty"`
	Id *string `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	VariantAttributes []string `json:"variantAttributes,omitempty"`
	IsConfigurable *bool `json:"isConfigurable,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	IsGiftcard *bool `json:"isGiftcard,omitempty"`
	UrlKey *ProductLocalizedText `json:"urlKey,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	Variants *map[string]ProductProductVariant `json:"variants,omitempty"`
	MediaGallery *ProductMediaGallery `json:"mediaGallery,omitempty"`
	MaxSaleableQuantity *int64 `json:"maxSaleableQuantity,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewProductProductEntity instantiates a new ProductProductEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductProductEntity() *ProductProductEntity {
	this := ProductProductEntity{}
	return &this
}

// NewProductProductEntityWithDefaults instantiates a new ProductProductEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductProductEntityWithDefaults() *ProductProductEntity {
	this := ProductProductEntity{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductProductEntity) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductProductEntity) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductProductEntity) SetTenantId(v string) {
	o.TenantId = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductProductEntity) GetGrn() string {
	if o == nil || isNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetGrnOk() (*string, bool) {
	if o == nil || isNil(o.Grn) {
    return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductProductEntity) HasGrn() bool {
	if o != nil && !isNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductProductEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ProductProductEntity) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
    return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ProductProductEntity) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ProductProductEntity) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ProductProductEntity) GetEntityCode() string {
	if o == nil || isNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetEntityCodeOk() (*string, bool) {
	if o == nil || isNil(o.EntityCode) {
    return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ProductProductEntity) HasEntityCode() bool {
	if o != nil && !isNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ProductProductEntity) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductProductEntity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductProductEntity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductProductEntity) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductProductEntity) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductProductEntity) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductProductEntity) SetCode(v string) {
	o.Code = &v
}

// GetVariantAttributes returns the VariantAttributes field value if set, zero value otherwise.
func (o *ProductProductEntity) GetVariantAttributes() []string {
	if o == nil || isNil(o.VariantAttributes) {
		var ret []string
		return ret
	}
	return o.VariantAttributes
}

// GetVariantAttributesOk returns a tuple with the VariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetVariantAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.VariantAttributes) {
    return nil, false
	}
	return o.VariantAttributes, true
}

// HasVariantAttributes returns a boolean if a field has been set.
func (o *ProductProductEntity) HasVariantAttributes() bool {
	if o != nil && !isNil(o.VariantAttributes) {
		return true
	}

	return false
}

// SetVariantAttributes gets a reference to the given []string and assigns it to the VariantAttributes field.
func (o *ProductProductEntity) SetVariantAttributes(v []string) {
	o.VariantAttributes = v
}

// GetIsConfigurable returns the IsConfigurable field value if set, zero value otherwise.
func (o *ProductProductEntity) GetIsConfigurable() bool {
	if o == nil || isNil(o.IsConfigurable) {
		var ret bool
		return ret
	}
	return *o.IsConfigurable
}

// GetIsConfigurableOk returns a tuple with the IsConfigurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetIsConfigurableOk() (*bool, bool) {
	if o == nil || isNil(o.IsConfigurable) {
    return nil, false
	}
	return o.IsConfigurable, true
}

// HasIsConfigurable returns a boolean if a field has been set.
func (o *ProductProductEntity) HasIsConfigurable() bool {
	if o != nil && !isNil(o.IsConfigurable) {
		return true
	}

	return false
}

// SetIsConfigurable gets a reference to the given bool and assigns it to the IsConfigurable field.
func (o *ProductProductEntity) SetIsConfigurable(v bool) {
	o.IsConfigurable = &v
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *ProductProductEntity) GetIsVirtual() bool {
	if o == nil || isNil(o.IsVirtual) {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetIsVirtualOk() (*bool, bool) {
	if o == nil || isNil(o.IsVirtual) {
    return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *ProductProductEntity) HasIsVirtual() bool {
	if o != nil && !isNil(o.IsVirtual) {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *ProductProductEntity) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

// GetIsGiftcard returns the IsGiftcard field value if set, zero value otherwise.
func (o *ProductProductEntity) GetIsGiftcard() bool {
	if o == nil || isNil(o.IsGiftcard) {
		var ret bool
		return ret
	}
	return *o.IsGiftcard
}

// GetIsGiftcardOk returns a tuple with the IsGiftcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetIsGiftcardOk() (*bool, bool) {
	if o == nil || isNil(o.IsGiftcard) {
    return nil, false
	}
	return o.IsGiftcard, true
}

// HasIsGiftcard returns a boolean if a field has been set.
func (o *ProductProductEntity) HasIsGiftcard() bool {
	if o != nil && !isNil(o.IsGiftcard) {
		return true
	}

	return false
}

// SetIsGiftcard gets a reference to the given bool and assigns it to the IsGiftcard field.
func (o *ProductProductEntity) SetIsGiftcard(v bool) {
	o.IsGiftcard = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductProductEntity) GetUrlKey() ProductLocalizedText {
	if o == nil || isNil(o.UrlKey) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetUrlKeyOk() (*ProductLocalizedText, bool) {
	if o == nil || isNil(o.UrlKey) {
    return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductProductEntity) HasUrlKey() bool {
	if o != nil && !isNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductLocalizedText and assigns it to the UrlKey field.
func (o *ProductProductEntity) SetUrlKey(v ProductLocalizedText) {
	o.UrlKey = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductProductEntity) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductProductEntity) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductProductEntity) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ProductProductEntity) GetVariants() map[string]ProductProductVariant {
	if o == nil || isNil(o.Variants) {
		var ret map[string]ProductProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetVariantsOk() (*map[string]ProductProductVariant, bool) {
	if o == nil || isNil(o.Variants) {
    return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ProductProductEntity) HasVariants() bool {
	if o != nil && !isNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]ProductProductVariant and assigns it to the Variants field.
func (o *ProductProductEntity) SetVariants(v map[string]ProductProductVariant) {
	o.Variants = &v
}

// GetMediaGallery returns the MediaGallery field value if set, zero value otherwise.
func (o *ProductProductEntity) GetMediaGallery() ProductMediaGallery {
	if o == nil || isNil(o.MediaGallery) {
		var ret ProductMediaGallery
		return ret
	}
	return *o.MediaGallery
}

// GetMediaGalleryOk returns a tuple with the MediaGallery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetMediaGalleryOk() (*ProductMediaGallery, bool) {
	if o == nil || isNil(o.MediaGallery) {
    return nil, false
	}
	return o.MediaGallery, true
}

// HasMediaGallery returns a boolean if a field has been set.
func (o *ProductProductEntity) HasMediaGallery() bool {
	if o != nil && !isNil(o.MediaGallery) {
		return true
	}

	return false
}

// SetMediaGallery gets a reference to the given ProductMediaGallery and assigns it to the MediaGallery field.
func (o *ProductProductEntity) SetMediaGallery(v ProductMediaGallery) {
	o.MediaGallery = &v
}

// GetMaxSaleableQuantity returns the MaxSaleableQuantity field value if set, zero value otherwise.
func (o *ProductProductEntity) GetMaxSaleableQuantity() int64 {
	if o == nil || isNil(o.MaxSaleableQuantity) {
		var ret int64
		return ret
	}
	return *o.MaxSaleableQuantity
}

// GetMaxSaleableQuantityOk returns a tuple with the MaxSaleableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetMaxSaleableQuantityOk() (*int64, bool) {
	if o == nil || isNil(o.MaxSaleableQuantity) {
    return nil, false
	}
	return o.MaxSaleableQuantity, true
}

// HasMaxSaleableQuantity returns a boolean if a field has been set.
func (o *ProductProductEntity) HasMaxSaleableQuantity() bool {
	if o != nil && !isNil(o.MaxSaleableQuantity) {
		return true
	}

	return false
}

// SetMaxSaleableQuantity gets a reference to the given int64 and assigns it to the MaxSaleableQuantity field.
func (o *ProductProductEntity) SetMaxSaleableQuantity(v int64) {
	o.MaxSaleableQuantity = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductProductEntity) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductProductEntity) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ProductProductEntity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductProductEntity) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductEntity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductProductEntity) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ProductProductEntity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ProductProductEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.VariantAttributes) {
		toSerialize["variantAttributes"] = o.VariantAttributes
	}
	if !isNil(o.IsConfigurable) {
		toSerialize["isConfigurable"] = o.IsConfigurable
	}
	if !isNil(o.IsVirtual) {
		toSerialize["isVirtual"] = o.IsVirtual
	}
	if !isNil(o.IsGiftcard) {
		toSerialize["isGiftcard"] = o.IsGiftcard
	}
	if !isNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	if !isNil(o.MediaGallery) {
		toSerialize["mediaGallery"] = o.MediaGallery
	}
	if !isNil(o.MaxSaleableQuantity) {
		toSerialize["maxSaleableQuantity"] = o.MaxSaleableQuantity
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProductProductEntity struct {
	value *ProductProductEntity
	isSet bool
}

func (v NullableProductProductEntity) Get() *ProductProductEntity {
	return v.value
}

func (v *NullableProductProductEntity) Set(val *ProductProductEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductProductEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductProductEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductProductEntity(val *ProductProductEntity) *NullableProductProductEntity {
	return &NullableProductProductEntity{value: val, isSet: true}
}

func (v NullableProductProductEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductProductEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


