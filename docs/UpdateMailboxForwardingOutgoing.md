# UpdateMailboxForwardingOutgoing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка исходящих писем. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | 
**OutgoingTo** | Pointer to **string** | Адрес для пересылки исходящих писем. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewUpdateMailboxForwardingOutgoing

`func NewUpdateMailboxForwardingOutgoing(isEnabled bool, ) *UpdateMailboxForwardingOutgoing`

NewUpdateMailboxForwardingOutgoing instantiates a new UpdateMailboxForwardingOutgoing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxForwardingOutgoingWithDefaults

`func NewUpdateMailboxForwardingOutgoingWithDefaults() *UpdateMailboxForwardingOutgoing`

NewUpdateMailboxForwardingOutgoingWithDefaults instantiates a new UpdateMailboxForwardingOutgoing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateMailboxForwardingOutgoing) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateMailboxForwardingOutgoing) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateMailboxForwardingOutgoing) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetOutgoingTo

`func (o *UpdateMailboxForwardingOutgoing) GetOutgoingTo() string`

GetOutgoingTo returns the OutgoingTo field if non-nil, zero value otherwise.

### GetOutgoingToOk

`func (o *UpdateMailboxForwardingOutgoing) GetOutgoingToOk() (*string, bool)`

GetOutgoingToOk returns a tuple with the OutgoingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTo

`func (o *UpdateMailboxForwardingOutgoing) SetOutgoingTo(v string)`

SetOutgoingTo sets OutgoingTo field to given value.

### HasOutgoingTo

`func (o *UpdateMailboxForwardingOutgoing) HasOutgoingTo() bool`

HasOutgoingTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


