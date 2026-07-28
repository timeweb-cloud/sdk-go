# GetModelsV3200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | [**[]ModelV3**](ModelV3.md) |  | 
**Meta** | [**GetModelsV3200ResponseMeta**](GetModelsV3200ResponseMeta.md) |  | 

## Methods

### NewGetModelsV3200Response

`func NewGetModelsV3200Response(models []ModelV3, meta GetModelsV3200ResponseMeta, ) *GetModelsV3200Response`

NewGetModelsV3200Response instantiates a new GetModelsV3200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetModelsV3200ResponseWithDefaults

`func NewGetModelsV3200ResponseWithDefaults() *GetModelsV3200Response`

NewGetModelsV3200ResponseWithDefaults instantiates a new GetModelsV3200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *GetModelsV3200Response) GetModels() []ModelV3`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GetModelsV3200Response) GetModelsOk() (*[]ModelV3, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GetModelsV3200Response) SetModels(v []ModelV3)`

SetModels sets Models field to given value.


### GetMeta

`func (o *GetModelsV3200Response) GetMeta() GetModelsV3200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetModelsV3200Response) GetMetaOk() (*GetModelsV3200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetModelsV3200Response) SetMeta(v GetModelsV3200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


