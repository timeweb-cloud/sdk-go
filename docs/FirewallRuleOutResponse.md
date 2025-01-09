# FirewallRuleOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Rule** | [**FirewallRule**](FirewallRule.md) |  | 

## Methods

### NewFirewallRuleOutResponse

`func NewFirewallRuleOutResponse(rule FirewallRule, ) *FirewallRuleOutResponse`

NewFirewallRuleOutResponse instantiates a new FirewallRuleOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleOutResponseWithDefaults

`func NewFirewallRuleOutResponseWithDefaults() *FirewallRuleOutResponse`

NewFirewallRuleOutResponseWithDefaults instantiates a new FirewallRuleOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallRuleOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallRuleOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallRuleOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallRuleOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetRule

`func (o *FirewallRuleOutResponse) GetRule() FirewallRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallRuleOutResponse) GetRuleOk() (*FirewallRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallRuleOutResponse) SetRule(v FirewallRule)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


