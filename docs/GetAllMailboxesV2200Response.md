# GetAllMailboxesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Mailboxes** | [**[]MailboxV2**](MailboxV2.md) |  | 

## Methods

### NewGetAllMailboxesV2200Response

`func NewGetAllMailboxesV2200Response(meta Meta, mailboxes []MailboxV2, ) *GetAllMailboxesV2200Response`

NewGetAllMailboxesV2200Response instantiates a new GetAllMailboxesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllMailboxesV2200ResponseWithDefaults

`func NewGetAllMailboxesV2200ResponseWithDefaults() *GetAllMailboxesV2200Response`

NewGetAllMailboxesV2200ResponseWithDefaults instantiates a new GetAllMailboxesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAllMailboxesV2200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllMailboxesV2200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllMailboxesV2200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetMailboxes

`func (o *GetAllMailboxesV2200Response) GetMailboxes() []MailboxV2`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *GetAllMailboxesV2200Response) GetMailboxesOk() (*[]MailboxV2, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *GetAllMailboxesV2200Response) SetMailboxes(v []MailboxV2)`

SetMailboxes sets Mailboxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


