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

// EntitymanagerAttributeOption struct for EntitymanagerAttributeOption
type EntitymanagerAttributeOption struct {
	Option   *map[string]string                   `json:"option,omitempty"`
	Sort     *int64                               `json:"sort,omitempty"`
	Id       *string                              `json:"id,omitempty"`
	Code     *string                              `json:"code,omitempty"`
	Value    *ProductentitymanagerLocalizedText   `json:"value,omitempty"`
	Swatches []EntitymanagerAttributeOptionSwatch `json:"swatches,omitempty"`
}

// NewEntitymanagerAttributeOption instantiates a new EntitymanagerAttributeOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerAttributeOption() *EntitymanagerAttributeOption {
	this := EntitymanagerAttributeOption{}
	return &this
}

// NewEntitymanagerAttributeOptionWithDefaults instantiates a new EntitymanagerAttributeOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerAttributeOptionWithDefaults() *EntitymanagerAttributeOption {
	this := EntitymanagerAttributeOption{}
	return &this
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetOption() map[string]string {
	if o == nil || isNil(o.Option) {
		var ret map[string]string
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetOptionOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasOption() bool {
	if o != nil && !isNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given map[string]string and assigns it to the Option field.
func (o *EntitymanagerAttributeOption) SetOption(v map[string]string) {
	o.Option = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetSort() int64 {
	if o == nil || isNil(o.Sort) {
		var ret int64
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetSortOk() (*int64, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int64 and assigns it to the Sort field.
func (o *EntitymanagerAttributeOption) SetSort(v int64) {
	o.Sort = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitymanagerAttributeOption) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerAttributeOption) SetCode(v string) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetValue() ProductentitymanagerLocalizedText {
	if o == nil || isNil(o.Value) {
		var ret ProductentitymanagerLocalizedText
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetValueOk() (*ProductentitymanagerLocalizedText, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ProductentitymanagerLocalizedText and assigns it to the Value field.
func (o *EntitymanagerAttributeOption) SetValue(v ProductentitymanagerLocalizedText) {
	o.Value = &v
}

// GetSwatches returns the Swatches field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOption) GetSwatches() []EntitymanagerAttributeOptionSwatch {
	if o == nil || isNil(o.Swatches) {
		var ret []EntitymanagerAttributeOptionSwatch
		return ret
	}
	return o.Swatches
}

// GetSwatchesOk returns a tuple with the Swatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOption) GetSwatchesOk() ([]EntitymanagerAttributeOptionSwatch, bool) {
	if o == nil || isNil(o.Swatches) {
		return nil, false
	}
	return o.Swatches, true
}

// HasSwatches returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOption) HasSwatches() bool {
	if o != nil && !isNil(o.Swatches) {
		return true
	}

	return false
}

// SetSwatches gets a reference to the given []EntitymanagerAttributeOptionSwatch and assigns it to the Swatches field.
func (o *EntitymanagerAttributeOption) SetSwatches(v []EntitymanagerAttributeOptionSwatch) {
	o.Swatches = v
}

func (o EntitymanagerAttributeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Swatches) {
		toSerialize["swatches"] = o.Swatches
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerAttributeOption struct {
	value *EntitymanagerAttributeOption
	isSet bool
}

func (v NullableEntitymanagerAttributeOption) Get() *EntitymanagerAttributeOption {
	return v.value
}

func (v *NullableEntitymanagerAttributeOption) Set(val *EntitymanagerAttributeOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerAttributeOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerAttributeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerAttributeOption(val *EntitymanagerAttributeOption) *NullableEntitymanagerAttributeOption {
	return &NullableEntitymanagerAttributeOption{value: val, isSet: true}
}

func (v NullableEntitymanagerAttributeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerAttributeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
