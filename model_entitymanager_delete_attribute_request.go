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

// checks if the EntitymanagerDeleteAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitymanagerDeleteAttributeRequest{}

// EntitymanagerDeleteAttributeRequest struct for EntitymanagerDeleteAttributeRequest
type EntitymanagerDeleteAttributeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	EntityData *EntitymanagerEntityIdentifier `json:"entityData,omitempty"`
	EntityId *string `json:"entityId,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewEntitymanagerDeleteAttributeRequest instantiates a new EntitymanagerDeleteAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerDeleteAttributeRequest() *EntitymanagerDeleteAttributeRequest {
	this := EntitymanagerDeleteAttributeRequest{}
	return &this
}

// NewEntitymanagerDeleteAttributeRequestWithDefaults instantiates a new EntitymanagerDeleteAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerDeleteAttributeRequestWithDefaults() *EntitymanagerDeleteAttributeRequest {
	this := EntitymanagerDeleteAttributeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerDeleteAttributeRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerDeleteAttributeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerDeleteAttributeRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerDeleteAttributeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *EntitymanagerDeleteAttributeRequest) GetEntityData() EntitymanagerEntityIdentifier {
	if o == nil || IsNil(o.EntityData) {
		var ret EntitymanagerEntityIdentifier
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerDeleteAttributeRequest) GetEntityDataOk() (*EntitymanagerEntityIdentifier, bool) {
	if o == nil || IsNil(o.EntityData) {
		return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *EntitymanagerDeleteAttributeRequest) HasEntityData() bool {
	if o != nil && !IsNil(o.EntityData) {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given EntitymanagerEntityIdentifier and assigns it to the EntityData field.
func (o *EntitymanagerDeleteAttributeRequest) SetEntityData(v EntitymanagerEntityIdentifier) {
	o.EntityData = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntitymanagerDeleteAttributeRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerDeleteAttributeRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntitymanagerDeleteAttributeRequest) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntitymanagerDeleteAttributeRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EntitymanagerDeleteAttributeRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerDeleteAttributeRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EntitymanagerDeleteAttributeRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EntitymanagerDeleteAttributeRequest) SetCode(v string) {
	o.Code = &v
}

func (o EntitymanagerDeleteAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitymanagerDeleteAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.EntityData) {
		toSerialize["entityData"] = o.EntityData
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableEntitymanagerDeleteAttributeRequest struct {
	value *EntitymanagerDeleteAttributeRequest
	isSet bool
}

func (v NullableEntitymanagerDeleteAttributeRequest) Get() *EntitymanagerDeleteAttributeRequest {
	return v.value
}

func (v *NullableEntitymanagerDeleteAttributeRequest) Set(val *EntitymanagerDeleteAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerDeleteAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerDeleteAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerDeleteAttributeRequest(val *EntitymanagerDeleteAttributeRequest) *NullableEntitymanagerDeleteAttributeRequest {
	return &NullableEntitymanagerDeleteAttributeRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerDeleteAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerDeleteAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


