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

// checks if the ProductMediaGalleryEntryMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMediaGalleryEntryMetadata{}

// ProductMediaGalleryEntryMetadata struct for ProductMediaGalleryEntryMetadata
type ProductMediaGalleryEntryMetadata struct {
	Roles []string `json:"roles,omitempty"`
	Alt *ProductLocalizedText `json:"alt,omitempty"`
}

// NewProductMediaGalleryEntryMetadata instantiates a new ProductMediaGalleryEntryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMediaGalleryEntryMetadata() *ProductMediaGalleryEntryMetadata {
	this := ProductMediaGalleryEntryMetadata{}
	return &this
}

// NewProductMediaGalleryEntryMetadataWithDefaults instantiates a new ProductMediaGalleryEntryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMediaGalleryEntryMetadataWithDefaults() *ProductMediaGalleryEntryMetadata {
	this := ProductMediaGalleryEntryMetadata{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ProductMediaGalleryEntryMetadata) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMediaGalleryEntryMetadata) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ProductMediaGalleryEntryMetadata) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ProductMediaGalleryEntryMetadata) SetRoles(v []string) {
	o.Roles = v
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *ProductMediaGalleryEntryMetadata) GetAlt() ProductLocalizedText {
	if o == nil || IsNil(o.Alt) {
		var ret ProductLocalizedText
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMediaGalleryEntryMetadata) GetAltOk() (*ProductLocalizedText, bool) {
	if o == nil || IsNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *ProductMediaGalleryEntryMetadata) HasAlt() bool {
	if o != nil && !IsNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given ProductLocalizedText and assigns it to the Alt field.
func (o *ProductMediaGalleryEntryMetadata) SetAlt(v ProductLocalizedText) {
	o.Alt = &v
}

func (o ProductMediaGalleryEntryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMediaGalleryEntryMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	return toSerialize, nil
}

type NullableProductMediaGalleryEntryMetadata struct {
	value *ProductMediaGalleryEntryMetadata
	isSet bool
}

func (v NullableProductMediaGalleryEntryMetadata) Get() *ProductMediaGalleryEntryMetadata {
	return v.value
}

func (v *NullableProductMediaGalleryEntryMetadata) Set(val *ProductMediaGalleryEntryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMediaGalleryEntryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMediaGalleryEntryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMediaGalleryEntryMetadata(val *ProductMediaGalleryEntryMetadata) *NullableProductMediaGalleryEntryMetadata {
	return &NullableProductMediaGalleryEntryMetadata{value: val, isSet: true}
}

func (v NullableProductMediaGalleryEntryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMediaGalleryEntryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


