# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра правила для балансировщика. Автоматически генерируется при создании. | 
**BalancerProto** | **string** | Протокол балансировщика. | 
**BalancerPort** | **float32** | Порт балансировщика. | 
**ServerProto** | **string** | Протокол сервера. | 
**ServerPort** | **float32** | Порт сервера. | 

## Methods

### NewRule

`func NewRule(id float32, balancerProto string, balancerPort float32, serverProto string, serverPort float32, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v float32)`

SetId sets Id field to given value.


### GetBalancerProto

`func (o *Rule) GetBalancerProto() string`

GetBalancerProto returns the BalancerProto field if non-nil, zero value otherwise.

### GetBalancerProtoOk

`func (o *Rule) GetBalancerProtoOk() (*string, bool)`

GetBalancerProtoOk returns a tuple with the BalancerProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancerProto

`func (o *Rule) SetBalancerProto(v string)`

SetBalancerProto sets BalancerProto field to given value.


### GetBalancerPort

`func (o *Rule) GetBalancerPort() float32`

GetBalancerPort returns the BalancerPort field if non-nil, zero value otherwise.

### GetBalancerPortOk

`func (o *Rule) GetBalancerPortOk() (*float32, bool)`

GetBalancerPortOk returns a tuple with the BalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancerPort

`func (o *Rule) SetBalancerPort(v float32)`

SetBalancerPort sets BalancerPort field to given value.


### GetServerProto

`func (o *Rule) GetServerProto() string`

GetServerProto returns the ServerProto field if non-nil, zero value otherwise.

### GetServerProtoOk

`func (o *Rule) GetServerProtoOk() (*string, bool)`

GetServerProtoOk returns a tuple with the ServerProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProto

`func (o *Rule) SetServerProto(v string)`

SetServerProto sets ServerProto field to given value.


### GetServerPort

`func (o *Rule) GetServerPort() float32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *Rule) GetServerPortOk() (*float32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *Rule) SetServerPort(v float32)`

SetServerPort sets ServerPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


