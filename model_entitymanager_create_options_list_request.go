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

// checks if the EntitymanagerCreateOptionsListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerCreateOptionsListRequest{}

// EntitymanagerCreateOptionsListRequest struct for EntitymanagerCreateOptionsListRequest
type EntitymanagerCreateOptionsListRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OptionList *EntitymanagerOptionsList `json:"optionList,omitempty"`
}

// NewEntitymanagerCreateOptionsListRequest instantiates a new EntitymanagerCreateOptionsListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateOptionsListRequest() *EntitymanagerCreateOptionsListRequest {
	this := EntitymanagerCreateOptionsListRequest{}
	return &this
}

// NewEntitymanagerCreateOptionsListRequestWithDefaults instantiates a new EntitymanagerCreateOptionsListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateOptionsListRequestWithDefaults() *EntitymanagerCreateOptionsListRequest {
	this := EntitymanagerCreateOptionsListRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerCreateOptionsListRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateOptionsListRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerCreateOptionsListRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerCreateOptionsListRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOptionList returns the OptionList field value if set, zero value otherwise.
func (o *EntitymanagerCreateOptionsListRequest) GetOptionList() EntitymanagerOptionsList {
	if o == nil || IsNil(o.OptionList) {
		var ret EntitymanagerOptionsList
		return ret
	}
	return *o.OptionList
}

// GetOptionListOk returns a tuple with the OptionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateOptionsListRequest) GetOptionListOk() (*EntitymanagerOptionsList, bool) {
	if o == nil || IsNil(o.OptionList) {
		return nil, false
	}
	return o.OptionList, true
}

// HasOptionList returns a boolean if a field has been set.
func (o *EntitymanagerCreateOptionsListRequest) HasOptionList() bool {
	if o != nil && !IsNil(o.OptionList) {
		return true
	}

	return false
}

// SetOptionList gets a reference to the given EntitymanagerOptionsList and assigns it to the OptionList field.
func (o *EntitymanagerCreateOptionsListRequest) SetOptionList(v EntitymanagerOptionsList) {
	o.OptionList = &v
}

func (o EntitymanagerCreateOptionsListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerCreateOptionsListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.OptionList) {
		toSerialize["optionList"] = o.OptionList
	}
	return toSerialize, nil
}

type NullableEntitymanagerCreateOptionsListRequest struct {
	value *EntitymanagerCreateOptionsListRequest
	isSet bool
}

func (v NullableEntitymanagerCreateOptionsListRequest) Get() *EntitymanagerCreateOptionsListRequest {
	return v.value
}

func (v *NullableEntitymanagerCreateOptionsListRequest) Set(val *EntitymanagerCreateOptionsListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateOptionsListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateOptionsListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateOptionsListRequest(val *EntitymanagerCreateOptionsListRequest) *NullableEntitymanagerCreateOptionsListRequest {
	return &NullableEntitymanagerCreateOptionsListRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateOptionsListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateOptionsListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


