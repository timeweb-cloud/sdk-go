# MailboxForwardingIncoming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка входящик писем | 
**IsDeleteMessages** | **bool** | Удалять ли входящие письма | 
**IncomingList** | **[]string** | Список адресов для пересылки | 

## Methods

### NewMailboxForwardingIncoming

`func NewMailboxForwardingIncoming(isEnabled bool, isDeleteMessages bool, incomingList []string, ) *MailboxForwardingIncoming`

NewMailboxForwardingIncoming instantiates a new MailboxForwardingIncoming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxForwardingIncomingWithDefaults

`func NewMailboxForwardingIncomingWithDefaults() *MailboxForwardingIncoming`

NewMailboxForwardingIncomingWithDefaults instantiates a new MailboxForwardingIncoming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MailboxForwardingIncoming) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MailboxForwardingIncoming) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MailboxForwardingIncoming) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsDeleteMessages

`func (o *MailboxForwardingIncoming) GetIsDeleteMessages() bool`

GetIsDeleteMessages returns the IsDeleteMessages field if non-nil, zero value otherwise.

### GetIsDeleteMessagesOk

`func (o *MailboxForwardingIncoming) GetIsDeleteMessagesOk() (*bool, bool)`

GetIsDeleteMessagesOk returns a tuple with the IsDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteMessages

`func (o *MailboxForwardingIncoming) SetIsDeleteMessages(v bool)`

SetIsDeleteMessages sets IsDeleteMessages field to given value.


### GetIncomingList

`func (o *MailboxForwardingIncoming) GetIncomingList() []string`

GetIncomingList returns the IncomingList field if non-nil, zero value otherwise.

### GetIncomingListOk

`func (o *MailboxForwardingIncoming) GetIncomingListOk() (*[]string, bool)`

GetIncomingListOk returns a tuple with the IncomingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingList

`func (o *MailboxForwardingIncoming) SetIncomingList(v []string)`

SetIncomingList sets IncomingList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


