# MailboxV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdnName** | **string** | IDN домен почтового ящика | 
**AutoreplyMessage** | **string** | Сообщение автоответчика на входящие письма | 
**AutoreplyStatus** | **bool** | Включен ли автоответчик на входящие письма | 
**AutoreplySubject** | **string** | Тема сообщения автоответчика на входящие письма | 
**Comment** | **string** | Комментарий к почтовому ящику | 
**FilterAction** | **string** | Что делать с письмами, которые попадают в спам | 
**FilterStatus** | **bool** | Включен ли спам-фильтр | 
**ForwardList** | **[]string** | Список адресов для пересылки входящих писем | 
**ForwardStatus** | **bool** | Включена ли пересылка входящих писем | 
**OutgoingControl** | **bool** | Включена ли пересылка исходящих писем | 
**OutgoingEmail** | **string** | Адрес для пересылки исходящих писем | 
**Password** | **string** | Пароль почтового ящика (всегда возвращается пустой строкой) | 
**Spambox** | **string** | Адрес для пересылки спама при выбранном действии forward | 
**WhiteList** | **[]string** | Белый список адресов от которых письма не будут попадать в спам | 
**Webmail** | **bool** | Доступен ли Webmail | 
**Dovecot** | **bool** | Есть ли доступ через dovecot | 
**Fqdn** | **string** | Домен почты | 
**LeaveMessages** | **bool** | Оставлять ли сообщения на сервере при пересылке | 
**Mailbox** | **string** | Название почтового ящика | 
**OwnerFullName** | **string** | ФИО владельца почтового ящика | 

## Methods

### NewMailboxV2

`func NewMailboxV2(idnName string, autoreplyMessage string, autoreplyStatus bool, autoreplySubject string, comment string, filterAction string, filterStatus bool, forwardList []string, forwardStatus bool, outgoingControl bool, outgoingEmail string, password string, spambox string, whiteList []string, webmail bool, dovecot bool, fqdn string, leaveMessages bool, mailbox string, ownerFullName string, ) *MailboxV2`

NewMailboxV2 instantiates a new MailboxV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxV2WithDefaults

`func NewMailboxV2WithDefaults() *MailboxV2`

NewMailboxV2WithDefaults instantiates a new MailboxV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdnName

`func (o *MailboxV2) GetIdnName() string`

GetIdnName returns the IdnName field if non-nil, zero value otherwise.

### GetIdnNameOk

`func (o *MailboxV2) GetIdnNameOk() (*string, bool)`

GetIdnNameOk returns a tuple with the IdnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnName

`func (o *MailboxV2) SetIdnName(v string)`

SetIdnName sets IdnName field to given value.


### GetAutoreplyMessage

`func (o *MailboxV2) GetAutoreplyMessage() string`

GetAutoreplyMessage returns the AutoreplyMessage field if non-nil, zero value otherwise.

### GetAutoreplyMessageOk

`func (o *MailboxV2) GetAutoreplyMessageOk() (*string, bool)`

GetAutoreplyMessageOk returns a tuple with the AutoreplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplyMessage

`func (o *MailboxV2) SetAutoreplyMessage(v string)`

SetAutoreplyMessage sets AutoreplyMessage field to given value.


### GetAutoreplyStatus

`func (o *MailboxV2) GetAutoreplyStatus() bool`

GetAutoreplyStatus returns the AutoreplyStatus field if non-nil, zero value otherwise.

### GetAutoreplyStatusOk

`func (o *MailboxV2) GetAutoreplyStatusOk() (*bool, bool)`

GetAutoreplyStatusOk returns a tuple with the AutoreplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplyStatus

`func (o *MailboxV2) SetAutoreplyStatus(v bool)`

SetAutoreplyStatus sets AutoreplyStatus field to given value.


### GetAutoreplySubject

`func (o *MailboxV2) GetAutoreplySubject() string`

GetAutoreplySubject returns the AutoreplySubject field if non-nil, zero value otherwise.

### GetAutoreplySubjectOk

`func (o *MailboxV2) GetAutoreplySubjectOk() (*string, bool)`

GetAutoreplySubjectOk returns a tuple with the AutoreplySubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplySubject

`func (o *MailboxV2) SetAutoreplySubject(v string)`

SetAutoreplySubject sets AutoreplySubject field to given value.


### GetComment

`func (o *MailboxV2) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MailboxV2) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MailboxV2) SetComment(v string)`

SetComment sets Comment field to given value.


### GetFilterAction

`func (o *MailboxV2) GetFilterAction() string`

GetFilterAction returns the FilterAction field if non-nil, zero value otherwise.

### GetFilterActionOk

`func (o *MailboxV2) GetFilterActionOk() (*string, bool)`

GetFilterActionOk returns a tuple with the FilterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAction

