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

// EntitymanagerListOptionsListsResponse struct for EntitymanagerListOptionsListsResponse
type EntitymanagerListOptionsListsResponse struct {
	Options []EntitymanagerOptionsList `json:"options,omitempty"`
}

// NewEntitymanagerListOptionsListsResponse instantiates a new EntitymanagerListOptionsListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListOptionsListsResponse() *EntitymanagerListOptionsListsResponse {
	this := EntitymanagerListOptionsListsResponse{}
	return &this
}

// NewEntitymanagerListOptionsListsResponseWithDefaults instantiates a new EntitymanagerListOptionsListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListOptionsListsResponseWithDefaults() *EntitymanagerListOptionsListsResponse {
	this := EntitymanagerListOptionsListsResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerListOptionsListsResponse) GetOptions() []EntitymanagerOptionsList {
	if o == nil || isNil(o.Options) {
		var ret []EntitymanagerOptionsList
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListOptionsListsResponse) GetOptionsOk() ([]EntitymanagerOptionsList, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerListOptionsListsResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerOptionsList and assigns it to the Options field.
func (o *EntitymanagerListOptionsListsResponse) SetOptions(v []EntitymanagerOptionsList) {
	o.Options = v
}

func (o EntitymanagerListOptionsListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerListOptionsListsResponse struct {
	value *EntitymanagerListOptionsListsResponse
	isSet bool
}

func (v NullableEntitymanagerListOptionsListsResponse) Get() *EntitymanagerListOptionsListsResponse {
	return v.value
}

func (v *NullableEntitymanagerListOptionsListsResponse) Set(val *EntitymanagerListOptionsListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListOptionsListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListOptionsListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListOptionsListsResponse(val *EntitymanagerListOptionsListsResponse) *NullableEntitymanagerListOptionsListsResponse {
	return &NullableEntitymanagerListOptionsListsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerListOptionsListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListOptionsListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


