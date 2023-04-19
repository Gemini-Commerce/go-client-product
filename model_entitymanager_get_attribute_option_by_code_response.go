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

// EntitymanagerGetAttributeOptionByCodeResponse struct for EntitymanagerGetAttributeOptionByCodeResponse
type EntitymanagerGetAttributeOptionByCodeResponse struct {
	Option *EntitymanagerAttributeOption `json:"option,omitempty"`
}

// NewEntitymanagerGetAttributeOptionByCodeResponse instantiates a new EntitymanagerGetAttributeOptionByCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerGetAttributeOptionByCodeResponse() *EntitymanagerGetAttributeOptionByCodeResponse {
	this := EntitymanagerGetAttributeOptionByCodeResponse{}
	return &this
}

// NewEntitymanagerGetAttributeOptionByCodeResponseWithDefaults instantiates a new EntitymanagerGetAttributeOptionByCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerGetAttributeOptionByCodeResponseWithDefaults() *EntitymanagerGetAttributeOptionByCodeResponse {
	this := EntitymanagerGetAttributeOptionByCodeResponse{}
	return &this
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *EntitymanagerGetAttributeOptionByCodeResponse) GetOption() EntitymanagerAttributeOption {
	if o == nil || isNil(o.Option) {
		var ret EntitymanagerAttributeOption
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerGetAttributeOptionByCodeResponse) GetOptionOk() (*EntitymanagerAttributeOption, bool) {
	if o == nil || isNil(o.Option) {
    return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *EntitymanagerGetAttributeOptionByCodeResponse) HasOption() bool {
	if o != nil && !isNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given EntitymanagerAttributeOption and assigns it to the Option field.
func (o *EntitymanagerGetAttributeOptionByCodeResponse) SetOption(v EntitymanagerAttributeOption) {
	o.Option = &v
}

func (o EntitymanagerGetAttributeOptionByCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerGetAttributeOptionByCodeResponse struct {
	value *EntitymanagerGetAttributeOptionByCodeResponse
	isSet bool
}

func (v NullableEntitymanagerGetAttributeOptionByCodeResponse) Get() *EntitymanagerGetAttributeOptionByCodeResponse {
	return v.value
}

func (v *NullableEntitymanagerGetAttributeOptionByCodeResponse) Set(val *EntitymanagerGetAttributeOptionByCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerGetAttributeOptionByCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerGetAttributeOptionByCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerGetAttributeOptionByCodeResponse(val *EntitymanagerGetAttributeOptionByCodeResponse) *NullableEntitymanagerGetAttributeOptionByCodeResponse {
	return &NullableEntitymanagerGetAttributeOptionByCodeResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerGetAttributeOptionByCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerGetAttributeOptionByCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


