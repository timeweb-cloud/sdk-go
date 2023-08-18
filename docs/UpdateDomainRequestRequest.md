# UpdateDomainRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoneySource** | **string** | Тип создаваемой заявки. | 
**PersonId** | Pointer to **float32** | Идентификатор администратора, на которого зарегистрирован домен. | [optional] 
**PaymentType** | **string** | Тип платежной системы. | 
**PayerId** | **float32** | Идентификационный номер плательщика | 
**AuthCode** | **string** | Код авторизации для переноса домена. | 
**BonusId** | **float32** | Идентификатор бонуса. | 

## Methods

### NewUpdateDomainRequestRequest

`func NewUpdateDomainRequestRequest(moneySource string, paymentType string, payerId float32, authCode string, bonusId float32, ) *UpdateDomainRequestRequest`

NewUpdateDomainRequestRequest instantiates a new UpdateDomainRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainRequestRequestWithDefaults

`func NewUpdateDomainRequestRequestWithDefaults() *UpdateDomainRequestRequest`

NewUpdateDomainRequestRequestWithDefaults instantiates a new UpdateDomainRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoneySource

`func (o *UpdateDomainRequestRequest) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *UpdateDomainRequestRequest) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *UpdateDomainRequestRequest) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### GetPersonId

`func (o *UpdateDomainRequestRequest) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *UpdateDomainRequestRequest) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *UpdateDomainRequestRequest) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *UpdateDomainRequestRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPaymentType

`func (o *UpdateDomainRequestRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *UpdateDomainRequestRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *UpdateDomainRequestRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPayerId

`func (o *UpdateDomainRequestRequest) GetPayerId() float32`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *UpdateDomainRequestRequest) GetPayerIdOk() (*float32, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *UpdateDomainRequestRequest) SetPayerId(v float32)`

SetPayerId sets PayerId field to given value.


### GetAuthCode

`func (o *UpdateDomainRequestRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *UpdateDomainRequestRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *UpdateDomainRequestRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetBonusId

`func (o *UpdateDomainRequestRequest) GetBonusId() float32`

GetBonusId returns the BonusId field if non-nil, zero value otherwise.

### GetBonusIdOk

`func (o *UpdateDomainRequestRequest) GetBonusIdOk() (*float32, bool)`

GetBonusIdOk returns a tuple with the BonusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusId

`func (o *UpdateDomainRequestRequest) SetBonusId(v float32)`

SetBonusId sets BonusId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


