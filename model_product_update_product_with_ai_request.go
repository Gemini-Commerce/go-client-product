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

// checks if the ProductUpdateProductWithAIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateProductWithAIRequest{}

// ProductUpdateProductWithAIRequest struct for ProductUpdateProductWithAIRequest
type ProductUpdateProductWithAIRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	Locale *string `json:"locale,omitempty"`
	SkipReview *bool `json:"skipReview,omitempty"`
	AttributesToEnrich []ProductAttributeToEnrich `json:"attributesToEnrich,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductUpdateProductWithAIRequest ProductUpdateProductWithAIRequest

// NewProductUpdateProductWithAIRequest instantiates a new ProductUpdateProductWithAIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateProductWithAIRequest() *ProductUpdateProductWithAIRequest {
	this := ProductUpdateProductWithAIRequest{}
	return &this
}

// NewProductUpdateProductWithAIRequestWithDefaults instantiates a new ProductUpdateProductWithAIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateProductWithAIRequestWithDefaults() *ProductUpdateProductWithAIRequest {
	this := ProductUpdateProductWithAIRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductUpdateProductWithAIRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductWithAIRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// &#39;Has&#39;TenantId returns a boolean if a field has been set.
func (o *ProductUpdateProductWithAIRequest) &#39;Has&#39;TenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductUpdateProductWithAIRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductUpdateProductWithAIRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductWithAIRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// &#39;Has&#39;Id returns a boolean if a field has been set.
func (o *ProductUpdateProductWithAIRequest) &#39;Has&#39;Id() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductUpdateProductWithAIRequest) SetId(v string) {
	o.Id = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ProductUpdateProductWithAIRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductWithAIRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// &#39;Has&#39;Locale returns a boolean if a field has been set.
func (o *ProductUpdateProductWithAIRequest) &#39;Has&#39;Locale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ProductUpdateProductWithAIRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetSkipReview returns the SkipReview field value if set, zero value otherwise.
func (o *ProductUpdateProductWithAIRequest) GetSkipReview() bool {
	if o == nil || IsNil(o.SkipReview) {
		var ret bool
		return ret
	}
	return *o.SkipReview
}

// GetSkipReviewOk returns a tuple with the SkipReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductWithAIRequest) GetSkipReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipReview) {
		return nil, false
	}
	return o.SkipReview, true
}

// &#39;Has&#39;SkipReview returns a boolean if a field has been set.
func (o *ProductUpdateProductWithAIRequest) &#39;Has&#39;SkipReview() bool {
	if o != nil && !IsNil(o.SkipReview) {
		return true
	}

	return false
}

// SetSkipReview gets a reference to the given bool and assigns it to the SkipReview field.
func (o *ProductUpdateProductWithAIRequest) SetSkipReview(v bool) {
	o.SkipReview = &v
}

// GetAttributesToEnrich returns the AttributesToEnrich field value if set, zero value otherwise.
func (o *ProductUpdateProductWithAIRequest) GetAttributesToEnrich() []ProductAttributeToEnrich {
	if o == nil || IsNil(o.AttributesToEnrich) {
		var ret []ProductAttributeToEnrich
		return ret
	}
	return o.AttributesToEnrich
}

// GetAttributesToEnrichOk returns a tuple with the AttributesToEnrich field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateProductWithAIRequest) GetAttributesToEnrichOk() ([]ProductAttributeToEnrich, bool) {
	if o == nil || IsNil(o.AttributesToEnrich) {
		return nil, false
	}
	return o.AttributesToEnrich, true
}

// &#39;Has&#39;AttributesToEnrich returns a boolean if a field has been set.
func (o *ProductUpdateProductWithAIRequest) &#39;Has&#39;AttributesToEnrich() bool {
	if o != nil && !IsNil(o.AttributesToEnrich) {
		return true
	}

	return false
}

// SetAttributesToEnrich gets a reference to the given []ProductAttributeToEnrich and assigns it to the AttributesToEnrich field.
func (o *ProductUpdateProductWithAIRequest) SetAttributesToEnrich(v []ProductAttributeToEnrich) {
	o.AttributesToEnrich = v
}

func (o ProductUpdateProductWithAIRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateProductWithAIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
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

func (o *ProductUpdateProductWithAIRequest) UnmarshalJSON(data []byte) (err error) {
	varProductUpdateProductWithAIRequest := _ProductUpdateProductWithAIRequest{}

	err = json.Unmarshal(data, &varProductUpdateProductWithAIRequest)

	if err != nil {
		return err
	}

	*o = ProductUpdateProductWithAIRequest(varProductUpdateProductWithAIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "skipReview")
		delete(additionalProperties, "attributesToEnrich")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductUpdateProductWithAIRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductUpdateProductWithAIRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductUpdateProductWithAIRequest struct {
	value *ProductUpdateProductWithAIRequest
	isSet bool
}

func (v NullableProductUpdateProductWithAIRequest) Get() *ProductUpdateProductWithAIRequest {
	return v.value
}

func (v *NullableProductUpdateProductWithAIRequest) Set(val *ProductUpdateProductWithAIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateProductWithAIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateProductWithAIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateProductWithAIRequest(val *ProductUpdateProductWithAIRequest) *NullableProductUpdateProductWithAIRequest {
	return &NullableProductUpdateProductWithAIRequest{value: val, isSet: true}
}

func (v NullableProductUpdateProductWithAIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateProductWithAIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


