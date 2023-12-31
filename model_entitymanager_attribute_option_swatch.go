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

// checks if the EntitymanagerAttributeOptionSwatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerAttributeOptionSwatch{}

// EntitymanagerAttributeOptionSwatch struct for EntitymanagerAttributeOptionSwatch
type EntitymanagerAttributeOptionSwatch struct {
	Grn *string `json:"grn,omitempty"`
	Hex *string `json:"hex,omitempty"`
}

// NewEntitymanagerAttributeOptionSwatch instantiates a new EntitymanagerAttributeOptionSwatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerAttributeOptionSwatch() *EntitymanagerAttributeOptionSwatch {
	this := EntitymanagerAttributeOptionSwatch{}
	return &this
}

// NewEntitymanagerAttributeOptionSwatchWithDefaults instantiates a new EntitymanagerAttributeOptionSwatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerAttributeOptionSwatchWithDefaults() *EntitymanagerAttributeOptionSwatch {
	this := EntitymanagerAttributeOptionSwatch{}
	return &this
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOptionSwatch) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOptionSwatch) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOptionSwatch) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *EntitymanagerAttributeOptionSwatch) SetGrn(v string) {
	o.Grn = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *EntitymanagerAttributeOptionSwatch) GetHex() string {
	if o == nil || IsNil(o.Hex) {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerAttributeOptionSwatch) GetHexOk() (*string, bool) {
	if o == nil || IsNil(o.Hex) {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *EntitymanagerAttributeOptionSwatch) HasHex() bool {
	if o != nil && !IsNil(o.Hex) {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *EntitymanagerAttributeOptionSwatch) SetHex(v string) {
	o.Hex = &v
}

func (o EntitymanagerAttributeOptionSwatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerAttributeOptionSwatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Hex) {
		toSerialize["hex"] = o.Hex
	}
	return toSerialize, nil
}

type NullableEntitymanagerAttributeOptionSwatch struct {
	value *EntitymanagerAttributeOptionSwatch
	isSet bool
}

func (v NullableEntitymanagerAttributeOptionSwatch) Get() *EntitymanagerAttributeOptionSwatch {
	return v.value
}

func (v *NullableEntitymanagerAttributeOptionSwatch) Set(val *EntitymanagerAttributeOptionSwatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerAttributeOptionSwatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerAttributeOptionSwatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerAttributeOptionSwatch(val *EntitymanagerAttributeOptionSwatch) *NullableEntitymanagerAttributeOptionSwatch {
	return &NullableEntitymanagerAttributeOptionSwatch{value: val, isSet: true}
}

func (v NullableEntitymanagerAttributeOptionSwatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerAttributeOptionSwatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


