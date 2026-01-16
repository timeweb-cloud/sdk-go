# MailboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdnName** | Pointer to **string** | IDN имя домена | [optional] 
**AutoreplyMessage** | Pointer to **string** | Сообщение автоответчика | [optional] 
**AutoreplyStatus** | Pointer to **bool** | Статус автоответчика | [optional] 
**AutoreplySubject** | Pointer to **string** | Тема автоответчика | [optional] 
**Comment** | Pointer to **string** | Комментарий | [optional] 
**FilterAction** | Pointer to **string** | Действие фильтра спама | [optional] 
**FilterStatus** | Pointer to **bool** | Статус фильтра спама | [optional] 
**ForwardList** | Pointer to **[]string** | Список адресов для пересылки | [optional] 
**ForwardStatus** | Pointer to **bool** | Статус пересылки | [optional] 
**OutgoingControl** | Pointer to **bool** | Контроль исходящей почты | [optional] 
**OutgoingEmail** | Pointer to **string** | Email для исходящих писем | [optional] 
**Password** | Pointer to **string** | Пароль (пустая строка в ответе) | [optional] 
**WhiteList** | Pointer to **[]string** | Белый список адресов | [optional] 
**Webmail** | Pointer to **bool** | Доступ к веб-почте | [optional] 
**Dovecot** | Pointer to **bool** | Использование Dovecot | [optional] 
**Fqdn** | Pointer to **string** | Полное доменное имя | [optional] 
**LeaveMessages** | Pointer to **bool** | Оставлять копии писем при пересылке | [optional] 
**Mailbox** | Pointer to **string** | Имя почтового ящика | [optional] 
**OwnerFullName** | Pointer to **string** | ФИО владельца | [optional] 

## Methods

### NewMailboxResponse

`func NewMailboxResponse() *MailboxResponse`

NewMailboxResponse instantiates a new MailboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxResponseWithDefaults

`func NewMailboxResponseWithDefaults() *MailboxResponse`

NewMailboxResponseWithDefaults instantiates a new MailboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdnName

`func (o *MailboxResponse) GetIdnName() string`

GetIdnName returns the IdnName field if non-nil, zero value otherwise.

### GetIdnNameOk

`func (o *MailboxResponse) GetIdnNameOk() (*string, bool)`

GetIdnNameOk returns a tuple with the IdnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnName

`func (o *MailboxResponse) SetIdnName(v string)`

SetIdnName sets IdnName field to given value.

### HasIdnName

`func (o *MailboxResponse) HasIdnName() bool`

HasIdnName returns a boolean if a field has been set.

### GetAutoreplyMessage

`func (o *MailboxResponse) GetAutoreplyMessage() string`

GetAutoreplyMessage returns the AutoreplyMessage field if non-nil, zero value otherwise.

### GetAutoreplyMessageOk

`func (o *MailboxResponse) GetAutoreplyMessageOk() (*string, bool)`

GetAutoreplyMessageOk returns a tuple with the AutoreplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplyMessage

`func (o *MailboxResponse) SetAutoreplyMessage(v string)`

SetAutoreplyMessage sets AutoreplyMessage field to given value.

### HasAutoreplyMessage

`func (o *MailboxResponse) HasAutoreplyMessage() bool`

HasAutoreplyMessage returns a boolean if a field has been set.

### GetAutoreplyStatus

`func (o *MailboxResponse) GetAutoreplyStatus() bool`

GetAutoreplyStatus returns the AutoreplyStatus field if non-nil, zero value otherwise.

### GetAutoreplyStatusOk

`func (o *MailboxResponse) GetAutoreplyStatusOk() (*bool, bool)`

GetAutoreplyStatusOk returns a tuple with the AutoreplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplyStatus

`func (o *MailboxResponse) SetAutoreplyStatus(v bool)`

SetAutoreplyStatus sets AutoreplyStatus field to given value.

### HasAutoreplyStatus

`func (o *MailboxResponse) HasAutoreplyStatus() bool`

HasAutoreplyStatus returns a boolean if a field has been set.

### GetAutoreplySubject

`func (o *MailboxResponse) GetAutoreplySubject() string`

GetAutoreplySubject returns the AutoreplySubject field if non-nil, zero value otherwise.

### GetAutoreplySubjectOk

`func (o *MailboxResponse) GetAutoreplySubjectOk() (*string, bool)`

GetAutoreplySubjectOk returns a tuple with the AutoreplySubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplySubject

`func (o *MailboxResponse) SetAutoreplySubject(v string)`

SetAutoreplySubject sets AutoreplySubject field to given value.

### HasAutoreplySubject

`func (o *MailboxResponse) HasAutoreplySubject() bool`

HasAutoreplySubject returns a boolean if a field has been set.

### GetComment

