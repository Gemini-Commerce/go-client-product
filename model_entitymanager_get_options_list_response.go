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

// checks if the EntitymanagerGetOptionsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerGetOptionsListResponse{}

// EntitymanagerGetOptionsListResponse struct for EntitymanagerGetOptionsListResponse
type EntitymanagerGetOptionsListResponse struct {
	OptionList *EntitymanagerOptionsList `json:"optionList,omitempty"`
}

// NewEntitymanagerGetOptionsListResponse instantiates a new EntitymanagerGetOptionsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerGetOptionsListResponse() *EntitymanagerGetOptionsListResponse {
	this := EntitymanagerGetOptionsListResponse{}
	return &this
}

// NewEntitymanagerGetOptionsListResponseWithDefaults instantiates a new EntitymanagerGetOptionsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerGetOptionsListResponseWithDefaults() *EntitymanagerGetOptionsListResponse {
	this := EntitymanagerGetOptionsListResponse{}
	return &this
}

// GetOptionList returns the OptionList field value if set, zero value otherwise.
func (o *EntitymanagerGetOptionsListResponse) GetOptionList() EntitymanagerOptionsList {
	if o == nil || IsNil(o.OptionList) {
		var ret EntitymanagerOptionsList
		return ret
	}
	return *o.OptionList
}

// GetOptionListOk returns a tuple with the OptionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerGetOptionsListResponse) GetOptionListOk() (*EntitymanagerOptionsList, bool) {
	if o == nil || IsNil(o.OptionList) {
		return nil, false
	}
	return o.OptionList, true
}

// HasOptionList returns a boolean if a field has been set.
func (o *EntitymanagerGetOptionsListResponse) HasOptionList() bool {
	if o != nil && !IsNil(o.OptionList) {
		return true
	}

	return false
}

// SetOptionList gets a reference to the given EntitymanagerOptionsList and assigns it to the OptionList field.
func (o *EntitymanagerGetOptionsListResponse) SetOptionList(v EntitymanagerOptionsList) {
	o.OptionList = &v
}

func (o EntitymanagerGetOptionsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerGetOptionsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionList) {
		toSerialize["optionList"] = o.OptionList
	}
	return toSerialize, nil
}

type NullableEntitymanagerGetOptionsListResponse struct {
	value *EntitymanagerGetOptionsListResponse
	isSet bool
}

func (v NullableEntitymanagerGetOptionsListResponse) Get() *EntitymanagerGetOptionsListResponse {
	return v.value
}

func (v *NullableEntitymanagerGetOptionsListResponse) Set(val *EntitymanagerGetOptionsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerGetOptionsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerGetOptionsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerGetOptionsListResponse(val *EntitymanagerGetOptionsListResponse) *NullableEntitymanagerGetOptionsListResponse {
	return &NullableEntitymanagerGetOptionsListResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerGetOptionsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerGetOptionsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


