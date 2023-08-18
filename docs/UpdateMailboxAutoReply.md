# UpdateMailboxAutoReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли автоответчик на входящие письма | 
**Message** | Pointer to **string** | Сообщение автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**Subject** | Pointer to **string** | Тема сообщения автоответчика на входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewUpdateMailboxAutoReply

`func NewUpdateMailboxAutoReply(isEnabled bool, ) *UpdateMailboxAutoReply`

NewUpdateMailboxAutoReply instantiates a new UpdateMailboxAutoReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxAutoReplyWithDefaults

`func NewUpdateMailboxAutoReplyWithDefaults() *UpdateMailboxAutoReply`

NewUpdateMailboxAutoReplyWithDefaults instantiates a new UpdateMailboxAutoReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateMailboxAutoReply) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateMailboxAutoReply) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateMailboxAutoReply) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *UpdateMailboxAutoReply) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateMailboxAutoReply) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateMailboxAutoReply) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateMailboxAutoReply) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubject

`func (o *UpdateMailboxAutoReply) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateMailboxAutoReply) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateMailboxAutoReply) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateMailboxAutoReply) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


