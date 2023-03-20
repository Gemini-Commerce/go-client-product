/*
product/product.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product

import (
	"encoding/json"
)

// ProductCreateProductRequest struct for ProductCreateProductRequest
type ProductCreateProductRequest struct {
	TenantId          *string                           `json:"tenantId,omitempty"`
	EntityType        *string                           `json:"entityType,omitempty"`
	EntityCode        *string                           `json:"entityCode,omitempty"`
	Code              *string                           `json:"code,omitempty"`
	IsConfigurable    *bool                             `json:"isConfigurable,omitempty"`
	VariantAttributes []string                          `json:"variantAttributes,omitempty"`
	IsVirtual         *bool                             `json:"isVirtual,omitempty"`
	IsGiftcard        *bool                             `json:"isGiftcard,omitempty"`
	UrlKey            *ProductLocalizedText             `json:"urlKey,omitempty"`
	Attributes        *map[string]ProtobufAny           `json:"attributes,omitempty"`
	Variants          *map[string]ProductProductVariant `json:"variants,omitempty"`
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
