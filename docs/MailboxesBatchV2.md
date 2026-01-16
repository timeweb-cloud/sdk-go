# MailboxesBatchV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailboxes** | [**[]MailboxV2**](MailboxV2.md) | Массив успешно созданных почтовых ящиков | 
**Errors** | **[]map[string]interface{}** | Массив ошибок при создании почтовых ящиков | 

## Methods

### NewMailboxesBatchV2

`func NewMailboxesBatchV2(mailboxes []MailboxV2, errors []map[string]interface{}, ) *MailboxesBatchV2`

NewMailboxesBatchV2 instantiates a new MailboxesBatchV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxesBatchV2WithDefaults

`func NewMailboxesBatchV2WithDefaults() *MailboxesBatchV2`

NewMailboxesBatchV2WithDefaults instantiates a new MailboxesBatchV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxes

`func (o *MailboxesBatchV2) GetMailboxes() []MailboxV2`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *MailboxesBatchV2) GetMailboxesOk() (*[]MailboxV2, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *MailboxesBatchV2) SetMailboxes(v []MailboxV2)`

SetMailboxes sets Mailboxes field to given value.


### GetErrors

`func (o *MailboxesBatchV2) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MailboxesBatchV2) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MailboxesBatchV2) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