`func (o *MailboxV2) SetFilterAction(v string)`

SetFilterAction sets FilterAction field to given value.


### GetFilterStatus

`func (o *MailboxV2) GetFilterStatus() bool`

GetFilterStatus returns the FilterStatus field if non-nil, zero value otherwise.

### GetFilterStatusOk

`func (o *MailboxV2) GetFilterStatusOk() (*bool, bool)`

GetFilterStatusOk returns a tuple with the FilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterStatus

`func (o *MailboxV2) SetFilterStatus(v bool)`

SetFilterStatus sets FilterStatus field to given value.


### GetForwardList

`func (o *MailboxV2) GetForwardList() []string`

GetForwardList returns the ForwardList field if non-nil, zero value otherwise.

### GetForwardListOk

`func (o *MailboxV2) GetForwardListOk() (*[]string, bool)`

GetForwardListOk returns a tuple with the ForwardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardList

`func (o *MailboxV2) SetForwardList(v []string)`

SetForwardList sets ForwardList field to given value.


### GetForwardStatus

`func (o *MailboxV2) GetForwardStatus() bool`

GetForwardStatus returns the ForwardStatus field if non-nil, zero value otherwise.

### GetForwardStatusOk

`func (o *MailboxV2) GetForwardStatusOk() (*bool, bool)`

GetForwardStatusOk returns a tuple with the ForwardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardStatus

`func (o *MailboxV2) SetForwardStatus(v bool)`

SetForwardStatus sets ForwardStatus field to given value.


### GetOutgoingControl

`func (o *MailboxV2) GetOutgoingControl() bool`

GetOutgoingControl returns the OutgoingControl field if non-nil, zero value otherwise.

### GetOutgoingControlOk

`func (o *MailboxV2) GetOutgoingControlOk() (*bool, bool)`

GetOutgoingControlOk returns a tuple with the OutgoingControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingControl

`func (o *MailboxV2) SetOutgoingControl(v bool)`

SetOutgoingControl sets OutgoingControl field to given value.


### GetOutgoingEmail

`func (o *MailboxV2) GetOutgoingEmail() string`

GetOutgoingEmail returns the OutgoingEmail field if non-nil, zero value otherwise.

### GetOutgoingEmailOk

`func (o *MailboxV2) GetOutgoingEmailOk() (*string, bool)`

GetOutgoingEmailOk returns a tuple with the OutgoingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingEmail

`func (o *MailboxV2) SetOutgoingEmail(v string)`

SetOutgoingEmail sets OutgoingEmail field to given value.


### GetPassword

`func (o *MailboxV2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailboxV2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailboxV2) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSpambox

`func (o *MailboxV2) GetSpambox() string`

GetSpambox returns the Spambox field if non-nil, zero value otherwise.

### GetSpamboxOk

`func (o *MailboxV2) GetSpamboxOk() (*string, bool)`

GetSpamboxOk returns a tuple with the Spambox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpambox

`func (o *MailboxV2) SetSpambox(v string)`

SetSpambox sets Spambox field to given value.


### GetWhiteList

`func (o *MailboxV2) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *MailboxV2) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *MailboxV2) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.


### GetWebmail

`func (o *MailboxV2) GetWebmail() bool`

GetWebmail returns the Webmail field if non-nil, zero value otherwise.

### GetWebmailOk

`func (o *MailboxV2) GetWebmailOk() (*bool, bool)`

GetWebmailOk returns a tuple with the Webmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebmail

`func (o *MailboxV2) SetWebmail(v bool)`

SetWebmail sets Webmail field to given value.


### GetDovecot

`func (o *MailboxV2) GetDovecot() bool`

GetDovecot returns the Dovecot field if non-nil, zero value otherwise.

### GetDovecotOk

`func (o *MailboxV2) GetDovecotOk() (*bool, bool)`

GetDovecotOk returns a tuple with the Dovecot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDovecot

`func (o *MailboxV2) SetDovecot(v bool)`

SetDovecot sets Dovecot field to given value.


### GetFqdn

`func (o *MailboxV2) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MailboxV2) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MailboxV2) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetLeaveMessages

`func (o *MailboxV2) GetLeaveMessages() bool`

GetLeaveMessages returns the LeaveMessages field if non-nil, zero value otherwise.

### GetLeaveMessagesOk

`func (o *MailboxV2) GetLeaveMessagesOk() (*bool, bool)`

GetLeaveMessagesOk returns a tuple with the LeaveMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveMessages

`func (o *MailboxV2) SetLeaveMessages(v bool)`

SetLeaveMessages sets LeaveMessages field to given value.


### GetMailbox

`func (o *MailboxV2) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *MailboxV2) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *MailboxV2) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetOwnerFullName

`func (o *MailboxV2) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *MailboxV2) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *MailboxV2) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


