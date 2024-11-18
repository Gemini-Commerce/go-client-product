/*
Product Service

Introducing our revolutionary Product Management Service! Designed to streamline your product inventory and elevate customer experiences, our cutting-edge protobuf service is a game-changer in the world of efficient product management.  With our service, you can effortlessly create new products, allowing you to quickly bring your ideas to life and expand your catalog. Retrieve product information in a snap, providing accurate and personalized details to your customers based on their specific needs and preferences.  Stay ahead of the competition by easily updating product information, ensuring your catalog is always up-to-date and optimized. Seamlessly remove products from your inventory when needed, maintaining a clean and relevant product selection.  Enhance the visual appeal of your products with our advanced media gallery functionalities. Effortlessly add and update captivating images and videos to showcase your products in the best possible light, engaging your customers and driving conversions.  Personalization is key in today's market, and our service enables you to offer unique options to your customers. Easily create and manage lists of customizable options for your products, providing flexibility and tailoring to individual preferences.  Attributes play a vital role in defining products, and our service empowers you to effectively manage them. From bulk attribute creation to listing and retrieving attribute options, our service ensures your product information is rich and comprehensive.  Our service extends its capabilities to entity management, allowing you to effortlessly handle different entities and create customized options lists associated with them. This provides further flexibility and customization options for your product offerings.  When it comes to bulk updates, our service has you covered. Effortlessly update multiple products simultaneously, saving you time and streamlining your operations.  Finding specific products and variants is a breeze with our service. Quickly locate products based on their unique stock keeping unit (SKU) values, ensuring efficient inventory management and smooth order fulfillment.  Experience a new level of efficiency and productivity with our Product Management Service. Unlock the full potential of streamlined product management and empower your business to thrive in today's competitive market. Try our service today and elevate your product management to new heights!

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product

import (
	"encoding/json"
	"time"
)

// checks if the ProductAttributeInReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttributeInReview{}

// ProductAttributeInReview struct for ProductAttributeInReview
type ProductAttributeInReview struct {
	Code                 *string                          `json:"code,omitempty"`
	String               *AttributeInReviewString         `json:"string,omitempty"`
	Int32                *int32                           `json:"int32,omitempty"`
	Int64                *string                          `json:"int64,omitempty"`
	Float32              *float32                         `json:"float32,omitempty"`
	Float64              *float64                         `json:"float64,omitempty"`
	Boolean              *bool                            `json:"boolean,omitempty"`
	Source               *ProductAttributeInReviewSource  `json:"source,omitempty"`
	CreatedAt            *time.Time                       `json:"createdAt,omitempty"`
	JobId                *string                          `json:"jobId,omitempty"`
	JobType              *ProductAttributeInReviewJobType `json:"jobType,omitempty"`
	Error                *ProductAttributeInReviewError   `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAttributeInReview ProductAttributeInReview

// NewProductAttributeInReview instantiates a new ProductAttributeInReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttributeInReview() *ProductAttributeInReview {
	this := ProductAttributeInReview{}
	var source ProductAttributeInReviewSource = PRODUCTATTRIBUTEINREVIEWSOURCE_UNKNOWN
	this.Source = &source
	var jobType ProductAttributeInReviewJobType = PRODUCTATTRIBUTEINREVIEWJOBTYPE_UNKNOWN
	this.JobType = &jobType
	return &this
}

// NewProductAttributeInReviewWithDefaults instantiates a new ProductAttributeInReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeInReviewWithDefaults() *ProductAttributeInReview {
	this := ProductAttributeInReview{}
	var source ProductAttributeInReviewSource = PRODUCTATTRIBUTEINREVIEWSOURCE_UNKNOWN
	this.Source = &source
	var jobType ProductAttributeInReviewJobType = PRODUCTATTRIBUTEINREVIEWJOBTYPE_UNKNOWN
	this.JobType = &jobType
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductAttributeInReview) SetCode(v string) {
	o.Code = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetString() AttributeInReviewString {
	if o == nil || IsNil(o.String) {
		var ret AttributeInReviewString
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetStringOk() (*AttributeInReviewString, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given AttributeInReviewString and assigns it to the String field.
func (o *ProductAttributeInReview) SetString(v AttributeInReviewString) {
	o.String = &v
}

// GetInt32 returns the Int32 field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetInt32() int32 {
	if o == nil || IsNil(o.Int32) {
		var ret int32
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetInt32Ok() (*int32, bool) {
	if o == nil || IsNil(o.Int32) {
		return nil, false
	}
	return o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasInt32() bool {
	if o != nil && !IsNil(o.Int32) {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *ProductAttributeInReview) SetInt32(v int32) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetInt64() string {
	if o == nil || IsNil(o.Int64) {
		var ret string
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetInt64Ok() (*string, bool) {
	if o == nil || IsNil(o.Int64) {
		return nil, false
	}
	return o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasInt64() bool {
	if o != nil && !IsNil(o.Int64) {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given string and assigns it to the Int64 field.
func (o *ProductAttributeInReview) SetInt64(v string) {
	o.Int64 = &v
}

// GetFloat32 returns the Float32 field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetFloat32() float32 {
	if o == nil || IsNil(o.Float32) {
		var ret float32
		return ret
	}
	return *o.Float32
}

// GetFloat32Ok returns a tuple with the Float32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetFloat32Ok() (*float32, bool) {
	if o == nil || IsNil(o.Float32) {
		return nil, false
	}
	return o.Float32, true
}

// HasFloat32 returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasFloat32() bool {
	if o != nil && !IsNil(o.Float32) {
		return true
	}

	return false
}

// SetFloat32 gets a reference to the given float32 and assigns it to the Float32 field.
func (o *ProductAttributeInReview) SetFloat32(v float32) {
	o.Float32 = &v
}

// GetFloat64 returns the Float64 field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetFloat64() float64 {
	if o == nil || IsNil(o.Float64) {
		var ret float64
		return ret
	}
	return *o.Float64
}

// GetFloat64Ok returns a tuple with the Float64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetFloat64Ok() (*float64, bool) {
	if o == nil || IsNil(o.Float64) {
		return nil, false
	}
	return o.Float64, true
}

// HasFloat64 returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasFloat64() bool {
	if o != nil && !IsNil(o.Float64) {
		return true
	}

	return false
}

// SetFloat64 gets a reference to the given float64 and assigns it to the Float64 field.
func (o *ProductAttributeInReview) SetFloat64(v float64) {
	o.Float64 = &v
}

// GetBoolean returns the Boolean field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetBoolean() bool {
	if o == nil || IsNil(o.Boolean) {
		var ret bool
		return ret
	}
	return *o.Boolean
}

// GetBooleanOk returns a tuple with the Boolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.Boolean) {
		return nil, false
	}
	return o.Boolean, true
}

// HasBoolean returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasBoolean() bool {
	if o != nil && !IsNil(o.Boolean) {
		return true
	}

	return false
}

// SetBoolean gets a reference to the given bool and assigns it to the Boolean field.
func (o *ProductAttributeInReview) SetBoolean(v bool) {
	o.Boolean = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetSource() ProductAttributeInReviewSource {
	if o == nil || IsNil(o.Source) {
		var ret ProductAttributeInReviewSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetSourceOk() (*ProductAttributeInReviewSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ProductAttributeInReviewSource and assigns it to the Source field.
func (o *ProductAttributeInReview) SetSource(v ProductAttributeInReviewSource) {
	o.Source = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProductAttributeInReview) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *ProductAttributeInReview) SetJobId(v string) {
	o.JobId = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetJobType() ProductAttributeInReviewJobType {
	if o == nil || IsNil(o.JobType) {
		var ret ProductAttributeInReviewJobType
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetJobTypeOk() (*ProductAttributeInReviewJobType, bool) {
	if o == nil || IsNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasJobType() bool {
	if o != nil && !IsNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given ProductAttributeInReviewJobType and assigns it to the JobType field.
func (o *ProductAttributeInReview) SetJobType(v ProductAttributeInReviewJobType) {
	o.JobType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ProductAttributeInReview) GetError() ProductAttributeInReviewError {
	if o == nil || IsNil(o.Error) {
		var ret ProductAttributeInReviewError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeInReview) GetErrorOk() (*ProductAttributeInReviewError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ProductAttributeInReview) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ProductAttributeInReviewError and assigns it to the Error field.
func (o *ProductAttributeInReview) SetError(v ProductAttributeInReviewError) {
	o.Error = &v
}

func (o ProductAttributeInReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAttributeInReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !IsNil(o.Int32) {
		toSerialize["int32"] = o.Int32
	}
	if !IsNil(o.Int64) {
		toSerialize["int64"] = o.Int64
	}
	if !IsNil(o.Float32) {
		toSerialize["float32"] = o.Float32
	}
	if !IsNil(o.Float64) {
		toSerialize["float64"] = o.Float64
	}
	if !IsNil(o.Boolean) {
		toSerialize["boolean"] = o.Boolean
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !IsNil(o.JobType) {
		toSerialize["jobType"] = o.JobType
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductAttributeInReview) UnmarshalJSON(data []byte) (err error) {
	varProductAttributeInReview := _ProductAttributeInReview{}

	err = json.Unmarshal(data, &varProductAttributeInReview)

	if err != nil {
		return err
	}

	*o = ProductAttributeInReview(varProductAttributeInReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "string")
		delete(additionalProperties, "int32")
		delete(additionalProperties, "int64")
		delete(additionalProperties, "float32")
		delete(additionalProperties, "float64")
		delete(additionalProperties, "boolean")
		delete(additionalProperties, "source")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "jobType")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductAttributeInReview) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populate the value of well-known types
func (o *ProductAttributeInReview) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductAttributeInReview struct {
	value *ProductAttributeInReview
	isSet bool
}

func (v NullableProductAttributeInReview) Get() *ProductAttributeInReview {
	return v.value
}

func (v *NullableProductAttributeInReview) Set(val *ProductAttributeInReview) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeInReview) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeInReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeInReview(val *ProductAttributeInReview) *NullableProductAttributeInReview {
	return &NullableProductAttributeInReview{value: val, isSet: true}
}

func (v NullableProductAttributeInReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeInReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
