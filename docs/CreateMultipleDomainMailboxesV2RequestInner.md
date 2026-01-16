# CreateMultipleDomainMailboxesV2RequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Название почтового ящика | 
**Password** | **string** | Пароль почтового ящика | 
**OwnerFullName** | Pointer to **string** | ФИО владельца почтового ящика | [optional] 
**Comment** | Pointer to **string** | Комментарий почтового ящика | [optional] 
**FilterStatus** | Pointer to **bool** | Статус фильтрации почты | [optional] 
**FilterAction** | Pointer to **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;directory&#x60; - переместить в папку спам; \\  &#x60;label&#x60; - пометить письмо; \\  Если передан параметр &#x60;filter_status&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewCreateMultipleDomainMailboxesV2RequestInner

`func NewCreateMultipleDomainMailboxesV2RequestInner(login string, password string, ) *CreateMultipleDomainMailboxesV2RequestInner`

NewCreateMultipleDomainMailboxesV2RequestInner instantiates a new CreateMultipleDomainMailboxesV2RequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMultipleDomainMailboxesV2RequestInnerWithDefaults

`func NewCreateMultipleDomainMailboxesV2RequestInnerWithDefaults() *CreateMultipleDomainMailboxesV2RequestInner`

NewCreateMultipleDomainMailboxesV2RequestInnerWithDefaults instantiates a new CreateMultipleDomainMailboxesV2RequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOwnerFullName

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.

### HasOwnerFullName

`func (o *CreateMultipleDomainMailboxesV2RequestInner) HasOwnerFullName() bool`

HasOwnerFullName returns a boolean if a field has been set.

### GetComment

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateMultipleDomainMailboxesV2RequestInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFilterStatus

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetFilterStatus() bool`

GetFilterStatus returns the FilterStatus field if non-nil, zero value otherwise.

### GetFilterStatusOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetFilterStatusOk() (*bool, bool)`

GetFilterStatusOk returns a tuple with the FilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterStatus

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetFilterStatus(v bool)`

SetFilterStatus sets FilterStatus field to given value.

### HasFilterStatus

`func (o *CreateMultipleDomainMailboxesV2RequestInner) HasFilterStatus() bool`

HasFilterStatus returns a boolean if a field has been set.

### GetFilterAction

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetFilterAction() string`

GetFilterAction returns the FilterAction field if non-nil, zero value otherwise.

### GetFilterActionOk

`func (o *CreateMultipleDomainMailboxesV2RequestInner) GetFilterActionOk() (*string, bool)`

GetFilterActionOk returns a tuple with the FilterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAction

`func (o *CreateMultipleDomainMailboxesV2RequestInner) SetFilterAction(v string)`

SetFilterAction sets FilterAction field to given value.

### HasFilterAction

`func (o *CreateMultipleDomainMailboxesV2RequestInner) HasFilterAction() bool`

HasFilterAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


