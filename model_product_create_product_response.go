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

// checks if the ProductCreateProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCreateProductResponse{}

// ProductCreateProductResponse The CreateProductResponse message is used to provide a response after creating a product within the system. It includes fields that indicate the success of the product creation and any errors encountered during the process.
type ProductCreateProductResponse struct {
	// Indicates whether the product creation was successful or not.
	Success *bool `json:"success,omitempty"`
	// Represents the ID of the created product.
	Id *string `json:"id,omitempty"`
	// Contains a list of ProductResponseError messages, indicating any errors related to the product creation.
	ProductErrors []ProductProductResponseError `json:"productErrors,omitempty"`
	// Contains a list of AttributeResponseError messages, indicating any errors related to the attributes of the product.
	AttributeErrors      []ProductAttributeResponseError `json:"attributeErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductCreateProductResponse ProductCreateProductResponse

// NewProductCreateProductResponse instantiates a new ProductCreateProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateProductResponse() *ProductCreateProductResponse {
	this := ProductCreateProductResponse{}
	return &this
}

// NewProductCreateProductResponseWithDefaults instantiates a new ProductCreateProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateProductResponseWithDefaults() *ProductCreateProductResponse {
	this := ProductCreateProductResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ProductCreateProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductCreateProductResponse) SetId(v string) {
	o.Id = &v
}

// GetProductErrors returns the ProductErrors field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetProductErrors() []ProductProductResponseError {
	if o == nil || IsNil(o.ProductErrors) {
		var ret []ProductProductResponseError
		return ret
	}
	return o.ProductErrors
}

// GetProductErrorsOk returns a tuple with the ProductErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetProductErrorsOk() ([]ProductProductResponseError, bool) {
	if o == nil || IsNil(o.ProductErrors) {
		return nil, false
	}
	return o.ProductErrors, true
}

// HasProductErrors returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasProductErrors() bool {
	if o != nil && !IsNil(o.ProductErrors) {
		return true
	}

	return false
}

// SetProductErrors gets a reference to the given []ProductProductResponseError and assigns it to the ProductErrors field.
func (o *ProductCreateProductResponse) SetProductErrors(v []ProductProductResponseError) {
	o.ProductErrors = v
}

// GetAttributeErrors returns the AttributeErrors field value if set, zero value otherwise.
func (o *ProductCreateProductResponse) GetAttributeErrors() []ProductAttributeResponseError {
	if o == nil || IsNil(o.AttributeErrors) {
		var ret []ProductAttributeResponseError
		return ret
	}
	return o.AttributeErrors
}

// GetAttributeErrorsOk returns a tuple with the AttributeErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateProductResponse) GetAttributeErrorsOk() ([]ProductAttributeResponseError, bool) {
	if o == nil || IsNil(o.AttributeErrors) {
		return nil, false
	}
	return o.AttributeErrors, true
}

// HasAttributeErrors returns a boolean if a field has been set.
func (o *ProductCreateProductResponse) HasAttributeErrors() bool {
	if o != nil && !IsNil(o.AttributeErrors) {
		return true
	}

	return false
}

// SetAttributeErrors gets a reference to the given []ProductAttributeResponseError and assigns it to the AttributeErrors field.
func (o *ProductCreateProductResponse) SetAttributeErrors(v []ProductAttributeResponseError) {
	o.AttributeErrors = v
}

func (o ProductCreateProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductCreateProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductErrors) {
		toSerialize["productErrors"] = o.ProductErrors
	}
	if !IsNil(o.AttributeErrors) {
		toSerialize["attributeErrors"] = o.AttributeErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductCreateProductResponse) UnmarshalJSON(data []byte) (err error) {
	varProductCreateProductResponse := _ProductCreateProductResponse{}

	err = json.Unmarshal(data, &varProductCreateProductResponse)

	if err != nil {
		return err
	}

	*o = ProductCreateProductResponse(varProductCreateProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "id")
		delete(additionalProperties, "productErrors")
		delete(additionalProperties, "attributeErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductCreateProductResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductCreateProductResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductCreateProductResponse struct {
	value *ProductCreateProductResponse
	isSet bool
}

func (v NullableProductCreateProductResponse) Get() *ProductCreateProductResponse {
	return v.value
}

func (v *NullableProductCreateProductResponse) Set(val *ProductCreateProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateProductResponse(val *ProductCreateProductResponse) *NullableProductCreateProductResponse {
	return &NullableProductCreateProductResponse{value: val, isSet: true}
}

func (v NullableProductCreateProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
