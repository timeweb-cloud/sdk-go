# CreatePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Сумма оплаты | 
**PaymentType** | [**PaymentType**](PaymentType.md) |  | 
**IsBindCard** | Pointer to **bool** | Привязать карту | [optional] 
**ReturnUrl** | Pointer to **string** | URL для перенаправления после успешной оплаты | [optional] 
**FailUrl** | Pointer to **string** | URL для перенаправления после неудачной оплаты | [optional] 

## Methods

### NewCreatePayment

`func NewCreatePayment(amount float32, paymentType PaymentType, ) *CreatePayment`

NewCreatePayment instantiates a new CreatePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentWithDefaults

`func NewCreatePaymentWithDefaults() *CreatePayment`

NewCreatePaymentWithDefaults instantiates a new CreatePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreatePayment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePayment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePayment) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetPaymentType

`func (o *CreatePayment) GetPaymentType() PaymentType`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CreatePayment) GetPaymentTypeOk() (*PaymentType, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CreatePayment) SetPaymentType(v PaymentType)`

SetPaymentType sets PaymentType field to given value.


### GetIsBindCard

`func (o *CreatePayment) GetIsBindCard() bool`

GetIsBindCard returns the IsBindCard field if non-nil, zero value otherwise.

### GetIsBindCardOk

`func (o *CreatePayment) GetIsBindCardOk() (*bool, bool)`

GetIsBindCardOk returns a tuple with the IsBindCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBindCard

`func (o *CreatePayment) SetIsBindCard(v bool)`

SetIsBindCard sets IsBindCard field to given value.

### HasIsBindCard

`func (o *CreatePayment) HasIsBindCard() bool`

HasIsBindCard returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreatePayment) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreatePayment) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreatePayment) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreatePayment) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetFailUrl

`func (o *CreatePayment) GetFailUrl() string`

GetFailUrl returns the FailUrl field if non-nil, zero value otherwise.

### GetFailUrlOk

`func (o *CreatePayment) GetFailUrlOk() (*string, bool)`

GetFailUrlOk returns a tuple with the FailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailUrl

`func (o *CreatePayment) SetFailUrl(v string)`

SetFailUrl sets FailUrl field to given value.

### HasFailUrl

`func (o *CreatePayment) HasFailUrl() bool`

HasFailUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


