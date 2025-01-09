# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID правила. | 
**Description** | **string** | Описание правила. | 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) |  | 
**Protocol** | [**FirewallRuleProtocol**](FirewallRuleProtocol.md) |  | 
**Port** | Pointer to **string** | Порт или диапазон портов, в случае tcp или udp. | [optional] 
**Cidr** | Pointer to **string** | Сетевой адрес или подсеть. Поддерживаются протоколы IPv4  и IPv6. | [optional] 
**GroupId** | **string** | ID группы правил. | 

## Methods

### NewFirewallRule

`func NewFirewallRule(id string, description string, direction FirewallRuleDirection, protocol FirewallRuleProtocol, groupId string, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *FirewallRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDirection

`func (o *FirewallRule) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRule) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRule) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *FirewallRule) GetProtocol() FirewallRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRule) GetProtocolOk() (*FirewallRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRule) SetProtocol(v FirewallRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *FirewallRule) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FirewallRule) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FirewallRule) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *FirewallRule) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *FirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FirewallRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGroupId

`func (o *FirewallRule) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FirewallRule) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FirewallRule) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