`func (o *MailboxResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MailboxResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MailboxResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MailboxResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFilterAction

`func (o *MailboxResponse) GetFilterAction() string`

GetFilterAction returns the FilterAction field if non-nil, zero value otherwise.

### GetFilterActionOk

`func (o *MailboxResponse) GetFilterActionOk() (*string, bool)`

GetFilterActionOk returns a tuple with the FilterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAction

`func (o *MailboxResponse) SetFilterAction(v string)`

SetFilterAction sets FilterAction field to given value.

### HasFilterAction

`func (o *MailboxResponse) HasFilterAction() bool`

HasFilterAction returns a boolean if a field has been set.

### GetFilterStatus

`func (o *MailboxResponse) GetFilterStatus() bool`

GetFilterStatus returns the FilterStatus field if non-nil, zero value otherwise.

### GetFilterStatusOk

`func (o *MailboxResponse) GetFilterStatusOk() (*bool, bool)`

GetFilterStatusOk returns a tuple with the FilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterStatus

`func (o *MailboxResponse) SetFilterStatus(v bool)`

SetFilterStatus sets FilterStatus field to given value.

### HasFilterStatus

`func (o *MailboxResponse) HasFilterStatus() bool`

HasFilterStatus returns a boolean if a field has been set.

### GetForwardList

`func (o *MailboxResponse) GetForwardList() []string`

GetForwardList returns the ForwardList field if non-nil, zero value otherwise.

### GetForwardListOk

`func (o *MailboxResponse) GetForwardListOk() (*[]string, bool)`

GetForwardListOk returns a tuple with the ForwardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardList

`func (o *MailboxResponse) SetForwardList(v []string)`

SetForwardList sets ForwardList field to given value.

### HasForwardList

`func (o *MailboxResponse) HasForwardList() bool`

HasForwardList returns a boolean if a field has been set.

### GetForwardStatus

`func (o *MailboxResponse) GetForwardStatus() bool`

GetForwardStatus returns the ForwardStatus field if non-nil, zero value otherwise.

### GetForwardStatusOk

`func (o *MailboxResponse) GetForwardStatusOk() (*bool, bool)`

GetForwardStatusOk returns a tuple with the ForwardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardStatus

`func (o *MailboxResponse) SetForwardStatus(v bool)`

SetForwardStatus sets ForwardStatus field to given value.

### HasForwardStatus

`func (o *MailboxResponse) HasForwardStatus() bool`

HasForwardStatus returns a boolean if a field has been set.

### GetOutgoingControl

`func (o *MailboxResponse) GetOutgoingControl() bool`

GetOutgoingControl returns the OutgoingControl field if non-nil, zero value otherwise.

### GetOutgoingControlOk

`func (o *MailboxResponse) GetOutgoingControlOk() (*bool, bool)`

GetOutgoingControlOk returns a tuple with the OutgoingControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingControl

`func (o *MailboxResponse) SetOutgoingControl(v bool)`

SetOutgoingControl sets OutgoingControl field to given value.

### HasOutgoingControl

`func (o *MailboxResponse) HasOutgoingControl() bool`

HasOutgoingControl returns a boolean if a field has been set.

### GetOutgoingEmail

`func (o *MailboxResponse) GetOutgoingEmail() string`

GetOutgoingEmail returns the OutgoingEmail field if non-nil, zero value otherwise.

### GetOutgoingEmailOk

`func (o *MailboxResponse) GetOutgoingEmailOk() (*string, bool)`

GetOutgoingEmailOk returns a tuple with the OutgoingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingEmail

`func (o *MailboxResponse) SetOutgoingEmail(v string)`

SetOutgoingEmail sets OutgoingEmail field to given value.

### HasOutgoingEmail

`func (o *MailboxResponse) HasOutgoingEmail() bool`

HasOutgoingEmail returns a boolean if a field has been set.

### GetPassword

`func (o *MailboxResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailboxResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailboxResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MailboxResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetWhiteList

`func (o *MailboxResponse) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *MailboxResponse) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *MailboxResponse) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *MailboxResponse) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetWebmail

`func (o *MailboxResponse) GetWebmail() bool`

GetWebmail returns the Webmail field if non-nil, zero value otherwise.

### GetWebmailOk

`func (o *MailboxResponse) GetWebmailOk() (*bool, bool)`

GetWebmailOk returns a tuple with the Webmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebmail

`func (o *MailboxResponse) SetWebmail(v bool)`

SetWebmail sets Webmail field to given value.

### HasWebmail

`func (o *MailboxResponse) HasWebmail() bool`

HasWebmail returns a boolean if a field has been set.

### GetDovecot

`func (o *MailboxResponse) GetDovecot() bool`

GetDovecot returns the Dovecot field if non-nil, zero value otherwise.

### GetDovecotOk

`func (o *MailboxResponse) GetDovecotOk() (*bool, bool)`

GetDovecotOk returns a tuple with the Dovecot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDovecot

`func (o *MailboxResponse) SetDovecot(v bool)`

SetDovecot sets Dovecot field to given value.

### HasDovecot

`func (o *MailboxResponse) HasDovecot() bool`

HasDovecot returns a boolean if a field has been set.

### GetFqdn

`func (o *MailboxResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MailboxResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MailboxResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *MailboxResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetLeaveMessages

`func (o *MailboxResponse) GetLeaveMessages() bool`

GetLeaveMessages returns the LeaveMessages field if non-nil, zero value otherwise.

### GetLeaveMessagesOk

`func (o *MailboxResponse) GetLeaveMessagesOk() (*bool, bool)`

GetLeaveMessagesOk returns a tuple with the LeaveMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveMessages

`func (o *MailboxResponse) SetLeaveMessages(v bool)`

SetLeaveMessages sets LeaveMessages field to given value.

### HasLeaveMessages

`func (o *MailboxResponse) HasLeaveMessages() bool`

HasLeaveMessages returns a boolean if a field has been set.

### GetMailbox

`func (o *MailboxResponse) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *MailboxResponse) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *MailboxResponse) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *MailboxResponse) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetOwnerFullName

`func (o *MailboxResponse) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *MailboxResponse) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *MailboxResponse) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.

### HasOwnerFullName

`func (o *MailboxResponse) HasOwnerFullName() bool`

HasOwnerFullName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


