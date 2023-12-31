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

// checks if the ProductAddMediaGalleryEntryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddMediaGalleryEntryResponse{}

// ProductAddMediaGalleryEntryResponse struct for ProductAddMediaGalleryEntryResponse
type ProductAddMediaGalleryEntryResponse struct {
	MediaGalleryEntry *ProductMediaGalleryEntry `json:"mediaGalleryEntry,omitempty"`
}

// NewProductAddMediaGalleryEntryResponse instantiates a new ProductAddMediaGalleryEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddMediaGalleryEntryResponse() *ProductAddMediaGalleryEntryResponse {
	this := ProductAddMediaGalleryEntryResponse{}
	return &this
}

// NewProductAddMediaGalleryEntryResponseWithDefaults instantiates a new ProductAddMediaGalleryEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddMediaGalleryEntryResponseWithDefaults() *ProductAddMediaGalleryEntryResponse {
	this := ProductAddMediaGalleryEntryResponse{}
	return &this
}

// GetMediaGalleryEntry returns the MediaGalleryEntry field value if set, zero value otherwise.
func (o *ProductAddMediaGalleryEntryResponse) GetMediaGalleryEntry() ProductMediaGalleryEntry {
	if o == nil || IsNil(o.MediaGalleryEntry) {
		var ret ProductMediaGalleryEntry
		return ret
	}
	return *o.MediaGalleryEntry
}

// GetMediaGalleryEntryOk returns a tuple with the MediaGalleryEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddMediaGalleryEntryResponse) GetMediaGalleryEntryOk() (*ProductMediaGalleryEntry, bool) {
	if o == nil || IsNil(o.MediaGalleryEntry) {
		return nil, false
	}
	return o.MediaGalleryEntry, true
}

// HasMediaGalleryEntry returns a boolean if a field has been set.
func (o *ProductAddMediaGalleryEntryResponse) HasMediaGalleryEntry() bool {
	if o != nil && !IsNil(o.MediaGalleryEntry) {
		return true
	}

	return false
}

// SetMediaGalleryEntry gets a reference to the given ProductMediaGalleryEntry and assigns it to the MediaGalleryEntry field.
func (o *ProductAddMediaGalleryEntryResponse) SetMediaGalleryEntry(v ProductMediaGalleryEntry) {
	o.MediaGalleryEntry = &v
}

func (o ProductAddMediaGalleryEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddMediaGalleryEntryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaGalleryEntry) {
		toSerialize["mediaGalleryEntry"] = o.MediaGalleryEntry
	}
	return toSerialize, nil
}

type NullableProductAddMediaGalleryEntryResponse struct {
	value *ProductAddMediaGalleryEntryResponse
	isSet bool
}

func (v NullableProductAddMediaGalleryEntryResponse) Get() *ProductAddMediaGalleryEntryResponse {
	return v.value
}

func (v *NullableProductAddMediaGalleryEntryResponse) Set(val *ProductAddMediaGalleryEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddMediaGalleryEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddMediaGalleryEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddMediaGalleryEntryResponse(val *ProductAddMediaGalleryEntryResponse) *NullableProductAddMediaGalleryEntryResponse {
	return &NullableProductAddMediaGalleryEntryResponse{value: val, isSet: true}
}

func (v NullableProductAddMediaGalleryEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddMediaGalleryEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


