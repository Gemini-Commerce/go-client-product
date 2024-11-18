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

// checks if the ProductUpdateProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateProductRequest{}

// ProductUpdateProductRequest struct for ProductUpdateProductRequest
type ProductUpdateProductRequest struct {
	TenantId               *string                           `json:"tenantId,omitempty"`
	Id                     *string                           `json:"id,omitempty"`
	Code                   *string                           `json:"code,omitempty"`
	AttributesMask         *ProductFieldMask                 `json:"attributesMask,omitempty"`
	UrlKey                 *ProductLocalizedText             `json:"urlKey,omitempty"`
	MaxSaleableQuantity    *int64                            `json:"maxSaleableQuantity,omitempty"`
	MediaVariantAttributes []string                          `json:"mediaVariantAttributes,omitempty"`
	Attributes             *map[string]ProtobufAny           `json:"attributes,omitempty"`
	Variants               *map[string]ProductProductVariant `json:"variants,omitempty"`
	InReview               *bool                             `json:"inReview,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _ProductUpdateProductRequest ProductUpdateProductRequest

// NewProductUpdateProductRequest instantiates a new ProductUpdateProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateProductRequest() *ProductUpdateProductRequest {
	this := ProductUpdateProductRequest{}
	return &this
}

// NewProductUpdateProductRequestWithDefaults instantiates a new ProductUpdateProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateProductRequestWithDefaults() *ProductUpdateProductRequest {
	this := ProductUpdateProductRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductUpdateProductRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductUpdateProductRequest) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductUpdateProductRequest) SetCode(v string) {
	o.Code = &v
}

// GetAttributesMask returns the AttributesMask field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetAttributesMask() ProductFieldMask {
	if o == nil || IsNil(o.AttributesMask) {
		var ret ProductFieldMask
		return ret
	}
	return *o.AttributesMask
}

// GetAttributesMaskOk returns a tuple with the AttributesMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetAttributesMaskOk() (*ProductFieldMask, bool) {
	if o == nil || IsNil(o.AttributesMask) {
		return nil, false
	}
	return o.AttributesMask, true
}

// HasAttributesMask returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasAttributesMask() bool {
	if o != nil && !IsNil(o.AttributesMask) {
		return true
	}

	return false
}

// SetAttributesMask gets a reference to the given ProductFieldMask and assigns it to the AttributesMask field.
func (o *ProductUpdateProductRequest) SetAttributesMask(v ProductFieldMask) {
	o.AttributesMask = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetUrlKey() ProductLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetUrlKeyOk() (*ProductLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductLocalizedText and assigns it to the UrlKey field.
func (o *ProductUpdateProductRequest) SetUrlKey(v ProductLocalizedText) {
	o.UrlKey = &v
}

// GetMaxSaleableQuantity returns the MaxSaleableQuantity field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetMaxSaleableQuantity() int64 {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		var ret int64
		return ret
	}
	return *o.MaxSaleableQuantity
}

// GetMaxSaleableQuantityOk returns a tuple with the MaxSaleableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetMaxSaleableQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxSaleableQuantity) {
		return nil, false
	}
	return o.MaxSaleableQuantity, true
}

// HasMaxSaleableQuantity returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasMaxSaleableQuantity() bool {
	if o != nil && !IsNil(o.MaxSaleableQuantity) {
		return true
	}

	return false
}

// SetMaxSaleableQuantity gets a reference to the given int64 and assigns it to the MaxSaleableQuantity field.
func (o *ProductUpdateProductRequest) SetMaxSaleableQuantity(v int64) {
	o.MaxSaleableQuantity = &v
}

// GetMediaVariantAttributes returns the MediaVariantAttributes field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetMediaVariantAttributes() []string {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		var ret []string
		return ret
	}
	return o.MediaVariantAttributes
}

// GetMediaVariantAttributesOk returns a tuple with the MediaVariantAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetMediaVariantAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.MediaVariantAttributes) {
		return nil, false
	}
	return o.MediaVariantAttributes, true
}

// HasMediaVariantAttributes returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasMediaVariantAttributes() bool {
	if o != nil && !IsNil(o.MediaVariantAttributes) {
		return true
	}

	return false
}

// SetMediaVariantAttributes gets a reference to the given []string and assigns it to the MediaVariantAttributes field.
func (o *ProductUpdateProductRequest) SetMediaVariantAttributes(v []string) {
	o.MediaVariantAttributes = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductUpdateProductRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetVariants() map[string]ProductProductVariant {
	if o == nil || IsNil(o.Variants) {
		var ret map[string]ProductProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetVariantsOk() (*map[string]ProductProductVariant, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]ProductProductVariant and assigns it to the Variants field.
func (o *ProductUpdateProductRequest) SetVariants(v map[string]ProductProductVariant) {
	o.Variants = &v
}

// GetInReview returns the InReview field value if set, zero value otherwise.
func (o *ProductUpdateProductRequest) GetInReview() bool {
	if o == nil || IsNil(o.InReview) {
		var ret bool
		return ret
	}
	return *o.InReview
}

// GetInReviewOk returns a tuple with the InReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductRequest) GetInReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.InReview) {
		return nil, false
	}
	return o.InReview, true
}

// HasInReview returns a boolean if a field has been set.
func (o *ProductUpdateProductRequest) HasInReview() bool {
	if o != nil && !IsNil(o.InReview) {
		return true
	}

	return false
}

// SetInReview gets a reference to the given bool and assigns it to the InReview field.
func (o *ProductUpdateProductRequest) SetInReview(v bool) {
	o.InReview = &v
}

func (o ProductUpdateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateProductRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MediaVariantAttributes) {
		toSerialize["mediaVariantAttributes"] = o.MediaVariantAttributes
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	if !IsNil(o.InReview) {
		toSerialize["inReview"] = o.InReview
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductUpdateProductRequest) UnmarshalJSON(data []byte) (err error) {
	varProductUpdateProductRequest := _ProductUpdateProductRequest{}

	err = json.Unmarshal(data, &varProductUpdateProductRequest)

	if err != nil {
		return err
	}

	*o = ProductUpdateProductRequest(varProductUpdateProductRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "attributesMask")
		delete(additionalProperties, "urlKey")
		delete(additionalProperties, "maxSaleableQuantity")
		delete(additionalProperties, "mediaVariantAttributes")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "variants")
		delete(additionalProperties, "inReview")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductUpdateProductRequest struct {
	value *ProductUpdateProductRequest
	isSet bool
}

func (v NullableProductUpdateProductRequest) Get() *ProductUpdateProductRequest {
	return v.value
}

func (v *NullableProductUpdateProductRequest) Set(val *ProductUpdateProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateProductRequest(val *ProductUpdateProductRequest) *NullableProductUpdateProductRequest {
	return &NullableProductUpdateProductRequest{value: val, isSet: true}
}

func (v NullableProductUpdateProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
