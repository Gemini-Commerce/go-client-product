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

// checks if the EntitymanagerCreateOptionsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerCreateOptionsListResponse{}

// EntitymanagerCreateOptionsListResponse struct for EntitymanagerCreateOptionsListResponse
type EntitymanagerCreateOptionsListResponse struct {
	OptionList *EntitymanagerOptionsList `json:"optionList,omitempty"`
}

// NewEntitymanagerCreateOptionsListResponse instantiates a new EntitymanagerCreateOptionsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateOptionsListResponse() *EntitymanagerCreateOptionsListResponse {
	this := EntitymanagerCreateOptionsListResponse{}
	return &this
}

// NewEntitymanagerCreateOptionsListResponseWithDefaults instantiates a new EntitymanagerCreateOptionsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateOptionsListResponseWithDefaults() *EntitymanagerCreateOptionsListResponse {
	this := EntitymanagerCreateOptionsListResponse{}
	return &this
}

// GetOptionList returns the OptionList field value if set, zero value otherwise.
func (o *EntitymanagerCreateOptionsListResponse) GetOptionList() EntitymanagerOptionsList {
	if o == nil || IsNil(o.OptionList) {
		var ret EntitymanagerOptionsList
		return ret
	}
	return *o.OptionList
}

// GetOptionListOk returns a tuple with the OptionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateOptionsListResponse) GetOptionListOk() (*EntitymanagerOptionsList, bool) {
	if o == nil || IsNil(o.OptionList) {
		return nil, false
	}
	return o.OptionList, true
}

// HasOptionList returns a boolean if a field has been set.
func (o *EntitymanagerCreateOptionsListResponse) HasOptionList() bool {
	if o != nil && !IsNil(o.OptionList) {
		return true
	}

	return false
}

// SetOptionList gets a reference to the given EntitymanagerOptionsList and assigns it to the OptionList field.
func (o *EntitymanagerCreateOptionsListResponse) SetOptionList(v EntitymanagerOptionsList) {
	o.OptionList = &v
}

func (o EntitymanagerCreateOptionsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerCreateOptionsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionList) {
		toSerialize["optionList"] = o.OptionList
	}
	return toSerialize, nil
}

type NullableEntitymanagerCreateOptionsListResponse struct {
	value *EntitymanagerCreateOptionsListResponse
	isSet bool
}

func (v NullableEntitymanagerCreateOptionsListResponse) Get() *EntitymanagerCreateOptionsListResponse {
	return v.value
}

func (v *NullableEntitymanagerCreateOptionsListResponse) Set(val *EntitymanagerCreateOptionsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateOptionsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateOptionsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateOptionsListResponse(val *EntitymanagerCreateOptionsListResponse) *NullableEntitymanagerCreateOptionsListResponse {
	return &NullableEntitymanagerCreateOptionsListResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateOptionsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateOptionsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


