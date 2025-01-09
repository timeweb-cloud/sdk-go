# CreateMultipleDomainMailboxesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailboxes** | [**[]CreateMultipleDomainMailboxesRequestMailboxesInner**](CreateMultipleDomainMailboxesRequestMailboxesInner.md) | Массив объектов с данными почтовых ящиков | 

## Methods

### NewCreateMultipleDomainMailboxesRequest

`func NewCreateMultipleDomainMailboxesRequest(mailboxes []CreateMultipleDomainMailboxesRequestMailboxesInner, ) *CreateMultipleDomainMailboxesRequest`

NewCreateMultipleDomainMailboxesRequest instantiates a new CreateMultipleDomainMailboxesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMultipleDomainMailboxesRequestWithDefaults

`func NewCreateMultipleDomainMailboxesRequestWithDefaults() *CreateMultipleDomainMailboxesRequest`

NewCreateMultipleDomainMailboxesRequestWithDefaults instantiates a new CreateMultipleDomainMailboxesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxes

`func (o *CreateMultipleDomainMailboxesRequest) GetMailboxes() []CreateMultipleDomainMailboxesRequestMailboxesInner`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *CreateMultipleDomainMailboxesRequest) GetMailboxesOk() (*[]CreateMultipleDomainMailboxesRequestMailboxesInner, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *CreateMultipleDomainMailboxesRequest) SetMailboxes(v []CreateMultipleDomainMailboxesRequestMailboxesInner)`

SetMailboxes sets Mailboxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


