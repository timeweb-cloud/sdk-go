# UpdateMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoReply** | Pointer to [**UpdateMailboxAutoReply**](UpdateMailboxAutoReply.md) |  | [optional] 
**SpamFilter** | Pointer to [**UpdateMailboxSpamFilter**](UpdateMailboxSpamFilter.md) |  | [optional] 
**ForwardingIncoming** | Pointer to [**UpdateMailboxForwardingIncoming**](UpdateMailboxForwardingIncoming.md) |  | [optional] 
**ForwardingOutgoing** | Pointer to [**UpdateMailboxForwardingOutgoing**](UpdateMailboxForwardingOutgoing.md) |  | [optional] 
**Comment** | Pointer to **string** | Комментарий к почтовому ящику | [optional] 
**Password** | Pointer to **string** | Пароль почтового ящика | [optional] 

## Methods

### NewUpdateMailbox

`func NewUpdateMailbox() *UpdateMailbox`

NewUpdateMailbox instantiates a new UpdateMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxWithDefaults

`func NewUpdateMailboxWithDefaults() *UpdateMailbox`

NewUpdateMailboxWithDefaults instantiates a new UpdateMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoReply

`func (o *UpdateMailbox) GetAutoReply() UpdateMailboxAutoReply`

GetAutoReply returns the AutoReply field if non-nil, zero value otherwise.

### GetAutoReplyOk

`func (o *UpdateMailbox) GetAutoReplyOk() (*UpdateMailboxAutoReply, bool)`

GetAutoReplyOk returns a tuple with the AutoReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReply

`func (o *UpdateMailbox) SetAutoReply(v UpdateMailboxAutoReply)`

SetAutoReply sets AutoReply field to given value.

### HasAutoReply

`func (o *UpdateMailbox) HasAutoReply() bool`

HasAutoReply returns a boolean if a field has been set.

### GetSpamFilter

`func (o *UpdateMailbox) GetSpamFilter() UpdateMailboxSpamFilter`

GetSpamFilter returns the SpamFilter field if non-nil, zero value otherwise.

### GetSpamFilterOk

`func (o *UpdateMailbox) GetSpamFilterOk() (*UpdateMailboxSpamFilter, bool)`

GetSpamFilterOk returns a tuple with the SpamFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamFilter

`func (o *UpdateMailbox) SetSpamFilter(v UpdateMailboxSpamFilter)`

SetSpamFilter sets SpamFilter field to given value.

### HasSpamFilter

`func (o *UpdateMailbox) HasSpamFilter() bool`

HasSpamFilter returns a boolean if a field has been set.

### GetForwardingIncoming

`func (o *UpdateMailbox) GetForwardingIncoming() UpdateMailboxForwardingIncoming`

GetForwardingIncoming returns the ForwardingIncoming field if non-nil, zero value otherwise.

### GetForwardingIncomingOk

`func (o *UpdateMailbox) GetForwardingIncomingOk() (*UpdateMailboxForwardingIncoming, bool)`

GetForwardingIncomingOk returns a tuple with the ForwardingIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingIncoming

`func (o *UpdateMailbox) SetForwardingIncoming(v UpdateMailboxForwardingIncoming)`

SetForwardingIncoming sets ForwardingIncoming field to given value.

### HasForwardingIncoming

`func (o *UpdateMailbox) HasForwardingIncoming() bool`

HasForwardingIncoming returns a boolean if a field has been set.

### GetForwardingOutgoing

`func (o *UpdateMailbox) GetForwardingOutgoing() UpdateMailboxForwardingOutgoing`

GetForwardingOutgoing returns the ForwardingOutgoing field if non-nil, zero value otherwise.

### GetForwardingOutgoingOk

`func (o *UpdateMailbox) GetForwardingOutgoingOk() (*UpdateMailboxForwardingOutgoing, bool)`

GetForwardingOutgoingOk returns a tuple with the ForwardingOutgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingOutgoing

`func (o *UpdateMailbox) SetForwardingOutgoing(v UpdateMailboxForwardingOutgoing)`

SetForwardingOutgoing sets ForwardingOutgoing field to given value.

### HasForwardingOutgoing

`func (o *UpdateMailbox) HasForwardingOutgoing() bool`

HasForwardingOutgoing returns a boolean if a field has been set.

### GetComment

`func (o *UpdateMailbox) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateMailbox) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateMailbox) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateMailbox) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateMailbox) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMailbox) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMailbox) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMailbox) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


