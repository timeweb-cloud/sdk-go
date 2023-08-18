# MailboxAutoReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли автоответчик на входящие письма | 
**Message** | **string** | Сообщение автоответчика на входящие письма | 
**Subject** | **string** | Тема сообщения автоответчика на входящие письма | 

## Methods

### NewMailboxAutoReply

`func NewMailboxAutoReply(isEnabled bool, message string, subject string, ) *MailboxAutoReply`

NewMailboxAutoReply instantiates a new MailboxAutoReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxAutoReplyWithDefaults

`func NewMailboxAutoReplyWithDefaults() *MailboxAutoReply`

NewMailboxAutoReplyWithDefaults instantiates a new MailboxAutoReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MailboxAutoReply) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MailboxAutoReply) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MailboxAutoReply) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *MailboxAutoReply) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MailboxAutoReply) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MailboxAutoReply) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSubject

`func (o *MailboxAutoReply) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MailboxAutoReply) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MailboxAutoReply) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


