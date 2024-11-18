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

// checks if the EntitymanagerAttributeGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerAttributeGroup{}

// EntitymanagerAttributeGroup struct for EntitymanagerAttributeGroup
type EntitymanagerAttributeGroup struct {
	Code       *string                            `json:"code,omitempty"`
	Label      *ProductentitymanagerLocalizedText `json:"label,omitempty"`
	Sort       *int32                             `json:"sort,omitempty"`
	Visibility []string                           `json:"visibility,omitempty"`
}

// NewEntitymanagerAttributeGroup instantiates a new EntitymanagerAttributeGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerAttributeGroup() *EntitymanagerAttributeGroup {
	this := EntitymanagerAttributeGroup{}
	return &this
}

// NewEntitymanagerAttributeGroupWithDefaults instantiates a new EntitymanagerAttributeGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerAttributeGroupWithDefaults() *EntitymanagerAttributeGroup {
	this := EntitymanagerAttributeGroup{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerAttributeGroup) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeGroup) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerAttributeGroup) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerAttributeGroup) SetCode(v string) {
	o.Code = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EntitymanagerAttributeGroup) GetLabel() ProductentitymanagerLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret ProductentitymanagerLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeGroup) GetLabelOk() (*ProductentitymanagerLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EntitymanagerAttributeGroup) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ProductentitymanagerLocalizedText and assigns it to the Label field.
func (o *EntitymanagerAttributeGroup) SetLabel(v ProductentitymanagerLocalizedText) {
	o.Label = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EntitymanagerAttributeGroup) GetSort() int32 {
	if o == nil || IsNil(o.Sort) {
		var ret int32
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeGroup) GetSortOk() (*int32, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EntitymanagerAttributeGroup) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int32 and assigns it to the Sort field.
func (o *EntitymanagerAttributeGroup) SetSort(v int32) {
	o.Sort = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *EntitymanagerAttributeGroup) GetVisibility() []string {
	if o == nil || IsNil(o.Visibility) {
		var ret []string
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeGroup) GetVisibilityOk() ([]string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *EntitymanagerAttributeGroup) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given []string and assigns it to the Visibility field.
func (o *EntitymanagerAttributeGroup) SetVisibility(v []string) {
	o.Visibility = v
}

func (o EntitymanagerAttributeGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerAttributeGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableEntitymanagerAttributeGroup struct {
	value *EntitymanagerAttributeGroup
	isSet bool
}

func (v NullableEntitymanagerAttributeGroup) Get() *EntitymanagerAttributeGroup {
	return v.value
}

func (v *NullableEntitymanagerAttributeGroup) Set(val *EntitymanagerAttributeGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerAttributeGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerAttributeGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerAttributeGroup(val *EntitymanagerAttributeGroup) *NullableEntitymanagerAttributeGroup {
	return &NullableEntitymanagerAttributeGroup{value: val, isSet: true}
}

func (v NullableEntitymanagerAttributeGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerAttributeGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
