# CreateDomainMailbox201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailbox** | [**Mailbox**](Mailbox.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateDomainMailbox201Response

`func NewCreateDomainMailbox201Response(mailbox Mailbox, responseId string, ) *CreateDomainMailbox201Response`

NewCreateDomainMailbox201Response instantiates a new CreateDomainMailbox201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainMailbox201ResponseWithDefaults

`func NewCreateDomainMailbox201ResponseWithDefaults() *CreateDomainMailbox201Response`

NewCreateDomainMailbox201ResponseWithDefaults instantiates a new CreateDomainMailbox201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailbox

`func (o *CreateDomainMailbox201Response) GetMailbox() Mailbox`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *CreateDomainMailbox201Response) GetMailboxOk() (*Mailbox, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *CreateDomainMailbox201Response) SetMailbox(v Mailbox)`

SetMailbox sets Mailbox field to given value.


### GetResponseId

`func (o *CreateDomainMailbox201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateDomainMailbox201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateDomainMailbox201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


