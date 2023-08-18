# UpdateMailboxForwardingIncoming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка входящик писем | 
**IsDeleteMessages** | Pointer to **bool** | Удалять ли входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**IncomingList** | Pointer to **[]string** | Список адресов для пересылки. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewUpdateMailboxForwardingIncoming

`func NewUpdateMailboxForwardingIncoming(isEnabled bool, ) *UpdateMailboxForwardingIncoming`

NewUpdateMailboxForwardingIncoming instantiates a new UpdateMailboxForwardingIncoming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxForwardingIncomingWithDefaults

`func NewUpdateMailboxForwardingIncomingWithDefaults() *UpdateMailboxForwardingIncoming`

NewUpdateMailboxForwardingIncomingWithDefaults instantiates a new UpdateMailboxForwardingIncoming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateMailboxForwardingIncoming) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateMailboxForwardingIncoming) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateMailboxForwardingIncoming) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsDeleteMessages

`func (o *UpdateMailboxForwardingIncoming) GetIsDeleteMessages() bool`

GetIsDeleteMessages returns the IsDeleteMessages field if non-nil, zero value otherwise.

### GetIsDeleteMessagesOk

`func (o *UpdateMailboxForwardingIncoming) GetIsDeleteMessagesOk() (*bool, bool)`

GetIsDeleteMessagesOk returns a tuple with the IsDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteMessages

`func (o *UpdateMailboxForwardingIncoming) SetIsDeleteMessages(v bool)`

SetIsDeleteMessages sets IsDeleteMessages field to given value.

### HasIsDeleteMessages

`func (o *UpdateMailboxForwardingIncoming) HasIsDeleteMessages() bool`

HasIsDeleteMessages returns a boolean if a field has been set.

### GetIncomingList

`func (o *UpdateMailboxForwardingIncoming) GetIncomingList() []string`

GetIncomingList returns the IncomingList field if non-nil, zero value otherwise.

### GetIncomingListOk

`func (o *UpdateMailboxForwardingIncoming) GetIncomingListOk() (*[]string, bool)`

GetIncomingListOk returns a tuple with the IncomingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingList

`func (o *UpdateMailboxForwardingIncoming) SetIncomingList(v []string)`

SetIncomingList sets IncomingList field to given value.

### HasIncomingList

`func (o *UpdateMailboxForwardingIncoming) HasIncomingList() bool`

HasIncomingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


