/*
Balance Microservice API

Balance Microservice API

API version: 1.0
Contact: nikita@kanash.in
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody{}

// InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody struct for InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody
type InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody struct {
	Amount *int32 `json:"amount,omitempty"`
	OrderId *int32 `json:"order_id,omitempty"`
	ProductId *int32 `json:"product_id,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
}

// NewInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody instantiates a new InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody() *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody {
	this := InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody{}
	return &this
}

// NewInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBodyWithDefaults instantiates a new InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBodyWithDefaults() *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody {
	this := InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) SetAmount(v int32) {
	o.Amount = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) SetProductId(v int32) {
	o.ProductId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) SetUserId(v int32) {
	o.UserId = &v
}

func (o InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody struct {
	value *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody
	isSet bool
}

func (v NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) Get() *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody {
	return v.value
}

func (v *NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) Set(val *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody(val *InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) *NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody {
	return &NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody{value: val, isSet: true}
}

func (v NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


