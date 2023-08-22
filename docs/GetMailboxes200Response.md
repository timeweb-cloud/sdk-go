# GetMailboxes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Mailboxes** | [**[]Mailbox**](Mailbox.md) |  | 

## Methods

### NewGetMailboxes200Response

`func NewGetMailboxes200Response(meta Meta, mailboxes []Mailbox, ) *GetMailboxes200Response`

NewGetMailboxes200Response instantiates a new GetMailboxes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMailboxes200ResponseWithDefaults

`func NewGetMailboxes200ResponseWithDefaults() *GetMailboxes200Response`

NewGetMailboxes200ResponseWithDefaults instantiates a new GetMailboxes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetMailboxes200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMailboxes200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMailboxes200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetMailboxes

`func (o *GetMailboxes200Response) GetMailboxes() []Mailbox`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *GetMailboxes200Response) GetMailboxesOk() (*[]Mailbox, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *GetMailboxes200Response) SetMailboxes(v []Mailbox)`

SetMailboxes sets Mailboxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


