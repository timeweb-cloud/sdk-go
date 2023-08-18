# Mailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoReply** | [**MailboxAutoReply**](MailboxAutoReply.md) |  | 
**SpamFilter** | [**MailboxSpamFilter**](MailboxSpamFilter.md) |  | 
**ForwardingIncoming** | [**MailboxForwardingIncoming**](MailboxForwardingIncoming.md) |  | 
**ForwardingOutgoing** | [**MailboxForwardingOutgoing**](MailboxForwardingOutgoing.md) |  | 
**Comment** | **string** | Комментарий к почтовому ящику | 
**Fqdn** | **string** | Домен почты | 
**Mailbox** | **string** | Название почтового ящика | 
**Password** | **string** | Пароль почтового ящика | 
**UsageSpace** | **float32** | Использованное место на почтовом ящике (в Мб) | 
**IsWebmail** | **bool** | Доступен ли Webmail | 
**IdnName** | **string** | IDN домен почтового ящика | 
**IsDovecot** | **bool** | Есть ли доступ через dovecot | 

## Methods

### NewMailbox

`func NewMailbox(autoReply MailboxAutoReply, spamFilter MailboxSpamFilter, forwardingIncoming MailboxForwardingIncoming, forwardingOutgoing MailboxForwardingOutgoing, comment string, fqdn string, mailbox string, password string, usageSpace float32, isWebmail bool, idnName string, isDovecot bool, ) *Mailbox`

NewMailbox instantiates a new Mailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxWithDefaults

`func NewMailboxWithDefaults() *Mailbox`

NewMailboxWithDefaults instantiates a new Mailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoReply

`func (o *Mailbox) GetAutoReply() MailboxAutoReply`

GetAutoReply returns the AutoReply field if non-nil, zero value otherwise.

### GetAutoReplyOk

`func (o *Mailbox) GetAutoReplyOk() (*MailboxAutoReply, bool)`

GetAutoReplyOk returns a tuple with the AutoReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReply

`func (o *Mailbox) SetAutoReply(v MailboxAutoReply)`

SetAutoReply sets AutoReply field to given value.


### GetSpamFilter

`func (o *Mailbox) GetSpamFilter() MailboxSpamFilter`

GetSpamFilter returns the SpamFilter field if non-nil, zero value otherwise.

### GetSpamFilterOk

`func (o *Mailbox) GetSpamFilterOk() (*MailboxSpamFilter, bool)`

GetSpamFilterOk returns a tuple with the SpamFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamFilter

`func (o *Mailbox) SetSpamFilter(v MailboxSpamFilter)`

SetSpamFilter sets SpamFilter field to given value.


### GetForwardingIncoming

`func (o *Mailbox) GetForwardingIncoming() MailboxForwardingIncoming`

GetForwardingIncoming returns the ForwardingIncoming field if non-nil, zero value otherwise.

### GetForwardingIncomingOk

`func (o *Mailbox) GetForwardingIncomingOk() (*MailboxForwardingIncoming, bool)`

GetForwardingIncomingOk returns a tuple with the ForwardingIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingIncoming

`func (o *Mailbox) SetForwardingIncoming(v MailboxForwardingIncoming)`

SetForwardingIncoming sets ForwardingIncoming field to given value.


### GetForwardingOutgoing

`func (o *Mailbox) GetForwardingOutgoing() MailboxForwardingOutgoing`

GetForwardingOutgoing returns the ForwardingOutgoing field if non-nil, zero value otherwise.

### GetForwardingOutgoingOk

`func (o *Mailbox) GetForwardingOutgoingOk() (*MailboxForwardingOutgoing, bool)`

GetForwardingOutgoingOk returns a tuple with the ForwardingOutgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingOutgoing

`func (o *Mailbox) SetForwardingOutgoing(v MailboxForwardingOutgoing)`

SetForwardingOutgoing sets ForwardingOutgoing field to given value.


### GetComment

`func (o *Mailbox) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Mailbox) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Mailbox) SetComment(v string)`

SetComment sets Comment field to given value.


### GetFqdn

`func (o *Mailbox) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Mailbox) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Mailbox) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetMailbox

`func (o *Mailbox) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *Mailbox) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *Mailbox) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetPassword

`func (o *Mailbox) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Mailbox) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Mailbox) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsageSpace

`func (o *Mailbox) GetUsageSpace() float32`

GetUsageSpace returns the UsageSpace field if non-nil, zero value otherwise.

### GetUsageSpaceOk

`func (o *Mailbox) GetUsageSpaceOk() (*float32, bool)`

GetUsageSpaceOk returns a tuple with the UsageSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSpace

`func (o *Mailbox) SetUsageSpace(v float32)`

SetUsageSpace sets UsageSpace field to given value.


### GetIsWebmail

`func (o *Mailbox) GetIsWebmail() bool`

GetIsWebmail returns the IsWebmail field if non-nil, zero value otherwise.

### GetIsWebmailOk

`func (o *Mailbox) GetIsWebmailOk() (*bool, bool)`

GetIsWebmailOk returns a tuple with the IsWebmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebmail

`func (o *Mailbox) SetIsWebmail(v bool)`

SetIsWebmail sets IsWebmail field to given value.


### GetIdnName

`func (o *Mailbox) GetIdnName() string`

GetIdnName returns the IdnName field if non-nil, zero value otherwise.

### GetIdnNameOk

`func (o *Mailbox) GetIdnNameOk() (*string, bool)`

GetIdnNameOk returns a tuple with the IdnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnName

`func (o *Mailbox) SetIdnName(v string)`

SetIdnName sets IdnName field to given value.


### GetIsDovecot

`func (o *Mailbox) GetIsDovecot() bool`

GetIsDovecot returns the IsDovecot field if non-nil, zero value otherwise.

### GetIsDovecotOk

`func (o *Mailbox) GetIsDovecotOk() (*bool, bool)`

GetIsDovecotOk returns a tuple with the IsDovecot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDovecot

`func (o *Mailbox) SetIsDovecot(v bool)`

SetIsDovecot sets IsDovecot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


