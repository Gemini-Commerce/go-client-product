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

// checks if the EntitymanagerAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerAttribute{}

// EntitymanagerAttribute struct for EntitymanagerAttribute
type EntitymanagerAttribute struct {
	EntityId *string `json:"entityId,omitempty"`
	Code *string `json:"code,omitempty"`
	Label *string `json:"label,omitempty"`
	Type *EntitymanagerTypes `json:"type,omitempty"`
	OptionList *string `json:"optionList,omitempty"`
	Entity *string `json:"entity,omitempty"`
	Default *string `json:"default,omitempty"`
	IsRequired *bool `json:"isRequired,omitempty"`
	IsSystem *bool `json:"isSystem,omitempty"`
	IsRepeated *bool `json:"isRepeated,omitempty"`
	Sort *int32 `json:"sort,omitempty"`
	GroupCode *string `json:"groupCode,omitempty"`
	Title *map[string]string `json:"title,omitempty"`
	RenderAs *EntitymanagerRenderAs `json:"renderAs,omitempty"`
	AiContext *EntitymanagerAiContext `json:"aiContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitymanagerAttribute EntitymanagerAttribute

// NewEntitymanagerAttribute instantiates a new EntitymanagerAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerAttribute() *EntitymanagerAttribute {
	this := EntitymanagerAttribute{}
	var type_ EntitymanagerTypes = ENTITYMANAGERTYPES_TEXT
	this.Type = &type_
	var renderAs EntitymanagerRenderAs = ENTITYMANAGERRENDERAS_DEFAULT
	this.RenderAs = &renderAs
	return &this
}

// NewEntitymanagerAttributeWithDefaults instantiates a new EntitymanagerAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerAttributeWithDefaults() *EntitymanagerAttribute {
	this := EntitymanagerAttribute{}
	var type_ EntitymanagerTypes = ENTITYMANAGERTYPES_TEXT
	this.Type = &type_
	var renderAs EntitymanagerRenderAs = ENTITYMANAGERRENDERAS_DEFAULT
	this.RenderAs = &renderAs
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntitymanagerAttribute) SetEntityId(v string) {
	o.EntityId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerAttribute) SetCode(v string) {
	o.Code = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EntitymanagerAttribute) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetType() EntitymanagerTypes {
	if o == nil || IsNil(o.Type) {
		var ret EntitymanagerTypes
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetTypeOk() (*EntitymanagerTypes, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EntitymanagerTypes and assigns it to the Type field.
func (o *EntitymanagerAttribute) SetType(v EntitymanagerTypes) {
	o.Type = &v
}

// GetOptionList returns the OptionList field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetOptionList() string {
	if o == nil || IsNil(o.OptionList) {
		var ret string
		return ret
	}
	return *o.OptionList
}

// GetOptionListOk returns a tuple with the OptionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetOptionListOk() (*string, bool) {
	if o == nil || IsNil(o.OptionList) {
		return nil, false
	}
	return o.OptionList, true
}

// HasOptionList returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasOptionList() bool {
	if o != nil && !IsNil(o.OptionList) {
		return true
	}

	return false
}

// SetOptionList gets a reference to the given string and assigns it to the OptionList field.
func (o *EntitymanagerAttribute) SetOptionList(v string) {
	o.OptionList = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *EntitymanagerAttribute) SetEntity(v string) {
	o.Entity = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *EntitymanagerAttribute) SetDefault(v string) {
	o.Default = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *EntitymanagerAttribute) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *EntitymanagerAttribute) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetIsRepeated returns the IsRepeated field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetIsRepeated() bool {
	if o == nil || IsNil(o.IsRepeated) {
		var ret bool
		return ret
	}
	return *o.IsRepeated
}

// GetIsRepeatedOk returns a tuple with the IsRepeated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetIsRepeatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRepeated) {
		return nil, false
	}
	return o.IsRepeated, true
}

// HasIsRepeated returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasIsRepeated() bool {
	if o != nil && !IsNil(o.IsRepeated) {
		return true
	}

	return false
}

// SetIsRepeated gets a reference to the given bool and assigns it to the IsRepeated field.
func (o *EntitymanagerAttribute) SetIsRepeated(v bool) {
	o.IsRepeated = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetSort() int32 {
	if o == nil || IsNil(o.Sort) {
		var ret int32
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetSortOk() (*int32, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int32 and assigns it to the Sort field.
func (o *EntitymanagerAttribute) SetSort(v int32) {
	o.Sort = &v
}

// GetGroupCode returns the GroupCode field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetGroupCode() string {
	if o == nil || IsNil(o.GroupCode) {
		var ret string
		return ret
	}
	return *o.GroupCode
}

// GetGroupCodeOk returns a tuple with the GroupCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetGroupCodeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupCode) {
		return nil, false
	}
	return o.GroupCode, true
}

// HasGroupCode returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasGroupCode() bool {
	if o != nil && !IsNil(o.GroupCode) {
		return true
	}

	return false
}

// SetGroupCode gets a reference to the given string and assigns it to the GroupCode field.
func (o *EntitymanagerAttribute) SetGroupCode(v string) {
	o.GroupCode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetTitle() map[string]string {
	if o == nil || IsNil(o.Title) {
		var ret map[string]string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetTitleOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given map[string]string and assigns it to the Title field.
func (o *EntitymanagerAttribute) SetTitle(v map[string]string) {
	o.Title = &v
}

// GetRenderAs returns the RenderAs field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetRenderAs() EntitymanagerRenderAs {
	if o == nil || IsNil(o.RenderAs) {
		var ret EntitymanagerRenderAs
		return ret
	}
	return *o.RenderAs
}

// GetRenderAsOk returns a tuple with the RenderAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetRenderAsOk() (*EntitymanagerRenderAs, bool) {
	if o == nil || IsNil(o.RenderAs) {
		return nil, false
	}
	return o.RenderAs, true
}

// HasRenderAs returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasRenderAs() bool {
	if o != nil && !IsNil(o.RenderAs) {
		return true
	}

	return false
}

// SetRenderAs gets a reference to the given EntitymanagerRenderAs and assigns it to the RenderAs field.
func (o *EntitymanagerAttribute) SetRenderAs(v EntitymanagerRenderAs) {
	o.RenderAs = &v
}

// GetAiContext returns the AiContext field value if set, zero value otherwise.
func (o *EntitymanagerAttribute) GetAiContext() EntitymanagerAiContext {
	if o == nil || IsNil(o.AiContext) {
		var ret EntitymanagerAiContext
		return ret
	}
	return *o.AiContext
}

// GetAiContextOk returns a tuple with the AiContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttribute) GetAiContextOk() (*EntitymanagerAiContext, bool) {
	if o == nil || IsNil(o.AiContext) {
		return nil, false
	}
	return o.AiContext, true
}

// HasAiContext returns a boolean if a field has been set.
func (o *EntitymanagerAttribute) HasAiContext() bool {
	if o != nil && !IsNil(o.AiContext) {
		return true
	}

	return false
}

// SetAiContext gets a reference to the given EntitymanagerAiContext and assigns it to the AiContext field.
func (o *EntitymanagerAttribute) SetAiContext(v EntitymanagerAiContext) {
	o.AiContext = &v
}

func (o EntitymanagerAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.OptionList) {
		toSerialize["optionList"] = o.OptionList
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !IsNil(o.IsSystem) {
		toSerialize["isSystem"] = o.IsSystem
	}
	if !IsNil(o.IsRepeated) {
		toSerialize["isRepeated"] = o.IsRepeated
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

func (o *EntitymanagerAttribute) UnmarshalJSON(data []byte) (err error) {
	varEntitymanagerAttribute := _EntitymanagerAttribute{}

	err = json.Unmarshal(data, &varEntitymanagerAttribute)

	if err != nil {
		return err
	}

	*o = EntitymanagerAttribute(varEntitymanagerAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entityId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "optionList")
		delete(additionalProperties, "entity")
		delete(additionalProperties, "default")
		delete(additionalProperties, "isRequired")
		delete(additionalProperties, "isSystem")
		delete(additionalProperties, "isRepeated")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "groupCode")
		delete(additionalProperties, "title")
		delete(additionalProperties, "renderAs")
		delete(additionalProperties, "aiContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EntitymanagerAttribute) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *EntitymanagerAttribute) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableEntitymanagerAttribute struct {
	value *EntitymanagerAttribute
	isSet bool
}

func (v NullableEntitymanagerAttribute) Get() *EntitymanagerAttribute {
	return v.value
}

func (v *NullableEntitymanagerAttribute) Set(val *EntitymanagerAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerAttribute(val *EntitymanagerAttribute) *NullableEntitymanagerAttribute {
	return &NullableEntitymanagerAttribute{value: val, isSet: true}
}

func (v NullableEntitymanagerAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


