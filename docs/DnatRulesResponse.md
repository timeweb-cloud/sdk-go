# DnatRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**DnatRules** | [**[]DnatRuleOut**](DnatRuleOut.md) | Правила проброса портов | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewDnatRulesResponse

`func NewDnatRulesResponse(dnatRules []DnatRuleOut, meta ComponentsSchemasMeta, ) *DnatRulesResponse`

NewDnatRulesResponse instantiates a new DnatRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatRulesResponseWithDefaults

`func NewDnatRulesResponseWithDefaults() *DnatRulesResponse`

NewDnatRulesResponseWithDefaults instantiates a new DnatRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *DnatRulesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *DnatRulesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *DnatRulesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *DnatRulesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetDnatRules

`func (o *DnatRulesResponse) GetDnatRules() []DnatRuleOut`

GetDnatRules returns the DnatRules field if non-nil, zero value otherwise.

### GetDnatRulesOk

`func (o *DnatRulesResponse) GetDnatRulesOk() (*[]DnatRuleOut, bool)`

GetDnatRulesOk returns a tuple with the DnatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatRules

`func (o *DnatRulesResponse) SetDnatRules(v []DnatRuleOut)`

SetDnatRules sets DnatRules field to given value.


### GetMeta

`func (o *DnatRulesResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DnatRulesResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DnatRulesResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


