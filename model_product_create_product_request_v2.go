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

// checks if the ProductCreateProductRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCreateProductRequestV2{}

// ProductCreateProductRequestV2 struct for ProductCreateProductRequestV2
type ProductCreateProductRequestV2 struct {
	TenantId               *string                           `json:"tenantId,omitempty"`
	EntityType             *string                           `json:"entityType,omitempty"`
	EntityCode             *string                           `json:"entityCode,omitempty"`
	Code                   *string                           `json:"code,omitempty"`
	IsConfigurable         *bool                             `json:"isConfigurable,omitempty"`
	VariantAttributes      []string                          `json:"variantAttributes,omitempty"`
	IsVirtual              *bool                             `json:"isVirtual,omitempty"`
	IsGiftcard             *bool                             `json:"isGiftcard,omitempty"`
	HasConfigurator        *bool                             `json:"hasConfigurator,omitempty"`
	UrlKey                 *ProductLocalizedText             `json:"urlKey,omitempty"`
	MaxSaleableQuantity    *int64                            `json:"maxSaleableQuantity,omitempty"`
	Attributes             *map[string]ProtobufAny           `json:"attributes,omitempty"`
	Variants               *map[string]ProductProductVariant `json:"variants,omitempty"`
	MediaVariantAttributes []string                          `json:"mediaVariantAttributes,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _ProductCreateProductRequestV2 ProductCreateProductRequestV2

// NewProductCreateProductRequestV2 instantiates a new ProductCreateProductRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateProductRequestV2() *ProductCreateProductRequestV2 {
	this := ProductCreateProductRequestV2{}
	return &this
}

// NewProductCreateProductRequestV2WithDefaults instantiates a new ProductCreateProductRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateProductRequestV2WithDefaults() *ProductCreateProductRequestV2 {
	this := ProductCreateProductRequestV2{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductCreateProductRequestV2) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ProductCreateProductRequestV2) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetEntityCode() string {
	if o == nil || IsNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetEntityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityCode) {
		return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasEntityCode() bool {
	if o != nil && !IsNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ProductCreateProductRequestV2) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductCreateProductRequestV2) SetCode(v string) {
	o.Code = &v
}

// GetIsConfigurable returns the IsConfigurable field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetIsConfigurable() bool {
	if o == nil || IsNil(o.IsConfigurable) {
		var ret bool
		return ret
	}
	return *o.IsConfigurable
}

// GetIsConfigurableOk returns a tuple with the IsConfigurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetIsConfigurableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfigurable) {
		return nil, false
	}
	return o.IsConfigurable, true
}

// HasIsConfigurable returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasIsConfigurable() bool {
	if o != nil && !IsNil(o.IsConfigurable) {
		return true
	}

	return false
}

// SetIsConfigurable gets a reference to the given bool and assigns it to the IsConfigurable field.
func (o *ProductCreateProductRequestV2) SetIsConfigurable(v bool) {
	o.IsConfigurable = &v
}

// GetVariantAttributes returns the VariantAttributes field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetVariantAttributes() []string {
	if o == nil || IsNil(o.VariantAttributes) {
		var ret []string
		return ret
	}
	return o.VariantAttributes
}

// GetVariantAttributesOk returns a tuple with the VariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetVariantAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.VariantAttributes) {
		return nil, false
	}
	return o.VariantAttributes, true
}

// HasVariantAttributes returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasVariantAttributes() bool {
	if o != nil && !IsNil(o.VariantAttributes) {
		return true
	}

	return false
}

// SetVariantAttributes gets a reference to the given []string and assigns it to the VariantAttributes field.
func (o *ProductCreateProductRequestV2) SetVariantAttributes(v []string) {
	o.VariantAttributes = v
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetIsVirtual() bool {
	if o == nil || IsNil(o.IsVirtual) {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetIsVirtualOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVirtual) {
		return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasIsVirtual() bool {
	if o != nil && !IsNil(o.IsVirtual) {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *ProductCreateProductRequestV2) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

// GetIsGiftcard returns the IsGiftcard field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetIsGiftcard() bool {
	if o == nil || IsNil(o.IsGiftcard) {
		var ret bool
		return ret
	}
	return *o.IsGiftcard
}

// GetIsGiftcardOk returns a tuple with the IsGiftcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetIsGiftcardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGiftcard) {
		return nil, false
	}
	return o.IsGiftcard, true
}

// HasIsGiftcard returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasIsGiftcard() bool {
	if o != nil && !IsNil(o.IsGiftcard) {
		return true
	}

	return false
}

// SetIsGiftcard gets a reference to the given bool and assigns it to the IsGiftcard field.
func (o *ProductCreateProductRequestV2) SetIsGiftcard(v bool) {
	o.IsGiftcard = &v
}

// GetHasConfigurator returns the HasConfigurator field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetHasConfigurator() bool {
	if o == nil || IsNil(o.HasConfigurator) {
		var ret bool
		return ret
	}
	return *o.HasConfigurator
}

// GetHasConfiguratorOk returns a tuple with the HasConfigurator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetHasConfiguratorOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfigurator) {
		return nil, false
	}
	return o.HasConfigurator, true
}

// HasHasConfigurator returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasHasConfigurator() bool {
	if o != nil && !IsNil(o.HasConfigurator) {
		return true
	}

	return false
}

// SetHasConfigurator gets a reference to the given bool and assigns it to the HasConfigurator field.
func (o *ProductCreateProductRequestV2) SetHasConfigurator(v bool) {
	o.HasConfigurator = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetUrlKey() ProductLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetUrlKeyOk() (*ProductLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductLocalizedText and assigns it to the UrlKey field.
func (o *ProductCreateProductRequestV2) SetUrlKey(v ProductLocalizedText) {
	o.UrlKey = &v
}

// GetMaxSaleableQuantity returns the MaxSaleableQuantity field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetMaxSaleableQuantity() int64 {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		var ret int64
		return ret
	}
	return *o.MaxSaleableQuantity
}

// GetMaxSaleableQuantityOk returns a tuple with the MaxSaleableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetMaxSaleableQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		return nil, false
	}
	return o.MaxSaleableQuantity, true
}

// HasMaxSaleableQuantity returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasMaxSaleableQuantity() bool {
	if o != nil && !IsNil(o.MaxSaleableQuantity) {
		return true
	}

	return false
}

// SetMaxSaleableQuantity gets a reference to the given int64 and assigns it to the MaxSaleableQuantity field.
func (o *ProductCreateProductRequestV2) SetMaxSaleableQuantity(v int64) {
	o.MaxSaleableQuantity = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductCreateProductRequestV2) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetVariants() map[string]ProductProductVariant {
	if o == nil || IsNil(o.Variants) {
		var ret map[string]ProductProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetVariantsOk() (*map[string]ProductProductVariant, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]ProductProductVariant and assigns it to the Variants field.
func (o *ProductCreateProductRequestV2) SetVariants(v map[string]ProductProductVariant) {
	o.Variants = &v
}

// GetMediaVariantAttributes returns the MediaVariantAttributes field value if set, zero value otherwise.
func (o *ProductCreateProductRequestV2) GetMediaVariantAttributes() []string {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		var ret []string
		return ret
	}
	return o.MediaVariantAttributes
}

// GetMediaVariantAttributesOk returns a tuple with the MediaVariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequestV2) GetMediaVariantAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		return nil, false
	}
	return o.MediaVariantAttributes, true
}

// HasMediaVariantAttributes returns a boolean if a field has been set.
func (o *ProductCreateProductRequestV2) HasMediaVariantAttributes() bool {
	if o != nil && !IsNil(o.MediaVariantAttributes) {
		return true
	}

	return false
}

// SetMediaVariantAttributes gets a reference to the given []string and assigns it to the MediaVariantAttributes field.
func (o *ProductCreateProductRequestV2) SetMediaVariantAttributes(v []string) {
	o.MediaVariantAttributes = v
}

func (o ProductCreateProductRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductCreateProductRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.IsConfigurable) {
		toSerialize["isConfigurable"] = o.IsConfigurable
	}
	if !IsNil(o.VariantAttributes) {
		toSerialize["variantAttributes"] = o.VariantAttributes
	}
	if !IsNil(o.IsVirtual) {
		toSerialize["isVirtual"] = o.IsVirtual
	}
	if !IsNil(o.IsGiftcard) {
		toSerialize["isGiftcard"] = o.IsGiftcard
	}
	if !IsNil(o.HasConfigurator) {
		toSerialize["hasConfigurator"] = o.HasConfigurator
	}
	if !IsNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !IsNil(o.MaxSaleableQuantity) {
		toSerialize["maxSaleableQuantity"] = o.MaxSaleableQuantity
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	if !IsNil(o.MediaVariantAttributes) {
		toSerialize["mediaVariantAttributes"] = o.MediaVariantAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductCreateProductRequestV2) UnmarshalJSON(data []byte) (err error) {
	varProductCreateProductRequestV2 := _ProductCreateProductRequestV2{}

	err = json.Unmarshal(data, &varProductCreateProductRequestV2)

	if err != nil {
		return err
	}

	*o = ProductCreateProductRequestV2(varProductCreateProductRequestV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "entityCode")
		delete(additionalProperties, "code")
		delete(additionalProperties, "isConfigurable")
		delete(additionalProperties, "variantAttributes")
		delete(additionalProperties, "isVirtual")
		delete(additionalProperties, "isGiftcard")
		delete(additionalProperties, "hasConfigurator")
		delete(additionalProperties, "urlKey")
		delete(additionalProperties, "maxSaleableQuantity")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "variants")
		delete(additionalProperties, "mediaVariantAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductCreateProductRequestV2 struct {
	value *ProductCreateProductRequestV2
	isSet bool
}

func (v NullableProductCreateProductRequestV2) Get() *ProductCreateProductRequestV2 {
	return v.value
}

func (v *NullableProductCreateProductRequestV2) Set(val *ProductCreateProductRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateProductRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateProductRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateProductRequestV2(val *ProductCreateProductRequestV2) *NullableProductCreateProductRequestV2 {
	return &NullableProductCreateProductRequestV2{value: val, isSet: true}
}

func (v NullableProductCreateProductRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateProductRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
