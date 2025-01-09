# FirewallRulesOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Rules** | [**[]FirewallRule**](FirewallRule.md) |  | 

## Methods

### NewFirewallRulesOutResponse

`func NewFirewallRulesOutResponse(meta Meta, rules []FirewallRule, ) *FirewallRulesOutResponse`

NewFirewallRulesOutResponse instantiates a new FirewallRulesOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRulesOutResponseWithDefaults

`func NewFirewallRulesOutResponseWithDefaults() *FirewallRulesOutResponse`

NewFirewallRulesOutResponseWithDefaults instantiates a new FirewallRulesOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallRulesOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallRulesOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallRulesOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallRulesOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *FirewallRulesOutResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FirewallRulesOutResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FirewallRulesOutResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetRules

`func (o *FirewallRulesOutResponse) GetRules() []FirewallRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallRulesOutResponse) GetRulesOk() (*[]FirewallRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallRulesOutResponse) SetRules(v []FirewallRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


