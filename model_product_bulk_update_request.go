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

// checks if the ProductBulkUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBulkUpdateRequest{}

// ProductBulkUpdateRequest struct for ProductBulkUpdateRequest
type ProductBulkUpdateRequest struct {
	TenantId             *string                          `json:"tenantId,omitempty"`
	ProductIds           []string                         `json:"productIds,omitempty"`
	Payload              *ProductBulkUpdateRequestPayload `json:"payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductBulkUpdateRequest ProductBulkUpdateRequest

// NewProductBulkUpdateRequest instantiates a new ProductBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkUpdateRequest() *ProductBulkUpdateRequest {
	this := ProductBulkUpdateRequest{}
	return &this
}

// NewProductBulkUpdateRequestWithDefaults instantiates a new ProductBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkUpdateRequestWithDefaults() *ProductBulkUpdateRequest {
	this := ProductBulkUpdateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductBulkUpdateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequest) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequest) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequest) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *ProductBulkUpdateRequest) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductBulkUpdateRequest) GetPayload() ProductBulkUpdateRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ProductBulkUpdateRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkUpdateRequest) GetPayloadOk() (*ProductBulkUpdateRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductBulkUpdateRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductBulkUpdateRequestPayload and assigns it to the Payload field.
func (o *ProductBulkUpdateRequest) SetPayload(v ProductBulkUpdateRequestPayload) {
	o.Payload = &v
}

func (o ProductBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBulkUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ProductIds) {
		toSerialize["productIds"] = o.ProductIds
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductBulkUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varProductBulkUpdateRequest := _ProductBulkUpdateRequest{}

	err = json.Unmarshal(data, &varProductBulkUpdateRequest)

	if err != nil {
		return err
	}

	*o = ProductBulkUpdateRequest(varProductBulkUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "productIds")
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductBulkUpdateRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populate the value of well-known types
func (o *ProductBulkUpdateRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductBulkUpdateRequest struct {
	value *ProductBulkUpdateRequest
	isSet bool
}

func (v NullableProductBulkUpdateRequest) Get() *ProductBulkUpdateRequest {
	return v.value
}

func (v *NullableProductBulkUpdateRequest) Set(val *ProductBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkUpdateRequest(val *ProductBulkUpdateRequest) *NullableProductBulkUpdateRequest {
	return &NullableProductBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableProductBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
