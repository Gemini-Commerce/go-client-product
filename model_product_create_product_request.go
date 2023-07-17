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

// ProductCreateProductRequest The CreateProductRequest message is used to create a new product within the system. It contains various fields that allow specifying the details and attributes of the product.
type ProductCreateProductRequest struct {
	// Represents the ID of the tenant associated with the product.
	TenantId *string `json:"tenantId,omitempty"`
	// Specifies the type of entity for the product.
	EntityType *string `json:"entityType,omitempty"`
	// Indicates the code of the entity associated with the product.
	EntityCode *string `json:"entityCode,omitempty"`
	// Represents the unique code or identifier for the product.
	Code *string `json:"code,omitempty"`
	// Specifies whether the product has variants or not.
	IsConfigurable *bool `json:"isConfigurable,omitempty"`
	// Contains a list of attributes specific to the product variants.
	VariantAttributes []string `json:"variantAttributes,omitempty"`
	// Indicates whether the product is virtual or not.
	IsVirtual *bool `json:"isVirtual,omitempty"`
	// Specifies whether the product is a gift card or not.
	IsGiftcard *bool `json:"isGiftcard,omitempty"`
	UrlKey *ProductLocalizedText `json:"urlKey,omitempty"`
	// Specifies the maximum quantity that can be sold for the product in each order.
	MaxSaleableQuantity *int64 `json:"maxSaleableQuantity,omitempty"`
	// Contains a map of additional attributes associated with the product, where the key is the attribute name and the value is any type of value.
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	// Represents a map of product variants associated with the product, where the key is the variant ID or code, and the value is a ProductVariant message.
	Variants *map[string]ProductProductVariant `json:"variants,omitempty"`
}

// NewProductCreateProductRequest instantiates a new ProductCreateProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateProductRequest() *ProductCreateProductRequest {
	this := ProductCreateProductRequest{}
	return &this
}

// NewProductCreateProductRequestWithDefaults instantiates a new ProductCreateProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateProductRequestWithDefaults() *ProductCreateProductRequest {
	this := ProductCreateProductRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductCreateProductRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
    return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ProductCreateProductRequest) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetEntityCode() string {
	if o == nil || isNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetEntityCodeOk() (*string, bool) {
	if o == nil || isNil(o.EntityCode) {
    return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasEntityCode() bool {
	if o != nil && !isNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ProductCreateProductRequest) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductCreateProductRequest) SetCode(v string) {
	o.Code = &v
}

// GetIsConfigurable returns the IsConfigurable field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetIsConfigurable() bool {
	if o == nil || isNil(o.IsConfigurable) {
		var ret bool
		return ret
	}
	return *o.IsConfigurable
}

// GetIsConfigurableOk returns a tuple with the IsConfigurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetIsConfigurableOk() (*bool, bool) {
	if o == nil || isNil(o.IsConfigurable) {
    return nil, false
	}
	return o.IsConfigurable, true
}

// HasIsConfigurable returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasIsConfigurable() bool {
	if o != nil && !isNil(o.IsConfigurable) {
		return true
	}

	return false
}

// SetIsConfigurable gets a reference to the given bool and assigns it to the IsConfigurable field.
func (o *ProductCreateProductRequest) SetIsConfigurable(v bool) {
	o.IsConfigurable = &v
}

// GetVariantAttributes returns the VariantAttributes field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetVariantAttributes() []string {
	if o == nil || isNil(o.VariantAttributes) {
		var ret []string
		return ret
	}
	return o.VariantAttributes
}

// GetVariantAttributesOk returns a tuple with the VariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetVariantAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.VariantAttributes) {
    return nil, false
	}
	return o.VariantAttributes, true
}

// HasVariantAttributes returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasVariantAttributes() bool {
	if o != nil && !isNil(o.VariantAttributes) {
		return true
	}

	return false
}

