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

// checks if the EntitymanagerOptionsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerOptionsList{}

// EntitymanagerOptionsList struct for EntitymanagerOptionsList
type EntitymanagerOptionsList struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewEntitymanagerOptionsList instantiates a new EntitymanagerOptionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerOptionsList() *EntitymanagerOptionsList {
	this := EntitymanagerOptionsList{}
	return &this
}

// NewEntitymanagerOptionsListWithDefaults instantiates a new EntitymanagerOptionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerOptionsListWithDefaults() *EntitymanagerOptionsList {
	this := EntitymanagerOptionsList{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerOptionsList) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerOptionsList) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerOptionsList) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerOptionsList) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitymanagerOptionsList) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerOptionsList) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitymanagerOptionsList) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitymanagerOptionsList) SetName(v string) {
	o.Name = &v
}

func (o EntitymanagerOptionsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerOptionsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEntitymanagerOptionsList struct {
	value *EntitymanagerOptionsList
	isSet bool
}

func (v NullableEntitymanagerOptionsList) Get() *EntitymanagerOptionsList {
	return v.value
}

func (v *NullableEntitymanagerOptionsList) Set(val *EntitymanagerOptionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerOptionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerOptionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerOptionsList(val *EntitymanagerOptionsList) *NullableEntitymanagerOptionsList {
	return &NullableEntitymanagerOptionsList{value: val, isSet: true}
}

func (v NullableEntitymanagerOptionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerOptionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


