# CreateDomainMailboxV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailbox** | **string** | Название почтового ящика | 
**Password** | **string** | Пароль почтового ящика | 
**Comment** | Pointer to **string** | Комментарий почтового ящика | [optional] 
**OwnerFullName** | Pointer to **string** | ФИО владельца почтового ящика | [optional] 
**FilterStatus** | Pointer to **bool** | Статус фильтрации почты | [optional] 
**FilterAction** | Pointer to **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;directory&#x60; - переместить в папку спам; \\  &#x60;label&#x60; - пометить письмо; \\  Если передан параметр &#x60;filter_status&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewCreateDomainMailboxV2Request

`func NewCreateDomainMailboxV2Request(mailbox string, password string, ) *CreateDomainMailboxV2Request`

NewCreateDomainMailboxV2Request instantiates a new CreateDomainMailboxV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainMailboxV2RequestWithDefaults

`func NewCreateDomainMailboxV2RequestWithDefaults() *CreateDomainMailboxV2Request`

NewCreateDomainMailboxV2RequestWithDefaults instantiates a new CreateDomainMailboxV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailbox

`func (o *CreateDomainMailboxV2Request) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *CreateDomainMailboxV2Request) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *CreateDomainMailboxV2Request) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetPassword

`func (o *CreateDomainMailboxV2Request) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDomainMailboxV2Request) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDomainMailboxV2Request) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetComment

`func (o *CreateDomainMailboxV2Request) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateDomainMailboxV2Request) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateDomainMailboxV2Request) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateDomainMailboxV2Request) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetOwnerFullName

`func (o *CreateDomainMailboxV2Request) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *CreateDomainMailboxV2Request) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *CreateDomainMailboxV2Request) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.

### HasOwnerFullName

`func (o *CreateDomainMailboxV2Request) HasOwnerFullName() bool`

HasOwnerFullName returns a boolean if a field has been set.

### GetFilterStatus

`func (o *CreateDomainMailboxV2Request) GetFilterStatus() bool`

GetFilterStatus returns the FilterStatus field if non-nil, zero value otherwise.

### GetFilterStatusOk

`func (o *CreateDomainMailboxV2Request) GetFilterStatusOk() (*bool, bool)`

GetFilterStatusOk returns a tuple with the FilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterStatus

`func (o *CreateDomainMailboxV2Request) SetFilterStatus(v bool)`

SetFilterStatus sets FilterStatus field to given value.

### HasFilterStatus

`func (o *CreateDomainMailboxV2Request) HasFilterStatus() bool`

HasFilterStatus returns a boolean if a field has been set.

### GetFilterAction

`func (o *CreateDomainMailboxV2Request) GetFilterAction() string`

GetFilterAction returns the FilterAction field if non-nil, zero value otherwise.

### GetFilterActionOk

`func (o *CreateDomainMailboxV2Request) GetFilterActionOk() (*string, bool)`

GetFilterActionOk returns a tuple with the FilterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAction

`func (o *CreateDomainMailboxV2Request) SetFilterAction(v string)`

SetFilterAction sets FilterAction field to given value.

### HasFilterAction

`func (o *CreateDomainMailboxV2Request) HasFilterAction() bool`

HasFilterAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