// SetVariantAttributes gets a reference to the given []string and assigns it to the VariantAttributes field.
func (o *ProductCreateProductRequest) SetVariantAttributes(v []string) {
	o.VariantAttributes = v
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetIsVirtual() bool {
	if o == nil || isNil(o.IsVirtual) {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetIsVirtualOk() (*bool, bool) {
	if o == nil || isNil(o.IsVirtual) {
    return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasIsVirtual() bool {
	if o != nil && !isNil(o.IsVirtual) {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *ProductCreateProductRequest) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

// GetIsGiftcard returns the IsGiftcard field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetIsGiftcard() bool {
	if o == nil || isNil(o.IsGiftcard) {
		var ret bool
		return ret
	}
	return *o.IsGiftcard
}

// GetIsGiftcardOk returns a tuple with the IsGiftcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetIsGiftcardOk() (*bool, bool) {
	if o == nil || isNil(o.IsGiftcard) {
    return nil, false
	}
	return o.IsGiftcard, true
}

// HasIsGiftcard returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasIsGiftcard() bool {
	if o != nil && !isNil(o.IsGiftcard) {
		return true
	}

	return false
}

// SetIsGiftcard gets a reference to the given bool and assigns it to the IsGiftcard field.
func (o *ProductCreateProductRequest) SetIsGiftcard(v bool) {
	o.IsGiftcard = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetUrlKey() ProductLocalizedText {
	if o == nil || isNil(o.UrlKey) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetUrlKeyOk() (*ProductLocalizedText, bool) {
	if o == nil || isNil(o.UrlKey) {
    return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasUrlKey() bool {
	if o != nil && !isNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductLocalizedText and assigns it to the UrlKey field.
func (o *ProductCreateProductRequest) SetUrlKey(v ProductLocalizedText) {
	o.UrlKey = &v
}

// GetMaxSaleableQuantity returns the MaxSaleableQuantity field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetMaxSaleableQuantity() int64 {
	if o == nil || isNil(o.MaxSaleableQuantity) {
		var ret int64
		return ret
	}
	return *o.MaxSaleableQuantity
}

// GetMaxSaleableQuantityOk returns a tuple with the MaxSaleableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetMaxSaleableQuantityOk() (*int64, bool) {
	if o == nil || isNil(o.MaxSaleableQuantity) {
    return nil, false
	}
	return o.MaxSaleableQuantity, true
}

// HasMaxSaleableQuantity returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasMaxSaleableQuantity() bool {
	if o != nil && !isNil(o.MaxSaleableQuantity) {
		return true
	}

	return false
}

// SetMaxSaleableQuantity gets a reference to the given int64 and assigns it to the MaxSaleableQuantity field.
func (o *ProductCreateProductRequest) SetMaxSaleableQuantity(v int64) {
	o.MaxSaleableQuantity = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductCreateProductRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ProductCreateProductRequest) GetVariants() map[string]ProductProductVariant {
	if o == nil || isNil(o.Variants) {
		var ret map[string]ProductProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductRequest) GetVariantsOk() (*map[string]ProductProductVariant, bool) {
	if o == nil || isNil(o.Variants) {
    return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ProductCreateProductRequest) HasVariants() bool {
	if o != nil && !isNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]ProductProductVariant and assigns it to the Variants field.
func (o *ProductCreateProductRequest) SetVariants(v map[string]ProductProductVariant) {
	o.Variants = &v
}

func (o ProductCreateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.IsConfigurable) {
		toSerialize["isConfigurable"] = o.IsConfigurable
	}
	if !isNil(o.VariantAttributes) {
		toSerialize["variantAttributes"] = o.VariantAttributes
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
	if !isNil(o.MaxSaleableQuantity) {
		toSerialize["maxSaleableQuantity"] = o.MaxSaleableQuantity
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return json.Marshal(toSerialize)
}

type NullableProductCreateProductRequest struct {
	value *ProductCreateProductRequest
	isSet bool
}

func (v NullableProductCreateProductRequest) Get() *ProductCreateProductRequest {
	return v.value
}

func (v *NullableProductCreateProductRequest) Set(val *ProductCreateProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateProductRequest(val *ProductCreateProductRequest) *NullableProductCreateProductRequest {
	return &NullableProductCreateProductRequest{value: val, isSet: true}
}

func (v NullableProductCreateProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


