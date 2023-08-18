# AutoReplyIsDisabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли автоответчик на входящие письма | 
**Message** | Pointer to **string** | Сообщение автоответчика на входящие письма | [optional] 
**Subject** | Pointer to **string** | Тема сообщения автоответчика на входящие письма. | [optional] 

## Methods

### NewAutoReplyIsDisabled

`func NewAutoReplyIsDisabled(isEnabled bool, ) *AutoReplyIsDisabled`

NewAutoReplyIsDisabled instantiates a new AutoReplyIsDisabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReplyIsDisabledWithDefaults

`func NewAutoReplyIsDisabledWithDefaults() *AutoReplyIsDisabled`

NewAutoReplyIsDisabledWithDefaults instantiates a new AutoReplyIsDisabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AutoReplyIsDisabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AutoReplyIsDisabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AutoReplyIsDisabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *AutoReplyIsDisabled) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AutoReplyIsDisabled) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AutoReplyIsDisabled) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AutoReplyIsDisabled) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubject

`func (o *AutoReplyIsDisabled) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AutoReplyIsDisabled) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AutoReplyIsDisabled) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AutoReplyIsDisabled) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


