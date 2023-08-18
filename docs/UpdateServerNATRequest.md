# UpdateServerNATRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatMode** | **string** | Правило для маршрутизации трафика. \\  Досутпные правила: &#x60;dnat_and_snat&#x60; – разрешен входящий и исходящий трафик, &#x60;snat&#x60; – разрешен только исходящий трафик, &#x60;no_nat&#x60; – разрешен трафик только в локальной сети. | 

## Methods

### NewUpdateServerNATRequest

`func NewUpdateServerNATRequest(natMode string, ) *UpdateServerNATRequest`

NewUpdateServerNATRequest instantiates a new UpdateServerNATRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerNATRequestWithDefaults

`func NewUpdateServerNATRequestWithDefaults() *UpdateServerNATRequest`

NewUpdateServerNATRequestWithDefaults instantiates a new UpdateServerNATRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatMode

`func (o *UpdateServerNATRequest) GetNatMode() string`

GetNatMode returns the NatMode field if non-nil, zero value otherwise.

### GetNatModeOk

`func (o *UpdateServerNATRequest) GetNatModeOk() (*string, bool)`

GetNatModeOk returns a tuple with the NatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatMode

`func (o *UpdateServerNATRequest) SetNatMode(v string)`

SetNatMode sets NatMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


