# ForwardingOutgoingIsDisabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка исходящих писем | 
**OutgoingTo** | Pointer to **string** | Адрес для пересылки исходящих писем | [optional] 

## Methods

### NewForwardingOutgoingIsDisabled

`func NewForwardingOutgoingIsDisabled(isEnabled bool, ) *ForwardingOutgoingIsDisabled`

NewForwardingOutgoingIsDisabled instantiates a new ForwardingOutgoingIsDisabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingOutgoingIsDisabledWithDefaults

`func NewForwardingOutgoingIsDisabledWithDefaults() *ForwardingOutgoingIsDisabled`

NewForwardingOutgoingIsDisabledWithDefaults instantiates a new ForwardingOutgoingIsDisabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ForwardingOutgoingIsDisabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ForwardingOutgoingIsDisabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ForwardingOutgoingIsDisabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetOutgoingTo

`func (o *ForwardingOutgoingIsDisabled) GetOutgoingTo() string`

GetOutgoingTo returns the OutgoingTo field if non-nil, zero value otherwise.

### GetOutgoingToOk

`func (o *ForwardingOutgoingIsDisabled) GetOutgoingToOk() (*string, bool)`

GetOutgoingToOk returns a tuple with the OutgoingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTo

`func (o *ForwardingOutgoingIsDisabled) SetOutgoingTo(v string)`

SetOutgoingTo sets OutgoingTo field to given value.

### HasOutgoingTo

`func (o *ForwardingOutgoingIsDisabled) HasOutgoingTo() bool`

HasOutgoingTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


