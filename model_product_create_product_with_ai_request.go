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

// checks if the ProductCreateProductWithAIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCreateProductWithAIRequest{}

// ProductCreateProductWithAIRequest struct for ProductCreateProductWithAIRequest
type ProductCreateProductWithAIRequest struct {
	TenantId             *string                        `json:"tenantId,omitempty"`
	Product              *ProductCreateProductRequestV2 `json:"product,omitempty"`
	Locale               *string                        `json:"locale,omitempty"`
	ProductBrand         *string                        `json:"productBrand,omitempty"`
	ProductCode          *string                        `json:"productCode,omitempty"`
	ProductName          *string                        `json:"productName,omitempty"`
	SkipReview           *bool                          `json:"skipReview,omitempty"`
	AttributesToEnrich   []ProductAttributeToEnrich     `json:"attributesToEnrich,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductCreateProductWithAIRequest ProductCreateProductWithAIRequest

// NewProductCreateProductWithAIRequest instantiates a new ProductCreateProductWithAIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateProductWithAIRequest() *ProductCreateProductWithAIRequest {
	this := ProductCreateProductWithAIRequest{}
	return &this
}

// NewProductCreateProductWithAIRequestWithDefaults instantiates a new ProductCreateProductWithAIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateProductWithAIRequestWithDefaults() *ProductCreateProductWithAIRequest {
	this := ProductCreateProductWithAIRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductCreateProductWithAIRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetProduct() ProductCreateProductRequestV2 {
	if o == nil || IsNil(o.Product) {
		var ret ProductCreateProductRequestV2
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetProductOk() (*ProductCreateProductRequestV2, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductCreateProductRequestV2 and assigns it to the Product field.
func (o *ProductCreateProductWithAIRequest) SetProduct(v ProductCreateProductRequestV2) {
	o.Product = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ProductCreateProductWithAIRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetProductBrand returns the ProductBrand field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetProductBrand() string {
	if o == nil || IsNil(o.ProductBrand) {
		var ret string
		return ret
	}
	return *o.ProductBrand
}

// GetProductBrandOk returns a tuple with the ProductBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetProductBrandOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBrand) {
		return nil, false
	}
	return o.ProductBrand, true
}

// HasProductBrand returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasProductBrand() bool {
	if o != nil && !IsNil(o.ProductBrand) {
		return true
	}

	return false
}

// SetProductBrand gets a reference to the given string and assigns it to the ProductBrand field.
func (o *ProductCreateProductWithAIRequest) SetProductBrand(v string) {
	o.ProductBrand = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *ProductCreateProductWithAIRequest) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *ProductCreateProductWithAIRequest) SetProductName(v string) {
	o.ProductName = &v
}

// GetSkipReview returns the SkipReview field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetSkipReview() bool {
	if o == nil || IsNil(o.SkipReview) {
		var ret bool
		return ret
	}
	return *o.SkipReview
}

// GetSkipReviewOk returns a tuple with the SkipReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetSkipReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipReview) {
		return nil, false
	}
	return o.SkipReview, true
}

// HasSkipReview returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasSkipReview() bool {
	if o != nil && !IsNil(o.SkipReview) {
		return true
	}

	return false
}

// SetSkipReview gets a reference to the given bool and assigns it to the SkipReview field.
func (o *ProductCreateProductWithAIRequest) SetSkipReview(v bool) {
	o.SkipReview = &v
}

// GetAttributesToEnrich returns the AttributesToEnrich field value if set, zero value otherwise.
func (o *ProductCreateProductWithAIRequest) GetAttributesToEnrich() []ProductAttributeToEnrich {
	if o == nil || IsNil(o.AttributesToEnrich) {
		var ret []ProductAttributeToEnrich
		return ret
	}
	return o.AttributesToEnrich
}

// GetAttributesToEnrichOk returns a tuple with the AttributesToEnrich field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductWithAIRequest) GetAttributesToEnrichOk() ([]ProductAttributeToEnrich, bool) {
	if o == nil || IsNil(o.AttributesToEnrich) {
		return nil, false
	}
	return o.AttributesToEnrich, true
}

// HasAttributesToEnrich returns a boolean if a field has been set.
func (o *ProductCreateProductWithAIRequest) HasAttributesToEnrich() bool {
	if o != nil && !IsNil(o.AttributesToEnrich) {
		return true
	}

	return false
}

// SetAttributesToEnrich gets a reference to the given []ProductAttributeToEnrich and assigns it to the AttributesToEnrich field.
func (o *ProductCreateProductWithAIRequest) SetAttributesToEnrich(v []ProductAttributeToEnrich) {
	o.AttributesToEnrich = v
}

func (o ProductCreateProductWithAIRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductCreateProductWithAIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.ProductBrand) {
		toSerialize["productBrand"] = o.ProductBrand
	}
	if !IsNil(o.ProductCode) {
		toSerialize["productCode"] = o.ProductCode
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.SkipReview) {
		toSerialize["skipReview"] = o.SkipReview
	}
	if !IsNil(o.AttributesToEnrich) {
		toSerialize["attributesToEnrich"] = o.AttributesToEnrich
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductCreateProductWithAIRequest) UnmarshalJSON(data []byte) (err error) {
	varProductCreateProductWithAIRequest := _ProductCreateProductWithAIRequest{}

	err = json.Unmarshal(data, &varProductCreateProductWithAIRequest)

	if err != nil {
		return err
	}

	*o = ProductCreateProductWithAIRequest(varProductCreateProductWithAIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "product")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "productBrand")
		delete(additionalProperties, "productCode")
		delete(additionalProperties, "productName")
		delete(additionalProperties, "skipReview")
		delete(additionalProperties, "attributesToEnrich")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductCreateProductWithAIRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductCreateProductWithAIRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductCreateProductWithAIRequest struct {
	value *ProductCreateProductWithAIRequest
	isSet bool
}

func (v NullableProductCreateProductWithAIRequest) Get() *ProductCreateProductWithAIRequest {
	return v.value
}

func (v *NullableProductCreateProductWithAIRequest) Set(val *ProductCreateProductWithAIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateProductWithAIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateProductWithAIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateProductWithAIRequest(val *ProductCreateProductWithAIRequest) *NullableProductCreateProductWithAIRequest {
	return &NullableProductCreateProductWithAIRequest{value: val, isSet: true}
}

func (v NullableProductCreateProductWithAIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateProductWithAIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
