# AutoreplyIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли автоответчик на входящие письма | 
**AutoreplyMessage** | Pointer to **string** | Сообщение автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**AutoreplySubject** | Pointer to **string** | Тема сообщения автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewAutoreplyIsEnabled

`func NewAutoreplyIsEnabled(isEnabled bool, ) *AutoreplyIsEnabled`

NewAutoreplyIsEnabled instantiates a new AutoreplyIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoreplyIsEnabledWithDefaults

`func NewAutoreplyIsEnabledWithDefaults() *AutoreplyIsEnabled`

NewAutoreplyIsEnabledWithDefaults instantiates a new AutoreplyIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AutoreplyIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AutoreplyIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AutoreplyIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAutoreplyMessage

`func (o *AutoreplyIsEnabled) GetAutoreplyMessage() string`

GetAutoreplyMessage returns the AutoreplyMessage field if non-nil, zero value otherwise.

### GetAutoreplyMessageOk

`func (o *AutoreplyIsEnabled) GetAutoreplyMessageOk() (*string, bool)`

GetAutoreplyMessageOk returns a tuple with the AutoreplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplyMessage

`func (o *AutoreplyIsEnabled) SetAutoreplyMessage(v string)`

SetAutoreplyMessage sets AutoreplyMessage field to given value.

### HasAutoreplyMessage

`func (o *AutoreplyIsEnabled) HasAutoreplyMessage() bool`

HasAutoreplyMessage returns a boolean if a field has been set.

### GetAutoreplySubject

`func (o *AutoreplyIsEnabled) GetAutoreplySubject() string`

GetAutoreplySubject returns the AutoreplySubject field if non-nil, zero value otherwise.

### GetAutoreplySubjectOk

`func (o *AutoreplyIsEnabled) GetAutoreplySubjectOk() (*string, bool)`

GetAutoreplySubjectOk returns a tuple with the AutoreplySubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplySubject

`func (o *AutoreplyIsEnabled) SetAutoreplySubject(v string)`

SetAutoreplySubject sets AutoreplySubject field to given value.

### HasAutoreplySubject

`func (o *AutoreplyIsEnabled) HasAutoreplySubject() bool`

HasAutoreplySubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


