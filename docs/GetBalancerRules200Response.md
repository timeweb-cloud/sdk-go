# GetBalancerRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Rules** | [**[]Rule**](Rule.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetBalancerRules200Response

`func NewGetBalancerRules200Response(meta Meta, rules []Rule, responseId string, ) *GetBalancerRules200Response`

NewGetBalancerRules200Response instantiates a new GetBalancerRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalancerRules200ResponseWithDefaults

`func NewGetBalancerRules200ResponseWithDefaults() *GetBalancerRules200Response`

NewGetBalancerRules200ResponseWithDefaults instantiates a new GetBalancerRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalancerRules200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalancerRules200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalancerRules200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetRules

`func (o *GetBalancerRules200Response) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetBalancerRules200Response) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetBalancerRules200Response) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### GetResponseId

`func (o *GetBalancerRules200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetBalancerRules200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetBalancerRules200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


