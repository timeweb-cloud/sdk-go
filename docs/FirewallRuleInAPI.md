# FirewallRuleInAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Описание правила. | [optional] 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) |  | 
**Port** | Pointer to **string** | Порт или диапазон портов, в случае tcp или udp. | [optional] 
**Protocol** | [**FirewallRuleProtocol**](FirewallRuleProtocol.md) |  | 
**Cidr** | Pointer to **string** | Сетевой адрес или подсеть. Поддерживаются протоколы IPv4  и IPv.6 | [optional] 

## Methods

### NewFirewallRuleInAPI

`func NewFirewallRuleInAPI(direction FirewallRuleDirection, protocol FirewallRuleProtocol, ) *FirewallRuleInAPI`

NewFirewallRuleInAPI instantiates a new FirewallRuleInAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleInAPIWithDefaults

`func NewFirewallRuleInAPIWithDefaults() *FirewallRuleInAPI`

NewFirewallRuleInAPIWithDefaults instantiates a new FirewallRuleInAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FirewallRuleInAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleInAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleInAPI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRuleInAPI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *FirewallRuleInAPI) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRuleInAPI) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRuleInAPI) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetPort

`func (o *FirewallRuleInAPI) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FirewallRuleInAPI) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FirewallRuleInAPI) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *FirewallRuleInAPI) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FirewallRuleInAPI) GetProtocol() FirewallRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleInAPI) GetProtocolOk() (*FirewallRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleInAPI) SetProtocol(v FirewallRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetCidr

`func (o *FirewallRuleInAPI) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FirewallRuleInAPI) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FirewallRuleInAPI) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FirewallRuleInAPI) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


