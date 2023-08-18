# AutoReplyIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Включен ли автоответчик на входящие письма | [optional] 
**Message** | Pointer to **string** | Сообщение автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**Subject** | Pointer to **string** | Тема сообщения автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewAutoReplyIsEnabled

`func NewAutoReplyIsEnabled() *AutoReplyIsEnabled`

NewAutoReplyIsEnabled instantiates a new AutoReplyIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReplyIsEnabledWithDefaults

`func NewAutoReplyIsEnabledWithDefaults() *AutoReplyIsEnabled`

NewAutoReplyIsEnabledWithDefaults instantiates a new AutoReplyIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AutoReplyIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AutoReplyIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AutoReplyIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AutoReplyIsEnabled) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMessage

`func (o *AutoReplyIsEnabled) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AutoReplyIsEnabled) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AutoReplyIsEnabled) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AutoReplyIsEnabled) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubject

`func (o *AutoReplyIsEnabled) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AutoReplyIsEnabled) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AutoReplyIsEnabled) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AutoReplyIsEnabled) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


