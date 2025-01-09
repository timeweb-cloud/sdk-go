# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoneySource** | **string** | Тип создаваемой заявки. | 
**PersonId** | Pointer to **float32** | ID администратора, на которого зарегистрирован домен. | [optional] 
**PaymentType** | **string** | Тип платежной системы. | 
**PayerId** | **float32** | Идентификационный номер плательщика | 

## Methods

### NewInvoice

`func NewInvoice(moneySource string, paymentType string, payerId float32, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoneySource

`func (o *Invoice) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *Invoice) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *Invoice) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### GetPersonId

`func (o *Invoice) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Invoice) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Invoice) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Invoice) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPaymentType

`func (o *Invoice) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *Invoice) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *Invoice) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPayerId

`func (o *Invoice) GetPayerId() float32`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *Invoice) GetPayerIdOk() (*float32, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *Invoice) SetPayerId(v float32)`

SetPayerId sets PayerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


