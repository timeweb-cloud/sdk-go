# GetAiModelParameters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelParameters** | [**[]ModelParameter**](ModelParameter.md) |  | 
**Meta** | [**GetAiModelParameters200ResponseMeta**](GetAiModelParameters200ResponseMeta.md) |  | 

## Methods

### NewGetAiModelParameters200Response

`func NewGetAiModelParameters200Response(modelParameters []ModelParameter, meta GetAiModelParameters200ResponseMeta, ) *GetAiModelParameters200Response`

NewGetAiModelParameters200Response instantiates a new GetAiModelParameters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiModelParameters200ResponseWithDefaults

`func NewGetAiModelParameters200ResponseWithDefaults() *GetAiModelParameters200Response`

NewGetAiModelParameters200ResponseWithDefaults instantiates a new GetAiModelParameters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelParameters

`func (o *GetAiModelParameters200Response) GetModelParameters() []ModelParameter`

GetModelParameters returns the ModelParameters field if non-nil, zero value otherwise.

### GetModelParametersOk

`func (o *GetAiModelParameters200Response) GetModelParametersOk() (*[]ModelParameter, bool)`

GetModelParametersOk returns a tuple with the ModelParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelParameters

`func (o *GetAiModelParameters200Response) SetModelParameters(v []ModelParameter)`

SetModelParameters sets ModelParameters field to given value.


### GetMeta

`func (o *GetAiModelParameters200Response) GetMeta() GetAiModelParameters200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiModelParameters200Response) GetMetaOk() (*GetAiModelParameters200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiModelParameters200Response) SetMeta(v GetAiModelParameters200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


