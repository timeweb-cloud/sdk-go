# CreateDomainMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailbox** | **string** | Название почтового ящика | 
**Password** | **string** | Пароль почтового ящика | 
**Comment** | Pointer to **string** | Комментарий почтового ящика | [optional] 

## Methods

### NewCreateDomainMailboxRequest

`func NewCreateDomainMailboxRequest(mailbox string, password string, ) *CreateDomainMailboxRequest`

NewCreateDomainMailboxRequest instantiates a new CreateDomainMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainMailboxRequestWithDefaults

`func NewCreateDomainMailboxRequestWithDefaults() *CreateDomainMailboxRequest`

NewCreateDomainMailboxRequestWithDefaults instantiates a new CreateDomainMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailbox

`func (o *CreateDomainMailboxRequest) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *CreateDomainMailboxRequest) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *CreateDomainMailboxRequest) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetPassword

`func (o *CreateDomainMailboxRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDomainMailboxRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDomainMailboxRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetComment

`func (o *CreateDomainMailboxRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateDomainMailboxRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateDomainMailboxRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateDomainMailboxRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


