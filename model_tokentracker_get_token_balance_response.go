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
	"time"
)

// checks if the TokentrackerGetTokenBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokentrackerGetTokenBalanceResponse{}

// TokentrackerGetTokenBalanceResponse struct for TokentrackerGetTokenBalanceResponse
type TokentrackerGetTokenBalanceResponse struct {
	Balance *string `json:"balance,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewTokentrackerGetTokenBalanceResponse instantiates a new TokentrackerGetTokenBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokentrackerGetTokenBalanceResponse() *TokentrackerGetTokenBalanceResponse {
	this := TokentrackerGetTokenBalanceResponse{}
	return &this
}

// NewTokentrackerGetTokenBalanceResponseWithDefaults instantiates a new TokentrackerGetTokenBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokentrackerGetTokenBalanceResponseWithDefaults() *TokentrackerGetTokenBalanceResponse {
	this := TokentrackerGetTokenBalanceResponse{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *TokentrackerGetTokenBalanceResponse) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokentrackerGetTokenBalanceResponse) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// IsSetBalance returns a boolean if a field has been set.
func (o *TokentrackerGetTokenBalanceResponse) IsSetBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *TokentrackerGetTokenBalanceResponse) SetBalance(v string) {
	o.Balance = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TokentrackerGetTokenBalanceResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokentrackerGetTokenBalanceResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// IsSetUpdatedAt returns a boolean if a field has been set.
func (o *TokentrackerGetTokenBalanceResponse) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TokentrackerGetTokenBalanceResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TokentrackerGetTokenBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokentrackerGetTokenBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTokentrackerGetTokenBalanceResponse struct {
	value *TokentrackerGetTokenBalanceResponse
	isSet bool
}

func (v NullableTokentrackerGetTokenBalanceResponse) Get() *TokentrackerGetTokenBalanceResponse {
	return v.value
}

func (v *NullableTokentrackerGetTokenBalanceResponse) Set(val *TokentrackerGetTokenBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokentrackerGetTokenBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokentrackerGetTokenBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokentrackerGetTokenBalanceResponse(val *TokentrackerGetTokenBalanceResponse) *NullableTokentrackerGetTokenBalanceResponse {
	return &NullableTokentrackerGetTokenBalanceResponse{value: val, isSet: true}
}

func (v NullableTokentrackerGetTokenBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokentrackerGetTokenBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


