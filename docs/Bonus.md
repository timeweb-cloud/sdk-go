# Bonus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoneySource** | **string** | Тип создаваемой заявки. | 
**PersonId** | Pointer to **float32** | ID администратора, на которого зарегистрирован домен. | [optional] 
**BonusId** | **float32** | ID бонуса. | 

## Methods

### NewBonus

`func NewBonus(moneySource string, bonusId float32, ) *Bonus`

NewBonus instantiates a new Bonus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBonusWithDefaults

`func NewBonusWithDefaults() *Bonus`

NewBonusWithDefaults instantiates a new Bonus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoneySource

`func (o *Bonus) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *Bonus) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *Bonus) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### GetPersonId

`func (o *Bonus) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Bonus) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Bonus) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Bonus) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetBonusId

`func (o *Bonus) GetBonusId() float32`

GetBonusId returns the BonusId field if non-nil, zero value otherwise.

### GetBonusIdOk

`func (o *Bonus) GetBonusIdOk() (*float32, bool)`

GetBonusIdOk returns a tuple with the BonusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusId

`func (o *Bonus) SetBonusId(v float32)`

SetBonusId sets BonusId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


