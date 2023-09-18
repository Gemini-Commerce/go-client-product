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

// EntitymanagerCreateAttributeGroupRequest struct for EntitymanagerCreateAttributeGroupRequest
type EntitymanagerCreateAttributeGroupRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
	Label *ProductentitymanagerLocalizedText `json:"label,omitempty"`
	Sort *int32 `json:"sort,omitempty"`
}

// NewEntitymanagerCreateAttributeGroupRequest instantiates a new EntitymanagerCreateAttributeGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateAttributeGroupRequest() *EntitymanagerCreateAttributeGroupRequest {
	this := EntitymanagerCreateAttributeGroupRequest{}
	return &this
}

// NewEntitymanagerCreateAttributeGroupRequestWithDefaults instantiates a new EntitymanagerCreateAttributeGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateAttributeGroupRequestWithDefaults() *EntitymanagerCreateAttributeGroupRequest {
	this := EntitymanagerCreateAttributeGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeGroupRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerCreateAttributeGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeGroupRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerCreateAttributeGroupRequest) SetCode(v string) {
	o.Code = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeGroupRequest) GetLabel() ProductentitymanagerLocalizedText {
	if o == nil || isNil(o.Label) {
		var ret ProductentitymanagerLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) GetLabelOk() (*ProductentitymanagerLocalizedText, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ProductentitymanagerLocalizedText and assigns it to the Label field.
func (o *EntitymanagerCreateAttributeGroupRequest) SetLabel(v ProductentitymanagerLocalizedText) {
	o.Label = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *EntitymanagerCreateAttributeGroupRequest) GetSort() int32 {
	if o == nil || isNil(o.Sort) {
		var ret int32
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) GetSortOk() (*int32, bool) {
	if o == nil || isNil(o.Sort) {
    return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *EntitymanagerCreateAttributeGroupRequest) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int32 and assigns it to the Sort field.
func (o *EntitymanagerCreateAttributeGroupRequest) SetSort(v int32) {
	o.Sort = &v
}

func (o EntitymanagerCreateAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerCreateAttributeGroupRequest struct {
	value *EntitymanagerCreateAttributeGroupRequest
	isSet bool
}

func (v NullableEntitymanagerCreateAttributeGroupRequest) Get() *EntitymanagerCreateAttributeGroupRequest {
	return v.value
}

func (v *NullableEntitymanagerCreateAttributeGroupRequest) Set(val *EntitymanagerCreateAttributeGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateAttributeGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateAttributeGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateAttributeGroupRequest(val *EntitymanagerCreateAttributeGroupRequest) *NullableEntitymanagerCreateAttributeGroupRequest {
	return &NullableEntitymanagerCreateAttributeGroupRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateAttributeGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateAttributeGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

