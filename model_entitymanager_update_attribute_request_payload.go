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

// checks if the EntitymanagerUpdateAttributeRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerUpdateAttributeRequestPayload{}

// EntitymanagerUpdateAttributeRequestPayload struct for EntitymanagerUpdateAttributeRequestPayload
type EntitymanagerUpdateAttributeRequestPayload struct {
	Label                *string                 `json:"label,omitempty"`
	Default              *string                 `json:"default,omitempty"`
	Sort                 *int32                  `json:"sort,omitempty"`
	GroupCode            *string                 `json:"groupCode,omitempty"`
	Title                *map[string]string      `json:"title,omitempty"`
	RenderAs             *EntitymanagerRenderAs  `json:"renderAs,omitempty"`
	AiContext            *EntitymanagerAiContext `json:"aiContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerUpdateAttributeRequestPayload EntitymanagerUpdateAttributeRequestPayload

// NewEntitymanagerUpdateAttributeRequestPayload instantiates a new EntitymanagerUpdateAttributeRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerUpdateAttributeRequestPayload() *EntitymanagerUpdateAttributeRequestPayload {
	this := EntitymanagerUpdateAttributeRequestPayload{}
	var renderAs EntitymanagerRenderAs = ENTITYMANAGERRENDERAS_DEFAULT
	this.RenderAs = &renderAs
	return &this
}

// NewEntitymanagerUpdateAttributeRequestPayloadWithDefaults instantiates a new EntitymanagerUpdateAttributeRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerUpdateAttributeRequestPayloadWithDefaults() *EntitymanagerUpdateAttributeRequestPayload {
	this := EntitymanagerUpdateAttributeRequestPayload{}
	var renderAs EntitymanagerRenderAs = ENTITYMANAGERRENDERAS_DEFAULT
	this.RenderAs = &renderAs
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetLabel(v string) {
	o.Label = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetDefault(v string) {
	o.Default = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetSort() int32 {
	if o == nil || IsNil(o.Sort) {
		var ret int32
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetSortOk() (*int32, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int32 and assigns it to the Sort field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetSort(v int32) {
	o.Sort = &v
}

// GetGroupCode returns the GroupCode field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetGroupCode() string {
	if o == nil || IsNil(o.GroupCode) {
		var ret string
		return ret
	}
	return *o.GroupCode
}

// GetGroupCodeOk returns a tuple with the GroupCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetGroupCodeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupCode) {
		return nil, false
	}
	return o.GroupCode, true
}

// HasGroupCode returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasGroupCode() bool {
	if o != nil && !IsNil(o.GroupCode) {
		return true
	}

	return false
}

// SetGroupCode gets a reference to the given string and assigns it to the GroupCode field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetGroupCode(v string) {
	o.GroupCode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetTitle() map[string]string {
	if o == nil || IsNil(o.Title) {
		var ret map[string]string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetTitleOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given map[string]string and assigns it to the Title field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetTitle(v map[string]string) {
	o.Title = &v
}

// GetRenderAs returns the RenderAs field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetRenderAs() EntitymanagerRenderAs {
	if o == nil || IsNil(o.RenderAs) {
		var ret EntitymanagerRenderAs
		return ret
	}
	return *o.RenderAs
}

// GetRenderAsOk returns a tuple with the RenderAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetRenderAsOk() (*EntitymanagerRenderAs, bool) {
	if o == nil || IsNil(o.RenderAs) {
		return nil, false
	}
	return o.RenderAs, true
}

// HasRenderAs returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasRenderAs() bool {
	if o != nil && !IsNil(o.RenderAs) {
		return true
	}

	return false
}

// SetRenderAs gets a reference to the given EntitymanagerRenderAs and assigns it to the RenderAs field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetRenderAs(v EntitymanagerRenderAs) {
	o.RenderAs = &v
}

// GetAiContext returns the AiContext field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetAiContext() EntitymanagerAiContext {
	if o == nil || IsNil(o.AiContext) {
		var ret EntitymanagerAiContext
		return ret
	}
	return *o.AiContext
}

// GetAiContextOk returns a tuple with the AiContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) GetAiContextOk() (*EntitymanagerAiContext, bool) {
	if o == nil || IsNil(o.AiContext) {
		return nil, false
	}
	return o.AiContext, true
}

// HasAiContext returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeRequestPayload) HasAiContext() bool {
	if o != nil && !IsNil(o.AiContext) {
		return true
	}

	return false
}

// SetAiContext gets a reference to the given EntitymanagerAiContext and assigns it to the AiContext field.
func (o *EntitymanagerUpdateAttributeRequestPayload) SetAiContext(v EntitymanagerAiContext) {
	o.AiContext = &v
}

func (o EntitymanagerUpdateAttributeRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerUpdateAttributeRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.GroupCode) {
		toSerialize["groupCode"] = o.GroupCode
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.RenderAs) {
		toSerialize["renderAs"] = o.RenderAs
	}
	if !IsNil(o.AiContext) {
		toSerialize["aiContext"] = o.AiContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitymanagerUpdateAttributeRequestPayload) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerUpdateAttributeRequestPayload := _EntitymanagerUpdateAttributeRequestPayload{}

	err = json.Unmarshal(data, &varEntitymanagerUpdateAttributeRequestPayload)

	if err != nil {
		return err
	}

	*o = EntitymanagerUpdateAttributeRequestPayload(varEntitymanagerUpdateAttributeRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "default")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "groupCode")
		delete(additionalProperties, "title")
		delete(additionalProperties, "renderAs")
		delete(additionalProperties, "aiContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitymanagerUpdateAttributeRequestPayload struct {
	value *EntitymanagerUpdateAttributeRequestPayload
	isSet bool
}

func (v NullableEntitymanagerUpdateAttributeRequestPayload) Get() *EntitymanagerUpdateAttributeRequestPayload {
	return v.value
}

func (v *NullableEntitymanagerUpdateAttributeRequestPayload) Set(val *EntitymanagerUpdateAttributeRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerUpdateAttributeRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerUpdateAttributeRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerUpdateAttributeRequestPayload(val *EntitymanagerUpdateAttributeRequestPayload) *NullableEntitymanagerUpdateAttributeRequestPayload {
	return &NullableEntitymanagerUpdateAttributeRequestPayload{value: val, isSet: true}
}

func (v NullableEntitymanagerUpdateAttributeRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerUpdateAttributeRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
