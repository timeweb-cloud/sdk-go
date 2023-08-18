# CreateRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancerProto** | **string** | Протокол балансировщика. | 
**BalancerPort** | **float32** | Порт балансировщика. | 
**ServerProto** | **string** | Протокол сервера. | 
**ServerPort** | **float32** | Порт сервера. | 

## Methods

### NewCreateRule

`func NewCreateRule(balancerProto string, balancerPort float32, serverProto string, serverPort float32, ) *CreateRule`

NewCreateRule instantiates a new CreateRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleWithDefaults

`func NewCreateRuleWithDefaults() *CreateRule`

NewCreateRuleWithDefaults instantiates a new CreateRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancerProto

`func (o *CreateRule) GetBalancerProto() string`

GetBalancerProto returns the BalancerProto field if non-nil, zero value otherwise.

### GetBalancerProtoOk

`func (o *CreateRule) GetBalancerProtoOk() (*string, bool)`

GetBalancerProtoOk returns a tuple with the BalancerProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancerProto

`func (o *CreateRule) SetBalancerProto(v string)`

SetBalancerProto sets BalancerProto field to given value.


### GetBalancerPort

`func (o *CreateRule) GetBalancerPort() float32`

GetBalancerPort returns the BalancerPort field if non-nil, zero value otherwise.

### GetBalancerPortOk

`func (o *CreateRule) GetBalancerPortOk() (*float32, bool)`

GetBalancerPortOk returns a tuple with the BalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancerPort

`func (o *CreateRule) SetBalancerPort(v float32)`

SetBalancerPort sets BalancerPort field to given value.


### GetServerProto

`func (o *CreateRule) GetServerProto() string`

GetServerProto returns the ServerProto field if non-nil, zero value otherwise.

### GetServerProtoOk

`func (o *CreateRule) GetServerProtoOk() (*string, bool)`

GetServerProtoOk returns a tuple with the ServerProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProto

`func (o *CreateRule) SetServerProto(v string)`

SetServerProto sets ServerProto field to given value.


### GetServerPort

`func (o *CreateRule) GetServerPort() float32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *CreateRule) GetServerPortOk() (*float32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *CreateRule) SetServerPort(v float32)`

SetServerPort sets ServerPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


