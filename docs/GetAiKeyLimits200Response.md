# GetAiKeyLimits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKeyEntityLimits** | [**[]AiKeyLimit**](AiKeyLimit.md) |  | 
**Meta** | [**GetAiKeyLimits200ResponseMeta**](GetAiKeyLimits200ResponseMeta.md) |  | 

## Methods

### NewGetAiKeyLimits200Response

`func NewGetAiKeyLimits200Response(aiKeyEntityLimits []AiKeyLimit, meta GetAiKeyLimits200ResponseMeta, ) *GetAiKeyLimits200Response`

NewGetAiKeyLimits200Response instantiates a new GetAiKeyLimits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiKeyLimits200ResponseWithDefaults

`func NewGetAiKeyLimits200ResponseWithDefaults() *GetAiKeyLimits200Response`

NewGetAiKeyLimits200ResponseWithDefaults instantiates a new GetAiKeyLimits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKeyEntityLimits

`func (o *GetAiKeyLimits200Response) GetAiKeyEntityLimits() []AiKeyLimit`

GetAiKeyEntityLimits returns the AiKeyEntityLimits field if non-nil, zero value otherwise.

### GetAiKeyEntityLimitsOk

`func (o *GetAiKeyLimits200Response) GetAiKeyEntityLimitsOk() (*[]AiKeyLimit, bool)`

GetAiKeyEntityLimitsOk returns a tuple with the AiKeyEntityLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyEntityLimits

`func (o *GetAiKeyLimits200Response) SetAiKeyEntityLimits(v []AiKeyLimit)`

SetAiKeyEntityLimits sets AiKeyEntityLimits field to given value.


### GetMeta

`func (o *GetAiKeyLimits200Response) GetMeta() GetAiKeyLimits200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiKeyLimits200Response) GetMetaOk() (*GetAiKeyLimits200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiKeyLimits200Response) SetMeta(v GetAiKeyLimits200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


