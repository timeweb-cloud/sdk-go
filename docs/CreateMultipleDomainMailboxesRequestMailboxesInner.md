# CreateMultipleDomainMailboxesRequestMailboxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Название почтового ящика | 
**Password** | **string** | Пароль почтового ящика | 
**OwnerFullName** | Pointer to **string** | ФИО владельца почтового ящика | [optional] 
**Comment** | Pointer to **string** | Комментарий почтового ящика | [optional] 

## Methods

### NewCreateMultipleDomainMailboxesRequestMailboxesInner

`func NewCreateMultipleDomainMailboxesRequestMailboxesInner(login string, password string, ) *CreateMultipleDomainMailboxesRequestMailboxesInner`

NewCreateMultipleDomainMailboxesRequestMailboxesInner instantiates a new CreateMultipleDomainMailboxesRequestMailboxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMultipleDomainMailboxesRequestMailboxesInnerWithDefaults

`func NewCreateMultipleDomainMailboxesRequestMailboxesInnerWithDefaults() *CreateMultipleDomainMailboxesRequestMailboxesInner`

NewCreateMultipleDomainMailboxesRequestMailboxesInnerWithDefaults instantiates a new CreateMultipleDomainMailboxesRequestMailboxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOwnerFullName

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.

### HasOwnerFullName

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) HasOwnerFullName() bool`

HasOwnerFullName returns a boolean if a field has been set.

### GetComment

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateMultipleDomainMailboxesRequestMailboxesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


