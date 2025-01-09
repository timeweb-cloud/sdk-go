# Use

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoneySource** | **string** | Тип создаваемой заявки. | 
**PersonId** | Pointer to **float32** | ID администратора, на которого зарегистрирован домен. | [optional] 

## Methods

### NewUse

`func NewUse(moneySource string, ) *Use`

NewUse instantiates a new Use object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUseWithDefaults

`func NewUseWithDefaults() *Use`

NewUseWithDefaults instantiates a new Use object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoneySource

`func (o *Use) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *Use) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *Use) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### GetPersonId

`func (o *Use) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Use) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Use) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Use) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


