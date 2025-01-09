# Free

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoneySource** | **string** | Тип создаваемой заявки. | 
**PersonId** | Pointer to **float32** | ID администратора, на которого зарегистрирован домен. | [optional] 
**AuthCode** | **string** | Код авторизации для переноса домена. | 

## Methods

### NewFree

`func NewFree(moneySource string, authCode string, ) *Free`

NewFree instantiates a new Free object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeWithDefaults

`func NewFreeWithDefaults() *Free`

NewFreeWithDefaults instantiates a new Free object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoneySource

`func (o *Free) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *Free) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *Free) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### GetPersonId

`func (o *Free) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Free) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Free) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Free) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetAuthCode

`func (o *Free) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *Free) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *Free) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


