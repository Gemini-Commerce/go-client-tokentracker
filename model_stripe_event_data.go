/*
Token Tracker Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokentracker

import (
	"encoding/json"
)

// checks if the StripeEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeEventData{}

// StripeEventData struct for StripeEventData
type StripeEventData struct {
	Object map[string]interface{} `json:"object,omitempty"`
	PreviousAttributes map[string]interface{} `json:"previousAttributes,omitempty"`
	Raw *string `json:"raw,omitempty"`
}

// NewStripeEventData instantiates a new StripeEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeEventData() *StripeEventData {
	this := StripeEventData{}
	return &this
}

// NewStripeEventDataWithDefaults instantiates a new StripeEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeEventDataWithDefaults() *StripeEventData {
	this := StripeEventData{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *StripeEventData) GetObject() map[string]interface{} {
	if o == nil || IsNil(o.Object) {
		var ret map[string]interface{}
		return ret
	}
	return o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeEventData) GetObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return map[string]interface{}{}, false
	}
	return o.Object, true
}

// IsSetObject returns a boolean if a field has been set.
func (o *StripeEventData) IsSetObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given map[string]interface{} and assigns it to the Object field.
func (o *StripeEventData) SetObject(v map[string]interface{}) {
	o.Object = v
}

// GetPreviousAttributes returns the PreviousAttributes field value if set, zero value otherwise.
func (o *StripeEventData) GetPreviousAttributes() map[string]interface{} {
	if o == nil || IsNil(o.PreviousAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.PreviousAttributes
}

// GetPreviousAttributesOk returns a tuple with the PreviousAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeEventData) GetPreviousAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PreviousAttributes) {
		return map[string]interface{}{}, false
	}
	return o.PreviousAttributes, true
}

// IsSetPreviousAttributes returns a boolean if a field has been set.
func (o *StripeEventData) IsSetPreviousAttributes() bool {
	if o != nil && !IsNil(o.PreviousAttributes) {
		return true
	}

	return false
}

// SetPreviousAttributes gets a reference to the given map[string]interface{} and assigns it to the PreviousAttributes field.
func (o *StripeEventData) SetPreviousAttributes(v map[string]interface{}) {
	o.PreviousAttributes = v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *StripeEventData) GetRaw() string {
	if o == nil || IsNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeEventData) GetRawOk() (*string, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// IsSetRaw returns a boolean if a field has been set.
func (o *StripeEventData) IsSetRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *StripeEventData) SetRaw(v string) {
	o.Raw = &v
}

func (o StripeEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.PreviousAttributes) {
		toSerialize["previousAttributes"] = o.PreviousAttributes
	}
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableStripeEventData struct {
	value *StripeEventData
	isSet bool
}

func (v NullableStripeEventData) Get() *StripeEventData {
	return v.value
}

func (v *NullableStripeEventData) Set(val *StripeEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeEventData(val *StripeEventData) *NullableStripeEventData {
	return &NullableStripeEventData{value: val, isSet: true}
}

func (v NullableStripeEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


