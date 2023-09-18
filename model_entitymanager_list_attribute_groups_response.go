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

// EntitymanagerListAttributeGroupsResponse struct for EntitymanagerListAttributeGroupsResponse
type EntitymanagerListAttributeGroupsResponse struct {
	AttributeGroups []EntitymanagerAttributeGroup `json:"attributeGroups,omitempty"`
}

// NewEntitymanagerListAttributeGroupsResponse instantiates a new EntitymanagerListAttributeGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListAttributeGroupsResponse() *EntitymanagerListAttributeGroupsResponse {
	this := EntitymanagerListAttributeGroupsResponse{}
	return &this
}

// NewEntitymanagerListAttributeGroupsResponseWithDefaults instantiates a new EntitymanagerListAttributeGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListAttributeGroupsResponseWithDefaults() *EntitymanagerListAttributeGroupsResponse {
	this := EntitymanagerListAttributeGroupsResponse{}
	return &this
}

// GetAttributeGroups returns the AttributeGroups field value if set, zero value otherwise.
func (o *EntitymanagerListAttributeGroupsResponse) GetAttributeGroups() []EntitymanagerAttributeGroup {
	if o == nil || isNil(o.AttributeGroups) {
		var ret []EntitymanagerAttributeGroup
		return ret
	}
	return o.AttributeGroups
}

// GetAttributeGroupsOk returns a tuple with the AttributeGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListAttributeGroupsResponse) GetAttributeGroupsOk() ([]EntitymanagerAttributeGroup, bool) {
	if o == nil || isNil(o.AttributeGroups) {
    return nil, false
	}
	return o.AttributeGroups, true
}

// HasAttributeGroups returns a boolean if a field has been set.
func (o *EntitymanagerListAttributeGroupsResponse) HasAttributeGroups() bool {
	if o != nil && !isNil(o.AttributeGroups) {
		return true
	}

	return false
}

// SetAttributeGroups gets a reference to the given []EntitymanagerAttributeGroup and assigns it to the AttributeGroups field.
func (o *EntitymanagerListAttributeGroupsResponse) SetAttributeGroups(v []EntitymanagerAttributeGroup) {
	o.AttributeGroups = v
}

func (o EntitymanagerListAttributeGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AttributeGroups) {
		toSerialize["attributeGroups"] = o.AttributeGroups
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerListAttributeGroupsResponse struct {
	value *EntitymanagerListAttributeGroupsResponse
	isSet bool
}

func (v NullableEntitymanagerListAttributeGroupsResponse) Get() *EntitymanagerListAttributeGroupsResponse {
	return v.value
}

func (v *NullableEntitymanagerListAttributeGroupsResponse) Set(val *EntitymanagerListAttributeGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListAttributeGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListAttributeGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListAttributeGroupsResponse(val *EntitymanagerListAttributeGroupsResponse) *NullableEntitymanagerListAttributeGroupsResponse {
	return &NullableEntitymanagerListAttributeGroupsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerListAttributeGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListAttributeGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


