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

// checks if the ProductUpdateProductRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateProductRequestV2{}

// ProductUpdateProductRequestV2 struct for ProductUpdateProductRequestV2
type ProductUpdateProductRequestV2 struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	AttributesMask *ProductFieldMask `json:"attributesMask,omitempty"`
	UrlKey *ProductLocalizedText `json:"urlKey,omitempty"`
	MaxSaleableQuantity *int64 `json:"maxSaleableQuantity,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	Variants *map[string]ProductProductVariant `json:"variants,omitempty"`
	MediaVariantAttributes []string `json:"mediaVariantAttributes,omitempty"`
}

// NewProductUpdateProductRequestV2 instantiates a new ProductUpdateProductRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateProductRequestV2() *ProductUpdateProductRequestV2 {
	this := ProductUpdateProductRequestV2{}
	return &this
}

// NewProductUpdateProductRequestV2WithDefaults instantiates a new ProductUpdateProductRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateProductRequestV2WithDefaults() *ProductUpdateProductRequestV2 {
	this := ProductUpdateProductRequestV2{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductUpdateProductRequestV2) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductUpdateProductRequestV2) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductUpdateProductRequestV2) SetCode(v string) {
	o.Code = &v
}

// GetAttributesMask returns the AttributesMask field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetAttributesMask() ProductFieldMask {
	if o == nil || IsNil(o.AttributesMask) {
		var ret ProductFieldMask
		return ret
	}
	return *o.AttributesMask
}

// GetAttributesMaskOk returns a tuple with the AttributesMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetAttributesMaskOk() (*ProductFieldMask, bool) {
	if o == nil || IsNil(o.AttributesMask) {
		return nil, false
	}
	return o.AttributesMask, true
}

// HasAttributesMask returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasAttributesMask() bool {
	if o != nil && !IsNil(o.AttributesMask) {
		return true
	}

	return false
}

// SetAttributesMask gets a reference to the given ProductFieldMask and assigns it to the AttributesMask field.
func (o *ProductUpdateProductRequestV2) SetAttributesMask(v ProductFieldMask) {
	o.AttributesMask = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetUrlKey() ProductLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetUrlKeyOk() (*ProductLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductLocalizedText and assigns it to the UrlKey field.
func (o *ProductUpdateProductRequestV2) SetUrlKey(v ProductLocalizedText) {
	o.UrlKey = &v
}

// GetMaxSaleableQuantity returns the MaxSaleableQuantity field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetMaxSaleableQuantity() int64 {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		var ret int64
		return ret
	}
	return *o.MaxSaleableQuantity
}

// GetMaxSaleableQuantityOk returns a tuple with the MaxSaleableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetMaxSaleableQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		return nil, false
	}
	return o.MaxSaleableQuantity, true
}

// HasMaxSaleableQuantity returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasMaxSaleableQuantity() bool {
	if o != nil && !IsNil(o.MaxSaleableQuantity) {
		return true
	}

	return false
}

// SetMaxSaleableQuantity gets a reference to the given int64 and assigns it to the MaxSaleableQuantity field.
func (o *ProductUpdateProductRequestV2) SetMaxSaleableQuantity(v int64) {
	o.MaxSaleableQuantity = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductUpdateProductRequestV2) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetVariants() map[string]ProductProductVariant {
	if o == nil || IsNil(o.Variants) {
		var ret map[string]ProductProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetVariantsOk() (*map[string]ProductProductVariant, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]ProductProductVariant and assigns it to the Variants field.
func (o *ProductUpdateProductRequestV2) SetVariants(v map[string]ProductProductVariant) {
	o.Variants = &v
}

// GetMediaVariantAttributes returns the MediaVariantAttributes field value if set, zero value otherwise.
func (o *ProductUpdateProductRequestV2) GetMediaVariantAttributes() []string {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		var ret []string
		return ret
	}
	return o.MediaVariantAttributes
}

// GetMediaVariantAttributesOk returns a tuple with the MediaVariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequestV2) GetMediaVariantAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		return nil, false
	}
	return o.MediaVariantAttributes, true
}

// HasMediaVariantAttributes returns a boolean if a field has been set.
func (o *ProductUpdateProductRequestV2) HasMediaVariantAttributes() bool {
	if o != nil && !IsNil(o.MediaVariantAttributes) {
		return true
	}

	return false
}

// SetMediaVariantAttributes gets a reference to the given []string and assigns it to the MediaVariantAttributes field.
func (o *ProductUpdateProductRequestV2) SetMediaVariantAttributes(v []string) {
	o.MediaVariantAttributes = v
}

func (o ProductUpdateProductRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateProductRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.AttributesMask) {
		toSerialize["attributesMask"] = o.AttributesMask
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
	return toSerialize, nil
}

type NullableProductUpdateProductRequestV2 struct {
	value *ProductUpdateProductRequestV2
	isSet bool
}

func (v NullableProductUpdateProductRequestV2) Get() *ProductUpdateProductRequestV2 {
	return v.value
}

func (v *NullableProductUpdateProductRequestV2) Set(val *ProductUpdateProductRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateProductRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateProductRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateProductRequestV2(val *ProductUpdateProductRequestV2) *NullableProductUpdateProductRequestV2 {
	return &NullableProductUpdateProductRequestV2{value: val, isSet: true}
}

func (v NullableProductUpdateProductRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateProductRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


