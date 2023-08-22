# FirewallRuleOutAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор правила | 
**Description** | **string** | Описание правила | 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) |  | 
**Protocol** | [**FirewallRuleProtocol**](FirewallRuleProtocol.md) |  | 
**Port** | Pointer to **string** | Порт или диапазон портов, в случае tcp или udp | [optional] 
**Cidr** | Pointer to **string** | Сетевой адрес или подсеть. Поддерживаются протоколы IPv4  и IPv6 | [optional] 
**GroupId** | **string** | Идентификатор группы правил | 

## Methods

### NewFirewallRuleOutAPI

`func NewFirewallRuleOutAPI(id string, description string, direction FirewallRuleDirection, protocol FirewallRuleProtocol, groupId string, ) *FirewallRuleOutAPI`

NewFirewallRuleOutAPI instantiates a new FirewallRuleOutAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleOutAPIWithDefaults

`func NewFirewallRuleOutAPIWithDefaults() *FirewallRuleOutAPI`

NewFirewallRuleOutAPIWithDefaults instantiates a new FirewallRuleOutAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRuleOutAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRuleOutAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRuleOutAPI) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *FirewallRuleOutAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleOutAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleOutAPI) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDirection

`func (o *FirewallRuleOutAPI) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRuleOutAPI) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRuleOutAPI) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *FirewallRuleOutAPI) GetProtocol() FirewallRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleOutAPI) GetProtocolOk() (*FirewallRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleOutAPI) SetProtocol(v FirewallRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *FirewallRuleOutAPI) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FirewallRuleOutAPI) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FirewallRuleOutAPI) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *FirewallRuleOutAPI) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *FirewallRuleOutAPI) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FirewallRuleOutAPI) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FirewallRuleOutAPI) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FirewallRuleOutAPI) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGroupId

`func (o *FirewallRuleOutAPI) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FirewallRuleOutAPI) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FirewallRuleOutAPI) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


