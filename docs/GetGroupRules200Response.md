# GetGroupRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Rules** | [**[]FirewallRuleOutAPI**](FirewallRuleOutAPI.md) | Массив объектов Firewall правил | 

## Methods

### NewGetGroupRules200Response

`func NewGetGroupRules200Response(responseId string, meta Meta, rules []FirewallRuleOutAPI, ) *GetGroupRules200Response`

NewGetGroupRules200Response instantiates a new GetGroupRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupRules200ResponseWithDefaults

`func NewGetGroupRules200ResponseWithDefaults() *GetGroupRules200Response`

NewGetGroupRules200ResponseWithDefaults instantiates a new GetGroupRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetGroupRules200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetGroupRules200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetGroupRules200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetGroupRules200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroupRules200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroupRules200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetRules

`func (o *GetGroupRules200Response) GetRules() []FirewallRuleOutAPI`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetGroupRules200Response) GetRulesOk() (*[]FirewallRuleOutAPI, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetGroupRules200Response) SetRules(v []FirewallRuleOutAPI)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


